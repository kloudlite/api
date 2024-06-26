package main

import (
	"context"

	proto_rpc "github.com/kloudlite/api/apps/tenant-agent/internal/proto-rpc"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/operator/pkg/logging"
	"google.golang.org/grpc/metadata"
)

type vectorGrpcProxyServer struct {
	proto_rpc.UnimplementedVectorServer
	realVectorClient proto_rpc.VectorClient
	logger           logging.Logger

	errCh chan error

	accessToken string
	accountName string
	clusterName string

	pushEventsCounter  int
	healthCheckCounter int
}

func (v *vectorGrpcProxyServer) PushEvents(ctx context.Context, msg *proto_rpc.PushEventsRequest) (*proto_rpc.PushEventsResponse, error) {
	if v.realVectorClient == nil {
		return nil, errors.Newf("vector client is not yet connected to message-office")
	}

	outgoingCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", v.accessToken))

	v.pushEventsCounter++
	v.logger.Debugf("[%v] received push-events message", v.pushEventsCounter)
	defer v.logger.Debugf("[%v] dispatched push-events message", v.pushEventsCounter)

	per, err := v.realVectorClient.PushEvents(outgoingCtx, msg)
	if err != nil {
		v.logger.Error(err)
		if v.errCh != nil {
			v.errCh <- err
		}
		return nil, errors.NewE(err)
	}
	return per, nil
}

func (v *vectorGrpcProxyServer) HealthCheck(ctx context.Context, msg *proto_rpc.HealthCheckRequest) (*proto_rpc.HealthCheckResponse, error) {
	if v.realVectorClient == nil {
		return nil, errors.Newf("vector client is not yet connected to message-office")
	}

	outgoingCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", v.accessToken))

	v.healthCheckCounter++
	v.logger.Debugf("[%v] received health-check message", v.healthCheckCounter)
	defer v.logger.Debugf("[%v] dispatched health-check message", v.healthCheckCounter)
	hcr, err := v.realVectorClient.HealthCheck(outgoingCtx, msg)
	if err != nil {
		v.logger.Error(err)
		if v.errCh != nil {
			v.errCh <- err
		}
		return nil, errors.NewE(err)
	}
	return hcr, nil
}
