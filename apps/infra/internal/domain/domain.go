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
	cluster_managed_service "github.com/kloudlite/api/apps/infra/internal/domain/cluster-managed-service"
	"github.com/kloudlite/api/apps/infra/internal/domain/clusters"
	commonDomain "github.com/kloudlite/api/apps/infra/internal/domain/common"
	global_vpn "github.com/kloudlite/api/apps/infra/internal/domain/global-vpn"
	global_vpn_connection "github.com/kloudlite/api/apps/infra/internal/domain/global-vpn-connection"
	global_vpn_device "github.com/kloudlite/api/apps/infra/internal/domain/global-vpn-device"
	helm_releases "github.com/kloudlite/api/apps/infra/internal/domain/helm-releases"
	ingress_hosts "github.com/kloudlite/api/apps/infra/internal/domain/ingress-hosts"
	managed_svc_templates "github.com/kloudlite/api/apps/infra/internal/domain/managed-svc-templates"
	name_check "github.com/kloudlite/api/apps/infra/internal/domain/name-check"
	namespace "github.com/kloudlite/api/apps/infra/internal/domain/namespaces"
	"github.com/kloudlite/api/apps/infra/internal/domain/nodepool"
	persistent_volume "github.com/kloudlite/api/apps/infra/internal/domain/persistent-volume"
	persistent_volume_claim "github.com/kloudlite/api/apps/infra/internal/domain/persistent-volume-claim"
	"github.com/kloudlite/api/apps/infra/internal/domain/ports"
	provider_secret "github.com/kloudlite/api/apps/infra/internal/domain/provider-secret"
	"github.com/kloudlite/api/apps/infra/internal/domain/types"
	volume_attachment "github.com/kloudlite/api/apps/infra/internal/domain/volume-attachment"
	"github.com/kloudlite/api/apps/infra/internal/entities"

	fc "github.com/kloudlite/api/apps/infra/internal/entities/field-constants"
	"github.com/kloudlite/api/apps/infra/internal/env"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/iam"
	message_office_internal "github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/message-office-internal"
	fn "github.com/kloudlite/api/pkg/functions"
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
	GlobalVPNDomainImpl             struct{ *global_vpn.Domain }
	GlobalVPNConnectionDomainImpl   struct{ *global_vpn_connection.Domain }
	GlobalVPNDeviceDomainImpl       struct{ *global_vpn_device.Domain }
	IngressHostsDomainImpl          struct{ *ingress_hosts.Domain }
	ClusterManagedServiceDomainImpl struct {
		*cluster_managed_service.Domain
	}
	ManagedSvcTemplatesDomainImpl struct{ *managed_svc_templates.Domain }
	NamespaceDomainImpl           struct{ *namespace.Domain }
	VolumeAttachmentDomainImpl    struct{ *volume_attachment.Domain }
	NameCheckDomainImpl           struct{ *name_check.Domain }
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

	IngressHostsDomainImpl
	ClusterManagedServiceDomainImpl

	ManagedSvcTemplatesDomainImpl
	NamespaceDomainImpl

	VolumeAttachmentDomainImpl
	NameCheckDomainImpl
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

			resourceDispatcher ports.ResourceDispatcher,

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

			accountsSvc ports.AccountsSvc,
			msgOfficeInternalClient message_office_internal.MessageOfficeInternalClient,
			logger logging.Logger,

			resourceEventPublisher ports.ResourceEventPublisher,
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
				ResDispatcher:          resourceDispatcher,
				ResourceEventPublisher: resourceEventPublisher,
			}

			gvpnDomain := &global_vpn.Domain{
				Domain:        cd,
				GlobalVPNRepo: gvpnRepo,
				GlobalVPNDevice: global_vpn.GlobalVPNDevice{
					CreateGlobalVPNDevice: func(ctx types.InfraContext, gvpn entities.GlobalVPNDevice) (*entities.GlobalVPNDevice, error) {
						return global_vpn_device.CreateGlobalVPNDevice(ctx, global_vpn_device.CreateGlobalVPNDeviceArgs{
							Logger:              logger,
							Device:              gvpn,
							FreeDeviceIPRepo:    freeDeviceIpRepo,
							ClaimDeviceIPRepo:   claimDeviceIPRepo,
							GlobalVPNDeviceRepo: gvpnDevicesRepo,

							GetGlobalVPN: func(ctx types.InfraContext, gvpnName string) (*entities.GlobalVPN, error) {
								return global_vpn.GetGlobalVPN(ctx, global_vpn.GetGlobalVPNArgs{
									GlobalVPNName: gvpnName,
									GlobalVPNRepo: gvpnRepo,
								})
							},
							IncrementGlobalVPNAllocatedDevices: func(ctx types.InfraContext, gvpnId repos.ID, incrementBy int) error {
								return global_vpn.IncrementAllocatedGlobalVPNDevices(ctx, global_vpn.IncrementAllocatedGlobalVPNDevicesArgs{
									IncrementBy:   incrementBy,
									GlobalVPNId:   gvpnId,
									GlobalVPNRepo: gvpnRepo,
								})
							},
							SyncGlobalVPNConnections: func(ctx types.InfraContext, gvpnName string) error {
								return global_vpn_connection.SyncGlobalVPNConnections(ctx, global_vpn_connection.SyncGlobalVPNConnectionsArgs{
									GlobalVPNName:           gvpnName,
									GlobalVPNConnectionRepo: gvpnConnRepo,
									GeneratePeersFromGlobalVPNDevices: func(ctx types.InfraContext, vpnName string) ([]wgv1.Peer, []wgv1.Peer, error) {
										return global_vpn_device.GeneratePeersFromGlobalVPNDevices(ctx, global_vpn_device.GeneratePeersFromGlobalVPNDevicesArgs{
											GlobalVPNName:       vpnName,
											GlobalVPNDeviceRepo: gvpnDevicesRepo,
										})
									},
									ResDispatcher: resourceDispatcher,
								})
							},
						})
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
					IncrementGlobalVPNAllocatedDevices: func(ctx types.InfraContext, gvpnId repos.ID, value int) error {
						return global_vpn.IncrementAllocatedGlobalVPNDevices(ctx, global_vpn.IncrementAllocatedGlobalVPNDevicesArgs{
							IncrementBy:   value,
							GlobalVPNId:   gvpnId,
							GlobalVPNRepo: gvpnRepo,
						})
					},
				},
				GlobalVPNConnectionDomain: global_vpn_device.GlobalVPNConnectionDomain{
					SyncGlobalVPNConnections: func(ctx types.InfraContext, gvpnName string) error {
						return global_vpn_connection.SyncGlobalVPNConnections(ctx, global_vpn_connection.SyncGlobalVPNConnectionsArgs{
							GlobalVPNName:           gvpnName,
							GlobalVPNConnectionRepo: gvpnConnRepo,
							GeneratePeersFromGlobalVPNDevices: func(ctx types.InfraContext, vpnName string) ([]wgv1.Peer, []wgv1.Peer, error) {
								return global_vpn_device.GeneratePeersFromGlobalVPNDevices(ctx, global_vpn_device.GeneratePeersFromGlobalVPNDevicesArgs{
									GlobalVPNName:       vpnName,
									GlobalVPNDeviceRepo: gvpnDevicesRepo,
								})
							},
							ResDispatcher: resourceDispatcher,
						})
					},
					ListGlobalVPNConnections: func(ctx types.InfraContext, gvpnName string) ([]*entities.GlobalVPNConnection, error) {
						return global_vpn_connection.ListGlobalVPNConnections(ctx, global_vpn_connection.ListGlobalVPNConnectionsArgs{
							GlobalVPNName:           gvpnName,
							GlobalVPNConnectionRepo: gvpnConnRepo,
						})
					},
					BuildGlobalVPNConnectionPeers: func(ctx types.InfraContext, vpns []*entities.GlobalVPNConnection) ([]wgv1.Peer, error) {
						return global_vpn_connection.GenerateGlobalVPNConnectionPeers(ctx, vpns)
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
						return gvpnDeviceDomain.BuildPeersFromGlobalVPNDevices(ctx, gvpnName)
					},
					SyncKloudliteDeviceOnPlatform: func(ctx types.InfraContext, gvpnName string) error {
						panic("not implemented")
					},
				},
				BYOKDomain: global_vpn_connection.BYOKDomain{
					IsBYOKCluster: func(ctx types.InfraContext, name string) bool {
						ok, err := byok_clusters.IsBYOKCluster(ctx, byok_clusters.IsBYOKClusterArgs{
							ClusterName:     name,
							BYOKClusterRepo: byokClusterRepo,
						})
						if err != nil {
							return false
						}
						return ok
					},
					MarkBYOKClusterReady: func(ctx types.InfraContext, name string, timestamp time.Time) error {
						_, err := byok_clusters.MarkBYOKClusterReady(ctx, byok_clusters.MarkBYOKClusterReadyArgs{
							ClusterName:     name,
							BYOKClusterRepo: byokClusterRepo,
							Time:            timestamp,
						})
						return err
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
				GlobalVPNConnectionDomain: clusters.GlobalVPNConnectionDomain{
					EnsureGlobalVPNConnection: func(ctx types.InfraContext, clusterName string, groupName string, clusterPublicEndpoint string) (*entities.GlobalVPNConnection, error) {
					},
					FindGlobalVPNConnection: func(ctx types.InfraContext, gvpnName string, clusterName string) (*entities.GlobalVPNConnection, error) {
					},
					DeleteGlobalVPNConnection: func(ctx types.InfraContext, gvpnName string, clusterName string) error {
					},
				},
				GlobalVPNDeviceDomain: clusters.GlobalVPNDeviceDomain{},
				BYOKClusterDomain:     clusters.BYOKClusterDomain{},
				ProviderSecretDomain:  clusters.ProviderSecretDomain{},
				NodepoolDomain:        clusters.NodepoolDomain{},
				PVDomain:              clusters.PVDomain{},
				HelmReleaseDomain:     clusters.HelmReleaseDomain{},
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
						return clusters.GetCluster(ctx, clusters.GetClusterArgs{
							ClusterRepo: clusterRepo,
							ClusterName: name,
						})
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

			clusterManagedServiceDomain := &cluster_managed_service.Domain{
				Domain:                    cd,
				ClusterManagedServiceRepo: clustermanagedserviceRepo,
			}

			ingressHostsDomain := &ingress_hosts.Domain{
				Domain:          cd,
				DomainEntryRepo: domainNameRepo,
			}

			msvcTemplatesDomain := &managed_svc_templates.Domain{
				ManagedSvcTemplates:    templates,
				ManagedSvcTemplatesMap: msvcTemplatesMap,
			}

			namespaceDomain := &namespace.Domain{
				NamespaceRepo: namespaceRepo,
			}

			volumeAttachDomain := &volume_attachment.Domain{
				Domain:               cd,
				VolumeAttachmentRepo: volumeAttachmentRepo,
			}

			nameCheckDomain := &name_check.Domain{
				ClusterDomain: name_check.ClusterDomain{
					IsClusterNameAvailable: func(ctx types.InfraContext, name string) (bool, error) {
						c, err := clusters.GetCluster(ctx, clusters.GetClusterArgs{
							ClusterRepo: clusterRepo,
							ClusterName: name,
						})
						if err != nil {
							return false, err
						}
						return fn.IsNil(c), nil
					},
				},
				GlobalVPNDeviceDomain: name_check.GlobalVPNDeviceDomain{
					IsGlobalVPNDeviceNameAvailable: func(ctx types.InfraContext, name string) (bool, error) {
						c, err := gvpnDomain.FindGlobalVPN(ctx, name)
						if err != nil {
							return false, err
						}
						return fn.IsNil(c), nil
					},
				},
				NodepoolDomain: name_check.NodepoolDomain{
					IsNodepoolNameAvailable: func(ctx types.InfraContext, clusterName string, name string) (bool, error) {
						c, err := nodepoolDomain.FindNodePool(ctx, clusterName, name)
						if err != nil {
							return false, err
						}
						return fn.IsNil(c), nil
					},
				},
				HelmReleaseDomain: name_check.HelmReleaseDomain{
					IsHelmReleaseNameAvailable: func(ctx types.InfraContext, clusterName string, name string) (bool, error) {
						c, err := helmReleaseDomain.FindHelmRelease(ctx, clusterName, name)
						if err != nil {
							return false, err
						}
						return fn.IsNil(c), nil
					},
				},
				ClusterManagedServiceDomain: name_check.ClusterManagedServiceDomain{
					IsClusterManagedSvcNameAvailable: func(ctx types.InfraContext, clusterName string, name string) (bool, error) {
						c, err := clusterManagedServiceDomain.FindClusterManagedService(ctx, name)
						if err != nil {
							return false, err
						}
						return fn.IsNil(c), nil
					},
				},
				ProviderSecretDomain: name_check.ProviderSecretDomain{
					IsProviderSecretNameAvailable: func(ctx types.InfraContext, name string) (bool, error) {
						c, err := providerSecretDomain.FindProviderSecret(ctx, name)
						if err != nil {
							return false, err
						}
						return fn.IsNil(c), nil
					},
				},
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

				ClusterManagedServiceDomainImpl: ClusterManagedServiceDomainImpl{Domain: clusterManagedServiceDomain},
				IngressHostsDomainImpl:          IngressHostsDomainImpl{Domain: ingressHostsDomain},
				ManagedSvcTemplatesDomainImpl:   ManagedSvcTemplatesDomainImpl{Domain: msvcTemplatesDomain},
				NamespaceDomainImpl:             NamespaceDomainImpl{Domain: namespaceDomain},
				VolumeAttachmentDomainImpl:      VolumeAttachmentDomainImpl{Domain: volumeAttachDomain},
				NameCheckDomainImpl:             NameCheckDomainImpl{Domain: nameCheckDomain},
			}, nil
		}),
)
