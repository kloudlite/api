package app

import (
	"context"

	"github.com/kloudlite/api/pkg/errors"

	"github.com/kloudlite/api/apps/finance/internal/domain"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/accounts"
	"github.com/kloudlite/api/pkg/grpc"
	"github.com/kloudlite/api/pkg/repos"
)

type FinanceGrpcServer grpc.Server

type financeGrpcServer struct {
	accounts.UnimplementedAccountsServer
	d domain.Domain
}

// GetAccount implements finance.AccountsServer.
func (s *financeGrpcServer) GetAccount(ctx context.Context, in *accounts.GetAccountIn) (*accounts.GetAccountOut, error) {
	acc, err := s.d.GetAccount(domain.UserContext{
		Context: ctx,
		UserId:  repos.ID(in.UserId),
	}, in.AccountName)
	if err != nil {
		return nil, errors.NewE(err)
	}

	isActive := false
	if acc.IsActive != nil {
		isActive = *acc.IsActive
	}

	return &accounts.GetAccountOut{
		IsActive:               isActive,
		TargetNamespace:        acc.TargetNamespace,
		AccountId:              string(acc.Id),
		KloudliteGatewayRegion: acc.KloudliteGatewayRegion,
	}, nil
}

func registerAccountsGRPCServer(server FinanceGrpcServer, d domain.Domain) accounts.AccountsServer {
	accountsSvc := &financeGrpcServer{d: d}
	accounts.RegisterAccountsServer(server, accountsSvc)
	return accountsSvc
}
