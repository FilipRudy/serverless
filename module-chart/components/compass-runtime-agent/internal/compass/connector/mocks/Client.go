// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	externalschema "github.com/kyma-incubator/compass/components/connector/pkg/graphql/externalschema"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Configuration provides a mock function with given fields: ctx, headers
func (_m *Client) Configuration(ctx context.Context, headers map[string]string) (externalschema.Configuration, error) {
	ret := _m.Called(ctx, headers)

	var r0 externalschema.Configuration
	if rf, ok := ret.Get(0).(func(context.Context, map[string]string) externalschema.Configuration); ok {
		r0 = rf(ctx, headers)
	} else {
		r0 = ret.Get(0).(externalschema.Configuration)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, map[string]string) error); ok {
		r1 = rf(ctx, headers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignCSR provides a mock function with given fields: ctx, csr, headers
func (_m *Client) SignCSR(ctx context.Context, csr string, headers map[string]string) (externalschema.CertificationResult, error) {
	ret := _m.Called(ctx, csr, headers)

	var r0 externalschema.CertificationResult
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) externalschema.CertificationResult); ok {
		r0 = rf(ctx, csr, headers)
	} else {
		r0 = ret.Get(0).(externalschema.CertificationResult)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, map[string]string) error); ok {
		r1 = rf(ctx, csr, headers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
