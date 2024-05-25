package global_vpn_connection

import (
	"fmt"
	"math"

	"github.com/kloudlite/api/pkg/iputils"

	"github.com/kloudlite/api/apps/infra/internal/domain/ports"
	domainT "github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	fc "github.com/kloudlite/api/apps/infra/internal/entities/field-constants"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
	t "github.com/kloudlite/api/pkg/types"
	ct "github.com/kloudlite/operator/apis/common-types"
	wgv1 "github.com/kloudlite/operator/apis/wireguard/v1"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (d *Domain) getGlobalVPNConnectionPeers(vpns []*entities.GlobalVPNConnection) ([]wgv1.Peer, error) {
	peers := make([]wgv1.Peer, 0, len(vpns))
	for _, c := range vpns {
		if c.ParsedWgParams != nil {
			if c.ParsedWgParams.WgPublicKey == "" {
				continue
			}

			if c.ParsedWgParams.NodePort == nil {
				d.Logger.Infof("nodeport not available for gvpn %s", c.Name)
				continue
			}

			peers = append(peers, wgv1.Peer{
				ClusterName: c.ClusterName,
				IP:          c.ParsedWgParams.IP,
				PublicKey:   c.ParsedWgParams.WgPublicKey,
				Endpoint:    fmt.Sprintf("%s:%s", c.ClusterPublicEndpoint, *c.ParsedWgParams.NodePort),
				AllowedIPs:  []string{c.ClusterSvcCIDR},
			})
		}
	}

	return peers, nil
}

func (d *Domain) listGlobalVPNConnections(ctx domainT.InfraContext, vpnName string) ([]*entities.GlobalVPNConnection, error) {
	return d.GlobalVPNClusterConnectionRepo.Find(ctx, repos.Query{
		Filter: repos.Filter{
			fields.AccountName:  ctx.AccountName,
			fields.MetadataName: vpnName,
		},
	})
}

func (d *Domain) reconGlobalVPNConnections(ctx domainT.InfraContext, vpnName string) error {
	vpns, err := d.listGlobalVPNConnections(ctx, vpnName)
	if err != nil {
		return errors.NewE(err)
	}

	peers, err := d.getGlobalVPNConnectionPeers(vpns)
	if err != nil {
		return err
	}

	publicPeers, privatePeers, err := d.BuildPeersFromGlobalVPNDevices(ctx, vpnName)
	if err != nil {
		return err
	}

	peers = append(peers, publicPeers...)
	peers = append(peers, privatePeers...)

	for _, xcc := range vpns {
		if fmt.Sprintf("%#v", xcc.Spec.Peers) == fmt.Sprintf("%#v", peers) {
			continue
		}

		xcc.Spec.Peers = peers
		unp, err := d.GlobalVPNClusterConnectionRepo.Patch(
			ctx,
			repos.Filter{
				fields.AccountName:  ctx.AccountName,
				fields.ClusterName:  xcc.ClusterName,
				fields.MetadataName: xcc.Name,
			},
			common.PatchForUpdate(ctx, xcc, common.PatchOpts{XPatch: map[string]any{fc.GlobalVPNConnectionSpecPeers: peers}}),
		)
		if err != nil {
			return errors.NewE(err)
		}

		if err := d.applyGlobalVPNConnection(ctx, unp); err != nil {
			return errors.NewE(err)
		}
	}

	return nil
}

func (d *Domain) claimNextClusterSvcCIDR(ctx domainT.InfraContext, clusterName string, gvpnName string) (string, error) {
	for {
		freeCIDR, err := d.FreeClusterSvcCIDRRepo.FindOne(ctx, repos.Filter{
			fc.AccountName:                     ctx.AccountName,
			fc.FreeClusterSvcCIDRGlobalVPNName: gvpnName,
		})
		if err != nil {
			return "", err
		}

		if freeCIDR == nil {
			gvpn, err := d.FindGlobalVPN(ctx, gvpnName)
			if err != nil {
				return "", err
			}

			numIPsPerCluster := int(math.Pow(2, float64(32-gvpn.AllocatableCIDRSuffix)))

			ipv4StartingAddr, err := iputils.GenIPAddr(gvpn.CIDR, gvpn.NumReservedIPsForNonClusterUse+gvpn.NumAllocatedClusterCIDRs*numIPsPerCluster)
			if err != nil {
				if errors.Is(err, iputils.ErrIPsMaxedOut) {
					return "", errors.NewEf(err, "max IPs exceeded, won't be able to allocate any more IPs")
				}
				return "", err
			}

			if _, err := d.FreeClusterSvcCIDRRepo.Create(ctx, &entities.FreeClusterSvcCIDR{
				AccountName:    ctx.AccountName,
				GlobalVPNName:  gvpnName,
				ClusterSvcCIDR: fmt.Sprintf("%s/%d", ipv4StartingAddr, gvpn.AllocatableCIDRSuffix),
			}); err != nil {
				// FIXME: handle gracefully
				continue
			}
			if _, err := d.GlobalVPNClusterConnectionRepo.PatchById(ctx, gvpn.Id, repos.Document{"$inc": map[string]any{fc.GlobalVPNNumAllocatedClusterCIDRs: 1}}); err != nil {
				continue
			}
			continue
		}

		cidr := freeCIDR.ClusterSvcCIDR

		if _, err := d.ClaimClusterSvcCIDRRepo.Create(ctx, &entities.ClaimClusterSvcCIDR{
			AccountName:      ctx.AccountName,
			GlobalVPNName:    gvpnName,
			ClusterSvcCIDR:   cidr,
			ClaimedByCluster: clusterName,
		}); err != nil {
			d.Logger.Warnf("cluster svc CIDR %s, already claimed, trying another", cidr)
			continue
		}

		if err := d.FreeClusterSvcCIDRRepo.DeleteById(ctx, freeCIDR.Id); err != nil {
			return "", err
		}

		return cidr, nil
	}
}

func (d *Domain) createGlobalVPNConnection(ctx domainT.InfraContext, gvpnConn entities.GlobalVPNConnection) (*entities.GlobalVPNConnection, error) {
	gvpnConn.ResourceMetadata.CreatedBy = common.CreatedOrUpdatedByKloudlite
	gvpnConn.ResourceMetadata.LastUpdatedBy = common.CreatedOrUpdatedByKloudlite

	var err error
	gvpnConn.ClusterSvcCIDR, err = d.claimNextClusterSvcCIDR(ctx, gvpnConn.ClusterName, gvpnConn.GlobalVPNName)
	if err != nil {
		return nil, err
	}

	gvpn, err := d.FindGlobalVPN(ctx, gvpnConn.GlobalVPNName)
	if err != nil {
		return nil, err
	}

	if gvpnConn.Spec.WgInterface == nil {
		gvpnConn.Spec.WgInterface = &gvpn.WgInterface
	}

	gvpnConn.SyncStatus = t.GenSyncStatus(t.SyncActionApply, 0)

	gvpnDevice, err := d.CreateGlobalVPNDevice(ctx, entities.GlobalVPNDevice{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("cluster-gateway-%s", gvpnConn.ClusterName),
		},
		AccountName:    ctx.AccountName,
		GlobalVPNName:  gvpnConn.Name,
		CreationMethod: domainT.GlobalVPNConnectionDeviceMethod,
	})
	if err != nil {
		return nil, err
	}

	gvpnConn.DeviceRef = entities.GlobalVPNConnDeviceRef{
		Name:   gvpnDevice.Name,
		IPAddr: gvpnDevice.IPAddr,
	}

	gv, err := d.GlobalVPNClusterConnectionRepo.Create(ctx, &gvpnConn)
	if err != nil {
		return nil, err
	}

	if err := d.applyGlobalVPNConnection(ctx, gv); err != nil {
		return nil, err
	}

	return gv, nil
}

