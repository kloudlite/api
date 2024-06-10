package nodepool

import (
	"fmt"
	"strings"

	iamT "github.com/kloudlite/api/apps/iam/types"
	fc "github.com/kloudlite/api/apps/infra/internal/entities/field-constants"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/errors"
	clustersv1 "github.com/kloudlite/operator/apis/clusters/v1"
	ct "github.com/kloudlite/operator/apis/common-types"
	"github.com/kloudlite/operator/operators/resource-watcher/types"

	"github.com/kloudlite/api/apps/infra/internal/domain/ports"
	domainT "github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/repos"
	t "github.com/kloudlite/api/pkg/types"
)

func (d *Domain) applyNodePool(ctx domainT.InfraContext, np *entities.NodePool) error {
	d.AddTrackingId(&np.NodePool, np.Id)
	return d.ResDispatcher.ApplyToTargetCluster(ctx, ctx.AccountName, np.ClusterName, &np.NodePool, np.RecordVersion)
}

func (d *Domain) CreateNodePool(ctx domainT.InfraContext, clusterName string, nodepool entities.NodePool) (*entities.NodePool, error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.CreateNodepool); err != nil {
		return nil, errors.NewE(err)
	}

	nodepool.IncrementRecordVersion()
	nodepool.CreatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}
	nodepool.LastUpdatedBy = nodepool.CreatedBy

	cluster, err := d.FindCluster(ctx, clusterName)
	if err != nil {
		return nil, errors.NewE(err)
	}

	switch nodepool.Spec.CloudProvider {
	case ct.CloudProviderAWS:
		{
			ps, err := d.FindProviderSecret(ctx, cluster.Spec.AWS.Credentials.SecretRef.Name)
			if err != nil {
				return nil, errors.NewE(err)
			}

			awsSubnetID := cluster.Spec.AWS.VPC.GetSubnetId(nodepool.Spec.AWS.AvailabilityZone)
			if awsSubnetID == "" {
				return nil, errors.Newf("kloudlite VPC has no subnet configured for this availability zone (%s), please select another availability zone in your cluster's region (%s)", nodepool.Spec.AWS.AvailabilityZone, cluster.Spec.AWS.Region)
			}

			nodepool.Spec.AWS = &clustersv1.AWSNodePoolConfig{
				VPCId:       cluster.Spec.AWS.VPC.ID,
				VPCSubnetID: awsSubnetID,

				AvailabilityZone: nodepool.Spec.AWS.AvailabilityZone,
				NvidiaGpuEnabled: nodepool.Spec.AWS.NvidiaGpuEnabled,
				RootVolumeType:   "gp3",
				RootVolumeSize: func() int {
					if nodepool.Spec.AWS.NvidiaGpuEnabled {
						return 80
					}
					return 50
				}(),
				IAMInstanceProfileRole: &ps.AWS.CfParamInstanceProfileName,
				PoolType:               nodepool.Spec.AWS.PoolType,
				EC2Pool:                nodepool.Spec.AWS.EC2Pool,
				SpotPool: func() *clustersv1.AwsSpotPoolConfig {
					if nodepool.Spec.AWS.SpotPool == nil {
						return nil
					}
					return &clustersv1.AwsSpotPoolConfig{
						SpotFleetTaggingRoleName: ps.AWS.CfParamRoleName,
						CpuNode:                  nodepool.Spec.AWS.SpotPool.CpuNode,
						GpuNode:                  nodepool.Spec.AWS.SpotPool.GpuNode,
						Nodes:                    nodepool.Spec.AWS.SpotPool.Nodes,
					}
				}(),
			}
		}
	case ct.CloudProviderGCP:
		{

			k8sSecret, err := d.GetProviderSecretAsK8sSecret(ctx, clusterName)
			if err != nil {
				return nil, errors.NewE(err)
			}

			if k8sSecret == nil {
				return nil, errors.Newf("failed to get provider secret as a k8s secret")
			}

			// INFO: because kube-system is omnipresent on k8s
			k8sSecret.Namespace = "kube-system"

			if err := d.ResDispatcher.ApplyToTargetCluster(ctx, ctx.AccountName, clusterName, k8sSecret, nodepool.RecordVersion); err != nil {
				return nil, errors.NewE(err)
			}

			nodepool.Spec.GCP = &clustersv1.GCPNodePoolConfig{
				Region: cluster.Spec.GCP.Region,
				AvailabilityZone: func() string {
					if strings.TrimSpace(nodepool.Spec.GCP.AvailabilityZone) != "" {
						return nodepool.Spec.GCP.AvailabilityZone
					}
					return fmt.Sprintf("%s-a", cluster.Spec.GCP.Region)
				}(),
				GCPProjectID: cluster.Spec.GCP.GCPProjectID,
				VPC:          &clustersv1.GcpVPCParams{Name: cluster.Spec.GCP.VPC.Name},
				Credentials: ct.SecretRef{
					Name:      k8sSecret.Name,
					Namespace: k8sSecret.Namespace,
				},
				// FIXME: once, we allow gcp service account for nodepools via UI
				ServiceAccount: clustersv1.GCPServiceAccount{
					Enabled: false,
					Email:   nil,
					Scopes:  nil,
				},
				PoolType:       nodepool.Spec.GCP.PoolType,
				MachineType:    nodepool.Spec.GCP.MachineType,
				BootVolumeType: "pd-ssd",
				BootVolumeSize: 50,
				Nodes:          map[string]clustersv1.NodeProps{},
			}
		}
	default:
		{
			return nil, errors.Newf("cloudprovider: %s, currently not supported", nodepool.Spec.CloudProvider)
		}
	}

	nodepool.AccountName = ctx.AccountName
	nodepool.ClusterName = clusterName
	nodepool.SyncStatus = t.GenSyncStatus(t.SyncActionApply, nodepool.RecordVersion)

	nodepool.EnsureGVK()
	if err := d.K8sClient.ValidateObject(ctx, &nodepool.NodePool); err != nil {
		return nil, errors.NewE(err)
	}
	nodepool.IncrementRecordVersion()

	np, err := d.NodepoolRepo.Create(ctx, &nodepool)
	if err != nil {
		if d.NodepoolRepo.ErrAlreadyExists(err) {
			return nil, errors.Newf("nodepool with name %q already exists", nodepool.Name)
		}
		return nil, errors.NewE(err)
	}
	d.ResourceEventPublisher.PublishResourceEvent(ctx, clusterName, ports.ResourceTypeNodePool, np.Name, ports.PublishAdd)

	if err := d.applyNodePool(ctx, np); err != nil {
		return nil, errors.NewE(err)
	}

	return np, nil
}

