package clusters

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/apps/infra/internal/domain/ports"
	"github.com/kloudlite/api/apps/infra/internal/domain/templates"
	domainT "github.com/kloudlite/api/apps/infra/internal/domain/types"
	fc "github.com/kloudlite/api/apps/infra/internal/entities/field-constants"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	ct "github.com/kloudlite/operator/apis/common-types"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
	"sigs.k8s.io/yaml"

	"github.com/kloudlite/api/apps/infra/internal/entities"
	clustersv1 "github.com/kloudlite/operator/apis/clusters/v1"

	"github.com/kloudlite/api/pkg/errors"
	fn "github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/repos"
	t "github.com/kloudlite/api/pkg/types"
	crdsv1 "github.com/kloudlite/operator/apis/crds/v1"
	corev1 "k8s.io/api/core/v1"
	apiErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	keyClusterToken = "cluster-token"
)

type ErrClusterAlreadyExists struct {
	ClusterName string
	AccountName string
}

var ErrClusterNotFound error = fmt.Errorf("cluster not found")

func (e ErrClusterAlreadyExists) Error() string {
	return fmt.Sprintf("cluster with name %q already exists for account: %s", e.ClusterName, e.AccountName)
}

const (
	DefaultGlobalVPNName = "default"
)

func (d *Domain) createTokenSecret(ctx domainT.InfraContext, clusterName string, clusterNamespace string) (*corev1.Secret, error) {
	secret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterName,
			Namespace: clusterNamespace,
		},
	}

	clusterToken, err := d.MsgOfficeSvc.GenerateClusterToken(ctx, ctx.AccountName, clusterName)
	if err != nil {
		return nil, errors.NewE(err)
	}

	secret.Data = map[string][]byte{
		keyClusterToken: []byte(clusterToken),
	}

	return secret, nil
}

func (d *Domain) GetClusterAdminKubeconfig(ctx domainT.InfraContext, clusterName string) (*string, error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.UpdateCluster); err != nil {
		return nil, errors.NewE(err)
	}

	cluster, err := d.FindCluster(ctx, clusterName)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if cluster.Spec.Output == nil {
		return fn.New(""), nil
	}

	kscrt := corev1.Secret{}
	if err := d.K8sClient.Get(ctx.Context, fn.NN(cluster.Namespace, cluster.Spec.Output.SecretName), &kscrt); err != nil {
		return nil, errors.NewE(err)
	}

	kubeconfig, ok := kscrt.Data[cluster.Spec.Output.KeyKubeconfig]
	if !ok {
		return nil, errors.Newf("kubeconfig key %q not found in secret %q", cluster.Spec.Output.KeyKubeconfig, cluster.Spec.Output.SecretName)
	}

	return fn.New(string(kubeconfig)), nil
}

func (d *Domain) applyCluster(ctx domainT.InfraContext, cluster *entities.Cluster) error {
	d.AddTrackingId(&cluster.Cluster, cluster.Id)
	return d.ApplyK8sResource(ctx, &cluster.Cluster, cluster.RecordVersion)
}

