// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	dberrors "github.com/kyma-incubator/compass/components/provisioner/internal/persistence/dberrors"

	mock "github.com/stretchr/testify/mock"

	model "github.com/kyma-incubator/compass/components/provisioner/internal/model"

	time "time"
)

// ReadWriteSession is an autogenerated mock type for the ReadWriteSession type
type ReadWriteSession struct {
	mock.Mock
}

// DeleteCluster provides a mock function with given fields: runtimeID
func (_m *ReadWriteSession) DeleteCluster(runtimeID string) dberrors.Error {
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

// FixShootProvisioningStage provides a mock function with given fields: message, newStage, transitionTime
func (_m *ReadWriteSession) FixShootProvisioningStage(message string, newStage model.OperationStage, transitionTime time.Time) dberrors.Error {
	ret := _m.Called(message, newStage, transitionTime)

	var r0 dberrors.Error
	if rf, ok := ret.Get(0).(func(string, model.OperationStage, time.Time) dberrors.Error); ok {
		r0 = rf(message, newStage, transitionTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dberrors.Error)
		}
	}

	return r0
}

// GetCluster provides a mock function with given fields: runtimeID
func (_m *ReadWriteSession) GetCluster(runtimeID string) (model.Cluster, dberrors.Error) {
	ret := _m.Called(runtimeID)

	var r0 model.Cluster
	if rf, ok := ret.Get(0).(func(string) model.Cluster); ok {
		r0 = rf(runtimeID)
	} else {
		r0 = ret.Get(0).(model.Cluster)
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

// GetGardenerClusterByName provides a mock function with given fields: name
func (_m *ReadWriteSession) GetGardenerClusterByName(name string) (model.Cluster, dberrors.Error) {
	ret := _m.Called(name)

	var r0 model.Cluster
	if rf, ok := ret.Get(0).(func(string) model.Cluster); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(model.Cluster)
	}

	var r1 dberrors.Error
	if rf, ok := ret.Get(1).(func(string) dberrors.Error); ok {
		r1 = rf(name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(dberrors.Error)
		}
	}

	return r0, r1
}

// GetLastOperation provides a mock function with given fields: runtimeID
func (_m *ReadWriteSession) GetLastOperation(runtimeID string) (model.Operation, dberrors.Error) {
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

// GetOperation provides a mock function with given fields: operationID
func (_m *ReadWriteSession) GetOperation(operationID string) (model.Operation, dberrors.Error) {
	ret := _m.Called(operationID)

	var r0 model.Operation
	if rf, ok := ret.Get(0).(func(string) model.Operation); ok {
		r0 = rf(operationID)
	} else {
		r0 = ret.Get(0).(model.Operation)
	}

	var r1 dberrors.Error
	if rf, ok := ret.Get(1).(func(string) dberrors.Error); ok {
		r1 = rf(operationID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(dberrors.Error)
		}
	}

	return r0, r1
}

// GetRuntimeUpgrade provides a mock function with given fields: operationId
func (_m *ReadWriteSession) GetRuntimeUpgrade(operationId string) (model.RuntimeUpgrade, dberrors.Error) {
	ret := _m.Called(operationId)

	var r0 model.RuntimeUpgrade
	if rf, ok := ret.Get(0).(func(string) model.RuntimeUpgrade); ok {
		r0 = rf(operationId)
	} else {
		r0 = ret.Get(0).(model.RuntimeUpgrade)
	}

	var r1 dberrors.Error
	if rf, ok := ret.Get(1).(func(string) dberrors.Error); ok {
		r1 = rf(operationId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(dberrors.Error)
		}
	}

	return r0, r1
}

// GetTenant provides a mock function with given fields: runtimeID
func (_m *ReadWriteSession) GetTenant(runtimeID string) (string, dberrors.Error) {
	ret := _m.Called(runtimeID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(runtimeID)
	} else {
		r0 = ret.Get(0).(string)
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

// GetTenantForOperation provides a mock function with given fields: operationID
func (_m *ReadWriteSession) GetTenantForOperation(operationID string) (string, dberrors.Error) {
	ret := _m.Called(operationID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(operationID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 dberrors.Error
	if rf, ok := ret.Get(1).(func(string) dberrors.Error); ok {
		r1 = rf(operationID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(dberrors.Error)
		}
	}

	return r0, r1
}

// InsertCluster provides a mock function with given fields: cluster
func (_m *ReadWriteSession) InsertCluster(cluster model.Cluster) dberrors.Error {
	ret := _m.Called(cluster)

	var r0 dberrors.Error
	if rf, ok := ret.Get(0).(func(model.Cluster) dberrors.Error); ok {
		r0 = rf(cluster)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dberrors.Error)
		}
	}

	return r0
}

// InsertGCPConfig provides a mock function with given fields: config
func (_m *ReadWriteSession) InsertGCPConfig(config model.GCPConfig) dberrors.Error {
	ret := _m.Called(config)

	var r0 dberrors.Error
	if rf, ok := ret.Get(0).(func(model.GCPConfig) dberrors.Error); ok {
		r0 = rf(config)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dberrors.Error)
		}
	}

	return r0
}

// InsertGardenerConfig provides a mock function with given fields: config
func (_m *ReadWriteSession) InsertGardenerConfig(config model.GardenerConfig) dberrors.Error {
	ret := _m.Called(config)

	var r0 dberrors.Error
	if rf, ok := ret.Get(0).(func(model.GardenerConfig) dberrors.Error); ok {
		r0 = rf(config)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dberrors.Error)
		}
	}

	return r0
}

