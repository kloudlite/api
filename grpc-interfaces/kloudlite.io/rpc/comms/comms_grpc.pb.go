// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: comms.proto

package comms

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

// CommsClient is the client API for Comms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommsClient interface {
	SendVerificationEmail(ctx context.Context, in *VerificationEmailInput, opts ...grpc.CallOption) (*Void, error)
	SendPasswordResetEmail(ctx context.Context, in *PasswordResetEmailInput, opts ...grpc.CallOption) (*Void, error)
	SendAccountMemberInviteEmail(ctx context.Context, in *AccountMemberInviteEmailInput, opts ...grpc.CallOption) (*Void, error)
	SendProjectMemberInviteEmail(ctx context.Context, in *ProjectMemberInviteEmailInput, opts ...grpc.CallOption) (*Void, error)
	SendWelcomeEmail(ctx context.Context, in *WelcomeEmailInput, opts ...grpc.CallOption) (*Void, error)
}

type commsClient struct {
	cc grpc.ClientConnInterface
}

func NewCommsClient(cc grpc.ClientConnInterface) CommsClient {
	return &commsClient{cc}
}

func (c *commsClient) SendVerificationEmail(ctx context.Context, in *VerificationEmailInput, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/Comms/SendVerificationEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commsClient) SendPasswordResetEmail(ctx context.Context, in *PasswordResetEmailInput, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/Comms/SendPasswordResetEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commsClient) SendAccountMemberInviteEmail(ctx context.Context, in *AccountMemberInviteEmailInput, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/Comms/SendAccountMemberInviteEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commsClient) SendProjectMemberInviteEmail(ctx context.Context, in *ProjectMemberInviteEmailInput, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/Comms/SendProjectMemberInviteEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commsClient) SendWelcomeEmail(ctx context.Context, in *WelcomeEmailInput, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/Comms/SendWelcomeEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommsServer is the server API for Comms service.
// All implementations must embed UnimplementedCommsServer
// for forward compatibility
type CommsServer interface {
	SendVerificationEmail(context.Context, *VerificationEmailInput) (*Void, error)
	SendPasswordResetEmail(context.Context, *PasswordResetEmailInput) (*Void, error)
	SendAccountMemberInviteEmail(context.Context, *AccountMemberInviteEmailInput) (*Void, error)
	SendProjectMemberInviteEmail(context.Context, *ProjectMemberInviteEmailInput) (*Void, error)
	SendWelcomeEmail(context.Context, *WelcomeEmailInput) (*Void, error)
	mustEmbedUnimplementedCommsServer()
}

// UnimplementedCommsServer must be embedded to have forward compatible implementations.
type UnimplementedCommsServer struct {
}

func (UnimplementedCommsServer) SendVerificationEmail(context.Context, *VerificationEmailInput) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendVerificationEmail not implemented")
}
func (UnimplementedCommsServer) SendPasswordResetEmail(context.Context, *PasswordResetEmailInput) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPasswordResetEmail not implemented")
}
func (UnimplementedCommsServer) SendAccountMemberInviteEmail(context.Context, *AccountMemberInviteEmailInput) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAccountMemberInviteEmail not implemented")
}
func (UnimplementedCommsServer) SendProjectMemberInviteEmail(context.Context, *ProjectMemberInviteEmailInput) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendProjectMemberInviteEmail not implemented")
}
func (UnimplementedCommsServer) SendWelcomeEmail(context.Context, *WelcomeEmailInput) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendWelcomeEmail not implemented")
}
func (UnimplementedCommsServer) mustEmbedUnimplementedCommsServer() {}

// UnsafeCommsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommsServer will
// result in compilation errors.
type UnsafeCommsServer interface {
	mustEmbedUnimplementedCommsServer()
}

func RegisterCommsServer(s grpc.ServiceRegistrar, srv CommsServer) {
	s.RegisterService(&Comms_ServiceDesc, srv)
}

func _Comms_SendVerificationEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerificationEmailInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommsServer).SendVerificationEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Comms/SendVerificationEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommsServer).SendVerificationEmail(ctx, req.(*VerificationEmailInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comms_SendPasswordResetEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordResetEmailInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommsServer).SendPasswordResetEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Comms/SendPasswordResetEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommsServer).SendPasswordResetEmail(ctx, req.(*PasswordResetEmailInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comms_SendAccountMemberInviteEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountMemberInviteEmailInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommsServer).SendAccountMemberInviteEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Comms/SendAccountMemberInviteEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommsServer).SendAccountMemberInviteEmail(ctx, req.(*AccountMemberInviteEmailInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comms_SendProjectMemberInviteEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectMemberInviteEmailInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommsServer).SendProjectMemberInviteEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Comms/SendProjectMemberInviteEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommsServer).SendProjectMemberInviteEmail(ctx, req.(*ProjectMemberInviteEmailInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comms_SendWelcomeEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WelcomeEmailInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommsServer).SendWelcomeEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Comms/SendWelcomeEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommsServer).SendWelcomeEmail(ctx, req.(*WelcomeEmailInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Comms_ServiceDesc is the grpc.ServiceDesc for Comms service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comms_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Comms",
	HandlerType: (*CommsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendVerificationEmail",
			Handler:    _Comms_SendVerificationEmail_Handler,
		},
		{
			MethodName: "SendPasswordResetEmail",
			Handler:    _Comms_SendPasswordResetEmail_Handler,
		},
		{
			MethodName: "SendAccountMemberInviteEmail",
			Handler:    _Comms_SendAccountMemberInviteEmail_Handler,
		},
		{
			MethodName: "SendProjectMemberInviteEmail",
			Handler:    _Comms_SendProjectMemberInviteEmail_Handler,
		},
		{
			MethodName: "SendWelcomeEmail",
			Handler:    _Comms_SendWelcomeEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comms.proto",
}
