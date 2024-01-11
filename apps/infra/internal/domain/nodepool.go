package domain

import (
	"fmt"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/operator/operators/resource-watcher/types"

	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/common"
	clustersv1 "github.com/kloudlite/operator/apis/clusters/v1"
	ct "github.com/kloudlite/operator/apis/common-types"

	"github.com/kloudlite/api/apps/infra/internal/entities"
	fn "github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/repos"
	t "github.com/kloudlite/api/pkg/types"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const tenantControllerNamespace = "kloudlite"

func (d *domain) applyNodePool(ctx InfraContext, np *entities.NodePool) error {
	addTrackingId(&np.NodePool, np.Id)
	return d.resDispatcher.ApplyToTargetCluster(ctx, np.ClusterName, &np.NodePool, np.RecordVersion)
}

func (d *domain) CreateNodePool(ctx InfraContext, clusterName string, nodepool entities.NodePool) (*entities.NodePool, error) {
	if err := d.canPerformActionInAccount(ctx, iamT.CreateNodepool); err != nil {
		return nil, errors.NewE(err)
	}

	nodepool.IncrementRecordVersion()
	nodepool.CreatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}
	nodepool.LastUpdatedBy = nodepool.CreatedBy

	out, err := d.accountsSvc.GetAccount(ctx, string(ctx.UserId), ctx.AccountName)
	if err != nil {
		return nil, errors.NewE(err)
	}

	cluster, err := d.findCluster(ctx, clusterName)
	if err != nil {
		return nil, errors.NewE(err)
	}

	// fetch cloud provider credentials, access key, and ps key
	credsSecret := &corev1.Secret{}
	if err := d.k8sClient.Get(ctx, fn.NN(cluster.Spec.CredentialsRef.Namespace, cluster.Spec.CredentialsRef.Name), credsSecret); err != nil {
		return nil, errors.NewE(err)
	}

	providerSecret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "provider-creds",
			Namespace: tenantControllerNamespace,
		},
		Data: map[string][]byte{
			"access_key": credsSecret.Data[cluster.Spec.CredentialKeys.KeyAccessKey],
			"secret_key": credsSecret.Data[cluster.Spec.CredentialKeys.KeySecretKey],
		},
	}

	if err := d.resDispatcher.ApplyToTargetCluster(ctx, clusterName, providerSecret, 1); err != nil {
		return nil, errors.NewE(err)
	}

	nodepool.Spec.IAC = clustersv1.InfrastuctureAsCode{
		StateS3BucketName:     fmt.Sprintf("kl-%s", out.AccountId),
		StateS3BucketRegion:   "ap-south-1",
		StateS3BucketFilePath: fmt.Sprintf("iac/kl-account-%s/cluster-%s/nodepool-%s.tfstate", ctx.AccountName, clusterName, nodepool.Name),
		CloudProviderAccessKey: ct.SecretKeyRef{
			Name:      providerSecret.Name,
			Namespace: providerSecret.Namespace,
			Key:       "access_key",
		},
		CloudProviderSecretKey: ct.SecretKeyRef{
			Name:      providerSecret.Name,
			Namespace: providerSecret.Namespace,
			Key:       "secret_key",
		},
	}

	ps, err := d.findProviderSecret(ctx, cluster.Spec.CredentialsRef.Name)
	if err != nil {
		return nil, errors.NewE(err)
	}

	switch nodepool.Spec.CloudProvider {
	case ct.CloudProviderAWS:
		{
			nodepool.Spec.AWS = &clustersv1.AWSNodePoolConfig{
				ImageId:          "ami-06d146e85d1709abb",
				ImageSSHUsername: "ubuntu",
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
	}

	if nodepool.Spec.TargetCount < nodepool.Spec.MinCount {
		nodepool.Spec.TargetCount = nodepool.Spec.MinCount
	}

	nodepool.AccountName = ctx.AccountName
	nodepool.ClusterName = clusterName
	nodepool.SyncStatus = t.GenSyncStatus(t.SyncActionApply, nodepool.RecordVersion)

	nodepool.EnsureGVK()
	if err := d.k8sClient.ValidateObject(ctx, &nodepool.NodePool); err != nil {
		return nil, errors.NewE(err)
	}
	nodepool.IncrementRecordVersion()

	np, err := d.nodePoolRepo.Create(ctx, &nodepool)
	if err != nil {
		if d.nodePoolRepo.ErrAlreadyExists(err) {
			return nil, errors.Newf("nodepool with name %q already exists", nodepool.Name)
		}
		return nil, errors.NewE(err)
	}
	d.resourceEventPublisher.PublishNodePoolEvent(&nodepool, PublishAdd)

	if err := d.applyNodePool(ctx, np); err != nil {
		return nil, errors.NewE(err)
	}

	return np, nil
}

func (d *domain) UpdateNodePool(ctx InfraContext, clusterName string, nodePool entities.NodePool) (*entities.NodePool, error) {
	if err := d.canPerformActionInAccount(ctx, iamT.UpdateNodepool); err != nil {
		return nil, errors.NewE(err)
	}

	nodePool.EnsureGVK()
	if err := d.k8sClient.ValidateObject(ctx, &nodePool.NodePool); err != nil {
		return nil, errors.NewE(err)
	}

	np, err := d.findNodePool(ctx, clusterName, nodePool.Name)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if np.IsMarkedForDeletion() {
		return nil, errors.Newf("nodepool %q (clusterName=%q) is marked for deletion, aborting update", nodePool.Name, clusterName)
	}

	np.IncrementRecordVersion()
	np.LastUpdatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}

	np.Labels = nodePool.Labels
	np.Annotations = nodePool.Annotations
	np.Spec.MinCount = nodePool.Spec.MinCount
	np.Spec.MaxCount = nodePool.Spec.MaxCount
	np.Spec.TargetCount = nodePool.Spec.TargetCount

	if np.Spec.TargetCount < np.Spec.MinCount {
		np.Spec.TargetCount = np.Spec.MinCount
	}

	np.SyncStatus = t.GenSyncStatus(t.SyncActionApply, np.RecordVersion)

	unp, err := d.nodePoolRepo.UpdateById(ctx, np.Id, np)
	if err != nil {
		return nil, errors.NewE(err)
	}

	d.resourceEventPublisher.PublishNodePoolEvent(unp, PublishUpdate)
	d.resourceEventPublisher.PublishNodePoolEvent(unp, PublishDelete)

	if err := d.applyNodePool(ctx, unp); err != nil {
		return nil, errors.NewE(err)
	}

	return unp, nil
}

