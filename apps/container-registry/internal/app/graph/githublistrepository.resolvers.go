package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"github.com/kloudlite/api/pkg/errors"
	"time"

	"github.com/kloudlite/api/apps/container-registry/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/container-registry/internal/app/graph/model"
	"github.com/kloudlite/api/apps/container-registry/internal/domain/entities"
	fn "github.com/kloudlite/api/pkg/functions"
)

// Repositories is the resolver for the repositories field.
func (r *githubListRepositoryResolver) Repositories(ctx context.Context, obj *entities.GithubListRepository) ([]*model.GithubComKloudliteAPIAppsContainerRegistryInternalDomainEntitiesGithubRepository, error) {
	if obj == nil {
		return nil, errors.Newf("Repositories: obj is nil")
	}

	repositories := make([]*model.GithubComKloudliteAPIAppsContainerRegistryInternalDomainEntitiesGithubRepository, len(obj.Repositories))

	for i, gr := range obj.Repositories {
		repositories[i] = &model.GithubComKloudliteAPIAppsContainerRegistryInternalDomainEntitiesGithubRepository{
			Archived:          gr.Archived,
			CloneURL:          gr.CloneURL,
			CreatedAt:         fn.New(gr.CreatedAt.Format(time.RFC3339)),
			DefaultBranch:     gr.DefaultBranch,
			Description:       gr.Description,
			Disabled:          gr.Disabled,
			FullName:          gr.FullName,
			GitURL:            gr.GitURL,
			GitignoreTemplate: gr.GitignoreTemplate,
			HTMLURL:           gr.HTMLURL,
			ID:                fn.New(int(fn.DefaultIfNil(gr.ID))),
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
			PushedAt:   fn.New(gr.PushedAt.Format(time.RFC3339)),
			Size:       gr.Size,
			TeamID:     fn.New(int(fn.DefaultIfNil(gr.TeamID))),
			UpdatedAt:  fn.New(gr.UpdatedAt.Format(time.RFC3339)),
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