func (d *Domain) UpdateNodePool(ctx domainT.InfraContext, clusterName string, nodePoolIn entities.NodePool) (*entities.NodePool, error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.UpdateNodepool); err != nil {
		return nil, errors.NewE(err)
	}

	nodePoolIn.EnsureGVK()
	if err := d.K8sClient.ValidateObject(ctx, &nodePoolIn.NodePool); err != nil {
		return nil, errors.NewE(err)
	}

	patchForUpdate := common.PatchForUpdate(
		ctx,
		&nodePoolIn,
		common.PatchOpts{
			XPatch: repos.Document{
				fc.NodePoolSpecMinCount: nodePoolIn.Spec.MinCount,
				fc.NodePoolSpecMaxCount: nodePoolIn.Spec.MaxCount,
			},
		})

	unp, err := d.NodepoolRepo.Patch(
		ctx,
		repos.Filter{
			fields.ClusterName:  clusterName,
			fields.AccountName:  ctx.AccountName,
			fields.MetadataName: nodePoolIn.Name,
		},
		patchForUpdate,
	)
	if err != nil {
		return nil, errors.NewE(err)
	}

	d.ResourceEventPublisher.PublishResourceEvent(ctx, clusterName, ports.ResourceTypeNodePool, unp.Name, ports.PublishUpdate)

	if err := d.applyNodePool(ctx, unp); err != nil {
		return nil, errors.NewE(err)
	}

	return unp, nil
}

func (d *Domain) DeleteNodePool(ctx domainT.InfraContext, clusterName string, poolName string) error {
	if err := d.CanPerformActionInAccount(ctx, iamT.DeleteNodepool); err != nil {
		return errors.NewE(err)
	}

	unp, err := d.NodepoolRepo.Patch(
		ctx,
		repos.Filter{
			fields.ClusterName:  clusterName,
			fields.AccountName:  ctx.AccountName,
			fields.MetadataName: poolName,
		},
		common.PatchForMarkDeletion(),
	)
	if err != nil {
		return errors.NewE(err)
	}

	d.ResourceEventPublisher.PublishResourceEvent(ctx, clusterName, ports.ResourceTypeNodePool, unp.Name, ports.PublishUpdate)
	return d.ResDispatcher.DeleteFromTargetCluster(ctx, ctx.AccountName, clusterName, &unp.NodePool)
}

