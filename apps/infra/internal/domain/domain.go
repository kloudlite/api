package domain

import (
	"io"
	"os"
	"time"

	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"

	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/k8s"

	byok_clusters "github.com/kloudlite/api/apps/infra/internal/domain/byok-clusters"
	"github.com/kloudlite/api/apps/infra/internal/domain/clusters"
	commonDomain "github.com/kloudlite/api/apps/infra/internal/domain/common"
	global_vpn "github.com/kloudlite/api/apps/infra/internal/domain/global-vpn"
	global_vpn_connection "github.com/kloudlite/api/apps/infra/internal/domain/global-vpn-connection"
	global_vpn_device "github.com/kloudlite/api/apps/infra/internal/domain/global-vpn-device"
	helm_releases "github.com/kloudlite/api/apps/infra/internal/domain/helm-releases"
	"github.com/kloudlite/api/apps/infra/internal/domain/nodepool"
	persistent_volume "github.com/kloudlite/api/apps/infra/internal/domain/persistent-volume"
	persistent_volume_claim "github.com/kloudlite/api/apps/infra/internal/domain/persistent-volume-claim"
	"github.com/kloudlite/api/apps/infra/internal/domain/ports"
	provider_secret "github.com/kloudlite/api/apps/infra/internal/domain/provider-secret"
	"github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"

	fc "github.com/kloudlite/api/apps/infra/internal/entities/field-constants"
	"github.com/kloudlite/api/apps/infra/internal/env"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/iam"
	message_office_internal "github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/message-office-internal"
	"github.com/kloudlite/api/pkg/logging"
	"github.com/kloudlite/api/pkg/repos"
	wgv1 "github.com/kloudlite/operator/apis/wireguard/v1"
	"go.uber.org/fx"
)

type (
	ClusterDomainImpl               struct{ *clusters.Domain }
	BYOKClusterDomainImpl           struct{ *byok_clusters.Domain }
	HelmReleaseDomainImpl           struct{ *helm_releases.Domain }
	CommonDomainImpl                struct{ *commonDomain.Domain }
	NodepoolDomainImpl              struct{ *nodepool.Domain }
	ProviderSecretDomainImpl        struct{ *provider_secret.Domain }
	PersistentVolumeDomainImpl      struct{ *persistent_volume.Domain }
	PersistentVolumeClaimDomainImpl struct {
		*persistent_volume_claim.Domain
	}
	GlobalVPNDomainImpl           struct{ *global_vpn.Domain }
	GlobalVPNConnectionDomainImpl struct{ *global_vpn_connection.Domain }
	GlobalVPNDeviceDomainImpl     struct{ *global_vpn_device.Domain }
	IngressHostsDomainImpl        struct{ *global_vpn_device.Domain }
)

type domain struct {
	CommonDomainImpl
	ClusterDomainImpl
	BYOKClusterDomainImpl
	HelmReleaseDomainImpl
	NodepoolDomainImpl
	ProviderSecretDomainImpl

	PersistentVolumeDomainImpl
	PersistentVolumeClaimDomainImpl

	GlobalVPNDomainImpl
	GlobalVPNConnectionDomainImpl
	GlobalVPNDeviceDomainImpl

	clusterRepo repos.DbRepo[*entities.Cluster]

	byokClusterRepo           repos.DbRepo[*entities.BYOKCluster]
	clusterManagedServiceRepo repos.DbRepo[*entities.ClusterManagedService]
	helmReleaseRepo           repos.DbRepo[*entities.HelmRelease]
	nodeRepo                  repos.DbRepo[*entities.Node]
	nodePoolRepo              repos.DbRepo[*entities.NodePool]

	gvpnConnRepo            repos.DbRepo[*entities.GlobalVPNConnection]
	freeClusterSvcCIDRRepo  repos.DbRepo[*entities.FreeClusterSvcCIDR]
	claimClusterSvcCIDRRepo repos.DbRepo[*entities.ClaimClusterSvcCIDR]

	gvpnRepo        repos.DbRepo[*entities.GlobalVPN]
	gvpnDevicesRepo repos.DbRepo[*entities.GlobalVPNDevice]

	// deviceAddressPoolRepo repos.DbRepo[*entities.GlobalVPNDeviceAddressPool]
	freeDeviceIpRepo  repos.DbRepo[*entities.FreeDeviceIP]
	claimDeviceIPRepo repos.DbRepo[*entities.ClaimDeviceIP]

	domainEntryRepo      repos.DbRepo[*entities.DomainEntry]
	secretRepo           repos.DbRepo[*entities.CloudProviderSecret]
	pvcRepo              repos.DbRepo[*entities.PersistentVolumeClaim]
	namespaceRepo        repos.DbRepo[*entities.Namespace]
	pvRepo               repos.DbRepo[*entities.PersistentVolume]
	volumeAttachmentRepo repos.DbRepo[*entities.VolumeAttachment]

	iamClient                   iam.IAMClient
	accountsSvc                 AccountsSvc
	messageOfficeInternalClient message_office_internal.MessageOfficeInternalClient
	resDispatcher               ResourceDispatcher
	k8sClient                   k8s.Client
	resourceEventPublisher      ResourceEventPublisher

	msvcTemplates    []*entities.MsvcTemplate
	msvcTemplatesMap map[string]map[string]*entities.MsvcTemplateEntry
}

