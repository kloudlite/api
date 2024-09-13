// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: infra.proto

package infra

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Infra_GetCluster_FullMethodName                = "/Infra/GetCluster"
	Infra_GetNodepool_FullMethodName               = "/Infra/GetNodepool"
	Infra_ClusterExists_FullMethodName             = "/Infra/ClusterExists"
	Infra_GetClusterKubeconfig_FullMethodName      = "/Infra/GetClusterKubeconfig"
	Infra_MarkClusterOnlineAt_FullMethodName       = "/Infra/MarkClusterOnlineAt"
	Infra_EnsureGlobalVPNConnection_FullMethodName = "/Infra/EnsureGlobalVPNConnection"
)

// InfraClient is the client API for Infra service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InfraClient interface {
	GetCluster(ctx context.Context, in *GetClusterIn, opts ...grpc.CallOption) (*GetClusterOut, error)
	GetNodepool(ctx context.Context, in *GetNodepoolIn, opts ...grpc.CallOption) (*GetNodepoolOut, error)
	ClusterExists(ctx context.Context, in *ClusterExistsIn, opts ...grpc.CallOption) (*ClusterExistsOut, error)
	GetClusterKubeconfig(ctx context.Context, in *GetClusterIn, opts ...grpc.CallOption) (*GetClusterKubeconfigOut, error)
	MarkClusterOnlineAt(ctx context.Context, in *MarkClusterOnlineAtIn, opts ...grpc.CallOption) (*MarkClusterOnlineAtOut, error)
	EnsureGlobalVPNConnection(ctx context.Context, in *EnsureGlobalVPNConnectionIn, opts ...grpc.CallOption) (*EnsureGlobalVPNConnectionOut, error)
}

type infraClient struct {
	cc grpc.ClientConnInterface
}

func NewInfraClient(cc grpc.ClientConnInterface) InfraClient {
	return &infraClient{cc}
}

