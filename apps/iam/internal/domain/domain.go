package domain

import (
	"context"
	"kloudlite.io/apps/iam/internal/entities"
	t "kloudlite.io/apps/iam/types"
	"kloudlite.io/pkg/repos"
)

type Domain interface {
	AddRoleBinding(ctx context.Context, rb entities.RoleBinding) (*entities.RoleBinding, error)
	RemoveRoleBinding(ctx context.Context, userId repos.ID, resourceRef string) error
	UpdateRoleBinding(ctx context.Context, userId repos.ID, resourceRef string, role t.Role) (*entities.RoleBinding, error)

	GetRoleBinding(ctx context.Context, userId repos.ID, resourceRef string) (*entities.RoleBinding, error)

	ListRoleBindingsForResource(ctx context.Context, resourceRef string) ([]*entities.RoleBinding, error)
	ListRoleBindingsForUser(ctx context.Context, userId repos.ID) ([]*entities.RoleBinding, error)

	Can(ctx context.Context, userId repos.ID, resourceRef string, action t.Action) (bool, error)
}

type domain struct {
	rbRepo         repos.DbRepo[*entities.RoleBinding]
	roleBindingMap map[t.Action][]t.Role
}

func (d domain) AddRoleBinding(ctx context.Context, rb entities.RoleBinding) (*entities.RoleBinding, error) {
}

func (d domain) RemoveRoleBinding(ctx context.Context, userId repos.ID, resourceRef string) error {
	//TODO implement me
	panic("implement me")
}

func (d domain) UpdateRoleBinding(ctx context.Context, userId repos.ID, resourceRef string, role t.Role) (*entities.RoleBinding, error) {
	//TODO implement me
	panic("implement me")
}

func (d domain) GetRoleBinding(ctx context.Context, userId repos.ID, resourceRef string) (*entities.RoleBinding, error) {
	//TODO implement me
	panic("implement me")
}

func (d domain) ListRoleBindingsForResource(ctx context.Context, resourceRef string) ([]*entities.RoleBinding, error) {
	//TODO implement me
	panic("implement me")
}

func (d domain) ListRoleBindingsForUser(ctx context.Context, userId repos.ID) ([]*entities.RoleBinding, error) {
	//TODO implement me
	panic("implement me")
}

func (d domain) Can(ctx context.Context, userId repos.ID, resourceRef string, action t.Action) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewDomain(rbRepo repos.DbRepo[*entities.RoleBinding], roleBindingMap map[t.Action][]t.Role) Domain {
	return &domain{
		rbRepo:         rbRepo,
		roleBindingMap: roleBindingMap,
	}
}
