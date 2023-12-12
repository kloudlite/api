package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"

	generated1 "github.com/kloudlite/api/apps/container-registry/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/container-registry/internal/app/graph/model"
	"github.com/kloudlite/api/apps/container-registry/internal/domain"
	"github.com/kloudlite/api/apps/container-registry/internal/domain/entities"
	fn "github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/repos"
	"github.com/kloudlite/api/pkg/types"
)

// CrCreateRepo is the resolver for the cr_createRepo field.
func (r *mutationResolver) CrCreateRepo(ctx context.Context, repository entities.Repository) (*entities.Repository, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.CreateRepository(cc, repository.Name)
}

// CrCreateCred is the resolver for the cr_createCred field.
func (r *mutationResolver) CrCreateCred(ctx context.Context, credential entities.Credential) (*entities.Credential, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.CreateCredential(cc, credential)
}

// CrDeleteRepo is the resolver for the cr_deleteRepo field.
func (r *mutationResolver) CrDeleteRepo(ctx context.Context, name string) (bool, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return false, err
	}
	if err := r.Domain.DeleteRepository(cc, name); err != nil {
		return false, err
	}

	return true, nil
}

// CrDeleteCred is the resolver for the cr_deleteCred field.
func (r *mutationResolver) CrDeleteCred(ctx context.Context, username string) (bool, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return false, err
	}
	if err := r.Domain.DeleteCredential(cc, username); err != nil {
		return false, err
	}
	return true, nil
}

// CrDeleteDigest is the resolver for the cr_deleteDigest field.
func (r *mutationResolver) CrDeleteDigest(ctx context.Context, repoName string, digest string) (bool, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return false, err
	}
	if err := r.Domain.DeleteRepositoryDigest(cc, repoName, digest); err != nil {
		return false, err
	}
	return true, nil
}

// CrAddBuild is the resolver for the cr_addBuild field.
func (r *mutationResolver) CrAddBuild(ctx context.Context, build entities.Build) (*entities.Build, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.AddBuild(cc, build)
}

// CrUpdateBuild is the resolver for the cr_updateBuild field.
func (r *mutationResolver) CrUpdateBuild(ctx context.Context, id repos.ID, build entities.Build) (*entities.Build, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.UpdateBuild(cc, id, build)
}

// CrDeleteBuild is the resolver for the cr_deleteBuild field.
func (r *mutationResolver) CrDeleteBuild(ctx context.Context, id repos.ID) (bool, error) {
	cc, err := toRegistryContext(ctx)

	if err != nil {
		return false, err
	}

	if err := r.Domain.DeleteBuild(cc, id); err != nil {
		return false, err
	}
	return true, nil
}

// CrTriggerBuild is the resolver for the cr_triggerBuild field.
func (r *mutationResolver) CrTriggerBuild(ctx context.Context, id repos.ID) (bool, error) {
	cc, err := toRegistryContext(ctx)

	if err != nil {
		return false, err
	}

	if err := r.Domain.TriggerBuild(cc, id); err != nil {
		return false, err
	}
	return true, nil
}

// CrAddBuildCacheKey is the resolver for the cr_addBuildCacheKey field.
func (r *mutationResolver) CrAddBuildCacheKey(ctx context.Context, buildCacheKey entities.BuildCacheKey) (*entities.BuildCacheKey, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.AddBuildCache(cc, buildCacheKey)
}

// CrDeleteBuildCacheKey is the resolver for the cr_deleteBuildCacheKey field.
func (r *mutationResolver) CrDeleteBuildCacheKey(ctx context.Context, id repos.ID) (bool, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return false, err
	}

	if err := r.Domain.DeleteBuildCache(cc, id); err != nil {
		return false, err
	}
	return true, nil
}

// CrUpdateBuildCacheKey is the resolver for the cr_updateBuildCacheKey field.
func (r *mutationResolver) CrUpdateBuildCacheKey(ctx context.Context, id repos.ID, buildCacheKey entities.BuildCacheKey) (*entities.BuildCacheKey, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.UpdateBuildCache(cc, id, buildCacheKey)
}

// CrListBuildsByBuildCacheID is the resolver for the cr_listBuildsByBuildCacheId field.
func (r *mutationResolver) CrListBuildsByBuildCacheID(ctx context.Context, buildCacheKeyID repos.ID, pagination *repos.CursorPagination) (*model.BuildPaginatedRecords, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return nil, err
	}

	rr, err := r.Domain.ListBuildsByCache(cc, buildCacheKeyID, fn.DefaultIfNil(pagination, repos.DefaultCursorPagination))
	if err != nil {
		return nil, err
	}

	records := make([]*model.BuildEdge, len(rr.Edges))

	for i := range rr.Edges {
		records[i] = &model.BuildEdge{
			Node:   rr.Edges[i].Node,
			Cursor: rr.Edges[i].Cursor,
		}
	}

	m := &model.BuildPaginatedRecords{
		Edges: records,
		PageInfo: &model.PageInfo{
			HasNextPage:     rr.PageInfo.HasNextPage,
			HasPreviousPage: rr.PageInfo.HasPrevPage,
			StartCursor:     &rr.PageInfo.StartCursor,
			EndCursor:       &rr.PageInfo.EndCursor,
		},
	}

	return m, nil
}