func (d *Domain) CreateCluster(ctx domainT.InfraContext, cluster entities.Cluster) (*entities.Cluster, error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.CreateCluster); err != nil {
		return nil, errors.NewE(err)
	}

	exists, err := d.clusterAlreadyExists(ctx, cluster.Name)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if exists != nil && *exists {
		return nil, errors.Newf("cluster/byok cluster with name (%s) already exists", cluster.Name)
	}

	accNs, err := d.GetAccNamespace(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if cluster.GlobalVPN == nil {
		cluster.GlobalVPN = fn.New(DefaultGlobalVPNName)
	}

	if _, err := d.EnsureGlobalVPN(ctx, *cluster.GlobalVPN); err != nil {
		return nil, errors.NewE(err)
	}

	cluster.EnsureGVK()
	cluster.Namespace = accNs

	existing, err := d.ClusterRepo.FindOne(ctx, repos.Filter{
		fields.MetadataName:      cluster.Name,
		fields.MetadataNamespace: cluster.Namespace,
		fields.AccountName:       ctx.AccountName,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if existing != nil {
		return nil, ErrClusterAlreadyExists{ClusterName: cluster.Name, AccountName: ctx.AccountName}
	}

	cluster.AccountName = ctx.AccountName
	out, err := d.AccountsSvc.GetAccount(ctx, string(ctx.UserId), ctx.AccountName)
	if err != nil {
		return nil, errors.NewEf(err, "failed to get account %q", ctx.AccountName)
	}

	cluster.Spec.AccountId = out.AccountId

	tokenScrt, err := d.createTokenSecret(ctx, cluster.Name, cluster.Namespace)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if err := d.ApplyK8sResource(ctx, tokenScrt, 1); err != nil {
		return nil, errors.NewE(err)
	}

	cluster.Spec = clustersv1.ClusterSpec{
		AccountName: ctx.AccountName,
		AccountId:   out.AccountId,
		ClusterTokenRef: ct.SecretKeyRef{
			Name:      tokenScrt.Name,
			Namespace: tokenScrt.Namespace,
			Key:       keyClusterToken,
		},
		AvailabilityMode: cluster.Spec.AvailabilityMode,

		// PublicDNSHost is <cluster-name>.<account-name>.tenants.<public-dns-host-suffix>
		PublicDNSHost:          fmt.Sprintf("%s.%s.tenants.%s", cluster.Name, ctx.AccountName, d.Env.PublicDNSHostSuffix),
		ClusterInternalDnsHost: fn.New("cluster.local"),
		CloudflareEnabled:      fn.New(true),
		TaintMasterNodes:       true,
		BackupToS3Enabled:      false,

		CloudProvider: cluster.Spec.CloudProvider,
		AWS: func() *clustersv1.AWSClusterConfig {
			if cluster.Spec.CloudProvider != ct.CloudProviderAWS {
				return nil
			}

			cps, err := d.FindProviderSecret(ctx, cluster.Spec.AWS.Credentials.SecretRef.Name)
			if err != nil {
				return nil
			}

			return &clustersv1.AWSClusterConfig{
				Credentials: clustersv1.AwsCredentials{
					AuthMechanism: cps.AWS.AuthMechanism,
					SecretRef: ct.SecretRef{
						Name:      cps.Name,
						Namespace: cps.Namespace,
					},
				},

				Region: cluster.Spec.AWS.Region,
				K3sMasters: clustersv1.AWSK3sMastersConfig{
					InstanceType:     cluster.Spec.AWS.K3sMasters.InstanceType,
					NvidiaGpuEnabled: cluster.Spec.AWS.K3sMasters.NvidiaGpuEnabled,
					RootVolumeType:   "gp3",
					RootVolumeSize: func() int {
						if cluster.Spec.AWS.K3sMasters.NvidiaGpuEnabled {
							return 80
						}
						return 50
					}(),
					IAMInstanceProfileRole: &cps.AWS.CfParamInstanceProfileName,
					Nodes: func() map[string]clustersv1.MasterNodeProps {
						if cluster.Spec.AvailabilityMode == "dev" {
							return map[string]clustersv1.MasterNodeProps{
								"master-1": {
									Role:             "primary-master",
									KloudliteRelease: d.Env.KloudliteRelease,
								},
							}
						}
						return map[string]clustersv1.MasterNodeProps{
							"master-1": {
								Role:             "primary-master",
								KloudliteRelease: d.Env.KloudliteRelease,
							},
							"master-2": {
								Role:             "secondary-master",
								KloudliteRelease: d.Env.KloudliteRelease,
							},
							"master-3": {
								Role:             "secondary-master",
								KloudliteRelease: d.Env.KloudliteRelease,
							},
						}
					}(),
				},
			}
		}(),
		GCP: func() *clustersv1.GCPClusterConfig {
			if cluster.Spec.CloudProvider != ct.CloudProviderGCP {
				return nil
			}

			cps, err := d.FindProviderSecret(ctx, cluster.Spec.GCP.CredentialsRef.Name)
			if err != nil {
				return nil
			}

			var gcpServiceAccountJSON struct {
				ProjectID string `json:"project_id"`
			}

			if cps.GCP != nil {
				if err := json.Unmarshal([]byte(cps.GCP.ServiceAccountJSON), &gcpServiceAccountJSON); err != nil {
					return nil
				}
			}

			return &clustersv1.GCPClusterConfig{
				Region:       cluster.Spec.GCP.Region,
				GCPProjectID: gcpServiceAccountJSON.ProjectID,
				CredentialsRef: ct.SecretRef{
					Name:      cps.Name,
					Namespace: cps.Namespace,
				},
				// FIXME: once, we allow gcp service account for clusters via UI
				ServiceAccount: clustersv1.GCPServiceAccount{
					Enabled: false,
					Email:   nil,
					Scopes:  nil,
				},
				MasterNodes: clustersv1.GCPMasterNodesConfig{
					RootVolumeType: "pd-ssd",
					RootVolumeSize: 50,
					Nodes: func() map[string]clustersv1.MasterNodeProps {
						if cluster.Spec.AvailabilityMode == "dev" {
							return map[string]clustersv1.MasterNodeProps{
								"master-1": {
									Role:             "primary-master",
									AvailabilityZone: fmt.Sprintf("%s-a", cluster.Spec.GCP.Region), // defaults to {{.region}}-a zone
									KloudliteRelease: d.Env.KloudliteRelease,
								},
							}
						}
						return map[string]clustersv1.MasterNodeProps{
							"master-1": {
								Role:             "primary-master",
								AvailabilityZone: fmt.Sprintf("%s-a", cluster.Spec.GCP.Region),
								KloudliteRelease: d.Env.KloudliteRelease,
							},
							"master-2": {
								Role:             "secondary-master",
								AvailabilityZone: fmt.Sprintf("%s-a", cluster.Spec.GCP.Region),
								KloudliteRelease: d.Env.KloudliteRelease,
							},
							"master-3": {
								Role:             "secondary-master",
								AvailabilityZone: fmt.Sprintf("%s-a", cluster.Spec.GCP.Region),
								KloudliteRelease: d.Env.KloudliteRelease,
							},
						}
					}(),
				},
			}
		}(),
		MessageQueueTopicName: common.GetTenantClusterMessagingTopic(ctx.AccountName, cluster.Name),
		KloudliteRelease:      d.Env.KloudliteRelease,
		Output:                nil,
	}

	cluster.IncrementRecordVersion()
	cluster.CreatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}
	cluster.LastUpdatedBy = cluster.CreatedBy

	cluster.Spec.AccountId = out.AccountId
	cluster.Spec.AccountName = ctx.AccountName
	cluster.SyncStatus = t.GenSyncStatus(t.SyncActionApply, 0)

	// clusterSvcCIDR, err := d.claimNextClusterSvcCIDR(ctx, cluster.Name, gvpn.Name)
	// if err != nil {
	// 	return nil, err
	// }

	gvpnConn, err := d.EnsureGlobalVPNConnection(ctx, cluster.Name, *cluster.GlobalVPN, cluster.Spec.PublicDNSHost)
	if err != nil {
		return nil, errors.NewE(err)
	}

	cluster.Spec.ClusterServiceCIDR = gvpnConn.ClusterSvcCIDR

	if err := d.K8sClient.ValidateObject(ctx, &cluster.Cluster); err != nil {
		return nil, errors.NewE(err)
	}

	nCluster, err := d.ClusterRepo.Create(ctx, &cluster)
	if err != nil {
		if d.ClusterRepo.ErrAlreadyExists(err) {
			return nil, errors.Newf("cluster with name %q already exists in namespace %q", cluster.Name, cluster.Namespace)
		}
		return nil, errors.NewE(err)
	}

	if err := d.applyCluster(ctx, nCluster); err != nil {
		return nil, errors.NewE(err)
	}

	d.ResourceEventPublisher.PublishInfraEvent(ctx, ports.ResourceTypeCluster, nCluster.Name, ports.PublishAdd)

	if err := d.applyHelmKloudliteAgent(ctx, string(tokenScrt.Data[keyClusterToken]), nCluster); err != nil {
		return nil, errors.NewE(err)
	}

	return nCluster, nil
}

func (d *Domain) syncKloudliteDeviceOnCluster(ctx domainT.InfraContext, gvpnName string) error {
	// 1. parse deployment template
	b, err := templates.Read(templates.GlobalVPNKloudliteDeviceTemplate)
	if err != nil {
		return errors.NewE(err)
	}
	accNs, err := d.GetAccNamespace(ctx)
	if err != nil {
		return errors.NewE(err)
	}

	gv, err := d.FindGlobalVPN(ctx, gvpnName)
	if err != nil {
		return err
	}

	if gv.KloudliteDevice.Name == "" {
		return nil
	}

	// 2. Grab wireguard config from that device
	// wgConfig, err := d.getGlobalVPNDeviceWgConfig(ctx, gv.Name, gv.KloudliteDevice.Name, nil)
	wgConfig, err := d.GetGlobalVPNDeviceWgConfig(ctx, gv.Name, gv.KloudliteDevice.Name)
	if err != nil {
		return err
	}

	deploymentBytes, err := templates.ParseBytes(b, templates.GVPNKloudliteDeviceTemplateVars{
		Name:                  fmt.Sprintf("kloudlite-device-proxy-%s", gv.Name),
		Namespace:             accNs,
		WgConfig:              wgConfig,
		KubeReverseProxyImage: d.Env.GlobalVPNKubeReverseProxyImage,
	})
	if err != nil {
		return err
	}

	if err := d.K8sClient.ApplyYAML(ctx, deploymentBytes); err != nil {
		return errors.NewE(err)
	}

	return nil
}

func (d *Domain) applyHelmKloudliteAgent(ctx domainT.InfraContext, clusterToken string, cluster *entities.Cluster) error {
	b, err := templates.Read(templates.HelmKloudliteAgent)
	if err != nil {
		return errors.NewE(err)
	}

	values := map[string]any{
		"account-name": ctx.AccountName,

		"cluster-name":  cluster.Name,
		"cluster-token": clusterToken,

		"kloudlite-release":        d.Env.KloudliteRelease,
		"message-office-grpc-addr": d.Env.MessageOfficeExternalGrpcAddr,

		"public-dns-host": cluster.Spec.PublicDNSHost,
		"cloudprovider":   cluster.Spec.CloudProvider,
	}

	if cluster.Spec.CloudProvider == ct.CloudProviderGCP {
		var credsSecret corev1.Secret
		if err := d.K8sClient.Get(ctx, fn.NN(cluster.Spec.GCP.CredentialsRef.Namespace, cluster.Spec.GCP.CredentialsRef.Name), &credsSecret); err != nil {
			return err
		}

		m := make(map[string]string)
		for k, v := range credsSecret.Data {
			m[k] = string(v)
		}

		gcpCreds, err := fn.JsonConvert[clustersv1.GCPCredentials](m)
		if err != nil {
			return err
		}

		values["gcp-service-account-json"] = base64.StdEncoding.EncodeToString([]byte(gcpCreds.ServiceAccountJSON))
	}

	b2, err := templates.ParseBytes(b, values)
	if err != nil {
		return errors.NewE(err)
	}

	var m map[string]any
	if err := yaml.Unmarshal(b2, &m); err != nil {
		return errors.NewE(err)
	}

	helmChart, err := fn.JsonConvert[crdsv1.HelmChart](m)
	if err != nil {
		return errors.NewE(err)
	}

	hr := entities.HelmRelease{
		HelmChart: helmChart,
		ResourceMetadata: common.ResourceMetadata{
			DisplayName: fmt.Sprintf("kloudlite agent %s", d.Env.KloudliteRelease),
			CreatedBy: common.CreatedOrUpdatedBy{
				UserId:    "kloudlite-platform",
				UserName:  "kloudlite-platform",
				UserEmail: "kloudlite-platform",
			},
			LastUpdatedBy: common.CreatedOrUpdatedBy{
				UserId:    "kloudlite-platform",
				UserName:  "kloudlite-platform",
				UserEmail: "kloudlite-platform",
			},
		},
		AccountName: ctx.AccountName,
		ClusterName: cluster.Name,
		SyncStatus:  t.GenSyncStatus(t.SyncActionApply, 0),
	}

	hr.IncrementRecordVersion()

	uhr, err := d.UpsertKloudliteHelmRelease(ctx, cluster.Name, &hr)
	if err != nil {
		return errors.NewE(err)
	}

	if err := d.ResDispatcher.ApplyToTargetCluster(ctx, cluster.AccountName, cluster.Name, &uhr.HelmChart, uhr.RecordVersion); err != nil {
		return errors.NewE(err)
	}

	return nil
}

func (d *Domain) UpgradeHelmKloudliteAgent(ctx domainT.InfraContext, clusterName string) error {
	clusterToken, err := d.MsgOfficeSvc.GetClusterToken(ctx, ctx.AccountName, clusterName)
	if err != nil {
		return errors.NewE(err)
	}

	cluster, err := d.FindCluster(ctx, clusterName)
	if err != nil {
		return errors.NewE(err)
	}

	if err := d.applyHelmKloudliteAgent(ctx, clusterToken, cluster); err != nil {
		return errors.NewE(err)
	}

	// if cluster.GlobalVPN != nil {
	// 	gvpn, err := d.FindGlobalVPNConnection(ctx, cluster.Name, *cluster.GlobalVPN)
	// 	if err != nil {
	// 		return errors.NewE(err)
	// 	}
	// 	if err := d.applyGlobalVPNConnection(ctx, gvpn); err != nil {
	// 		return errors.NewE(err)
	// 	}
	// }

	return nil
}

func (d *Domain) ListClusters(ctx domainT.InfraContext, mf map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.Cluster], error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.ListClusters); err != nil {
		return nil, errors.NewE(err)
	}

	accNs, err := d.GetAccNamespace(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	f := repos.Filter{
		fields.AccountName:       ctx.AccountName,
		fields.MetadataNamespace: accNs,
	}

	pr, err := d.ClusterRepo.FindPaginated(ctx, d.ClusterRepo.MergeMatchFilters(f, mf), pagination)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return pr, nil
}

