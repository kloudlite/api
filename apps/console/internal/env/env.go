package env

import (
	"github.com/codingconcepts/env"
	"github.com/kloudlite/api/pkg/errors"
)

type Env struct {
	Port     uint16 `env:"HTTP_PORT" required:"true"`
	GrpcPort uint16 `env:"GRPC_PORT" required:"true"`

	DNSAddr            string `env:"DNS_ADDR" required:"true"`
	KloudliteDNSSuffix string `env:"KLOUDLITE_DNS_SUFFIX" required:"true"`

	ConsoleDBUri  string `env:"MONGO_URI" required:"true"`
	ConsoleDBName string `env:"MONGO_DB_NAME" required:"true"`

	AccountCookieName string `env:"ACCOUNT_COOKIE_NAME" required:"true"`
	ClusterCookieName string `env:"CLUSTER_COOKIE_NAME" required:"true"`

	NatsURL                    string `env:"NATS_URL" required:"true"`
	NatsReceiveFromAgentStream string `env:"NATS_RECEIVE_FROM_AGENT_STREAM" required:"true"`

	IAMGrpcAddr   string `env:"IAM_GRPC_ADDR" required:"true"`
	InfraGrpcAddr string `env:"INFRA_GRPC_ADDR" required:"true"`

	PromHttpAddr         string `env:"PROM_HTTP_ADDR" required:"true"`
	SessionKVBucket      string `env:"SESSION_KV_BUCKET" required:"true"`
	ConsoleCacheKVBucket string `env:"CONSOLE_CACHE_KV_BUCKET" required:"true"`
	IsDev                bool
	KubernetesApiProxy   string `env:"KUBERNETES_API_PROXY"`

	DeviceNamespace string `env:"DEVICE_NAMESPACE" required:"true"`
}

func LoadEnv() (*Env, error) {
	var e Env
	if err := env.Set(&e); err != nil {
		return nil, errors.NewE(err)
	}
	return &e, nil
}
