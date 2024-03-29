package env

import "github.com/codingconcepts/env"

type Env struct {
	HttpPort        uint16 `env:"HTTP_PORT" required:"true"`
	HttpCorsOrigins string `env:"HTTP_CORS_ORIGINS" required:"false"`

	AccountCookieName string `env:"ACCOUNT_COOKIE_NAME" required:"true"`

	NatsURL         string `env:"NATS_URL" required:"true"`
	SessionKVBucket string `env:"SESSION_KV_BUCKET" required:"true"`

	IAMGrpcAddr string `env:"IAM_GRPC_ADDR" required:"true"`

	PromHttpAddr string `env:"PROM_HTTP_ADDR" required:"true"`

	IsDev bool
}

func LoadEnv() (*Env, error) {
	var ev Env
	if err := env.Set(&ev); err != nil {
		return nil, err
	}
	return &ev, nil
}