var Module = fx.Module("domain",
	fx.Provide(
		func(
			ev *env.Env,
			clusterRepo repos.DbRepo[*entities.Cluster],
			byokClusterRepo repos.DbRepo[*entities.BYOKCluster],
			clustermanagedserviceRepo repos.DbRepo[*entities.ClusterManagedService],
			nodeRepo repos.DbRepo[*entities.Node],
			nodePoolRepo repos.DbRepo[*entities.NodePool],
			secretRepo repos.DbRepo[*entities.CloudProviderSecret],
			domainNameRepo repos.DbRepo[*entities.DomainEntry],

			resourceDispatcher ResourceDispatcher,
			resourceDispatcherV2 ports.ResourceDispatcher,

			helmReleaseRepo repos.DbRepo[*entities.HelmRelease],

			gvpnConnRepo repos.DbRepo[*entities.GlobalVPNConnection],
			gvpnRepo repos.DbRepo[*entities.GlobalVPN],
			gvpnDevicesRepo repos.DbRepo[*entities.GlobalVPNDevice],

			freeDeviceIpRepo repos.DbRepo[*entities.FreeDeviceIP],
			claimDeviceIPRepo repos.DbRepo[*entities.ClaimDeviceIP],

			freeClusterSvcCIDRRepo repos.DbRepo[*entities.FreeClusterSvcCIDR],
			claimClusterSvcCIDRRepo repos.DbRepo[*entities.ClaimClusterSvcCIDR],

			pvcRepo repos.DbRepo[*entities.PersistentVolumeClaim],
			pvRepo repos.DbRepo[*entities.PersistentVolume],
			namespaceRepo repos.DbRepo[*entities.Namespace],
			volumeAttachmentRepo repos.DbRepo[*entities.VolumeAttachment],

			k8sClient k8s.Client,

			iamClient iam.IAMClient,

			iamSvc ports.IAMSvc,
			msgOfficeSvc ports.MessageOfficeInternalSvc,
			accountsSvcV2 ports.AccountsSvc,

			accountsSvc AccountsSvc,
			msgOfficeInternalClient message_office_internal.MessageOfficeInternalClient,
			logger logging.Logger,

			resourceEventPublisher ResourceEventPublisher,
			resourceEventPublisherV2 ports.ResourceEventPublisher,
		) (Domain, error) {
			open, err := os.Open(ev.MsvcTemplateFilePath)
			if err != nil {
				return nil, errors.NewE(err)
			}

			b, err := io.ReadAll(open)
			if err != nil {
				return nil, errors.NewE(err)
			}

			var templates []*entities.MsvcTemplate

			if err := yaml.Unmarshal(b, &templates); err != nil {
				return nil, errors.NewE(err)
			}

			msvcTemplatesMap := map[string]map[string]*entities.MsvcTemplateEntry{}

			for _, t := range templates {
				if _, ok := msvcTemplatesMap[t.Category]; !ok {
					msvcTemplatesMap[t.Category] = make(map[string]*entities.MsvcTemplateEntry, len(t.Items))
				}
				for i := range t.Items {
					msvcTemplatesMap[t.Category][t.Items[i].Name] = &t.Items[i]
				}
			}

			cd := &commonDomain.Domain{
				Logger:                 logger,
				Env:                    ev,
				K8sClient:              k8sClient,
				IAMSvc:                 iamSvc,
				AccountsSvc:            accountsSvc,
				ResDispatcher:          resourceDispatcherV2,
				ResourceEventPublisher: resourceEventPublisherV2,
			}

			gvpnDomain := &global_vpn.Domain{
				Domain:        cd,
				GlobalVPNRepo: gvpnRepo,
				GlobalVPNDevice: global_vpn.GlobalVPNDevice{
					CreateGlobalVPNDevice: func(ctx types.InfraContext, gvpn entities.GlobalVPNDevice) (*entities.GlobalVPNDevice, error) {
						panic("not implemented")
					},
				},
				ClusterDomain: global_vpn.ClusterDomain{
					CountClustersInGlobalVPN: func(ctx types.InfraContext, gvpnName string) (int64, error) {
						return clusterRepo.Count(ctx, repos.Filter{
							fc.AccountName:      ctx.AccountName,
							fc.ClusterGlobalVPN: gvpnName,
						})
					},
				},
			}

			gvpnDeviceDomain := &global_vpn_device.Domain{
				Domain:              cd,
				GlobalVPNDeviceRepo: gvpnDevicesRepo,
				FreeDeviceIPRepo:    freeDeviceIpRepo,
				ClaimDeviceIPRepo:   claimDeviceIPRepo,
				GlobalVPNDomain: global_vpn_device.GlobalVPNDomain{
					FindGlobalVPN: func(ctx types.InfraContext, gvpnName string) (*entities.GlobalVPN, error) {
						return gvpnDomain.FindGlobalVPN(ctx, gvpnName)
					},
					IncrementGlobalVPNAllocatedDevices: func(ctx types.InfraContext, value int) error {
						panic("not implmented")
					},
				},
				GlobalVPNConnectionDomain: global_vpn_device.GlobalVPNConnectionDomain{
					SyncGlobalVPNConnections: func(ctx types.InfraContext, gvpnName string) error {
						panic("not implemented")
					},
					ListGlobalVPNConnections: func(ctx types.InfraContext, gvpnName string) ([]*entities.GlobalVPNConnection, error) {
						panic("not implemented")
					},
					BuildGlobalVPNConnectionPeers: func(vpns []*entities.GlobalVPNConnection) ([]wgv1.Peer, error) {
						panic("not implemented")
					},
				},
			}

			gvpnConnDomain := &global_vpn_connection.Domain{
				Domain:                         cd,
				GlobalVPNClusterConnectionRepo: gvpnConnRepo,
				FreeClusterSvcCIDRRepo:         freeClusterSvcCIDRRepo,
				ClaimClusterSvcCIDRRepo:        claimClusterSvcCIDRRepo,
				GlobalVPNDomain: global_vpn_connection.GlobalVPNDomain{
					FindGlobalVPN: func(ctx types.InfraContext, gvpnName string) (*entities.GlobalVPN, error) {
						return gvpnDomain.FindGlobalVPN(ctx, gvpnName)
					},
				},
				GlobalVPNDeviceDomain: global_vpn_connection.GlobalVPNDeviceDomain{
					CreateGlobalVPNDevice: func(ctx types.InfraContext, device entities.GlobalVPNDevice) (*entities.GlobalVPNDevice, error) {
						return gvpnDeviceDomain.CreateGlobalVPNDevice(ctx, device)
					},
					DeleteGlobalVPNDevice: func(ctx types.InfraContext, gvpnName string, deviceName string) error {
						return gvpnDeviceDomain.DeleteGlobalVPNDevice(ctx, gvpnName, deviceName)
					},
					BuildPeersFromGlobalVPNDevices: func(ctx types.InfraContext, gvpnName string) ([]wgv1.Peer, []wgv1.Peer, error) {
						panic("not implemented")
					},
					SyncKloudliteDeviceOnPlatform: func(ctx types.InfraContext, gvpnName string) error {
						panic("not implemented")
					},
				},
				BYOKDomain: global_vpn_connection.BYOKDomain{
					IsBYOKCluster: func(ctx types.InfraContext, name string) bool {
						panic("not implemented")
					},
					MarkBYOKClusterReady: func(ctx types.InfraContext, name string, timestamp time.Time) error {
						panic("not implemented")
					},
				},
			}

			clusterDomain := &clusters.Domain{
				Domain:       cd,
				ClusterRepo:  clusterRepo,
				MsgOfficeSvc: msgOfficeSvc,
				GlobalVPNDomain: clusters.GlobalVPNDomain{
					EnsureGlobalVPN: func(ctx types.InfraContext, gvpnName string) (*entities.GlobalVPN, error) {
						return gvpnDomain.EnsureGlobalVPN(ctx, gvpnName)
					},
					FindGlobalVPN: func(ctx types.InfraContext, gvpnName string) (*entities.GlobalVPN, error) {
						return gvpnDomain.FindGlobalVPN(ctx, gvpnName)
					},
				},
				GlobalVPNConnectionDomain: clusters.GlobalVPNConnectionDomain{},
				GlobalVPNDeviceDomain:     clusters.GlobalVPNDeviceDomain{},
				BYOKClusterDomain:         clusters.BYOKClusterDomain{},
				ProviderSecretDomain:      clusters.ProviderSecretDomain{},
				NodepoolDomain:            clusters.NodepoolDomain{},
				PVDomain:                  clusters.PVDomain{},
				HelmReleaseDomain:         clusters.HelmReleaseDomain{},
			}

			helmReleaseDomain := &helm_releases.Domain{
				Domain:          cd,
				HelmReleaseRepo: helmReleaseRepo,
			}

			byokClusterDomain := &byok_clusters.Domain{
				Domain:          cd,
				BYOKClusterRepo: byokClusterRepo,
				MsgOfficeSvc:    msgOfficeSvc,
				GlobalVPNDomain: byok_clusters.GlobalVPNDomain{
					EnsureGlobalVPN: func(ctx types.InfraContext, gvpnName string) (*entities.GlobalVPN, error) {
						return gvpnDomain.EnsureGlobalVPN(ctx, gvpnName)
					},
					FindGlobalVPN: func(ctx types.InfraContext, gvpnName string) (*entities.GlobalVPN, error) {
						return gvpnDomain.FindGlobalVPN(ctx, gvpnName)
					},
				},
				GlobalVPNConnectionDomain: byok_clusters.GlobalVPNConnectionDomain{
					EnsureGlobalVPNConnection: func(ctx types.InfraContext, clusterName string, groupName string, clusterPublicEndpoint string) (*entities.GlobalVPNConnection, error) {
						return gvpnConnDomain.EnsureGlobalVPNConnection(ctx, clusterName, groupName, clusterPublicEndpoint)
					},
					FindGlobalVPNConnection: func(ctx types.InfraContext, gvpnName string, clusterName string) (*entities.GlobalVPNConnection, error) {
						return gvpnConnDomain.FindGlobalVPNConnection(ctx, gvpnName, clusterName)
					},
					DeleteGlobalVPNConnection: func(ctx types.InfraContext, gvpnName string, clusterName string) error {
						return gvpnConnDomain.DeleteGlobalVPNDevice(ctx, gvpnName, clusterName)
					},
				},
				ClusterDomain: byok_clusters.ClusterDomain{},
			}

			providerSecretDomain := &provider_secret.Domain{
				Domain:             cd,
				ProviderSecretRepo: secretRepo,
				ClusterDomain: provider_secret.ClusterDomain{
					CountClustersWithProviderSecret: func(ctx types.InfraContext, providerSecretName string) (int64, error) {
						return clusterRepo.Count(ctx, repos.Filter{
							fc.AccountName: ctx.AccountName,
							fc.ClusterSpecAwsCredentialsSecretRefName: providerSecretName,
						})
					},
				},
			}

			nodepoolDomain := &nodepool.Domain{
				Domain:       cd,
				NodepoolRepo: nodePoolRepo,
				ClusterDomain: nodepool.ClusterDomain{
					FindCluster: func(ctx types.InfraContext, name string) (*entities.Cluster, error) {
						return clusterDomain.FindCluster(ctx, name)
					},
				},
				ProviderSecretDomain: nodepool.ProviderSecretDomain{
					FindProviderSecret: func(ctx types.InfraContext, name string) (*entities.CloudProviderSecret, error) {
						return providerSecretDomain.FindProviderSecret(ctx, name)
					},
					GetProviderSecretAsK8sSecret: func(ctx types.InfraContext, name string) (*v1.Secret, error) {
						cps, err := providerSecretDomain.FindProviderSecret(ctx, name)
						if err != nil {
							return nil, err
						}
						return providerSecretDomain.CoreV1SecretFromProviderSecret(cps)
					},
				},
			}

			pvDomain := &persistent_volume.Domain{
				Domain:               cd,
				PersistentVolumeRepo: pvRepo,
			}

			pvcDomain := &persistent_volume_claim.Domain{
				Domain:                    cd,
				PersistentVolumeClaimRepo: pvcRepo,
			}

			return &domain{
				CommonDomainImpl:                CommonDomainImpl{Domain: cd},
				ClusterDomainImpl:               ClusterDomainImpl{Domain: clusterDomain},
				BYOKClusterDomainImpl:           BYOKClusterDomainImpl{Domain: byokClusterDomain},
				HelmReleaseDomainImpl:           HelmReleaseDomainImpl{Domain: helmReleaseDomain},
				NodepoolDomainImpl:              NodepoolDomainImpl{Domain: nodepoolDomain},
				ProviderSecretDomainImpl:        ProviderSecretDomainImpl{Domain: providerSecretDomain},
				PersistentVolumeDomainImpl:      PersistentVolumeDomainImpl{Domain: pvDomain},
				PersistentVolumeClaimDomainImpl: PersistentVolumeClaimDomainImpl{Domain: pvcDomain},
				GlobalVPNDomainImpl:             GlobalVPNDomainImpl{Domain: gvpnDomain},
				GlobalVPNDeviceDomainImpl:       GlobalVPNDeviceDomainImpl{Domain: gvpnDeviceDomain},

				msvcTemplatesMap: msvcTemplatesMap,
				msvcTemplates:    templates,
				clusterRepo:      clusterRepo,
				gvpnConnRepo:     gvpnConnRepo,
				// deviceAddressPoolRepo:   deviceAddressPoolRepo,

				claimDeviceIPRepo:       claimDeviceIPRepo,
				freeDeviceIpRepo:        freeDeviceIpRepo,
				freeClusterSvcCIDRRepo:  freeClusterSvcCIDRRepo,
				claimClusterSvcCIDRRepo: claimClusterSvcCIDRRepo,

				gvpnRepo:        gvpnRepo,
				gvpnDevicesRepo: gvpnDevicesRepo,

				byokClusterRepo:             byokClusterRepo,
				clusterManagedServiceRepo:   clustermanagedserviceRepo,
				nodeRepo:                    nodeRepo,
				nodePoolRepo:                nodePoolRepo,
				secretRepo:                  secretRepo,
				domainEntryRepo:             domainNameRepo,
				resDispatcher:               resourceDispatcher,
				k8sClient:                   k8sClient,
				iamClient:                   iamClient,
				accountsSvc:                 accountsSvc,
				messageOfficeInternalClient: msgOfficeInternalClient,
				resourceEventPublisher:      resourceEventPublisher,
				helmReleaseRepo:             helmReleaseRepo,

				pvcRepo:              pvcRepo,
				volumeAttachmentRepo: volumeAttachmentRepo,
				pvRepo:               pvRepo,
				namespaceRepo:        namespaceRepo,
			}, nil
		}),
)
