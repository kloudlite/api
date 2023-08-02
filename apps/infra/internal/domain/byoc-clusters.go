package domain

// import (
// 	"fmt"
// 	"kloudlite.io/apps/infra/internal/entities"
//
// 	redpandaMsvcv1 "github.com/kloudlite/operator/apis/redpanda.msvc/v1"
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	"kloudlite.io/common"
// 	"kloudlite.io/pkg/repos"
// 	t "kloudlite.io/pkg/types"
// )
//
// func (d *domain) findBYOCCluster(ctx InfraContext, clusterName string) (*entities.BYOCCluster, error) {
// 	cluster, err := d.byocClusterRepo.FindOne(ctx, repos.Filter{
// 		"spec.accountName": ctx.AccountName,
// 		"metadata.name":    clusterName,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	if cluster == nil {
// 		return nil, fmt.Errorf("BYOC cluster with name %q not found", clusterName)
// 	}
// 	return cluster, nil
// }
//
// func (d *domain) CreateBYOCCluster(ctx InfraContext, cluster entities.BYOCCluster) (*entities.BYOCCluster, error) {
// 	cluster.EnsureGVK()
// 	cluster.IncomingKafkaTopicName = common.GetKafkaTopicName(ctx.AccountName, cluster.Name)
//
// 	if err := d.k8sExtendedClient.ValidateStruct(ctx, &cluster.BYOC); err != nil {
// 		return nil, err
// 	}
//
// 	cluster.IncrementRecordVersion()
// 	cluster.IsConnected = false
// 	cluster.Spec.AccountName = ctx.AccountName
// 	cluster.SyncStatus = t.GenSyncStatus(t.SyncActionApply, cluster.RecordVersion)
//
// 	nCluster, err := d.byocClusterRepo.Create(ctx, &cluster)
// 	if err != nil {
// 		if d.clusterRepo.ErrAlreadyExists(err) {
// 			return nil, fmt.Errorf("cluster with name %q already exists", cluster.Name)
// 		}
// 	}
//
// 	if err := d.applyK8sResource(ctx, &nCluster.BYOC, nCluster.RecordVersion); err != nil {
// 		return nil, err
// 	}
//
// 	redpandaTopic := redpandaMsvcv1.Topic{
// 		TypeMeta:   metav1.TypeMeta{},
// 		ObjectMeta: metav1.ObjectMeta{Name: cluster.IncomingKafkaTopicName, Namespace: d.env.ProviderSecretNamespace},
// 	}
//
// 	redpandaTopic.EnsureGVK()
//
// 	if err := d.applyK8sResource(ctx, &redpandaTopic, nCluster.RecordVersion); err != nil {
// 		return nil, err
// 	}
//
// 	return nCluster, nil
// }
//
// func (d *domain) ListBYOCClusters(ctx InfraContext, pagination t.CursorPagination) (*repos.PaginatedRecord[*entities.BYOCCluster], error) {
// 	return d.byocClusterRepo.FindPaginated(ctx, repos.Filter{
// 		"spec.accountName": ctx.AccountName,
// 	}, pagination)
// }
//
// func (d *domain) GetBYOCCluster(ctx InfraContext, name string) (*entities.BYOCCluster, error) {
// 	return d.findBYOCCluster(ctx, name)
// }
//
// func (d *domain) UpdateBYOCCluster(ctx InfraContext, cluster entities.BYOCCluster) (*entities.BYOCCluster, error) {
// 	cluster.EnsureGVK()
// 	if err := d.k8sExtendedClient.ValidateStruct(ctx, &cluster.BYOC); err != nil {
// 		return nil, err
// 	}
//
// 	c, err := d.findBYOCCluster(ctx, cluster.Name)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	c.IncrementRecordVersion()
// 	c.BYOC = cluster.BYOC
// 	c.SyncStatus = t.GenSyncStatus(t.SyncActionApply, c.RecordVersion)
//
// 	// c.Spec.AccountName = ctx.AccountName
// 	// c.Spec.Region = cluster.Spec.Region
// 	// c.Spec.Provider = cluster.Spec.Provider
// 	uCluster, err := d.byocClusterRepo.UpdateById(ctx, c.Id, c)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	if err := d.applyK8sResource(ctx, &uCluster.BYOC, uCluster.RecordVersion); err != nil {
// 		return nil, err
// 	}
//
// 	return uCluster, nil
// }
//
// func (d *domain) DeleteBYOCCluster(ctx InfraContext, name string) error {
// 	clus, err := d.findBYOCCluster(ctx, name)
// 	if err != nil {
// 		return err
// 	}
//
// 	clus.SyncStatus = t.GetSyncStatusForDeletion(clus.Generation)
// 	upC, err := d.byocClusterRepo.UpdateById(ctx, clus.Id, clus)
// 	if err != nil {
// 		return err
// 	}
// 	return d.deleteK8sResource(ctx, &upC.BYOC)
// }
//
// func (d *domain) ResyncBYOCCluster(ctx InfraContext, name string) error {
// 	clus, err := d.findBYOCCluster(ctx, name)
// 	if err != nil {
// 		return err
// 	}
//
// 	if err := d.applyK8sResource(ctx, &clus.BYOC, clus.RecordVersion); err != nil {
// 		return err
// 	}
//
// 	redpandaTopic := redpandaMsvcv1.Topic{
// 		TypeMeta: metav1.TypeMeta{},
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      clus.IncomingKafkaTopicName,
// 			Namespace: d.env.ProviderSecretNamespace,
// 		},
// 	}
//
// 	redpandaTopic.EnsureGVK()
// 	return d.applyK8sResource(ctx, &redpandaTopic, clus.RecordVersion)
// }
//
// func (d *domain) OnDeleteBYOCClusterMessage(ctx InfraContext, cluster entities.BYOCCluster) error {
// 	return d.clusterRepo.DeleteOne(ctx, repos.Filter{
// 		"spec.accountName": ctx.AccountName,
// 		"metadata.name":    cluster.Name,
// 	})
// }
//
// func (d *domain) OnBYOCClusterHelmUpdates(ctx InfraContext, cluster entities.BYOCCluster) error {
// 	c, err := d.findBYOCCluster(ctx, cluster.Name)
// 	if err != nil {
// 		return err
// 	}
//
// 	c.SyncStatus.State = t.SyncStateReceivedUpdateFromAgent
//
// 	_, err = d.byocClusterRepo.UpdateById(ctx, c.Id, &cluster)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }