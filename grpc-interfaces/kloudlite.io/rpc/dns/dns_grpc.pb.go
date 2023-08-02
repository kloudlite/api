// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: dns.proto

package dns

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
	DNS_GetAccountDomains_FullMethodName = "/DNS/GetAccountDomains"
)

// DNSClient is the client API for DNS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DNSClient interface {
	GetAccountDomains(ctx context.Context, in *GetAccountDomainsIn, opts ...grpc.CallOption) (*GetAccountDomainsOut, error)
}

type dNSClient struct {
	cc grpc.ClientConnInterface
}

func NewDNSClient(cc grpc.ClientConnInterface) DNSClient {
	return &dNSClient{cc}
}

func (c *dNSClient) GetAccountDomains(ctx context.Context, in *GetAccountDomainsIn, opts ...grpc.CallOption) (*GetAccountDomainsOut, error) {
	out := new(GetAccountDomainsOut)
	err := c.cc.Invoke(ctx, DNS_GetAccountDomains_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DNSServer is the server API for DNS service.
// All implementations must embed UnimplementedDNSServer
// for forward compatibility
type DNSServer interface {
	GetAccountDomains(context.Context, *GetAccountDomainsIn) (*GetAccountDomainsOut, error)
	mustEmbedUnimplementedDNSServer()
}

// UnimplementedDNSServer must be embedded to have forward compatible implementations.
type UnimplementedDNSServer struct {
}

func (UnimplementedDNSServer) GetAccountDomains(context.Context, *GetAccountDomainsIn) (*GetAccountDomainsOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountDomains not implemented")
}
func (UnimplementedDNSServer) mustEmbedUnimplementedDNSServer() {}

// UnsafeDNSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DNSServer will
// result in compilation errors.
type UnsafeDNSServer interface {
	mustEmbedUnimplementedDNSServer()
}

func RegisterDNSServer(s grpc.ServiceRegistrar, srv DNSServer) {
	s.RegisterService(&DNS_ServiceDesc, srv)
}

func _DNS_GetAccountDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountDomainsIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSServer).GetAccountDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DNS_GetAccountDomains_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSServer).GetAccountDomains(ctx, req.(*GetAccountDomainsIn))
	}
	return interceptor(ctx, in, info, handler)
}

// DNS_ServiceDesc is the grpc.ServiceDesc for DNS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DNS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DNS",
	HandlerType: (*DNSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountDomains",
			Handler:    _DNS_GetAccountDomains_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dns.proto",
}
