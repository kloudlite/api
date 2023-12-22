package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"github.com/kloudlite/api/pkg/errors"
	"time"

	"github.com/kloudlite/api/apps/container-registry/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/container-registry/internal/app/graph/model"
	"github.com/kloudlite/api/apps/container-registry/internal/domain/entities"
)

// Repositories is the resolver for the repositories field.
func (r *githubListRepositoryResolver) Repositories(ctx context.Context, obj *entities.GithubListRepository) ([]*model.GithubComKloudliteAPIAppsContainerRegistryInternalDomainEntitiesGithubRepository, error) {
	if obj == nil {
		return nil, errors.Newf("Repositories: obj is nil")
	}

	repositories := make([]*model.KloudliteIoAppsContainerRegistryInternalDomainEntitiesGithubRepository, len(obj.Repositories))

	for i, gr := range obj.Repositories {
		repositories[i] = &model.KloudliteIoAppsContainerRegistryInternalDomainEntitiesGithubRepository{
			Archived:          gr.Archived,
			CloneURL:          gr.CloneURL,
			CreatedAt:         getStringPtr(gr.CreatedAt.Format(time.RFC3339)),
			DefaultBranch:     gr.DefaultBranch,
			Description:       gr.Description,
			Disabled:          gr.Disabled,
			FullName:          gr.FullName,
			GitURL:            gr.GitURL,
			GitignoreTemplate: gr.GitignoreTemplate,
			HTMLURL:           gr.HTMLURL,
			ID:                getInt(gr.ID),
			Language:          gr.Language,
			MasterBranch:      gr.MasterBranch,
			MirrorURL:         gr.MirrorURL,
			Name:              gr.Name,
			NodeID:            gr.NodeID,
			Permissions: (func() map[string]any {
				m := make(map[string]any)
				for k, v := range gr.Permissions {
					m[k] = v
				}
				return m
			})(),
			Private:    gr.Private,
			PushedAt:   getStringPtr(gr.PushedAt.Format(time.RFC3339)),
			Size:       gr.Size,
			TeamID:     getInt(gr.TeamID),
			UpdatedAt:  getStringPtr(gr.UpdatedAt.Format(time.RFC3339)),
			Visibility: gr.Visibility,
			URL:        gr.URL,
		}
	}

	return repositories, nil
}

// GithubListRepository returns generated.GithubListRepositoryResolver implementation.
func (r *Resolver) GithubListRepository() generated.GithubListRepositoryResolver {
	return &githubListRepositoryResolver{r}
}

type githubListRepositoryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func getStringPtr(s string) *string {
	return &s
}