func (d *Domain) GetCluster(ctx domainT.InfraContext, name string) (*entities.Cluster, error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.GetCluster); err != nil {
		return nil, errors.NewE(err)
	}

	c, err := d.FindCluster(ctx, name)
	if err != nil {
		if errors.Is(err, ErrClusterNotFound) {
			byokCluster, err := d.FindBYOKCluster(ctx, name)
			if err != nil {
				return nil, err
			}

			if byokCluster == nil {
				return nil, nil
			}

			return &entities.Cluster{
				Cluster: clustersv1.Cluster{
					ObjectMeta: metav1.ObjectMeta{
						Name: byokCluster.Name,
					},
					Spec: clustersv1.ClusterSpec{
						ClusterServiceCIDR: byokCluster.ClusterSvcCIDR,
						PublicDNSHost:      "",
					},
				},
				ResourceMetadata: byokCluster.ResourceMetadata,
				AccountName:      byokCluster.AccountName,
				GlobalVPN:        &byokCluster.GlobalVPN,
			}, nil
		}
		return nil, errors.NewE(err)
	}

	return c, nil
}

func (d *Domain) UpdateCluster(ctx domainT.InfraContext, clusterIn entities.Cluster) (*entities.Cluster, error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.UpdateCluster); err != nil {
		return nil, errors.NewE(err)
	}
	clusterIn.EnsureGVK()

	uCluster, err := d.ClusterRepo.Patch(
		ctx,
		repos.Filter{
			fields.AccountName:  ctx.AccountName,
			fields.MetadataName: clusterIn.Name,
		},
		common.PatchForUpdate(ctx, &clusterIn),
	)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if err := d.applyCluster(ctx, uCluster); err != nil {
		return nil, errors.NewE(err)
	}

	d.ResourceEventPublisher.PublishInfraEvent(ctx, ports.ResourceTypeCluster, uCluster.Name, ports.PublishUpdate)
	return uCluster, nil
}

