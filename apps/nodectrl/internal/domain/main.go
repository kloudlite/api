package domain

import (
	"fmt"

	"go.uber.org/fx"
	"kloudlite.io/apps/nodectrl/internal/env"
)

type domain struct {
	env *env.Env
}

// StartJob implements Domain
func (d domain) StartJob() error {
	switch d.env.CloudProvider {
	case "aws":
		if err := d.StartAwsJob(); err != nil {
			return err
		}
	case "azure":
		if err := d.StartAzureJob(); err != nil {
			return err
		}
	case "do":
		if err := d.StartDoJob(); err != nil {
			return err
		}
	case "gcp":
		if err := d.StartGCPJob(); err != nil {
			return err
		}
	case "":
		return fmt.Errorf("please provide cloud provider")
	default:
		return fmt.Errorf("provided cloud provider %s not available now", d.env.CloudProvider)

	}
	return nil
}

var Module = fx.Module("domain",
	fx.Provide(
		func(env *env.Env) Domain {
			return domain{
				env: env,
			}
		},
	),
)
