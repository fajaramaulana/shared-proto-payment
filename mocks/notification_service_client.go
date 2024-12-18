// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	notification "github.com/fajaramaulana/shared-proto-payment/proto/notification"
)

// NotificationServiceClient is an autogenerated mock type for the NotificationServiceClient type
type NotificationServiceClient struct {
	mock.Mock
}

// SendNotification provides a mock function with given fields: ctx, in, opts
func (_m *NotificationServiceClient) SendNotification(ctx context.Context, in *notification.NotificationRequest, opts ...grpc.CallOption) (*notification.NotificationResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SendNotification")
	}

	var r0 *notification.NotificationResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *notification.NotificationRequest, ...grpc.CallOption) (*notification.NotificationResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *notification.NotificationRequest, ...grpc.CallOption) *notification.NotificationResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*notification.NotificationResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *notification.NotificationRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewNotificationServiceClient creates a new instance of NotificationServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNotificationServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *NotificationServiceClient {
	mock := &NotificationServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
