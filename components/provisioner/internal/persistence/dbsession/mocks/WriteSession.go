// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import dberrors "github.com/kyma-incubator/compass/components/provisioner/internal/persistence/dberrors"

import mock "github.com/stretchr/testify/mock"
import model "github.com/kyma-incubator/compass/components/provisioner/internal/model"
import time "time"

// WriteSession is an autogenerated mock type for the WriteSession type
type WriteSession struct {
	mock.Mock
}

// DeleteCluster provides a mock function with given fields: runtimeID
func (_m *WriteSession) DeleteCluster(runtimeID string) dberrors.Error {
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

// InsertCluster provides a mock function with given fields: runtimeID, creationTimestamp, terraformState
func (_m *WriteSession) InsertCluster(runtimeID string, creationTimestamp time.Time, terraformState string) dberrors.Error {
	ret := _m.Called(runtimeID, creationTimestamp, terraformState)

	var r0 dberrors.Error
	if rf, ok := ret.Get(0).(func(string, time.Time, string) dberrors.Error); ok {
		r0 = rf(runtimeID, creationTimestamp, terraformState)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dberrors.Error)
		}
	}

	return r0
}

// InsertGCPConfig provides a mock function with given fields: runtimeID, config
func (_m *WriteSession) InsertGCPConfig(runtimeID string, config model.GCPConfig) dberrors.Error {
	ret := _m.Called(runtimeID, config)

	var r0 dberrors.Error
	if rf, ok := ret.Get(0).(func(string, model.GCPConfig) dberrors.Error); ok {
		r0 = rf(runtimeID, config)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dberrors.Error)
		}
	}

	return r0
}

// InsertGardenerConfig provides a mock function with given fields: runtimeID, config
func (_m *WriteSession) InsertGardenerConfig(runtimeID string, config model.GardenerConfig) dberrors.Error {
	ret := _m.Called(runtimeID, config)

	var r0 dberrors.Error
	if rf, ok := ret.Get(0).(func(string, model.GardenerConfig) dberrors.Error); ok {
		r0 = rf(runtimeID, config)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dberrors.Error)
		}
	}

	return r0
}

// InsertKymaConfig provides a mock function with given fields: runtimeID, version
func (_m *WriteSession) InsertKymaConfig(runtimeID string, version string) (string, dberrors.Error) {
	ret := _m.Called(runtimeID, version)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(runtimeID, version)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 dberrors.Error
	if rf, ok := ret.Get(1).(func(string, string) dberrors.Error); ok {
		r1 = rf(runtimeID, version)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(dberrors.Error)
		}
	}

	return r0, r1
}

// InsertKymaConfigModule provides a mock function with given fields: kymaConfigID, module
func (_m *WriteSession) InsertKymaConfigModule(kymaConfigID string, module model.KymaModule) dberrors.Error {
	ret := _m.Called(kymaConfigID, module)

	var r0 dberrors.Error
	if rf, ok := ret.Get(0).(func(string, model.KymaModule) dberrors.Error); ok {
		r0 = rf(kymaConfigID, module)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dberrors.Error)
		}
	}

	return r0
}

// InsertOperation provides a mock function with given fields: operation
func (_m *WriteSession) InsertOperation(operation model.Operation) (string, dberrors.Error) {
	ret := _m.Called(operation)

	var r0 string
	if rf, ok := ret.Get(0).(func(model.Operation) string); ok {
		r0 = rf(operation)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 dberrors.Error
	if rf, ok := ret.Get(1).(func(model.Operation) dberrors.Error); ok {
		r1 = rf(operation)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(dberrors.Error)
		}
	}

	return r0, r1
}

// UpdateCluster provides a mock function with given fields: runtimeID, kubeconfig, terraformState
func (_m *WriteSession) UpdateCluster(runtimeID string, kubeconfig string, terraformState string) dberrors.Error {
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

// UpdateOperationState provides a mock function with given fields: operationID, message, state
func (_m *WriteSession) UpdateOperationState(operationID string, message string, state model.OperationState) dberrors.Error {
	ret := _m.Called(operationID, message, state)

	var r0 dberrors.Error
	if rf, ok := ret.Get(0).(func(string, string, model.OperationState) dberrors.Error); ok {
		r0 = rf(operationID, message, state)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dberrors.Error)
		}
	}

	return r0
}
