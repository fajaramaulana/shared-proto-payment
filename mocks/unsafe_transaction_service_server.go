// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeTransactionServiceServer is an autogenerated mock type for the UnsafeTransactionServiceServer type
type UnsafeTransactionServiceServer struct {
	mock.Mock
}

// mustEmbedUnimplementedTransactionServiceServer provides a mock function with given fields:
func (_m *UnsafeTransactionServiceServer) mustEmbedUnimplementedTransactionServiceServer() {
	_m.Called()
}

// NewUnsafeTransactionServiceServer creates a new instance of UnsafeTransactionServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUnsafeTransactionServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *UnsafeTransactionServiceServer {
	mock := &UnsafeTransactionServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
