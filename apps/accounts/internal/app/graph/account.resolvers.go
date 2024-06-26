package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"github.com/kloudlite/api/pkg/errors"
	"time"

	"github.com/kloudlite/api/apps/accounts/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/accounts/internal/entities"
	fn "github.com/kloudlite/api/pkg/functions"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreationTime is the resolver for the creationTime field.
func (r *accountResolver) CreationTime(ctx context.Context, obj *entities.Account) (string, error) {
	if obj == nil {
		return "", errors.Newf("account is nil")
	}
	return obj.BaseEntity.CreationTime.Format(time.RFC3339), nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *accountResolver) UpdateTime(ctx context.Context, obj *entities.Account) (string, error) {
	if obj == nil {
		return "", errors.Newf("resource is nil")
	}
	return obj.UpdateTime.Format(time.RFC3339), nil
}

// Metadata is the resolver for the metadata field.
func (r *accountInResolver) Metadata(ctx context.Context, obj *entities.Account, data *v1.ObjectMeta) error {
	if obj == nil {
		return errors.Newf("obj is nil")
	}
	return fn.JsonConversion(data, &obj.ObjectMeta)
}

// Account returns generated.AccountResolver implementation.
func (r *Resolver) Account() generated.AccountResolver { return &accountResolver{r} }

// AccountIn returns generated.AccountInResolver implementation.
func (r *Resolver) AccountIn() generated.AccountInResolver { return &accountInResolver{r} }

type accountResolver struct{ *Resolver }
type accountInResolver struct{ *Resolver }
