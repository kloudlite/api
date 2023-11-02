package app

import (
	"context"

	"kloudlite.io/apps/finance_deprecated/internal/domain"
	"kloudlite.io/grpc-interfaces/kloudlite.io/rpc/finance"
)

type financeServerImpl struct {
	finance.UnimplementedFinanceServer
	d domain.Domain
}

func (f financeServerImpl) StartBillable(ctx context.Context, in *finance.StartBillableIn) (*finance.StartBillableOut, error) {
	// TODO implement me
	panic("implement me")
}

func (f financeServerImpl) StopBillable(ctx context.Context, in *finance.StopBillableIn) (*finance.StopBillableOut, error) {
	// TODO implement me
	panic("implement me")
}

func (f financeServerImpl) GetAttachedCluster(ctx context.Context, in *finance.GetAttachedClusterIn) (*finance.GetAttachedClusterOut, error) {
	account, err := f.d.GetAccount(domain.FinanceContext{Context: ctx, UserId: ""}, in.AccountName)
	if err != nil {
		return nil, err
	}
	return &finance.GetAttachedClusterOut{ClusterId: string(account.ClusterID)}, nil
}

func fxFinanceGrpcServer(domain domain.Domain) finance.FinanceServer {
	return &financeServerImpl{
		d: domain,
	}
}