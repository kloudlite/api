package domain

import (
	"context"
	"fmt"

	"kloudlite.io/apps/iam/internal/entities"
	t "kloudlite.io/apps/iam/types"
	"kloudlite.io/pkg/errors"
	"kloudlite.io/pkg/repos"
)

type Domain interface {
	AddRoleBinding(ctx context.Context, rb entities.RoleBinding) (*entities.RoleBinding, error)
	RemoveRoleBinding(ctx context.Context, userId repos.ID, resourceRef string) error
	RemoveRoleBindingsForResource(ctx context.Context, resourceRef string) error
	UpdateRoleBinding(ctx context.Context, rb entities.RoleBinding) (*entities.RoleBinding, error)

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
	nrb, err := d.rbRepo.Create(ctx, &rb)
	if err != nil {
		return nil, err
	}
	return nrb, nil
}

func (s domain) findRoleBinding(ctx context.Context, userId repos.ID, resourceRef string) (*entities.RoleBinding, error) {
	rb, err := s.rbRepo.FindOne(
		ctx, repos.Filter{
			"user_id":      userId,
			"resource_ref": resourceRef,
		},
	)
	if err != nil {
		return nil, err
	}

	if rb == nil {
		return nil, fmt.Errorf("role binding for (userId=%s,  ResourceRef=%s) not found", userId, resourceRef)
	}

	return rb, nil
}

func (d domain) RemoveRoleBinding(ctx context.Context, userId repos.ID, resourceRef string) error {
	rb, err := d.findRoleBinding(ctx, userId, resourceRef)
	if err != nil {
		return err
	}

	if err := d.rbRepo.DeleteById(ctx, rb.Id); err != nil {
		return errors.NewEf(err, "could not delete resource for (userId=%s, resourceRef=%s)", userId, resourceRef)
	}

	return nil
}

func (d domain) RemoveRoleBindingsForResource(ctx context.Context, resourceRef string) error {
	if err := d.rbRepo.DeleteMany(ctx, repos.Filter{"resource_ref": resourceRef}); err != nil {
		return errors.NewEf(err, "could not delete role bindings for (resourceRef=%s)", resourceRef)
	}
	return nil
}

// UpdateMembership updates only the role for a user on an already specified resource_ref
func (d domain) UpdateRoleBinding(ctx context.Context, rb entities.RoleBinding) (*entities.RoleBinding, error) {
	currRb, err := d.rbRepo.FindOne(
		ctx, repos.Filter{
			"user_id":       rb.UserId,
			"resource_ref":  rb.ResourceRef,
			"resource_type": rb.ResourceType,
		},
	)
	if err != nil {
		return nil, err
	}
	if currRb == nil {
		return nil, fmt.Errorf("role binding for (userId=%q,  ResourceRef=%q, ResourceType=%q) not found", rb.UserId, rb.ResourceRef, rb.ResourceType)
	}

	currRb.Role = rb.Role
	return d.rbRepo.UpdateById(ctx, currRb.Id, currRb)
}

func (d domain) GetRoleBinding(ctx context.Context, userId repos.ID, resourceRef string) (*entities.RoleBinding, error) {
	return d.findRoleBinding(ctx, userId, resourceRef)
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