func (d *Domain) GetNodePool(ctx domainT.InfraContext, clusterName string, poolName string) (*entities.NodePool, error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.GetNodepool); err != nil {
		return nil, errors.NewE(err)
	}
	np, err := d.findNodePool(ctx, clusterName, poolName)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return np, nil
}

func (d *Domain) ListNodePools(ctx domainT.InfraContext, clusterName string, matchFilters map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.NodePool], error) {
	if err := d.CanPerformActionInAccount(ctx, iamT.ListNodepools); err != nil {
		return nil, errors.NewE(err)
	}
	filter := repos.Filter{
		fields.AccountName: ctx.AccountName,
		fields.ClusterName: clusterName,
	}
	return d.NodepoolRepo.FindPaginated(ctx, d.NodepoolRepo.MergeMatchFilters(filter, matchFilters), pagination)
}

func (d *Domain) findNodePool(ctx domainT.InfraContext, clusterName string, poolName string) (*entities.NodePool, error) {
	np, err := d.NodepoolRepo.FindOne(ctx, repos.Filter{
		fields.AccountName:  ctx.AccountName,
		fields.ClusterName:  clusterName,
		fields.MetadataName: poolName,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}
	if np == nil {
		return nil, errors.Newf("nodepool with name %q not found", clusterName)
	}
	return np, nil
}

func (d *Domain) ResyncNodePool(ctx domainT.InfraContext, clusterName string, poolName string) error {
	if err := func() error {
		if err := d.CanPerformActionInAccount(ctx, iamT.UpdateNodepool); err != nil {
			return d.CanPerformActionInAccount(ctx, iamT.DeleteNodepool)
		}
		return nil
	}(); err != nil {
		return errors.NewE(err)
	}
	np, err := d.findNodePool(ctx, clusterName, poolName)
	if err != nil {
		return errors.NewE(err)
	}

	return d.ResyncToTargetCluster(ctx, np.SyncStatus.Action, clusterName, &np.NodePool, np.RecordVersion)
}

// on message events

func (d *Domain) OnNodePoolDeleteMessage(ctx domainT.InfraContext, clusterName string, nodePool entities.NodePool) error {
	err := d.NodepoolRepo.DeleteOne(
		ctx,
		repos.Filter{
			fields.AccountName:  ctx.AccountName,
			fields.ClusterName:  clusterName,
			fields.MetadataName: nodePool.Name,
		},
	)
	if err != nil {
		return errors.NewE(err)
	}
	d.ResourceEventPublisher.PublishResourceEvent(ctx, clusterName, ports.ResourceTypeNodePool, nodePool.Name, ports.PublishDelete)
	return err
}

func (d *Domain) OnNodePoolUpdateMessage(ctx domainT.InfraContext, clusterName string, nodePool entities.NodePool, status types.ResourceStatus, opts domainT.UpdateAndDeleteOpts) error {
	xnp, err := d.findNodePool(ctx, clusterName, nodePool.Name)
	if err != nil {
		return errors.NewE(err)
	}

	if xnp == nil {
		return errors.Newf("no nodepool found")
	}

	if _, err := d.MatchRecordVersion(nodePool.Annotations, xnp.RecordVersion); err != nil {
		return d.ResyncToTargetCluster(ctx, xnp.SyncStatus.Action, clusterName, &xnp.NodePool, xnp.RecordVersion)
	}

	recordVersion, err := d.MatchRecordVersion(nodePool.Annotations, xnp.RecordVersion)
	if err != nil {
		return errors.NewE(err)
	}

	unp, err := d.NodepoolRepo.PatchById(
		ctx,
		xnp.Id,
		common.PatchForSyncFromAgent(&nodePool,
			recordVersion, status,
			common.PatchOpts{
				MessageTimestamp: opts.MessageTimestamp,
			}))
	if err != nil {
		return errors.NewE(err)
	}

	d.ResourceEventPublisher.PublishResourceEvent(ctx, clusterName, ports.ResourceTypeNodePool, unp.GetName(), ports.PublishUpdate)
	return nil
}

// OnNodepoolApplyError implements Domain.
func (d *Domain) OnNodepoolApplyError(ctx domainT.InfraContext, clusterName string, name string, errMsg string, opts domainT.UpdateAndDeleteOpts) error {
	unp, err := d.NodepoolRepo.Patch(
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
	d.ResourceEventPublisher.PublishResourceEvent(ctx, clusterName, ports.ResourceTypeNodePool, unp.Name, ports.PublishUpdate)
	return errors.NewE(err)
}
