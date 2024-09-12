package framework

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/nats"

	"github.com/kloudlite/api/pkg/kv"
	"github.com/kloudlite/api/pkg/repos"

	"github.com/kloudlite/api/pkg/logging"

	"github.com/kloudlite/api/apps/finance/internal/app"
	"github.com/kloudlite/api/apps/finance/internal/env"
	"github.com/kloudlite/api/pkg/grpc"
	httpServer "github.com/kloudlite/api/pkg/http-server"
	"go.uber.org/fx"
)

type fm struct {
	env *env.Env
}

func (f *fm) GetMongoConfig() (url string, dbName string) {
	return f.env.DBUrl, f.env.DBName
}

var Module = fx.Module("framework",
	fx.Provide(func(ev *env.Env) *fm {
		return &fm{env: ev}
	}),

	fx.Provide(func(ev *env.Env, logger *slog.Logger) (*nats.JetstreamClient, error) {
		name := "finance:jetstream-client"
		nc, err := nats.NewClient(ev.NatsURL, nats.ClientOpts{
			Name:   name,
			Logger: logger,
		})
		if err != nil {
			return nil, errors.NewE(err)
		}

		return nats.NewJetstreamClient(nc)
	}),

	fx.Provide(
		func(ev *env.Env, jc *nats.JetstreamClient) (kv.Repo[*common.AuthSession], error) {
			cxt := context.TODO()
			return kv.NewNatsKVRepo[*common.AuthSession](cxt, ev.SessionKVBucket, jc)
		},
	),
	repos.NewMongoClientFx[*fm](),

	fx.Provide(func(ev *env.Env) (app.AuthClient, error) {
		return grpc.NewGrpcClient(ev.AuthGrpcAddr)
	}),

	fx.Provide(func(ev *env.Env) (app.IAMClient, error) {
		return grpc.NewGrpcClient(ev.IamGrpcAddr)
	}),

	fx.Provide(func(ev *env.Env) (app.CommsClient, error) {
		return grpc.NewGrpcClient(ev.CommsGrpcAddr)
	}),

	app.Module,

	fx.Invoke(func(c1 app.AuthClient, c2 app.IAMClient, c3 app.CommsClient, lf fx.Lifecycle) {
		lf.Append(fx.Hook{
			OnStop: func(context.Context) error {
				if err := c1.Close(); err != nil {
					return errors.NewE(err)
				}
				if err := c2.Close(); err != nil {
					return errors.NewE(err)
				}
				if err := c3.Close(); err != nil {
					return errors.NewE(err)
				}
				return nil
			},
		})
	}),

	fx.Provide(func(logger logging.Logger, e *env.Env) httpServer.Server {
		corsOrigins := "https://studio.apollographql.com"
		return httpServer.NewServer(httpServer.ServerArgs{Logger: logger, CorsAllowOrigins: &corsOrigins, IsDev: e.IsDev})
	}),

	fx.Invoke(func(lf fx.Lifecycle, server httpServer.Server, ev *env.Env) {
		lf.Append(fx.Hook{
			OnStart: func(context.Context) error {
				return server.Listen(fmt.Sprintf(":%d", ev.HttpPort))
			},
			OnStop: func(context.Context) error {
				return server.Close()
			},
		})
	}),
)
