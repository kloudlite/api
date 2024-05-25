package byok_clusters

import (
	"time"

	"github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	fc "github.com/kloudlite/api/apps/infra/internal/entities/field-constants"
	"github.com/kloudlite/api/pkg/repos"
	t "github.com/kloudlite/api/pkg/types"
)

type IsBYOKClusterArgs struct {
	ClusterName     string
	BYOKClusterRepo repos.DbRepo[*entities.BYOKCluster]
}

func IsBYOKCluster(ctx types.InfraContext, args IsBYOKClusterArgs) (bool, error) {
	return args.BYOKClusterRepo.Exists(ctx, entities.UniqueBYOKClusterFilter(ctx.AccountName, args.ClusterName))
}

type MarkBYOKClusterReadyArgs struct {
	ClusterName     string
	BYOKClusterRepo repos.DbRepo[*entities.BYOKCluster]
	Time            time.Time
}

func MarkBYOKClusterReady(ctx types.InfraContext, args MarkBYOKClusterReadyArgs) (*entities.BYOKCluster, error) {
	return args.BYOKClusterRepo.PatchOne(ctx, entities.UniqueBYOKClusterFilter(ctx.AccountName, args.ClusterName), repos.Document{
		fc.SyncStatusState:        t.SyncStateUpdatedAtAgent,
		fc.SyncStatusLastSyncedAt: args.Time,
		fc.SyncStatusError:        nil,
	})
}
