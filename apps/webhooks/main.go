package main

import (
	"flag"

	"go.uber.org/fx"
	"kloudlite.io/apps/webhooks/internal/env"
	"kloudlite.io/apps/webhooks/internal/framework"
	"kloudlite.io/pkg/config"
	"kloudlite.io/pkg/logging"
)

func main() {
	var isDev bool
	flag.BoolVar(&isDev, "dev", false, "--dev")
	flag.Parse()

	fx.New(
		fx.Provide(
			func() (logging.Logger, error) {
				return logging.New(&logging.Options{Name: "webhooks", Dev: isDev})
			},
		),
		config.EnvFx[env.Env](),
		framework.Module,
		func() fx.Option {
			if !isDev {
				return fx.NopLogger
			}
			return nil
		}(),
	).Run()
}
