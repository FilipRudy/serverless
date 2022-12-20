// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	"github.com/kyma-project/kyma/components/telemetry-operator/apis/telemetry/v1alpha1"
	"github.com/stretchr/testify/mock"
)

// OutputValidator is an autogenerated mock type for the OutputValidator type
type OutputValidator struct {
	mock.Mock
}

// Validate provides a mock function with given fields: logPipeline
func (_m *OutputValidator) Validate(logPipeline *v1alpha1.LogPipeline) error {
	ret := _m.Called(logPipeline)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.LogPipeline) error); ok {
		r0 = rf(logPipeline)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewOutputValidator interface {
	mock.TestingT
	Cleanup(func())
}

// NewOutputValidator creates a new instance of OutputValidator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOutputValidator(t mockConstructorTestingTNewOutputValidator) *OutputValidator {
	mock := &OutputValidator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
