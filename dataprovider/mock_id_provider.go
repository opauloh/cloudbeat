// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by mockery v2.24.0. DO NOT EDIT.

package dataprovider

import mock "github.com/stretchr/testify/mock"

// MockIdProvider is an autogenerated mock type for the IdProvider type
type MockIdProvider struct {
	mock.Mock
}

type MockIdProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIdProvider) EXPECT() *MockIdProvider_Expecter {
	return &MockIdProvider_Expecter{mock: &_m.Mock}
}

// GetId provides a mock function with given fields: resourceType, resourceId
func (_m *MockIdProvider) GetId(resourceType string, resourceId string) string {
	ret := _m.Called(resourceType, resourceId)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(resourceType, resourceId)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockIdProvider_GetId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetId'
type MockIdProvider_GetId_Call struct {
	*mock.Call
}

// GetId is a helper method to define mock.On call
//   - resourceType string
//   - resourceId string
func (_e *MockIdProvider_Expecter) GetId(resourceType interface{}, resourceId interface{}) *MockIdProvider_GetId_Call {
	return &MockIdProvider_GetId_Call{Call: _e.mock.On("GetId", resourceType, resourceId)}
}

func (_c *MockIdProvider_GetId_Call) Run(run func(resourceType string, resourceId string)) *MockIdProvider_GetId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockIdProvider_GetId_Call) Return(_a0 string) *MockIdProvider_GetId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIdProvider_GetId_Call) RunAndReturn(run func(string, string) string) *MockIdProvider_GetId_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockIdProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIdProvider creates a new instance of MockIdProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIdProvider(t mockConstructorTestingTNewMockIdProvider) *MockIdProvider {
	mock := &MockIdProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}