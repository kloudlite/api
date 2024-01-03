package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"fmt"

	"github.com/kloudlite/api/apps/infra/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/infra/internal/app/graph/model"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreationTime is the resolver for the creationTime field.
func (r *volumeAttachmentResolver) CreationTime(ctx context.Context, obj *entities.VolumeAttachment) (string, error) {
	panic(fmt.Errorf("not implemented: CreationTime - creationTime"))
}

// ID is the resolver for the id field.
func (r *volumeAttachmentResolver) ID(ctx context.Context, obj *entities.VolumeAttachment) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// Spec is the resolver for the spec field.
func (r *volumeAttachmentResolver) Spec(ctx context.Context, obj *entities.VolumeAttachment) (*model.K8sIoAPIStorageV1VolumeAttachmentSpec, error) {
	panic(fmt.Errorf("not implemented: Spec - spec"))
}

// Status is the resolver for the status field.
func (r *volumeAttachmentResolver) Status(ctx context.Context, obj *entities.VolumeAttachment) (*model.K8sIoAPIStorageV1VolumeAttachmentStatus, error) {
	panic(fmt.Errorf("not implemented: Status - status"))
}

// UpdateTime is the resolver for the updateTime field.
func (r *volumeAttachmentResolver) UpdateTime(ctx context.Context, obj *entities.VolumeAttachment) (string, error) {
	panic(fmt.Errorf("not implemented: UpdateTime - updateTime"))
}

// Metadata is the resolver for the metadata field.
func (r *volumeAttachmentInResolver) Metadata(ctx context.Context, obj *entities.VolumeAttachment, data *v1.ObjectMeta) error {
	panic(fmt.Errorf("not implemented: Metadata - metadata"))
}

// Spec is the resolver for the spec field.
func (r *volumeAttachmentInResolver) Spec(ctx context.Context, obj *entities.VolumeAttachment, data *model.K8sIoAPIStorageV1VolumeAttachmentSpecIn) error {
	panic(fmt.Errorf("not implemented: Spec - spec"))
}

// Status is the resolver for the status field.
func (r *volumeAttachmentInResolver) Status(ctx context.Context, obj *entities.VolumeAttachment, data *model.K8sIoAPIStorageV1VolumeAttachmentStatusIn) error {
	panic(fmt.Errorf("not implemented: Status - status"))
}

// VolumeAttachment returns generated.VolumeAttachmentResolver implementation.
func (r *Resolver) VolumeAttachment() generated.VolumeAttachmentResolver {
	return &volumeAttachmentResolver{r}
}

// VolumeAttachmentIn returns generated.VolumeAttachmentInResolver implementation.
func (r *Resolver) VolumeAttachmentIn() generated.VolumeAttachmentInResolver {
	return &volumeAttachmentInResolver{r}
}

type volumeAttachmentResolver struct{ *Resolver }
type volumeAttachmentInResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *volumeAttachmentResolver) BaseEntity(ctx context.Context, obj *entities.VolumeAttachment) (*model.GithubComKloudliteAPIPkgReposBaseEntity, error) {
	panic(fmt.Errorf("not implemented: BaseEntity - BaseEntity"))
}
func (r *volumeAttachmentInResolver) BaseEntity(ctx context.Context, obj *entities.VolumeAttachment, data *model.GithubComKloudliteAPIPkgReposBaseEntityIn) error {
	panic(fmt.Errorf("not implemented: BaseEntity - BaseEntity"))
}
func (r *volumeAttachmentResolver) objectMeta(ctx context.Context, obj *model.VolumeAttachment) (*v1.ObjectMeta, error) {
	panic(fmt.Errorf("not implemented: objectMeta - metadata"))
}