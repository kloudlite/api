package domain

import (
	"go.uber.org/fx"
	"kloudlite.io/pkg/config"
	"kloudlite.io/pkg/messaging"
)

type Domain interface {
	CreateCluster(action SetupClusterAction) error
	UpdateCluster(action UpdateClusterAction) error
}

type domain struct {
	infraCli        InfraClient
	messageProducer messaging.Producer[messaging.Json]
	messageTopic    string
}

func (d *domain) CreateCluster(action SetupClusterAction) error {
	err := d.infraCli.CreateKubernetes(action)
	err = d.infraCli.SetupCSI(action.ClusterID, action.Provider)
	d.messageProducer.SendMessage(d.messageTopic, action.ClusterID, messaging.Json{
		"cluster_id": action.ClusterID,
		"status":     "live",
	})
	return err
}

func (d *domain) UpdateCluster(action UpdateClusterAction) error {
	err := d.infraCli.UpdateKubernetes(action)
	return err
}
func makeDomain(
	env *Env,
	infraCli InfraClient,
	messageProducer messaging.Producer[messaging.Json],
) Domain {
	return &domain{
		infraCli:        infraCli,
		messageProducer: messageProducer,
		messageTopic:    env.KafkaInfraActionResulTopic,
	}
}

type Env struct {
	KafkaInfraActionResulTopic string `env:"KAFKA_INFRA_ACTION_RESULT_TOPIC", required:"true"`
}

var Module = fx.Module("domain",
	fx.Provide(config.LoadEnv[Env]()),
	fx.Provide(makeDomain))
