package adapters

import (
	"context"

	"github.com/kloudlite/api/apps/infra/internal/domain/ports"
	message_office_internal "github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/message-office-internal"
)

type messageOfficeInternalSvc struct {
	cli message_office_internal.MessageOfficeInternalClient
}

// GenerateClusterToken implements ports.MessageOfficeInternalSvc.
func (m *messageOfficeInternalSvc) GenerateClusterToken(ctx context.Context, accountName string, clusterName string) (string, error) {
	out, err := m.cli.GenerateClusterToken(ctx, &message_office_internal.GenerateClusterTokenIn{
		AccountName: accountName,
		ClusterName: clusterName,
	})
	if err != nil {
		return "", err
	}

	return out.ClusterToken, nil
}

// GetClusterToken implements ports.MessageOfficeInternalSvc.
func (m *messageOfficeInternalSvc) GetClusterToken(ctx context.Context, accountName string, clusterName string) (string, error) {
	out, err := m.cli.GetClusterToken(ctx, &message_office_internal.GetClusterTokenIn{
		AccountName: accountName,
		ClusterName: clusterName,
	})
	if err != nil {
		return "", err
	}

	return out.ClusterToken, nil
}

func NewMessageOfficeInternalSvc(cli message_office_internal.MessageOfficeInternalClient) ports.MessageOfficeInternalSvc {
	return &messageOfficeInternalSvc{
		cli: cli,
	}
}
