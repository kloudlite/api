package mocks

import (
	context "context"
	message_office_internal "github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/message-office-internal"
	grpc "google.golang.org/grpc"
)

type MessageOfficeInternalClientCallerInfo struct {
	Args []any
}

type MessageOfficeInternalClient struct {
	Calls                    map[string][]MessageOfficeInternalClientCallerInfo
	MockGenerateClusterToken func(ctx context.Context, in *message_office_internal.GenerateClusterTokenIn, opts ...grpc.CallOption) (*message_office_internal.GenerateClusterTokenOut, error)
	MockGetClusterToken      func(ctx context.Context, in *message_office_internal.GetClusterTokenIn, opts ...grpc.CallOption) (*message_office_internal.GetClusterTokenOut, error)
}

func (m *MessageOfficeInternalClient) registerCall(funcName string, args ...any) {
	if m.Calls == nil {
		m.Calls = map[string][]MessageOfficeInternalClientCallerInfo{}
	}
	m.Calls[funcName] = append(m.Calls[funcName], MessageOfficeInternalClientCallerInfo{Args: args})
}

func (mMock *MessageOfficeInternalClient) GenerateClusterToken(ctx context.Context, in *message_office_internal.GenerateClusterTokenIn, opts ...grpc.CallOption) (*message_office_internal.GenerateClusterTokenOut, error) {
	if mMock.MockGenerateClusterToken != nil {
		mMock.registerCall("GenerateClusterToken", ctx, in, opts)
		return mMock.MockGenerateClusterToken(ctx, in, opts...)
	}
	panic("MessageOfficeInternalClient: method 'GenerateClusterToken' not implemented, yet")
}

func (mMock *MessageOfficeInternalClient) GetClusterToken(ctx context.Context, in *message_office_internal.GetClusterTokenIn, opts ...grpc.CallOption) (*message_office_internal.GetClusterTokenOut, error) {
	if mMock.MockGetClusterToken != nil {
		mMock.registerCall("GetClusterToken", ctx, in, opts)
		return mMock.MockGetClusterToken(ctx, in, opts...)
	}
	panic("MessageOfficeInternalClient: method 'GetClusterToken' not implemented, yet")
}

func NewMessageOfficeInternalClient() *MessageOfficeInternalClient {
	return &MessageOfficeInternalClient{}
}
