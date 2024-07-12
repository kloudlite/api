package app

import (
	"encoding/json"
	"os"

	"github.com/kloudlite/api/apps/iam/internal/entities"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/grpc"
	"github.com/kloudlite/api/pkg/logging"

	"github.com/kloudlite/api/apps/iam/internal/env"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/accounts"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/iam"
	"github.com/kloudlite/api/pkg/repos"
	"go.uber.org/fx"
)

type AccountsClient grpc.Client

var Module = fx.Module(
	"app",
	fx.Provide(func(ev *env.Env) (RoleBindingMap, error) {
		if ev.ActionRoleMapFile != "" {
			b, err := os.ReadFile(ev.ActionRoleMapFile)
			if err != nil {
				return nil, errors.NewE(err)
			}
			var rbm RoleBindingMap
			if err := json.Unmarshal(b, &rbm); err != nil {
				return nil, errors.NewE(err)
			}
			return rbm, nil
		}

		return roleBindings, nil
	}),

	repos.NewFxMongoRepo[*entities.RoleBinding]("role_bindings", "rb", entities.RoleBindingIndices),

	fx.Provide(
		func(conn AccountsClient) accounts.AccountsClient {
			return accounts.NewAccountsClient(conn)
		},
	),

	fx.Provide(func(logger logging.Logger, rbRepo repos.DbRepo[*entities.RoleBinding], rbm RoleBindingMap) iam.IAMServer {
		return newIAMGrpcService(logger, rbRepo, rbm)
	}),

	fx.Invoke(
		func(server IAMGrpcServer, iamService iam.IAMServer) {
			iam.RegisterIAMServer(server, iamService)
		},
	),
)
