// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	dberrors "github.com/kyma-incubator/compass/components/provisioner/internal/persistence/dberrors"
	mock "github.com/stretchr/testify/mock"

	model "github.com/kyma-incubator/compass/components/provisioner/internal/model"
)

// RuntimeService is an autogenerated mock type for the RuntimeService type
type RuntimeService struct {
	mock.Mock
}

// CleanupData provides a mock function with given fields: runtimeID
func (_m *RuntimeService) CleanupData(runtimeID string) dberrors.Error {
	ret := _m.Called(runtimeID)

	var r0 dberrors.Error
	if rf, ok := ret.Get(0).(func(string) dberrors.Error); ok {
		r0 = rf(runtimeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dberrors.Error)
		}
	}

	return r0
}

// GetLastOperation provides a mock function with given fields: runtimeID
func (_m *RuntimeService) GetLastOperation(runtimeID string) (model.Operation, dberrors.Error) {
	ret := _m.Called(runtimeID)

	var r0 model.Operation
	if rf, ok := ret.Get(0).(func(string) model.Operation); ok {
		r0 = rf(runtimeID)
	} else {
		r0 = ret.Get(0).(model.Operation)
	}

	var r1 dberrors.Error
	if rf, ok := ret.Get(1).(func(string) dberrors.Error); ok {
		r1 = rf(runtimeID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(dberrors.Error)
		}
	}

	return r0, r1
}

// GetStatus provides a mock function with given fields: runtimeID
func (_m *RuntimeService) GetStatus(runtimeID string) (model.RuntimeStatus, dberrors.Error) {
	ret := _m.Called(runtimeID)

	var r0 model.RuntimeStatus
	if rf, ok := ret.Get(0).(func(string) model.RuntimeStatus); ok {
		r0 = rf(runtimeID)
	} else {
		r0 = ret.Get(0).(model.RuntimeStatus)
	}

	var r1 dberrors.Error
	if rf, ok := ret.Get(1).(func(string) dberrors.Error); ok {
		r1 = rf(runtimeID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(dberrors.Error)
		}
	}

	return r0, r1
}

// SetDeprovisioningStarted provides a mock function with given fields: runtimeID
func (_m *RuntimeService) SetDeprovisioningStarted(runtimeID string) (model.Operation, dberrors.Error) {
	ret := _m.Called(runtimeID)

	var r0 model.Operation
	if rf, ok := ret.Get(0).(func(string) model.Operation); ok {
		r0 = rf(runtimeID)
	} else {
		r0 = ret.Get(0).(model.Operation)
	}

	var r1 dberrors.Error
	if rf, ok := ret.Get(1).(func(string) dberrors.Error); ok {
		r1 = rf(runtimeID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(dberrors.Error)
		}
	}

	return r0, r1
}

// SetProvisioningStarted provides a mock function with given fields: runtimeID, runtimeConfig
func (_m *RuntimeService) SetProvisioningStarted(runtimeID string, runtimeConfig model.RuntimeConfig) (model.Operation, dberrors.Error) {
	ret := _m.Called(runtimeID, runtimeConfig)

	var r0 model.Operation
	if rf, ok := ret.Get(0).(func(string, model.RuntimeConfig) model.Operation); ok {
		r0 = rf(runtimeID, runtimeConfig)
	} else {
		r0 = ret.Get(0).(model.Operation)
	}

	var r1 dberrors.Error
	if rf, ok := ret.Get(1).(func(string, model.RuntimeConfig) dberrors.Error); ok {
		r1 = rf(runtimeID, runtimeConfig)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(dberrors.Error)
		}
	}

	return r0, r1
}

// SetUpgradeStarted provides a mock function with given fields: runtimeID
func (_m *RuntimeService) SetUpgradeStarted(runtimeID string) (model.Operation, dberrors.Error) {
	ret := _m.Called(runtimeID)

	var r0 model.Operation
	if rf, ok := ret.Get(0).(func(string) model.Operation); ok {
		r0 = rf(runtimeID)
	} else {
		r0 = ret.Get(0).(model.Operation)
	}

	var r1 dberrors.Error
	if rf, ok := ret.Get(1).(func(string) dberrors.Error); ok {
		r1 = rf(runtimeID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(dberrors.Error)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: runtimeID, kubeconfig, terraformState
func (_m *RuntimeService) Update(runtimeID string, kubeconfig string, terraformState string) dberrors.Error {
	ret := _m.Called(runtimeID, kubeconfig, terraformState)

	var r0 dberrors.Error
	if rf, ok := ret.Get(0).(func(string, string, string) dberrors.Error); ok {
		r0 = rf(runtimeID, kubeconfig, terraformState)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dberrors.Error)
		}
	}

	return r0
}
