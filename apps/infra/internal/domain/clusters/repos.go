package clusters

import (
	"github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/pkg/repos"
)

type GetClusterArgs struct {
	ClusterRepo repos.DbRepo[*entities.Cluster]
	ClusterName string
}

func GetCluster(ctx types.InfraContext, args GetClusterArgs) (*entities.Cluster, error) {
	return args.ClusterRepo.FindOne(ctx, entities.UniqueCluster(ctx.AccountName, args.ClusterName))
}
