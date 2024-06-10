package common

import (
	"fmt"
	"strconv"

	domainT "github.com/kloudlite/api/apps/infra/internal/domain/types"
	constant "github.com/kloudlite/api/constants"
	"github.com/kloudlite/api/pkg/errors"
	fn "github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/repos"
	types "github.com/kloudlite/api/pkg/types"
	"github.com/kloudlite/operator/pkg/constants"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (d *Domain) ResyncToTargetCluster(ctx domainT.InfraContext, action types.SyncAction, clusterName string, obj client.Object, recordVersion int) error {
	switch action {
	case types.SyncActionApply:
		return d.ResDispatcher.ApplyToTargetCluster(ctx, ctx.AccountName, clusterName, obj, recordVersion)
	case types.SyncActionDelete:
		return d.ResDispatcher.DeleteFromTargetCluster(ctx, ctx.AccountName, clusterName, obj)
	}
	return errors.Newf("unknown action: %q", action)
}

func (d *Domain) AddTrackingId(obj client.Object, id repos.ID) {
	ann := obj.GetAnnotations()
	if ann == nil {
		ann = make(map[string]string, 1)
	}
	ann[constant.ObservabilityTrackingKey] = string(id)

	labels := obj.GetLabels()
	if labels == nil {
		labels = make(map[string]string, 1)
	}
	labels[constant.ObservabilityTrackingKey] = string(id)
	obj.SetLabels(labels)
}

func (d *Domain) ApplyK8sResource(ctx domainT.InfraContext, obj client.Object, recordVersion int) error {
	if recordVersion > 0 {
		ann := obj.GetAnnotations()
		if ann == nil {
			ann = make(map[string]string, 1)
		}
		ann[constants.RecordVersionKey] = fmt.Sprintf("%d", recordVersion)
		obj.SetAnnotations(ann)
	}

	b, err := fn.K8sObjToYAML(obj)
	if err != nil {
		return errors.NewE(err)
	}

	if err := d.K8sClient.ApplyYAML(ctx, b); err != nil {
		return errors.NewE(err)
	}
	return nil
}

func (d *Domain) DeleteK8sResource(ctx domainT.InfraContext, obj client.Object) error {
	b, err := fn.K8sObjToYAML(obj)
	if err != nil {
		return errors.NewE(err)
	}

	if err := d.K8sClient.DeleteYAML(ctx, b); err != nil {
		return errors.NewE(err)
	}
	return nil
}

func (d *Domain) ParseRecordVersionFromAnnotations(annotations map[string]string) (int, error) {
	annotatedVersion, ok := annotations[constants.RecordVersionKey]
	if !ok {
		return 0, errors.Newf("no annotation with record version key (%s), found on the resource", constants.RecordVersionKey)
	}

	annVersion, err := strconv.ParseInt(annotatedVersion, 10, 32)
	if err != nil {
		return 0, errors.NewE(err)
	}

	return int(annVersion), nil
}

func (d *Domain) MatchRecordVersion(annotations map[string]string, rv int) (int, error) {
	annVersion, err := d.ParseRecordVersionFromAnnotations(annotations)
	if err != nil {
		return -1, errors.NewE(err)
	}

	if annVersion != rv {
		return -1, errors.Newf("record version mismatch, expected %d, got %d", rv, annVersion)
	}

	return annVersion, nil
}

func (d *Domain) GetAccNamespace(ctx domainT.InfraContext) (string, error) {
	acc, err := d.AccountsSvc.GetAccount(ctx, string(ctx.UserId), ctx.AccountName)
	if err != nil {
		return "", errors.NewE(err)
	}
	if !acc.IsActive {
		return "", errors.Newf("account %q is not active", ctx.AccountName)
	}

	return acc.TargetNamespace, nil
}