func (d *Domain) readClusterK8sResource(ctx domainT.InfraContext, namespace string, name string) (cluster *clustersv1.Cluster, found bool, err error) {
	var clus entities.Cluster
	if err := d.K8sClient.Get(ctx, fn.NN(namespace, name), &clus.Cluster); err != nil {
		if apiErrors.IsNotFound(err) {
			return nil, false, nil
		}
	}
	return &clus.Cluster, true, nil
}

func (d *Domain) DeleteCluster(ctx domainT.InfraContext, name string) error {
	if err := d.CanPerformActionInAccount(ctx, iamT.DeleteCluster); err != nil {
		return errors.NewE(err)
	}

	// filter := repos.Filter{
	// 	fields.AccountName: ctx.AccountName,
	// 	fields.ClusterName: name,
	// }

	npCount, err := d.CountNodepools(ctx, name)
	if err != nil {
		return errors.NewE(err)
	}
	if npCount != 0 {
		return errors.Newf("delete nodepool first, aborting cluster deletion")
	}

	pvCount, err := d.CountPVs(ctx, name)
	if err != nil {
		return errors.NewE(err)
	}
	if pvCount != 0 {
		return errors.Newf("delete pvs first, aborting cluster deletion")
	}

	ucluster, err := d.ClusterRepo.Patch(
		ctx,
		repos.Filter{
			fields.AccountName:  ctx.AccountName,
			fields.MetadataName: name,
		},
		common.PatchForMarkDeletion(),
	)
	if err != nil {
		return errors.NewE(err)
	}

	d.ResourceEventPublisher.PublishInfraEvent(ctx, ports.ResourceTypeCluster, ucluster.Name, ports.PublishUpdate)
	if err := d.DeleteK8sResource(ctx, &ucluster.Cluster); err != nil {
		if !apiErrors.IsNotFound(err) {
			return errors.NewE(err)
		}

		return d.OnClusterDeleteMessage(ctx, *ucluster)
	}

	return nil
}

