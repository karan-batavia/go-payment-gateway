// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// ApiRequestInterface is an autogenerated mock type for the ApiRequestInterface type
type ApiRequestInterface struct {
	mock.Mock
}

// Call provides a mock function with given fields: ctx, httpMethod, url, header, body, result
func (_m *ApiRequestInterface) Call(ctx context.Context, httpMethod string, url string, header http.Header, body interface{}, result interface{}) error {
	ret := _m.Called(ctx, httpMethod, url, header, body, result)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, http.Header, interface{}, interface{}) error); ok {
		r0 = rf(ctx, httpMethod, url, header, body, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DoRequest provides a mock function with given fields: req, result
func (_m *ApiRequestInterface) DoRequest(req *http.Request, result interface{}) error {
	ret := _m.Called(req, result)

	var r0 error
	if rf, ok := ret.Get(0).(func(*http.Request, interface{}) error); ok {
		r0 = rf(req, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewApiRequestInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewApiRequestInterface creates a new instance of ApiRequestInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewApiRequestInterface(t mockConstructorTestingTNewApiRequestInterface) *ApiRequestInterface {
	mock := &ApiRequestInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