// InsertKymaConfig provides a mock function with given fields: kymaConfig
func (_m *ReadWriteSession) InsertKymaConfig(kymaConfig model.KymaConfig) dberrors.Error {
	ret := _m.Called(kymaConfig)

	var r0 dberrors.Error
	if rf, ok := ret.Get(0).(func(model.KymaConfig) dberrors.Error); ok {
		r0 = rf(kymaConfig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dberrors.Error)
		}
	}

	return r0
}

// InsertOperation provides a mock function with given fields: operation
func (_m *ReadWriteSession) InsertOperation(operation model.Operation) dberrors.Error {
	ret := _m.Called(operation)

	var r0 dberrors.Error
	if rf, ok := ret.Get(0).(func(model.Operation) dberrors.Error); ok {
		r0 = rf(operation)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dberrors.Error)
		}
	}

	return r0
}

// InsertRuntimeUpgrade provides a mock function with given fields: runtimeUpgrade
func (_m *ReadWriteSession) InsertRuntimeUpgrade(runtimeUpgrade model.RuntimeUpgrade) dberrors.Error {
	ret := _m.Called(runtimeUpgrade)

	var r0 dberrors.Error
	if rf, ok := ret.Get(0).(func(model.RuntimeUpgrade) dberrors.Error); ok {
		r0 = rf(runtimeUpgrade)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dberrors.Error)
		}
	}

	return r0
}

// ListInProgressOperations provides a mock function with given fields:
func (_m *ReadWriteSession) ListInProgressOperations() ([]model.Operation, dberrors.Error) {
	ret := _m.Called()

	var r0 []model.Operation
	if rf, ok := ret.Get(0).(func() []model.Operation); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Operation)
		}
	}

	var r1 dberrors.Error
	if rf, ok := ret.Get(1).(func() dberrors.Error); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(dberrors.Error)
		}
	}

	return r0, r1
}

// MarkClusterAsDeleted provides a mock function with given fields: runtimeID
func (_m *ReadWriteSession) MarkClusterAsDeleted(runtimeID string) dberrors.Error {
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

// SetActiveKymaConfig provides a mock function with given fields: runtimeID, kymaConfigId
func (_m *ReadWriteSession) SetActiveKymaConfig(runtimeID string, kymaConfigId string) dberrors.Error {
	ret := _m.Called(runtimeID, kymaConfigId)

	var r0 dberrors.Error
	if rf, ok := ret.Get(0).(func(string, string) dberrors.Error); ok {
		r0 = rf(runtimeID, kymaConfigId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dberrors.Error)
		}
	}

	return r0
}

// TransitionOperation provides a mock function with given fields: operationID, message, stage, transitionTime
func (_m *ReadWriteSession) TransitionOperation(operationID string, message string, stage model.OperationStage, transitionTime time.Time) dberrors.Error {
	ret := _m.Called(operationID, message, stage, transitionTime)

	var r0 dberrors.Error
	if rf, ok := ret.Get(0).(func(string, string, model.OperationStage, time.Time) dberrors.Error); ok {
		r0 = rf(operationID, message, stage, transitionTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dberrors.Error)
		}
	}

	return r0
}

// UpdateCluster provides a mock function with given fields: runtimeID, kubeconfig, terraformState
func (_m *ReadWriteSession) UpdateCluster(runtimeID string, kubeconfig string, terraformState []byte) dberrors.Error {
	ret := _m.Called(runtimeID, kubeconfig, terraformState)

	var r0 dberrors.Error
	if rf, ok := ret.Get(0).(func(string, string, []byte) dberrors.Error); ok {
		r0 = rf(runtimeID, kubeconfig, terraformState)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dberrors.Error)
		}
	}

	return r0
}

// UpdateOperationState provides a mock function with given fields: operationID, message, state, endTime
func (_m *ReadWriteSession) UpdateOperationState(operationID string, message string, state model.OperationState, endTime time.Time) dberrors.Error {
	ret := _m.Called(operationID, message, state, endTime)

	var r0 dberrors.Error
	if rf, ok := ret.Get(0).(func(string, string, model.OperationState, time.Time) dberrors.Error); ok {
		r0 = rf(operationID, message, state, endTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dberrors.Error)
		}
	}

	return r0
}

// UpdateUpgradeState provides a mock function with given fields: operationID, upgradeState
func (_m *ReadWriteSession) UpdateUpgradeState(operationID string, upgradeState model.UpgradeState) dberrors.Error {
	ret := _m.Called(operationID, upgradeState)

	var r0 dberrors.Error
	if rf, ok := ret.Get(0).(func(string, model.UpgradeState) dberrors.Error); ok {
		r0 = rf(operationID, upgradeState)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dberrors.Error)
		}
	}

	return r0
}