func (d *Domain) OnClusterDeleteMessage(ctx domainT.InfraContext, cluster entities.Cluster) error {
	xcluster, err := d.FindCluster(ctx, cluster.Name)
	if err != nil {
		return errors.NewE(err)
	}

	if err = d.ClusterRepo.DeleteById(ctx, xcluster.Id); err != nil {
		return errors.NewE(err)
	}

	d.ResourceEventPublisher.PublishInfraEvent(ctx, ports.ResourceTypeCluster, cluster.Name, ports.PublishDelete)

	if xcluster.GlobalVPN != nil {
		// if err := d.claimClusterSvcCIDRRepo.DeleteOne(ctx, repos.Filter{
		// 	fc.ClaimClusterSvcCIDRClaimedByCluster: xcluster.Name,
		// 	fc.AccountName:                         ctx.AccountName,
		// 	fc.ClaimClusterSvcCIDRGlobalVPNName:    xcluster.GlobalVPN,
		// }); err != nil {
		// 	return errors.NewE(err)
		// }
		//
		// if _, err := d.freeClusterSvcCIDRRepo.Create(ctx, &entities.FreeClusterSvcCIDR{
		// 	AccountName:    ctx.AccountName,
		// 	GlobalVPNName:  *xcluster.GlobalVPN,
		// 	ClusterSvcCIDR: xcluster.Spec.ClusterServiceCIDR,
		// }); err != nil {
		// 	return errors.NewE(err)
		// }

		gv, err := d.FindGlobalVPNConnection(ctx, xcluster.Name, *xcluster.GlobalVPN)
		if err != nil {
			return errors.NewE(err)
		}

		// if err := d.OnGlobalVPNConnectionDeleteMessage(ctx, xcluster.Name, *gv); err != nil {
		if err := d.DeleteGlobalVPNConnection(ctx, gv.Name, xcluster.Name); err != nil {
			return errors.NewE(err)
		}
	}

	return nil
}

