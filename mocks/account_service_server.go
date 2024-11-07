// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	account "github.com/fajaramaulana/shared-proto-payment/proto/account"

	mock "github.com/stretchr/testify/mock"
)

// AccountServiceServer is an autogenerated mock type for the AccountServiceServer type
type AccountServiceServer struct {
	mock.Mock
}

// GetAccount provides a mock function with given fields: _a0, _a1
func (_m *AccountServiceServer) GetAccount(_a0 context.Context, _a1 *account.AccountRequest) (*account.AccountResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetAccount")
	}

	var r0 *account.AccountResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *account.AccountRequest) (*account.AccountResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *account.AccountRequest) *account.AccountResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*account.AccountResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *account.AccountRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountByUser provides a mock function with given fields: _a0, _a1
func (_m *AccountServiceServer) GetAccountByUser(_a0 context.Context, _a1 *account.AccountRequestByUser) (*account.AccountResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetAccountByUser")
	}

	var r0 *account.AccountResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *account.AccountRequestByUser) (*account.AccountResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *account.AccountRequestByUser) *account.AccountResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*account.AccountResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *account.AccountRequestByUser) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedAccountServiceServer provides a mock function with given fields:
func (_m *AccountServiceServer) mustEmbedUnimplementedAccountServiceServer() {
	_m.Called()
}

// NewAccountServiceServer creates a new instance of AccountServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountServiceServer {
	mock := &AccountServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}