func (c *infraClient) GetCluster(ctx context.Context, in *GetClusterIn, opts ...grpc.CallOption) (*GetClusterOut, error) {
	out := new(GetClusterOut)
	err := c.cc.Invoke(ctx, Infra_GetCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraClient) GetNodepool(ctx context.Context, in *GetNodepoolIn, opts ...grpc.CallOption) (*GetNodepoolOut, error) {
	out := new(GetNodepoolOut)
	err := c.cc.Invoke(ctx, Infra_GetNodepool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraClient) ClusterExists(ctx context.Context, in *ClusterExistsIn, opts ...grpc.CallOption) (*ClusterExistsOut, error) {
	out := new(ClusterExistsOut)
	err := c.cc.Invoke(ctx, Infra_ClusterExists_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraClient) GetClusterKubeconfig(ctx context.Context, in *GetClusterIn, opts ...grpc.CallOption) (*GetClusterKubeconfigOut, error) {
	out := new(GetClusterKubeconfigOut)
	err := c.cc.Invoke(ctx, Infra_GetClusterKubeconfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraClient) MarkClusterOnlineAt(ctx context.Context, in *MarkClusterOnlineAtIn, opts ...grpc.CallOption) (*MarkClusterOnlineAtOut, error) {
	out := new(MarkClusterOnlineAtOut)
	err := c.cc.Invoke(ctx, Infra_MarkClusterOnlineAt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraClient) EnsureGlobalVPNConnection(ctx context.Context, in *EnsureGlobalVPNConnectionIn, opts ...grpc.CallOption) (*EnsureGlobalVPNConnectionOut, error) {
	out := new(EnsureGlobalVPNConnectionOut)
	err := c.cc.Invoke(ctx, Infra_EnsureGlobalVPNConnection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfraServer is the server API for Infra service.
// All implementations must embed UnimplementedInfraServer
// for forward compatibility
type InfraServer interface {
	GetCluster(context.Context, *GetClusterIn) (*GetClusterOut, error)
	GetNodepool(context.Context, *GetNodepoolIn) (*GetNodepoolOut, error)
	ClusterExists(context.Context, *ClusterExistsIn) (*ClusterExistsOut, error)
	GetClusterKubeconfig(context.Context, *GetClusterIn) (*GetClusterKubeconfigOut, error)
	MarkClusterOnlineAt(context.Context, *MarkClusterOnlineAtIn) (*MarkClusterOnlineAtOut, error)
	EnsureGlobalVPNConnection(context.Context, *EnsureGlobalVPNConnectionIn) (*EnsureGlobalVPNConnectionOut, error)
	mustEmbedUnimplementedInfraServer()
}

// UnimplementedInfraServer must be embedded to have forward compatible implementations.
type UnimplementedInfraServer struct {
}

func (UnimplementedInfraServer) GetCluster(context.Context, *GetClusterIn) (*GetClusterOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCluster not implemented")
}
func (UnimplementedInfraServer) GetNodepool(context.Context, *GetNodepoolIn) (*GetNodepoolOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodepool not implemented")
}
func (UnimplementedInfraServer) ClusterExists(context.Context, *ClusterExistsIn) (*ClusterExistsOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClusterExists not implemented")
}
func (UnimplementedInfraServer) GetClusterKubeconfig(context.Context, *GetClusterIn) (*GetClusterKubeconfigOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterKubeconfig not implemented")
}
func (UnimplementedInfraServer) MarkClusterOnlineAt(context.Context, *MarkClusterOnlineAtIn) (*MarkClusterOnlineAtOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkClusterOnlineAt not implemented")
}
func (UnimplementedInfraServer) EnsureGlobalVPNConnection(context.Context, *EnsureGlobalVPNConnectionIn) (*EnsureGlobalVPNConnectionOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnsureGlobalVPNConnection not implemented")
}
func (UnimplementedInfraServer) mustEmbedUnimplementedInfraServer() {}

// UnsafeInfraServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InfraServer will
// result in compilation errors.
type UnsafeInfraServer interface {
	mustEmbedUnimplementedInfraServer()
}

func RegisterInfraServer(s grpc.ServiceRegistrar, srv InfraServer) {
	s.RegisterService(&Infra_ServiceDesc, srv)
}

func _Infra_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Infra_GetCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraServer).GetCluster(ctx, req.(*GetClusterIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Infra_GetNodepool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodepoolIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraServer).GetNodepool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Infra_GetNodepool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraServer).GetNodepool(ctx, req.(*GetNodepoolIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Infra_ClusterExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterExistsIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraServer).ClusterExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Infra_ClusterExists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraServer).ClusterExists(ctx, req.(*ClusterExistsIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Infra_GetClusterKubeconfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraServer).GetClusterKubeconfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Infra_GetClusterKubeconfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraServer).GetClusterKubeconfig(ctx, req.(*GetClusterIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Infra_MarkClusterOnlineAt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkClusterOnlineAtIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraServer).MarkClusterOnlineAt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Infra_MarkClusterOnlineAt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraServer).MarkClusterOnlineAt(ctx, req.(*MarkClusterOnlineAtIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Infra_EnsureGlobalVPNConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnsureGlobalVPNConnectionIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraServer).EnsureGlobalVPNConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Infra_EnsureGlobalVPNConnection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraServer).EnsureGlobalVPNConnection(ctx, req.(*EnsureGlobalVPNConnectionIn))
	}
	return interceptor(ctx, in, info, handler)
}

// Infra_ServiceDesc is the grpc.ServiceDesc for Infra service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Infra_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Infra",
	HandlerType: (*InfraServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCluster",
			Handler:    _Infra_GetCluster_Handler,
		},
		{
			MethodName: "GetNodepool",
			Handler:    _Infra_GetNodepool_Handler,
		},
		{
			MethodName: "ClusterExists",
			Handler:    _Infra_ClusterExists_Handler,
		},
		{
			MethodName: "GetClusterKubeconfig",
			Handler:    _Infra_GetClusterKubeconfig_Handler,
		},
		{
			MethodName: "MarkClusterOnlineAt",
			Handler:    _Infra_MarkClusterOnlineAt_Handler,
		},
		{
			MethodName: "EnsureGlobalVPNConnection",
			Handler:    _Infra_EnsureGlobalVPNConnection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infra.proto",
}