func (d *Domain) OnClusterUpdateMessage(ctx domainT.InfraContext, cluster entities.Cluster, status types.ResourceStatus, opts domainT.UpdateAndDeleteOpts) error {
	xCluster, err := d.FindCluster(ctx, cluster.Name)
	if err != nil {
		return errors.NewE(err)
	}

	recordVersion, err := d.MatchRecordVersion(cluster.Annotations, xCluster.RecordVersion)
	if err != nil {
		return nil
	}

	patchDoc := repos.Document{}
	if cluster.Spec.Output != nil {
		patchDoc[fc.ClusterSpecOutput] = cluster.Spec.Output
	}

	if cluster.Spec.AWS != nil && cluster.Spec.AWS.VPC != nil {
		patchDoc[fc.ClusterSpecAwsVpc] = cluster.Spec.AWS.VPC
	}

	if cluster.Spec.GCP != nil && cluster.Spec.GCP.VPC != nil {
		patchDoc[fc.ClusterSpecGcpVpc] = cluster.Spec.GCP.VPC
	}

	uCluster, err := d.ClusterRepo.PatchById(
		ctx,
		xCluster.Id,
		common.PatchForSyncFromAgent(&cluster, recordVersion, status, common.PatchOpts{
			MessageTimestamp: opts.MessageTimestamp,
			XPatch:           patchDoc,
		}))
	d.ResourceEventPublisher.PublishInfraEvent(ctx, ports.ResourceTypeCluster, uCluster.GetName(), ports.PublishUpdate)
	return errors.NewE(err)
}

func (d *Domain) FindCluster(ctx domainT.InfraContext, clusterName string) (*entities.Cluster, error) {
	accNs, err := d.GetAccNamespace(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	cluster, err := d.ClusterRepo.FindOne(ctx, repos.Filter{
		fields.AccountName:       ctx.AccountName,
		fields.MetadataName:      clusterName,
		fields.MetadataNamespace: accNs,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if cluster == nil {
		return nil, ErrClusterNotFound
	}
	return cluster, nil
}

func (d *Domain) clusterAlreadyExists(ctx domainT.InfraContext, name string) (*bool, error) {
	exists, err := d.ClusterRepo.FindOne(ctx, repos.Filter{
		fc.AccountName:  ctx.AccountName,
		fc.MetadataName: name,
	})
	if err != nil {
		return nil, err
	}
	if exists != nil {
		return fn.New(true), nil
	}

	existsBYOK, err := d.FindBYOKCluster(ctx, name)
	if err != nil {
		return nil, err
	}

	if existsBYOK != nil {
		return fn.New(true), nil
	}

	return fn.New(false), nil
}
