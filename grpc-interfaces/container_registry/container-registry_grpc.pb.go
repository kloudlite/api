// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: container-registry.proto

package container_registry

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
	ContainerRegistry_CreateReadOnlyCredential_FullMethodName = "/ContainerRegistry/CreateReadOnlyCredential"
)

// ContainerRegistryClient is the client API for ContainerRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContainerRegistryClient interface {
	// rpc CreateProjectForAccount(CreateProjectIn) returns (CreateProjectOut);
	// rpc GetSvcCredentials(GetSvcCredentialsIn) returns (GetSvcCredentialsOut);
	CreateReadOnlyCredential(ctx context.Context, in *CreateReadOnlyCredentialIn, opts ...grpc.CallOption) (*CreateReadOnlyCredentialOut, error)
}

type containerRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewContainerRegistryClient(cc grpc.ClientConnInterface) ContainerRegistryClient {
	return &containerRegistryClient{cc}
}

func (c *containerRegistryClient) CreateReadOnlyCredential(ctx context.Context, in *CreateReadOnlyCredentialIn, opts ...grpc.CallOption) (*CreateReadOnlyCredentialOut, error) {
	out := new(CreateReadOnlyCredentialOut)
	err := c.cc.Invoke(ctx, ContainerRegistry_CreateReadOnlyCredential_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContainerRegistryServer is the server API for ContainerRegistry service.
// All implementations must embed UnimplementedContainerRegistryServer
// for forward compatibility
type ContainerRegistryServer interface {
	// rpc CreateProjectForAccount(CreateProjectIn) returns (CreateProjectOut);
	// rpc GetSvcCredentials(GetSvcCredentialsIn) returns (GetSvcCredentialsOut);
	CreateReadOnlyCredential(context.Context, *CreateReadOnlyCredentialIn) (*CreateReadOnlyCredentialOut, error)
	mustEmbedUnimplementedContainerRegistryServer()
}

// UnimplementedContainerRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedContainerRegistryServer struct {
}

func (UnimplementedContainerRegistryServer) CreateReadOnlyCredential(context.Context, *CreateReadOnlyCredentialIn) (*CreateReadOnlyCredentialOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReadOnlyCredential not implemented")
}
func (UnimplementedContainerRegistryServer) mustEmbedUnimplementedContainerRegistryServer() {}

// UnsafeContainerRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContainerRegistryServer will
// result in compilation errors.
type UnsafeContainerRegistryServer interface {
	mustEmbedUnimplementedContainerRegistryServer()
}

func RegisterContainerRegistryServer(s grpc.ServiceRegistrar, srv ContainerRegistryServer) {
	s.RegisterService(&ContainerRegistry_ServiceDesc, srv)
}

func _ContainerRegistry_CreateReadOnlyCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReadOnlyCredentialIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerRegistryServer).CreateReadOnlyCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContainerRegistry_CreateReadOnlyCredential_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerRegistryServer).CreateReadOnlyCredential(ctx, req.(*CreateReadOnlyCredentialIn))
	}
	return interceptor(ctx, in, info, handler)
}

// ContainerRegistry_ServiceDesc is the grpc.ServiceDesc for ContainerRegistry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContainerRegistry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ContainerRegistry",
	HandlerType: (*ContainerRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReadOnlyCredential",
			Handler:    _ContainerRegistry_CreateReadOnlyCredential_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "container-registry.proto",
}
