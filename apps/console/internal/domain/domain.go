package domain

import (
	"encoding/json"
	"fmt"

	t "github.com/kloudlite/operator/agent/types"
	"github.com/kloudlite/operator/pkg/kubectl"
	"go.uber.org/fx"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"kloudlite.io/apps/console/internal/domain/entities"
	"kloudlite.io/apps/console/internal/env"
	iamT "kloudlite.io/apps/iam/types"
	"kloudlite.io/common"
	"kloudlite.io/grpc-interfaces/kloudlite.io/rpc/iam"
	fn "kloudlite.io/pkg/functions"
	"kloudlite.io/pkg/k8s"
	"kloudlite.io/pkg/redpanda"
	"kloudlite.io/pkg/repos"
	types "kloudlite.io/pkg/types"
)

type domain struct {
	k8sExtendedClient k8s.ExtendedK8sClient
	k8sYamlClient     *kubectl.YAMLClient

	producer redpanda.Producer

	iamClient iam.IAMClient

	projectRepo   repos.DbRepo[*entities.Project]
	workspaceRepo repos.DbRepo[*entities.Workspace]
	appRepo       repos.DbRepo[*entities.App]
	configRepo    repos.DbRepo[*entities.Config]
	secretRepo    repos.DbRepo[*entities.Secret]
	routerRepo    repos.DbRepo[*entities.Router]
	msvcRepo      repos.DbRepo[*entities.MSvc]
	mresRepo      repos.DbRepo[*entities.MRes]

	envVars *env.Env
}

func errAlreadyMarkedForDeletion(label, namespace, name string) error {
	return fmt.Errorf(
		"%s (namespace=%s, name=%s) already marked for deletion",
		label,
		namespace,
		name,
	)
}

func (d *domain) applyK8sResource(ctx ConsoleContext, obj client.Object) error {
	m, err := fn.K8sObjToMap(obj)
	if err != nil {
		return err
	}
	b, err := json.Marshal(t.AgentMessage{
		AccountName: ctx.AccountName,
		ClusterName: ctx.ClusterName,
		Action:      t.ActionApply,
		Object:      m,
	})
	if err != nil {
		return err
	}

	_, err = d.producer.Produce(
		ctx,
		common.GetKafkaTopicName(ctx.AccountName, ctx.ClusterName),
		obj.GetNamespace(),
		b,
	)
	return err
}

func (d *domain) deleteK8sResource(ctx ConsoleContext, obj client.Object) error {
	m, err := fn.K8sObjToMap(obj)
	if err != nil {
		return err
	}
	b, err := json.Marshal(t.AgentMessage{
		AccountName: ctx.AccountName,
		ClusterName: ctx.ClusterName,
		Action:      t.ActionDelete,
		Object:      m,
	})
	if err != nil {
		return err
	}
	_, err = d.producer.Produce(
		ctx,
		common.GetKafkaTopicName(ctx.AccountName, ctx.ClusterName),
		obj.GetNamespace(),
		b,
	)
	return err
}

func (d *domain) resyncK8sResource(ctx ConsoleContext, action types.SyncAction, obj client.Object) error {
	switch action {
	case types.SyncActionApply:
		{
			return d.applyK8sResource(ctx, obj)
		}
	case types.SyncActionDelete:
		{
			return d.deleteK8sResource(ctx, obj)
		}
	default:
		{
			return fmt.Errorf("unknown sync action %q", action)
		}
	}
}

func (d *domain) canMutateResourcesInProject(ctx ConsoleContext, targetNamespace string) error {
	prj, err := d.findProjectByTargetNs(ctx, targetNamespace)
	if err != nil {
		return err
	}

	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceProject, prj.Name),
		},
		Action: string(iamT.MutateResourcesInProject),
	})
	if err != nil {
		return err
	}
	if !co.Status {
		return fmt.Errorf("unauthorized to mutate resources in project %q", prj.Name)
	}
	return nil
}

func (d *domain) canMutateResourcesInWorkspace(ctx ConsoleContext, targetNamespace string) error {
	ws, err := d.findWorkspaceByTargetNs(ctx, targetNamespace)
	if err != nil {
		return err
	}

	wsp, err := d.findWorkspace(ctx, ws.Namespace, ws.Name)
	if err != nil {
		return err
	}

	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceProject, wsp.Spec.ProjectName),
		},
		Action: string(iamT.MutateResourcesInProject),
	})
	if err != nil {
		return err
	}
	if !co.Status {
		return fmt.Errorf("unauthorized to mutate resources in workspace %q", wsp.Name)
	}
	return nil
}

func (d *domain) canReadResourcesInWorkspace(ctx ConsoleContext, targetNamespace string) error {
	ws, err := d.findWorkspaceByTargetNs(ctx, targetNamespace)
	if err != nil {
		return err
	}

	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceProject, ws.Spec.ProjectName),
		},
		Action: string(iamT.GetProject),
	})
	if err != nil {
		return err
	}
	if !co.Status {
		return fmt.Errorf("unauthorized to read resources in project %q", ws.Spec.ProjectName)
	}
	return nil
}

func (d *domain) canReadResourcesInProject(ctx ConsoleContext, targetNamespace string) error {
	prj, err := d.findProjectByTargetNs(ctx, targetNamespace)
	if err != nil {
		return err
	}

	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceProject, prj.Name),
		},
		Action: string(iamT.GetProject),
	})
	if err != nil {
		return err
	}
	if !co.Status {
		return fmt.Errorf("unauthorized to read resources in project %q", prj.Name)
	}
	return nil
}

var Module = fx.Module("domain",
	fx.Provide(func(
		k8sYamlClient *kubectl.YAMLClient,
		k8sExtendedClient k8s.ExtendedK8sClient,

		producer redpanda.Producer,

		iamClient iam.IAMClient,

		projectRepo repos.DbRepo[*entities.Project],
		environmentRepo repos.DbRepo[*entities.Workspace],
		appRepo repos.DbRepo[*entities.App],
		configRepo repos.DbRepo[*entities.Config],
		secretRepo repos.DbRepo[*entities.Secret],
		routerRepo repos.DbRepo[*entities.Router],
		msvcRepo repos.DbRepo[*entities.MSvc],
		mresRepo repos.DbRepo[*entities.MRes],

		ev *env.Env,
	) Domain {
		return &domain{
			k8sExtendedClient: k8sExtendedClient,
			k8sYamlClient:     k8sYamlClient,

			producer: producer,

			iamClient: iamClient,

			projectRepo:   projectRepo,
			workspaceRepo: environmentRepo,
			appRepo:       appRepo,
			configRepo:    configRepo,
			routerRepo:    routerRepo,
			secretRepo:    secretRepo,
			msvcRepo:      msvcRepo,
			mresRepo:      mresRepo,

			envVars: ev,
		}
	}),
)
