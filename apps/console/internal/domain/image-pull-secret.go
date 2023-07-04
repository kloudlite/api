package domain

import (
	"fmt"
	"time"

	"kloudlite.io/apps/console/internal/domain/entities"
	"kloudlite.io/pkg/repos"
	t "kloudlite.io/pkg/types"
)

func (d *domain) ListImagePullSecrets(ctx ConsoleContext, namespace string, pagination t.CursorPagination) (*repos.PaginatedRecord[*entities.ImagePullSecret], error) {
	if err := d.canReadSecretsFromAccount(ctx, string(ctx.UserId), ctx.AccountName); err != nil {
		return nil, err
	}

	return d.ipsRepo.FindPaginated(ctx, repos.Filter{
		"accountName": ctx.AccountName,
		"clusterName": ctx.ClusterName,
		"namespace":   namespace,
	}, pagination)
}

func (d *domain) findImagePullSecret(ctx ConsoleContext, namespace, name string) (*entities.ImagePullSecret, error) {
	ips, err := d.ipsRepo.FindOne(ctx, repos.Filter{"accountName": ctx.AccountName, "name": name, "namespace": namespace})
	if err != nil {
		return nil, err
	}

	if ips == nil {
		return nil, fmt.Errorf("no image-pull-secret with name=%q, namespace=%q found", name, namespace)
	}
	return ips, nil
}

func (d *domain) GetImagePullSecret(ctx ConsoleContext, namespace, name string) (*entities.ImagePullSecret, error) {
	if err := d.canReadSecretsFromAccount(ctx, string(ctx.UserId), ctx.AccountName); err != nil {
		return nil, err
	}

	return d.findImagePullSecret(ctx, namespace, name)
}

func (d *domain) CreateImagePullSecret(ctx ConsoleContext, ips entities.ImagePullSecret) (*entities.ImagePullSecret, error) {
	if err := d.canMutateSecretsInAccount(ctx, string(ctx.UserId), ctx.AccountName); err != nil {
		return nil, err
	}

	ips.EnsureGVK()
	if err := d.k8sExtendedClient.ValidateStruct(ctx, &ips.ImagePullSecret); err != nil {
		return nil, err
	}

	ips.AccountName = ctx.AccountName
	ips.ClusterName = ctx.ClusterName
	ips.Generation = 1
	ips.SyncStatus = t.GenSyncStatus(t.SyncActionApply, ips.Generation)

	nIps, err := d.ipsRepo.Create(ctx, &ips)
	if err != nil {
		if d.ipsRepo.ErrAlreadyExists(err) {
			// TODO: better insights into error, when it is being caused by duplicated indexes
			return nil, err
		}
		return nil, err
	}

	return nIps, nil
}

func (d *domain) UpdateImagePullSecret(ctx ConsoleContext, ips entities.ImagePullSecret) (*entities.ImagePullSecret, error) {
	if err := d.canMutateSecretsInAccount(ctx, string(ctx.UserId), ctx.AccountName); err != nil {
		return nil, err
	}

	ips.EnsureGVK()

	exIps, err := d.findImagePullSecret(ctx, ips.Namespace, ips.Name)
	if err != nil {
		return nil, err
	}

	exIps.Spec.DockerConfigJson = ips.Spec.DockerConfigJson
	exIps.Spec.DockerUsername = ips.Spec.DockerUsername
	exIps.Spec.DockerPassword = ips.Spec.DockerPassword
	exIps.Spec.DockerUsername = ips.Spec.DockerUsername

	upIps, err := d.ipsRepo.UpdateById(ctx, exIps.Id, exIps)
	if err != nil {
		return nil, err
	}

	if err := d.applyK8sResource(ctx, &upIps.ImagePullSecret); err != nil {
		return nil, err
	}

	return upIps, err
}

func (d *domain) DeleteImagePullSecret(ctx ConsoleContext, namespace, name string) error {
	if err := d.canMutateSecretsInAccount(ctx, string(ctx.UserId), ctx.AccountName); err != nil {
		return err
	}

	ips, err := d.findImagePullSecret(ctx, namespace, name)
	if err != nil {
		return err
	}

	ips.SyncStatus = t.GenSyncStatus(t.SyncActionDelete, ips.Generation)

	if _, err := d.ipsRepo.UpdateById(ctx, ips.Id, ips); err != nil {
		return err
	}

	return d.deleteK8sResource(ctx, &ips.ImagePullSecret)
}

func (d *domain) OnUpdateImagePullSecretMessage(ctx ConsoleContext, ips entities.ImagePullSecret) error {
	a, err := d.findImagePullSecret(ctx, ips.Namespace, ips.Name)
	if err != nil {
		return err
	}

	a.Status = ips.Status
	a.SyncStatus.Error = nil
	a.SyncStatus.LastSyncedAt = time.Now()
	a.SyncStatus.Generation = ips.Generation
	a.SyncStatus.State = t.ParseSyncState(ips.Status.IsReady)

	_, err = d.ipsRepo.UpdateById(ctx, a.Id, a)
	return err
}

func (d *domain) OnDeleteImagePullSecretMessage(ctx ConsoleContext, ips entities.ImagePullSecret) error {
	a, err := d.findImagePullSecret(ctx, ips.Namespace, ips.Name)
	if err != nil {
		return err
	}

	return d.ipsRepo.DeleteById(ctx, a.Id)
}

func (d *domain) OnApplyImagePullSecretError(ctx ConsoleContext, errMsg string, namespace string, name string) error {
	a, err2 := d.findImagePullSecret(ctx, namespace, name)
	if err2 != nil {
		return err2
	}

	a.SyncStatus.Error = &errMsg
	_, err := d.ipsRepo.UpdateById(ctx, a.Id, a)
	return err
}

func (d *domain) ResyncImagePullSecret(ctx ConsoleContext, namespace, name string) error {
	if err := d.canMutateResourcesInWorkspace(ctx, namespace); err != nil {
		return err
	}

	a, err := d.findImagePullSecret(ctx, namespace, name)
	if err != nil {
		return err
	}
	return d.resyncK8sResource(ctx, a.SyncStatus.Action, &a.ImagePullSecret)
}
