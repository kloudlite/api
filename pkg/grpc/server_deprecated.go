package grpc

import (
	"context"
	"fmt"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/logging"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"net"
)

type ServerOptions interface {
	GetGRPCPort() uint16
}

func NewGrpcServerFx[T ServerOptions]() fx.Option {
	return fx.Module(
		"grpc-server",
		fx.Provide(func(logger logging.Logger) *grpc.Server {
			return grpc.NewServer(
				grpc.StreamInterceptor(func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
					p, ok := peer.FromContext(stream.Context())
					if ok {
						logger.Debugf("[Stream] New connection from %s", p.Addr.String())
					}
					return handler(srv, stream)
				}),
			)
		}),
		fx.Invoke(
			func(lf fx.Lifecycle, env T, server *grpc.Server, logger logging.Logger) {
				lf.Append(
					fx.Hook{
						OnStart: func(ctx context.Context) error {
							listen, err := net.Listen("tcp", fmt.Sprintf(":%d", env.GetGRPCPort()))
							defer func() {
								logger.Infof("[GRPC Server] started on port (%d)", env.GetGRPCPort())
							}()
							if err != nil {
								return errors.NewEf(err, "could not listen to net/tcp server")
							}
							go func() {
								err := server.Serve(listen)
								if err != nil {
									logger.Errorf(err, "could not start grpc server ")
								}
							}()
							return nil
						},
						OnStop: func(context.Context) error {
							server.Stop()
							return nil
						},
					},
				)
			},
		),
	)
}
