package app

import (
	"context"
	"fmt"

	"go.uber.org/fx"
	"kloudlite.io/apps/nodectrl/internal/domain"
	"kloudlite.io/apps/nodectrl/internal/env"
)

var Module = fx.Module("app",
	domain.Module,
	fx.Invoke(
		func(env *env.Env, pc domain.ProviderClient, shutdowner fx.Shutdowner, lifecycle fx.Lifecycle) {
			lifecycle.Append(fx.Hook{
				OnStart: func(context.Context) error {
					switch env.Action {
					case "create":

						fmt.Println("needs to create node")
						if err := pc.NewNode(); err != nil {
							return err
						}
					case "delete":
						fmt.Println("needs to delete node")
						if err := pc.DeleteNode(); err != nil {
							return err
						}

					case "":
						return fmt.Errorf("ACTION not provided, supported actions {create, delete} ")
					default:
						return fmt.Errorf("not supported actions '%s' please provide one of supported action like { create, delete }", env.Action)

					}
					shutdowner.Shutdown()
					return nil
				},
				OnStop: func(context.Context) error {
					return nil
				},
			})

		},
	),
)
