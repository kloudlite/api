package grpc

import (
	"github.com/kloudlite/api/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	grpc.ClientConnInterface
	Close() error
}

type grpcClient struct {
	*grpc.ClientConn
}

func (g *grpcClient) Close() error {
	return g.ClientConn.Close()
}

func NewGrpcClient(serverAddr string) (Client, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.NewE(err)
	}
	return &grpcClient{ClientConn: conn}, nil
}
