package env

import "github.com/codingconcepts/env"

type Env struct {
	CloudProvider string `env:"CLOUD_PROVIDER" required:"true"`
	Action        string `env:"ACTION" required:"true"`

	NodeConfig     string `env:"NODE_CONFIG" required:"true"`
	ProviderConfig string `env:"PROVIDER_CONFIG" required:"true"`

	AWSProviderConfig   string `env:"AWS_CONFIG"`
	GCPProviderConfig   string `env:"GCP_CONFIG"`
	AzureProviderConfig string `env:"AZURE_CONFIG"`
	DoProviderConfig    string `env:"DO_CONFIG"`

	DBUrl  string `env:"DB_URL" required:"true"`
	DBName string `env:"DB_NAME" required:"true"`
}

func LoadEnv() (*Env, error) {
	var e Env
	if err := env.Set(&e); err != nil {
		return nil, err
	}
	return &e, nil
}