func (d *domain) DeleteNodePool(ctx InfraContext, clusterName string, poolName string) error {
	if err := d.canPerformActionInAccount(ctx, iamT.DeleteNodepool); err != nil {
		return errors.NewE(err)
	}
	np, err := d.findNodePool(ctx, clusterName, poolName)
	if err != nil {
		return errors.NewE(err)
	}

	if np.IsMarkedForDeletion() {
		return errors.Newf("nodepool %q (clusterName=%q) is already marked for deletion", poolName, clusterName)
	}

	np.MarkedForDeletion = fn.New(true)
	np.SyncStatus = t.GetSyncStatusForDeletion(np.Generation)
	upC, err := d.nodePoolRepo.UpdateById(ctx, np.Id, np)
	if err != nil {
		return errors.NewE(err)
	}
	d.resourceEventPublisher.PublishNodePoolEvent(upC, PublishUpdate)
	return d.resDispatcher.DeleteFromTargetCluster(ctx, clusterName, &upC.NodePool)
}

func (d *domain) GetNodePool(ctx InfraContext, clusterName string, poolName string) (*entities.NodePool, error) {
	if err := d.canPerformActionInAccount(ctx, iamT.GetNodepool); err != nil {
		return nil, errors.NewE(err)
	}
	np, err := d.findNodePool(ctx, clusterName, poolName)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return np, nil
}

