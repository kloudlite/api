package graph

import (
	"context"

	"github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/pkg/errors"
)

func toInfraContext(ctx context.Context) (types.InfraContext, error) {
	if d, ok := ctx.Value("infra-ctx").(types.InfraContext); ok {
		return d, nil
	}
	return types.InfraContext{}, errors.Newf("infra context not found in gql context")
}