func (d *Domain) DeleteGlobalVPNConnection(ctx domainT.InfraContext, clusterName string, gvpnName string) error {
	gvpnConn, err := d.GlobalVPNClusterConnectionRepo.FindOne(ctx, repos.Filter{
		fields.AccountName:  ctx.AccountName,
		fields.ClusterName:  clusterName,
		fields.MetadataName: gvpnName,
	})
	if err != nil {
		return errors.NewE(err)
	}
	if gvpnConn == nil {
		return errors.Newf("no global vpn connection with name (%s) not found, for cluster (%s)", gvpnName, clusterName)
	}

	if err := d.DeleteGlobalVPNDevice(ctx, gvpnName, gvpnConn.DeviceRef.Name); err != nil {
		return errors.NewE(err)
	}

	return nil
}

func (d *Domain) EnsureGlobalVPNConnection(ctx domainT.InfraContext, clusterName string, groupName string, clusterPublicEndpoint string) (*entities.GlobalVPNConnection, error) {
	gvpn, err := d.GlobalVPNClusterConnectionRepo.FindOne(ctx, repos.Filter{
		fields.AccountName:  ctx.AccountName,
		fields.ClusterName:  clusterName,
		fields.MetadataName: groupName,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if gvpn != nil {
		if err := d.applyGlobalVPNConnection(ctx, gvpn); err != nil {
			return nil, err
		}
		return gvpn, nil
	}

	return d.createGlobalVPNConnection(ctx, entities.GlobalVPNConnection{
		GlobalVPN: wgv1.GlobalVPN{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "wireguard.kloudlite.io/v1",
				// FIXME: look into it
				Kind: "GlobalVPN",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: groupName,
			},
			Spec: wgv1.GlobVPNSpec{
				WgRef: ct.SecretRef{
					Name:      fmt.Sprintf("global-vpn-params-%s", groupName),
					Namespace: "kube-system",
				},
			},
		},
		GlobalVPNName:         groupName,
		ResourceMetadata:      common.ResourceMetadata{DisplayName: groupName, CreatedBy: common.CreatedOrUpdatedByKloudlite, LastUpdatedBy: common.CreatedOrUpdatedByKloudlite},
		AccountName:           ctx.AccountName,
		ClusterName:           clusterName,
		ClusterPublicEndpoint: clusterPublicEndpoint,
		ParsedWgParams:        nil,
	})
}

func (d *Domain) applyGlobalVPNConnection(ctx domainT.InfraContext, gvpn *entities.GlobalVPNConnection) error {
	if err := d.ResDispatcher.ApplyToTargetCluster(ctx, ctx.AccountName, gvpn.ClusterName, &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      gvpn.Spec.WgRef.Name,
			Namespace: gvpn.Spec.WgRef.Namespace,
		},
		StringData: map[string]string{
			"ip": gvpn.DeviceRef.IPAddr,
		},
	}, 0); err != nil {
		return err
	}
	return d.ResDispatcher.ApplyToTargetCluster(ctx, gvpn.AccountName, gvpn.ClusterName, &gvpn.GlobalVPN, gvpn.RecordVersion)
}

