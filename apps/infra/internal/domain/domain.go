package domain

import (
	"github.com/kloudlite/operator/pkg/kubectl"
	"go.uber.org/fx"
	"kloudlite.io/apps/infra/internal/domain/entities"
	"kloudlite.io/apps/infra/internal/env"
	"kloudlite.io/grpc-interfaces/kloudlite.io/rpc/finance"
	"kloudlite.io/pkg/agent"
	fn "kloudlite.io/pkg/functions"
	"kloudlite.io/pkg/k8s"
	"kloudlite.io/pkg/repos"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type domain struct {
	env *env.Env

	clusterRepo    repos.DbRepo[*entities.Cluster]
	edgeRepo       repos.DbRepo[*entities.Edge]
	providerRepo   repos.DbRepo[*entities.CloudProvider]
	k8sClient      client.Client
	masterNodeRepo repos.DbRepo[*entities.MasterNode]
	workerNodeRepo repos.DbRepo[*entities.WorkerNode]
	nodePoolRepo   repos.DbRepo[*entities.NodePool]
	secretRepo     repos.DbRepo[*entities.Secret]

	agentSender agent.Sender

	k8sYamlClient     *kubectl.YAMLClient
	k8sExtendedClient k8s.ExtendedK8sClient
}

func (d *domain) dispatchToTargetAgent(ctx InfraContext, action agent.Action, clusterName string, obj client.Object) error {
	b, err := fn.K8sObjToYAML(obj)
	if err != nil {
		return err
	}
	_, err = d.agentSender.Dispatch(ctx, clusterName+"-incoming", obj.GetNamespace(), agent.Message{
		Action: agent.Apply,
		Yamls:  b,
	})
	return err
}

func (d *domain) applyK8sResource(ctx InfraContext, obj client.Object) error {
	b, err := fn.K8sObjToYAML(obj)
	if err != nil {
		return err
	}
	if _, err := d.k8sYamlClient.ApplyYAML(ctx, b); err != nil {
		return err
	}

	return nil
}

func (d *domain) deleteK8sResource(ctx InfraContext, obj client.Object) error {
	b, err := fn.K8sObjToYAML(obj)
	if err != nil {
		return err
	}

	if err := d.k8sYamlClient.DeleteYAML(ctx, b); err != nil {
		return err
	}
	return nil
}

var Module = fx.Module("domain",
	fx.Provide(
		func(
			env *env.Env,
			clusterRepo repos.DbRepo[*entities.Cluster],
			providerRepo repos.DbRepo[*entities.CloudProvider],
			edgeRepo repos.DbRepo[*entities.Edge],
			masterNodeRepo repos.DbRepo[*entities.MasterNode],
			workerNodeRepo repos.DbRepo[*entities.WorkerNode],
			nodePoolRepo repos.DbRepo[*entities.NodePool],
			secretRepo repos.DbRepo[*entities.Secret],

			financeClient finance.FinanceClient,
			agentMessenger agent.Sender,

			k8sClient client.Client,
			k8sYamlClient *kubectl.YAMLClient,
			k8sExtendedClient k8s.ExtendedK8sClient,
		) Domain {
			return &domain{
				env: env,

				clusterRepo:    clusterRepo,
				providerRepo:   providerRepo,
				edgeRepo:       edgeRepo,
				masterNodeRepo: masterNodeRepo,
				workerNodeRepo: workerNodeRepo,
				nodePoolRepo:   nodePoolRepo,
				secretRepo:     secretRepo,

				agentSender: agentMessenger,

				k8sClient:         k8sClient,
				k8sYamlClient:     k8sYamlClient,
				k8sExtendedClient: k8sExtendedClient,
			}
		}),
)