func (d *domain) ListNodePools(ctx InfraContext, clusterName string, matchFilters map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.NodePool], error) {
	if err := d.canPerformActionInAccount(ctx, iamT.ListNodepools); err != nil {
		return nil, errors.NewE(err)
	}
	filter := repos.Filter{
		"accountName": ctx.AccountName,
		"clusterName": clusterName,
	}
	return d.nodePoolRepo.FindPaginated(ctx, d.nodePoolRepo.MergeMatchFilters(filter, matchFilters), pagination)
}

func (d *domain) findNodePool(ctx InfraContext, clusterName string, poolName string) (*entities.NodePool, error) {
	np, err := d.nodePoolRepo.FindOne(ctx, repos.Filter{
		"accountName":   ctx.AccountName,
		"clusterName":   clusterName,
		"metadata.name": poolName,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}
	if np == nil {
		return nil, errors.Newf("nodepool with name %q not found", clusterName)
	}
	return np, nil
}

func (d *domain) ResyncNodePool(ctx InfraContext, clusterName string, poolName string) error {
	if err := func() error {
		if err := d.canPerformActionInAccount(ctx, iamT.UpdateNodepool); err != nil {
			return d.canPerformActionInAccount(ctx, iamT.DeleteNodepool)
		}
		return nil
	}(); err != nil {
		return errors.NewE(err)
	}
	np, err := d.findNodePool(ctx, clusterName, poolName)
	if err != nil {
		return errors.NewE(err)
	}

	return d.resyncToTargetCluster(ctx, np.SyncStatus.Action, clusterName, &np.NodePool, np.RecordVersion)
}

// on message events

func (d *domain) OnDeleteNodePoolMessage(ctx InfraContext, clusterName string, nodePool entities.NodePool) error {
	np, _ := d.findNodePool(ctx, clusterName, nodePool.Name)
	if np == nil {
		// does not exist, (maybe already deleted)
		return nil
	}

	if err := d.matchRecordVersion(nodePool.Annotations, np.RecordVersion); err != nil {
		return d.resyncToTargetCluster(ctx, np.SyncStatus.Action, clusterName, &np.NodePool, np.RecordVersion)
	}

	err := d.nodePoolRepo.DeleteById(ctx, np.Id)
	d.resourceEventPublisher.PublishNodePoolEvent(np, PublishDelete)
	return err
}

func (d *domain) OnUpdateNodePoolMessage(ctx InfraContext, clusterName string, nodePool entities.NodePool, status types.ResourceStatus, opts UpdateAndDeleteOpts) error {
	np, err := d.findNodePool(ctx, clusterName, nodePool.Name)
	if err != nil {
		return errors.NewE(err)
	}

	if err := d.matchRecordVersion(nodePool.Annotations, np.RecordVersion); err != nil {
		return d.resyncToTargetCluster(ctx, np.SyncStatus.Action, clusterName, &np.NodePool, np.RecordVersion)
	}

	np.Status = nodePool.Status

	np.SyncStatus.State = func() t.SyncState {
		if status == types.ResourceStatusDeleting {
			return t.SyncStateDeletingAtAgent
		}
		return t.SyncStateUpdatedAtAgent
	}()
	np.SyncStatus.LastSyncedAt = opts.MessageTimestamp
	np.SyncStatus.Error = nil
	np.SyncStatus.RecordVersion = np.RecordVersion

	if _, err := d.nodePoolRepo.UpdateById(ctx, np.Id, np); err != nil {
		return errors.NewE(err)
	}
	d.resourceEventPublisher.PublishNodePoolEvent(np, PublishUpdate)
	return nil
}

// OnNodepoolApplyError implements Domain.
func (d *domain) OnNodepoolApplyError(ctx InfraContext, clusterName string, name string, errMsg string, opts UpdateAndDeleteOpts) error {
	np, err := d.findNodePool(ctx, clusterName, name)
	if err != nil {
		return errors.NewE(err)
	}

	np.SyncStatus.State = t.SyncStateErroredAtAgent
	np.SyncStatus.LastSyncedAt = opts.MessageTimestamp
	np.SyncStatus.Error = &errMsg

	_, err = d.nodePoolRepo.UpdateById(ctx, np.Id, np)
	d.resourceEventPublisher.PublishNodePoolEvent(np, PublishUpdate)
	return errors.NewE(err)
}
