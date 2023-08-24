// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	console "kloudlite.io/grpc-interfaces/kloudlite.io/rpc/console"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// console_client is an autogenerated mock type for the ConsoleClient type
type console_client struct {
	mock.Mock
}

type console_client_Expecter struct {
	mock *mock.Mock
}

func (_m *console_client) EXPECT() *console_client_Expecter {
	return &console_client_Expecter{mock: &_m.Mock}
}

// GetApp provides a mock function with given fields: ctx, in, opts
func (_m *console_client) GetApp(ctx context.Context, in *console.AppIn, opts ...grpc.CallOption) (*console.AppOut, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *console.AppOut
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *console.AppIn, ...grpc.CallOption) (*console.AppOut, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *console.AppIn, ...grpc.CallOption) *console.AppOut); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*console.AppOut)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *console.AppIn, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// console_client_GetApp_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetApp'
type console_client_GetApp_Call struct {
	*mock.Call
}

// GetApp is a helper method to define mock.On call
//   - ctx context.Context
//   - in *console.AppIn
//   - opts ...grpc.CallOption
func (_e *console_client_Expecter) GetApp(ctx interface{}, in interface{}, opts ...interface{}) *console_client_GetApp_Call {
	return &console_client_GetApp_Call{Call: _e.mock.On("GetApp",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *console_client_GetApp_Call) Run(run func(ctx context.Context, in *console.AppIn, opts ...grpc.CallOption)) *console_client_GetApp_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*console.AppIn), variadicArgs...)
	})
	return _c
}

func (_c *console_client_GetApp_Call) Return(_a0 *console.AppOut, _a1 error) *console_client_GetApp_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *console_client_GetApp_Call) RunAndReturn(run func(context.Context, *console.AppIn, ...grpc.CallOption) (*console.AppOut, error)) *console_client_GetApp_Call {
	_c.Call.Return(run)
	return _c
}

// GetManagedSvc provides a mock function with given fields: ctx, in, opts
func (_m *console_client) GetManagedSvc(ctx context.Context, in *console.MSvcIn, opts ...grpc.CallOption) (*console.MSvcOut, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *console.MSvcOut
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *console.MSvcIn, ...grpc.CallOption) (*console.MSvcOut, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *console.MSvcIn, ...grpc.CallOption) *console.MSvcOut); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*console.MSvcOut)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *console.MSvcIn, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// console_client_GetManagedSvc_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetManagedSvc'
type console_client_GetManagedSvc_Call struct {
	*mock.Call
}

// GetManagedSvc is a helper method to define mock.On call
//   - ctx context.Context
//   - in *console.MSvcIn
//   - opts ...grpc.CallOption
func (_e *console_client_Expecter) GetManagedSvc(ctx interface{}, in interface{}, opts ...interface{}) *console_client_GetManagedSvc_Call {
	return &console_client_GetManagedSvc_Call{Call: _e.mock.On("GetManagedSvc",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *console_client_GetManagedSvc_Call) Run(run func(ctx context.Context, in *console.MSvcIn, opts ...grpc.CallOption)) *console_client_GetManagedSvc_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*console.MSvcIn), variadicArgs...)
	})
	return _c
}

func (_c *console_client_GetManagedSvc_Call) Return(_a0 *console.MSvcOut, _a1 error) *console_client_GetManagedSvc_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *console_client_GetManagedSvc_Call) RunAndReturn(run func(context.Context, *console.MSvcIn, ...grpc.CallOption) (*console.MSvcOut, error)) *console_client_GetManagedSvc_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectName provides a mock function with given fields: ctx, in, opts
func (_m *console_client) GetProjectName(ctx context.Context, in *console.ProjectIn, opts ...grpc.CallOption) (*console.ProjectOut, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *console.ProjectOut
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *console.ProjectIn, ...grpc.CallOption) (*console.ProjectOut, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *console.ProjectIn, ...grpc.CallOption) *console.ProjectOut); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*console.ProjectOut)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *console.ProjectIn, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// console_client_GetProjectName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectName'
type console_client_GetProjectName_Call struct {
	*mock.Call
}

// GetProjectName is a helper method to define mock.On call
//   - ctx context.Context
//   - in *console.ProjectIn
//   - opts ...grpc.CallOption
func (_e *console_client_Expecter) GetProjectName(ctx interface{}, in interface{}, opts ...interface{}) *console_client_GetProjectName_Call {
	return &console_client_GetProjectName_Call{Call: _e.mock.On("GetProjectName",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *console_client_GetProjectName_Call) Run(run func(ctx context.Context, in *console.ProjectIn, opts ...grpc.CallOption)) *console_client_GetProjectName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*console.ProjectIn), variadicArgs...)
	})
	return _c
}

func (_c *console_client_GetProjectName_Call) Return(_a0 *console.ProjectOut, _a1 error) *console_client_GetProjectName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *console_client_GetProjectName_Call) RunAndReturn(run func(context.Context, *console.ProjectIn, ...grpc.CallOption) (*console.ProjectOut, error)) *console_client_GetProjectName_Call {
	_c.Call.Return(run)
	return _c
}

// SetupAccount provides a mock function with given fields: ctx, in, opts
func (_m *console_client) SetupAccount(ctx context.Context, in *console.AccountSetupIn, opts ...grpc.CallOption) (*console.AccountSetupVoid, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *console.AccountSetupVoid
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *console.AccountSetupIn, ...grpc.CallOption) (*console.AccountSetupVoid, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *console.AccountSetupIn, ...grpc.CallOption) *console.AccountSetupVoid); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*console.AccountSetupVoid)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *console.AccountSetupIn, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// console_client_SetupAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetupAccount'
type console_client_SetupAccount_Call struct {
	*mock.Call
}

// SetupAccount is a helper method to define mock.On call
//   - ctx context.Context
//   - in *console.AccountSetupIn
//   - opts ...grpc.CallOption
func (_e *console_client_Expecter) SetupAccount(ctx interface{}, in interface{}, opts ...interface{}) *console_client_SetupAccount_Call {
	return &console_client_SetupAccount_Call{Call: _e.mock.On("SetupAccount",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *console_client_SetupAccount_Call) Run(run func(ctx context.Context, in *console.AccountSetupIn, opts ...grpc.CallOption)) *console_client_SetupAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*console.AccountSetupIn), variadicArgs...)
	})
	return _c
}

func (_c *console_client_SetupAccount_Call) Return(_a0 *console.AccountSetupVoid, _a1 error) *console_client_SetupAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *console_client_SetupAccount_Call) RunAndReturn(run func(context.Context, *console.AccountSetupIn, ...grpc.CallOption) (*console.AccountSetupVoid, error)) *console_client_SetupAccount_Call {
	_c.Call.Return(run)
	return _c
}

// newConsole_client creates a new instance of console_client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newConsole_client(t interface {
	mock.TestingT
	Cleanup(func())
}) *console_client {
	mock := &console_client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}