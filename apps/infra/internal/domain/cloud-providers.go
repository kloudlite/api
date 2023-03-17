package domain

import (
	"fmt"
	"time"

	"kloudlite.io/apps/infra/internal/domain/entities"
	"kloudlite.io/pkg/repos"
	t "kloudlite.io/pkg/types"
)

func (d *domain) CreateProviderSecret(ctx InfraContext, ps *entities.Secret) (*entities.Secret, error) {
	d.secretRepo.SilentUpsert(ctx, repos.Filter{"metadata.name": ps.Name, "metadata.namespace": ps.Namespace}, ps)
	return nil, nil
}

func (d *domain) GetProviderSecret(ctx InfraContext, providerName string) (*entities.Secret, error) {
	cp, err := d.findCloudProvider(ctx, providerName)
	if err != nil {
		return nil, err
	}
	return d.findProviderSecret(ctx, cp.Spec.ProviderSecret.Name)
}

func (d *domain) CreateCloudProvider(ctx InfraContext, cloudProvider entities.CloudProvider, providerSecret entities.Secret) (*entities.CloudProvider, error) {
	cloudProvider.EnsureGVK()
	providerSecret.EnsureGVK()

	providerSecret.Namespace = d.env.ProviderSecretNamespace

	if err := d.k8sExtendedClient.ValidateStruct(ctx, &providerSecret.Secret); err != nil {
		return nil, err
	}

	cloudProvider.Spec.ProviderSecret.Name = providerSecret.Name
	cloudProvider.Spec.ProviderSecret.Namespace = providerSecret.Namespace

	if err := d.k8sExtendedClient.ValidateStruct(ctx, &cloudProvider.CloudProvider); err != nil {
		return nil, err
	}

	cloudProvider.AccountName = ctx.AccountName
	cloudProvider.SyncStatus = t.SyncStatus{
		LastSyncedAt: time.Now(),
		Action:       t.SyncActionApply,
		Generation:   int(cloudProvider.Generation) + 1,
		State:        t.SyncStateIdle,
	}

	cp, err := d.providerRepo.Create(ctx, &cloudProvider)
	if err != nil {
		if d.providerRepo.ErrAlreadyExists(err) {
			return nil, fmt.Errorf("cloud provider with name %q, already exists", cloudProvider.Name)
		}
		return nil, err
	}

	if err := d.applyK8sResource(ctx, &providerSecret.Secret); err != nil {
		return nil, err
	}

	if err := d.applyK8sResource(ctx, &cp.CloudProvider); err != nil {
		return nil, err
	}

	return cp, nil
}

func (d *domain) ListCloudProviders(ctx InfraContext) ([]*entities.CloudProvider, error) {
	return d.providerRepo.Find(ctx, repos.Query{Filter: repos.Filter{
		"accountName": ctx.AccountName,
	}})
}

func (d *domain) GetCloudProvider(ctx InfraContext, name string) (*entities.CloudProvider, error) {
	return d.findCloudProvider(ctx, name)
}

func (d *domain) findCloudProvider(ctx InfraContext, name string) (*entities.CloudProvider, error) {
	cp, err := d.providerRepo.FindOne(ctx, repos.Filter{
		"accountName":   ctx.AccountName,
		"metadata.name": name,
	})
	if err != nil {
		return nil, err
	}

	if cp == nil {
		return nil, fmt.Errorf("cloud provider with name %q not found", name)
	}

	return cp, nil
}

func (d *domain) UpdateCloudProvider(ctx InfraContext, cloudProvider entities.CloudProvider, providerSecret *entities.Secret) (*entities.CloudProvider, error) {
	cloudProvider.EnsureGVK()
	providerSecret.EnsureGVK()

	if err := d.k8sExtendedClient.ValidateStruct(ctx, &cloudProvider.CloudProvider); err != nil {
		return nil, err
	}

	cp, err := d.findCloudProvider(ctx, cloudProvider.Name)
	if err != nil {
		return nil, err
	}

	if providerSecret != nil {
		if err := d.k8sExtendedClient.ValidateStruct(ctx, &providerSecret.Secret); err != nil {
			return nil, err
		}

		if providerSecret.Name != cp.Spec.ProviderSecret.Name {
			return nil, fmt.Errorf("secret name mismatch b/w provider (%s) and provider secret(%s)", cp.Spec.ProviderSecret.Name, providerSecret.Name)
		}
	}

	cloudProvider.SyncStatus = t.SyncStatus{
		LastSyncedAt: time.Now(),
		Action:       t.SyncActionApply,
		Generation:   int(providerSecret.Generation) + 1,
		State:        t.SyncStateIdle,
	}

	uProvider, err := d.providerRepo.UpdateById(ctx, cp.Id, &cloudProvider)
	if err != nil {
		return nil, err
	}

	if err := d.applyK8sResource(ctx, &uProvider.CloudProvider); err != nil {
		return nil, err
	}

	upSecret, err := d.CreateProviderSecret(ctx, providerSecret)
	if err != nil {
		return nil, err
	}

	if err := d.applyK8sResource(ctx, &upSecret.Secret); err != nil {
		return nil, err
	}

	return nil, err
}

func (d *domain) DeleteCloudProvider(ctx InfraContext, name string) error {
	cp, err := d.findCloudProvider(ctx, name)
	if err != nil {
		return err
	}

	cp.SyncStatus = t.SyncStatus{
		SyncScheduledAt: time.Now(),
		Action:          t.SyncActionDelete,
		Generation:      int(cp.Generation),
		State:           t.SyncStateIdle,
	}

	uCp, err := d.providerRepo.UpdateById(ctx, cp.Id, cp)
	if err != nil {
		return err
	}

	return d.deleteK8sResource(ctx, &uCp.CloudProvider)
}

func (d *domain) OnDeleteCloudProviderMessage(ctx InfraContext, cloudProvider entities.CloudProvider) error {
	return d.providerRepo.DeleteOne(ctx, repos.Filter{
		"accountName":   ctx.AccountName,
		"metadata.name": cloudProvider.Name,
	})
}

func (d *domain) OnUpdateCloudProviderMessage(ctx InfraContext, cloudProvider entities.CloudProvider) error {
	cp, err := d.findCloudProvider(ctx, cloudProvider.Name)
	if err != nil {
		return err
	}

	cp.CloudProvider = cloudProvider.CloudProvider
	cp.SyncStatus.LastSyncedAt = time.Now()
	cp.SyncStatus.State = func() t.SyncState {
		if cloudProvider.Status.IsReady {
			return t.SyncStateReady
		}
		return t.SyncStateNotReady
	}()
	_, err = d.providerRepo.UpdateById(ctx, cp.Id, cp)
	return err
}

func (d *domain) findProviderSecret(ctx InfraContext, name string) (*entities.Secret, error) {
	ps, err := d.secretRepo.FindOne(ctx, repos.Filter{
		"accountName":        ctx.AccountName,
		"metadata.name":      name,
		"metadata.namespace": d.env.ProviderSecretNamespace,
	})

	if err != nil {
		return nil, err
	}

	if ps == nil {
		return nil, fmt.Errorf("provider secret with name %q not found", name)
	}

	return ps, nil
}
