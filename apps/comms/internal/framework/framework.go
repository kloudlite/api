package framework

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"kloudlite.io/apps/comms/internal/app"
	"kloudlite.io/apps/comms/internal/env"
	"kloudlite.io/pkg/grpc"
	"kloudlite.io/pkg/logging"
	"kloudlite.io/pkg/mail"
	"time"
)

var Module = fx.Module(
	"framework",

	fx.Provide(func(ev *env.Env) mail.Mailer {
		return mail.NewSendgridMailer(ev.SendgridApiKey)
	}),

	fx.Provide(func(logger logging.Logger) (app.CommsGrpcServer, error) {
		return grpc.NewGrpcServer(grpc.ServerOpts{Logger: logger})
	}),

	app.Module,

	fx.Invoke(func(lf fx.Lifecycle, server app.CommsGrpcServer, ev *env.Env, logger logging.Logger) {
		lf.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				errCh := make(chan error, 1)

				tctx, cf := context.WithTimeout(ctx, 2*time.Second)
				defer cf()

				go func() {
					err := server.Listen(fmt.Sprintf(":%d", ev.GrpcPort))
					if err != nil {
						errCh <- err
						logger.Errorf(err, "failed to start grpc server")
					}
				}()

				select {
				case <-tctx.Done():
				case err := <-errCh:
					return err
				}

				return nil
			},

			OnStop: func(context.Context) error {
				server.Stop()
				return nil
			},
		})
	}),
)