func (d *Domain) FindGlobalVPNConnection(ctx domainT.InfraContext, groupName string, clusterName string) (*entities.GlobalVPNConnection, error) {
	cc, err := d.GlobalVPNClusterConnectionRepo.FindOne(ctx, repos.Filter{
		fields.AccountName:                  ctx.AccountName,
		fields.ClusterName:                  clusterName,
		fc.GlobalVPNConnectionGlobalVPNName: groupName,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}
	if cc == nil {
		return nil, errors.Newf("global vpn with name (%s) not found, for cluster (%s)", groupName, clusterName)
	}
	return cc, nil
}

func (d *Domain) OnGlobalVPNConnectionDeleteMessage(ctx domainT.InfraContext, clusterName string, gvpnConn entities.GlobalVPNConnection) error {
	currRecord, err := d.GlobalVPNClusterConnectionRepo.FindOne(ctx, repos.Filter{
		fields.AccountName:  ctx.AccountName,
		fields.ClusterName:  clusterName,
		fields.MetadataName: gvpnConn.Name,
	})
	if err != nil {
		return err
	}

	if err := d.GlobalVPNClusterConnectionRepo.DeleteOne(ctx, repos.Filter{
		fields.AccountName:  ctx.AccountName,
		fields.ClusterName:  clusterName,
		fields.MetadataName: gvpnConn.Name,
	}); err != nil {
		return errors.NewE(err)
	}

	if currRecord.DeviceRef.Name != "" {
		if err := d.DeleteGlobalVPNDevice(ctx, currRecord.GlobalVPNName, currRecord.DeviceRef.Name); err != nil {
			return errors.NewE(err)
		}
	}

	d.ResourceEventPublisher.PublishResourceEvent(ctx, clusterName, ports.ResourceTypeClusterConnection, gvpnConn.Name, ports.PublishDelete)
	return err
}

func (d *Domain) OnGlobalVPNConnectionUpdateMessage(ctx domainT.InfraContext, clusterName string, gvpn entities.GlobalVPNConnection, status types.ResourceStatus, opts domainT.UpdateAndDeleteOpts) error {
	xconn, err := d.FindGlobalVPNConnection(ctx, gvpn.Name, clusterName)
	if err != nil {
		return errors.NewE(err)
	}

	// INFO: BYOK cluster does not have any status update message
	if d.IsBYOKCluster(ctx, xconn.ClusterName) {
		if err := d.MarkBYOKClusterReady(ctx, clusterName, opts.MessageTimestamp); err != nil {
			return errors.NewE(err)
		}
		// if _, err := d.byokClusterRepo.PatchOne(ctx, entities.UniqueBYOKClusterFilter(ctx.AccountName, clusterName), repos.Document{
		// 	fc.SyncStatusState:        t.SyncStateUpdatedAtAgent,
		// 	fc.SyncStatusLastSyncedAt: opts.MessageTimestamp,
		// 	fc.SyncStatusError:        nil,
		// }); err != nil {
		// 	return errors.NewE(err)
		// }
	}

	if _, err := d.MatchRecordVersion(gvpn.Annotations, xconn.RecordVersion); err != nil {
		return d.ResyncToTargetCluster(ctx, xconn.SyncStatus.Action, clusterName, &xconn.GlobalVPN, xconn.RecordVersion)
	}

	recordVersion, err := d.MatchRecordVersion(gvpn.Annotations, xconn.RecordVersion)
	if err != nil {
		return errors.NewE(err)
	}

	patchDoc := common.PatchForSyncFromAgent(&gvpn, recordVersion, status, common.PatchOpts{
		MessageTimestamp: opts.MessageTimestamp,
	})

	if gvpn.ParsedWgParams != nil {
		patchDoc[fc.GlobalVPNConnectionParsedWgParams] = gvpn.ParsedWgParams
	}

	ugvpn, err := d.GlobalVPNClusterConnectionRepo.PatchById(ctx, xconn.Id, patchDoc)
	if err != nil {
		return errors.NewE(err)
	}

	if err := d.reconGlobalVPNConnections(ctx, xconn.Name); err != nil {
		return errors.NewE(err)
	}

	if err := d.SyncKloudliteDeviceOnPlatform(ctx, xconn.GlobalVPNName); err != nil {
		return errors.NewE(err)
	}

	// if err := d.syncKloudliteDeviceOnCluster(ctx, xconn.GlobalVPNName); err != nil {
	// 	return errors.NewE(err)
	// }

	d.ResourceEventPublisher.PublishResourceEvent(ctx, clusterName, ports.ResourceTypeClusterConnection, ugvpn.Name, ports.PublishUpdate)
	return nil
}

func (d *Domain) OnGlobalVPNConnectionApplyError(ctx domainT.InfraContext, clusterName string, name string, errMsg string, opts domainT.UpdateAndDeleteOpts) error {
	unp, err := d.GlobalVPNClusterConnectionRepo.Patch(
		ctx,
		repos.Filter{
			fields.AccountName:  ctx.AccountName,
			fields.ClusterName:  clusterName,
			fields.MetadataName: name,
		},
		common.PatchForErrorFromAgent(
			errMsg,
			common.PatchOpts{
				MessageTimestamp: opts.MessageTimestamp,
			},
		),
	)
	if err != nil {
		return errors.NewE(err)
	}
	d.ResourceEventPublisher.PublishResourceEvent(ctx, clusterName, ports.ResourceTypeClusterConnection, unp.Name, ports.PublishUpdate)
	return errors.NewE(err)
}
