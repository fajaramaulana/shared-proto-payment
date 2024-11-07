// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	auth "github.com/fajaramaulana/shared-proto-payment/proto/auth"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// AuthServiceClient is an autogenerated mock type for the AuthServiceClient type
type AuthServiceClient struct {
	mock.Mock
}

// LoginUser provides a mock function with given fields: ctx, in, opts
func (_m *AuthServiceClient) LoginUser(ctx context.Context, in *auth.LoginRequest, opts ...grpc.CallOption) (*auth.LoginResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for LoginUser")
	}

	var r0 *auth.LoginResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *auth.LoginRequest, ...grpc.CallOption) (*auth.LoginResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *auth.LoginRequest, ...grpc.CallOption) *auth.LoginResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.LoginResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *auth.LoginRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RefreshToken provides a mock function with given fields: ctx, in, opts
func (_m *AuthServiceClient) RefreshToken(ctx context.Context, in *auth.RefreshTokenRequest, opts ...grpc.CallOption) (*auth.RefreshTokenResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RefreshToken")
	}

	var r0 *auth.RefreshTokenResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *auth.RefreshTokenRequest, ...grpc.CallOption) (*auth.RefreshTokenResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *auth.RefreshTokenRequest, ...grpc.CallOption) *auth.RefreshTokenResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.RefreshTokenResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *auth.RefreshTokenRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: ctx, in, opts
func (_m *AuthServiceClient) RegisterUser(ctx context.Context, in *auth.RegisterRequest, opts ...grpc.CallOption) (*auth.RegisterResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RegisterUser")
	}

	var r0 *auth.RegisterResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *auth.RegisterRequest, ...grpc.CallOption) (*auth.RegisterResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *auth.RegisterRequest, ...grpc.CallOption) *auth.RegisterResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.RegisterResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *auth.RegisterRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthServiceClient creates a new instance of AuthServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthServiceClient {
	mock := &AuthServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}