package domain

import (
	"fmt"
	"time"

	"kloudlite.io/apps/console/internal/domain/entities"
	iamT "kloudlite.io/apps/iam/types"
	"kloudlite.io/grpc-interfaces/kloudlite.io/rpc/iam"
	"kloudlite.io/pkg/repos"
	t "kloudlite.io/pkg/types"
)

func (d *domain) findEnvironment(ctx ConsoleContext, name string) (*entities.Environment, error) {
	env, err := d.environmentRepo.FindOne(ctx, repos.Filter{
		"accountName":   ctx.AccountName,
		"clusterName":   ctx.ClusterName,
		"metadata.name": name,
	})
	if err != nil {
		return nil, err
	}
	if env == nil {
		return nil, fmt.Errorf("no environment with name=%q found", name)
	}
	return env, nil
}

// environment:query

func (d *domain) GetEnvironment(ctx ConsoleContext, name string) (*entities.Environment, error) {
	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceProject, name),
		},
		Action: string(iamT.GetEnvironment),
	})
	if err != nil {
		return nil, err
	}

	if !co.Status {
		return nil, fmt.Errorf("unauthorized to get environment")
	}
	return d.findEnvironment(ctx, name)
}

// ListEnvironments implements Domain
func (d *domain) ListEnvironments(ctx ConsoleContext) ([]*entities.Environment, error) {
	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceProject, ctx.AccountName),
		},
		Action: string(iamT.ListEnvironments),
	})
	if err != nil {
		return nil, err
	}

	if !co.Status {
		return nil, fmt.Errorf("unauthorized to list environments")
	}

	filter := repos.Filter{"accountName": ctx.AccountName}
	if ctx.ClusterName != "" {
		filter["clusterName"] = ctx.ClusterName
	}
	return d.environmentRepo.Find(ctx, repos.Query{Filter: filter})
}

// CreateEnvironment implements Domain
func (d *domain) CreateEnvironment(ctx ConsoleContext, env entities.Environment) (*entities.Environment, error) {
	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceProject, ctx.AccountName),
		},
		Action: string(iamT.CreateEnvironment),
	})
	if err != nil {
		return nil, err
	}

	if !co.Status {
		return nil, fmt.Errorf("unauthorized to create Project")
	}

	env.EnsureGVK()
	if err := d.k8sExtendedClient.ValidateStruct(ctx, &env.Env); err != nil {
		return nil, err
	}

	env.AccountName = ctx.AccountName
	env.ClusterName = ctx.ClusterName
	env.SyncStatus = t.GetSyncStatusForCreation()
	prj, err := d.environmentRepo.Create(ctx, &env)
	if err != nil {
		if d.projectRepo.ErrAlreadyExists(err) {
			return nil, fmt.Errorf("environment with name %q, already exists", env.Name)
		}
		return nil, err
	}

	if err := d.applyK8sResource(ctx, &env.Env); err != nil {
		return nil, err
	}

	return prj, nil
}

// UpdateEnvironment implements Domain
func (d *domain) UpdateEnvironment(ctx ConsoleContext, env entities.Environment) (*entities.Environment, error) {
	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceProject, env.Spec.ProjectName),
		},
		Action: string(iamT.UpdateEnvironment),
	})
	if err != nil {
		return nil, err
	}

	if !co.Status {
		return nil, fmt.Errorf("unauthorized to update environment %q", env.Name)
	}

	env.EnsureGVK()
	if err := d.k8sExtendedClient.ValidateStruct(ctx, &env.Env); err != nil {
		return nil, err
	}

	exEnv, err := d.findEnvironment(ctx, env.Name)
	if err != nil {
		return nil, err
	}

	if exEnv.GetDeletionTimestamp() != nil {
		return nil, errAlreadyMarkedForDeletion("environment", "", env.Name)
	}

	exEnv.Spec = env.Spec
	exEnv.SyncStatus = t.GetSyncStatusForUpdation(exEnv.Generation)

	upEnv, err := d.environmentRepo.UpdateById(ctx, exEnv.Id, exEnv)
	if err != nil {
		return nil, err
	}

	if err := d.applyK8sResource(ctx, &upEnv.Env); err != nil {
		return nil, err
	}

	return upEnv, nil
}

// DeleteEnvironment implements Domain
func (d *domain) DeleteEnvironment(ctx ConsoleContext, name string) error {
	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceProject, ctx.AccountName),
		},
		Action: string(iamT.DeleteEnvironment),
	})
	if err != nil {
		return err
	}

	if !co.Status {
		return fmt.Errorf("unauthorized to delete environment")
	}

	env, err := d.findEnvironment(ctx, name)
	if err != nil {
		return err
	}

	env.SyncStatus = t.GetSyncStatusForDeletion(env.Generation)
	if _, err := d.environmentRepo.UpdateById(ctx, env.Id, env); err != nil {
		return err
	}

	return d.deleteK8sResource(ctx, &env.Env)
}

func (d *domain) OnApplyEnvironmentError(ctx ConsoleContext, err error, name string) error {
	env, err2 := d.findEnvironment(ctx, name)
	if err2 != nil {
		return err2
	}

	env.SyncStatus.Error = err.Error()
	_, err = d.environmentRepo.UpdateById(ctx, env.Id, env)
	return err
}

func (d *domain) OnDeleteEnvironmentMessage(ctx ConsoleContext, env entities.Environment) error {
	p, err := d.findEnvironment(ctx, env.Name)
	if err != nil {
		return err
	}

	return d.environmentRepo.DeleteById(ctx, p.Id)
}

// OnUpdateEnvironmentMessage implements Domain
func (d *domain) OnUpdateEnvironmentMessage(ctx ConsoleContext, env entities.Environment) error {
	e, err := d.findEnvironment(ctx, env.Name)
	if err != nil {
		return err
	}

	e.Status = env.Status
	e.SyncStatus.LastSyncedAt = time.Now()
	e.SyncStatus.State = t.ParseSyncState(env.Status.IsReady)

	_, err = d.environmentRepo.UpdateById(ctx, e.Id, e)
	return err
}