// CrListRepos is the resolver for the cr_listRepos field.
func (r *queryResolver) CrListRepos(ctx context.Context, search *model.SearchRepos, pagination *repos.CursorPagination) (*model.RepositoryPaginatedRecords, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return nil, err
	}

	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["name"] = *search.Text
		}
	}

	rr, err := r.Domain.ListRepositories(cc, filter, fn.DefaultIfNil(pagination, repos.DefaultCursorPagination))
	if err != nil {
		return nil, err
	}

	records := make([]*model.RepositoryEdge, len(rr.Edges))

	for i := range rr.Edges {
		records[i] = &model.RepositoryEdge{
			Node:   rr.Edges[i].Node,
			Cursor: rr.Edges[i].Cursor,
		}
	}

	m := &model.RepositoryPaginatedRecords{
		Edges: records,
		PageInfo: &model.PageInfo{
			HasNextPage:     rr.PageInfo.HasNextPage,
			HasPreviousPage: rr.PageInfo.HasPrevPage,
			StartCursor:     &rr.PageInfo.StartCursor,
			EndCursor:       &rr.PageInfo.EndCursor,
		},
		TotalCount: int(rr.TotalCount),
	}

	return m, nil
}

// CrListCreds is the resolver for the cr_listCreds field.
func (r *queryResolver) CrListCreds(ctx context.Context, search *model.SearchCreds, pagination *repos.CursorPagination) (*model.CredentialPaginatedRecords, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return nil, err
	}

	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["name"] = *search.Text
		}
	}

	rr, err := r.Domain.ListCredentials(cc, filter, fn.DefaultIfNil(pagination, repos.DefaultCursorPagination))
	if err != nil {
		return nil, err
	}

	records := make([]*model.CredentialEdge, len(rr.Edges))

	for i := range rr.Edges {
		records[i] = &model.CredentialEdge{
			Node:   rr.Edges[i].Node,
			Cursor: rr.Edges[i].Cursor,
		}
	}

	m := &model.CredentialPaginatedRecords{
		Edges: records,
		PageInfo: &model.PageInfo{
			HasNextPage:     rr.PageInfo.HasNextPage,
			HasPreviousPage: rr.PageInfo.HasPrevPage,
			StartCursor:     &rr.PageInfo.StartCursor,
			EndCursor:       &rr.PageInfo.EndCursor,
		},
		TotalCount: int(rr.TotalCount),
	}

	return m, nil
}

// CrListDigests is the resolver for the cr_listDigests field.
func (r *queryResolver) CrListDigests(ctx context.Context, repoName string, search *model.SearchRepos, pagination *repos.CursorPagination) (*model.DigestPaginatedRecords, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return nil, err
	}

	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["name"] = *search.Text
		}
	}

	rr, err := r.Domain.ListRepositoryDigests(cc, repoName, filter, fn.DefaultIfNil(pagination, repos.DefaultCursorPagination))
	if err != nil {
		return nil, err
	}

	records := make([]*model.DigestEdge, len(rr.Edges))

	for i := range rr.Edges {
		records[i] = &model.DigestEdge{
			Node:   rr.Edges[i].Node,
			Cursor: rr.Edges[i].Cursor,
		}
	}
	m := &model.DigestPaginatedRecords{
		Edges: records,
		PageInfo: &model.PageInfo{
			HasNextPage:     rr.PageInfo.HasNextPage,
			HasPreviousPage: rr.PageInfo.HasPrevPage,
			StartCursor:     &rr.PageInfo.StartCursor,
			EndCursor:       &rr.PageInfo.EndCursor,
		},
		TotalCount: int(rr.TotalCount),
	}

	return m, nil
}

// CrGetCredToken is the resolver for the cr_getCredToken field.
func (r *queryResolver) CrGetCredToken(ctx context.Context, username string) (string, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return "", err
	}

	token, err := r.Domain.GetToken(cc, username)
	if err != nil {
		return "", err
	}

	return token, nil
}

// CrCheckUserNameAvailability is the resolver for the cr_checkUserNameAvailability field.
func (r *queryResolver) CrCheckUserNameAvailability(ctx context.Context, name string) (*domain.CheckNameAvailabilityOutput, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.CheckUserNameAvailability(cc, name)
}

// CrGetBuild is the resolver for the cr_getBuild field.
func (r *queryResolver) CrGetBuild(ctx context.Context, id repos.ID) (*entities.Build, error) {
	cc, err := toRegistryContext(ctx)

	if err != nil {
		return nil, err
	}

	return r.Domain.GetBuild(cc, id)
}

