// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import hydroform "github.com/kyma-incubator/compass/components/provisioner/internal/hydroform"
import mock "github.com/stretchr/testify/mock"
import model "github.com/kyma-incubator/compass/components/provisioner/internal/model"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// DeprovisionCluster provides a mock function with given fields: runtimeConfig, secretName
func (_m *Client) DeprovisionCluster(runtimeConfig model.RuntimeConfig, secretName string) error {
	ret := _m.Called(runtimeConfig, secretName)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.RuntimeConfig, string) error); ok {
		r0 = rf(runtimeConfig, secretName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProvisionCluster provides a mock function with given fields: runtimeConfig, secretName
func (_m *Client) ProvisionCluster(runtimeConfig model.RuntimeConfig, secretName string) (hydroform.ClusterInfo, error) {
	ret := _m.Called(runtimeConfig, secretName)

	var r0 hydroform.ClusterInfo
	if rf, ok := ret.Get(0).(func(model.RuntimeConfig, string) hydroform.ClusterInfo); ok {
		r0 = rf(runtimeConfig, secretName)
	} else {
		r0 = ret.Get(0).(hydroform.ClusterInfo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.RuntimeConfig, string) error); ok {
		r1 = rf(runtimeConfig, secretName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
