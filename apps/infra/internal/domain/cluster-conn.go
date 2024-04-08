package domain

import (
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/operator/operators/resource-watcher/types"
)

func (d *domain) OnClusterConnDeleteMessage(ctx InfraContext, clusterName string, clusterConn entities.ClusterConnection) error {
	panic("implement me")
}
func (d *domain) OnClusterConnUpdateMessage(ctx InfraContext, clusterName string, clusterConn entities.ClusterConnection, status types.ResourceStatus, opts UpdateAndDeleteOpts) error {
	panic("implement me")
}
func (d *domain) OnClusterConnApplyError(ctx InfraContext, clusterName string, name string, errMsg string, opts UpdateAndDeleteOpts) error {
	panic("implement me")
}