// CrListBuilds is the resolver for the cr_listBuilds field.
func (r *queryResolver) CrListBuilds(ctx context.Context, repoName string, search *model.SearchBuilds, pagination *repos.CursorPagination) (*model.BuildPaginatedRecords, error) {
	cc, err := toRegistryContext(ctx)

	if err != nil {
		return nil, err
	}

	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["name"] = *search.Text
		}
	}

	rr, err := r.Domain.ListBuilds(cc, repoName, filter, fn.DefaultIfNil(pagination, repos.DefaultCursorPagination))

	if err != nil {
		return nil, err
	}

	records := make([]*model.BuildEdge, len(rr.Edges))

	for i := range rr.Edges {
		records[i] = &model.BuildEdge{
			Node:   rr.Edges[i].Node,
			Cursor: rr.Edges[i].Cursor,
		}
	}

	m := &model.BuildPaginatedRecords{
		Edges: records,
		PageInfo: &model.PageInfo{
			HasNextPage:     rr.PageInfo.HasNextPage,
			HasPreviousPage: rr.PageInfo.HasPrevPage,
			StartCursor:     &rr.PageInfo.StartCursor,
			EndCursor:       &rr.PageInfo.EndCursor,
		},
		TotalCount: int(rr.TotalCount),
	}

	return m, nil
}

// CrListGithubInstallations is the resolver for the cr_listGithubInstallations field.
func (r *queryResolver) CrListGithubInstallations(ctx context.Context, pagination *types.Pagination) ([]*entities.GithubInstallation, error) {
	userId, err := getUserId(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.GithubListInstallations(ctx, userId, pagination)
}

// CrListGithubRepos is the resolver for the cr_listGithubRepos field.
func (r *queryResolver) CrListGithubRepos(ctx context.Context, installationID int, pagination *types.Pagination) (*entities.GithubListRepository, error) {
	userId, err := getUserId(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.GithubListRepos(ctx, userId, int64(installationID), pagination)
}

// CrSearchGithubRepos is the resolver for the cr_searchGithubRepos field.
func (r *queryResolver) CrSearchGithubRepos(ctx context.Context, organization string, search string, pagination *types.Pagination) (*entities.GithubSearchRepository, error) {
	userId, err := getUserId(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.GithubSearchRepos(ctx, userId, search, organization, pagination)
}

// CrListGithubBranches is the resolver for the cr_listGithubBranches field.
func (r *queryResolver) CrListGithubBranches(ctx context.Context, repoURL string, pagination *types.Pagination) ([]*entities.GitBranch, error) {
	userId, err := getUserId(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.GithubListBranches(ctx, userId, repoURL, pagination)
}

// CrListGitlabGroups is the resolver for the cr_listGitlabGroups field.
func (r *queryResolver) CrListGitlabGroups(ctx context.Context, query *string, pagination *types.Pagination) ([]*entities.GitlabGroup, error) {
	userId, err := getUserId(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.GitlabListGroups(ctx, userId, query, pagination)
}

// CrListGitlabRepositories is the resolver for the cr_listGitlabRepositories field.
func (r *queryResolver) CrListGitlabRepositories(ctx context.Context, groupID string, query *string, pagination *types.Pagination) ([]*entities.GitlabProject, error) {
	userId, err := getUserId(ctx)
	if err != nil {
		return nil, err
	}
	return r.Domain.GitlabListRepos(ctx, userId, groupID, query, pagination)
}

// CrListGitlabBranches is the resolver for the cr_listGitlabBranches field.
func (r *queryResolver) CrListGitlabBranches(ctx context.Context, repoID string, query *string, pagination *types.Pagination) ([]*entities.GitBranch, error) {
	userId, err := getUserId(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.GitlabListBranches(ctx, userId, repoID, query, pagination)
}

// CrListBuildCacheKeys is the resolver for the cr_listBuildCacheKeys field.
func (r *queryResolver) CrListBuildCacheKeys(ctx context.Context, pq *repos.CursorPagination, search *model.SearchBuildCacheKeys) (*model.BuildCacheKeyPaginatedRecords, error) {
	cc, err := toRegistryContext(ctx)
	if err != nil {
		return nil, err
	}

	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["name"] = *search.Text
		}
	}

	rr, err := r.Domain.ListBuildCaches(cc, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, err
	}

	records := make([]*model.BuildCacheKeyEdge, len(rr.Edges))

	for i := range rr.Edges {
		records[i] = &model.BuildCacheKeyEdge{
			Node:   rr.Edges[i].Node,
			Cursor: rr.Edges[i].Cursor,
		}
	}

	m := &model.BuildCacheKeyPaginatedRecords{
		Edges: records,
		PageInfo: &model.PageInfo{
			HasNextPage:     rr.PageInfo.HasNextPage,
			HasPreviousPage: rr.PageInfo.HasPrevPage,
			StartCursor:     &rr.PageInfo.StartCursor,
			EndCursor:       &rr.PageInfo.EndCursor,
		},
	}

	return m, nil
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
