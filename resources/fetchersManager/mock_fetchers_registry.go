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

package fetchersManager

import (
	context "context"

	fetching "github.com/elastic/cloudbeat/resources/fetching"
	mock "github.com/stretchr/testify/mock"

	uniqueness "github.com/elastic/cloudbeat/uniqueness"
)

// MockFetchersRegistry is an autogenerated mock type for the FetchersRegistry type
type MockFetchersRegistry struct {
	mock.Mock
}

type MockFetchersRegistry_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFetchersRegistry) EXPECT() *MockFetchersRegistry_Expecter {
	return &MockFetchersRegistry_Expecter{mock: &_m.Mock}
}

// Keys provides a mock function with given fields:
func (_m *MockFetchersRegistry) Keys() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockFetchersRegistry_Keys_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Keys'
type MockFetchersRegistry_Keys_Call struct {
	*mock.Call
}

// Keys is a helper method to define mock.On call
func (_e *MockFetchersRegistry_Expecter) Keys() *MockFetchersRegistry_Keys_Call {
	return &MockFetchersRegistry_Keys_Call{Call: _e.mock.On("Keys")}
}

func (_c *MockFetchersRegistry_Keys_Call) Run(run func()) *MockFetchersRegistry_Keys_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFetchersRegistry_Keys_Call) Return(_a0 []string) *MockFetchersRegistry_Keys_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFetchersRegistry_Keys_Call) RunAndReturn(run func() []string) *MockFetchersRegistry_Keys_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields: fetcherName, f, c
func (_m *MockFetchersRegistry) Register(fetcherName string, f fetching.Fetcher, c ...fetching.Condition) error {
	_va := make([]interface{}, len(c))
	for _i := range c {
		_va[_i] = c[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, fetcherName, f)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, fetching.Fetcher, ...fetching.Condition) error); ok {
		r0 = rf(fetcherName, f, c...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFetchersRegistry_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type MockFetchersRegistry_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - fetcherName string
//   - f fetching.Fetcher
//   - c ...fetching.Condition
func (_e *MockFetchersRegistry_Expecter) Register(fetcherName interface{}, f interface{}, c ...interface{}) *MockFetchersRegistry_Register_Call {
	return &MockFetchersRegistry_Register_Call{Call: _e.mock.On("Register",
		append([]interface{}{fetcherName, f}, c...)...)}
}

func (_c *MockFetchersRegistry_Register_Call) Run(run func(fetcherName string, f fetching.Fetcher, c ...fetching.Condition)) *MockFetchersRegistry_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]fetching.Condition, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(fetching.Condition)
			}
		}
		run(args[0].(string), args[1].(fetching.Fetcher), variadicArgs...)
	})
	return _c
}

func (_c *MockFetchersRegistry_Register_Call) Return(_a0 error) *MockFetchersRegistry_Register_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFetchersRegistry_Register_Call) RunAndReturn(run func(string, fetching.Fetcher, ...fetching.Condition) error) *MockFetchersRegistry_Register_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterFetchers provides a mock function with given fields: fetchers, le
func (_m *MockFetchersRegistry) RegisterFetchers(fetchers []*ParsedFetcher, le uniqueness.Manager) error {
	ret := _m.Called(fetchers, le)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*ParsedFetcher, uniqueness.Manager) error); ok {
		r0 = rf(fetchers, le)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFetchersRegistry_RegisterFetchers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterFetchers'
type MockFetchersRegistry_RegisterFetchers_Call struct {
	*mock.Call
}

// RegisterFetchers is a helper method to define mock.On call
//   - fetchers []*ParsedFetcher
//   - le uniqueness.Manager
func (_e *MockFetchersRegistry_Expecter) RegisterFetchers(fetchers interface{}, le interface{}) *MockFetchersRegistry_RegisterFetchers_Call {
	return &MockFetchersRegistry_RegisterFetchers_Call{Call: _e.mock.On("RegisterFetchers", fetchers, le)}
}

func (_c *MockFetchersRegistry_RegisterFetchers_Call) Run(run func(fetchers []*ParsedFetcher, le uniqueness.Manager)) *MockFetchersRegistry_RegisterFetchers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]*ParsedFetcher), args[1].(uniqueness.Manager))
	})
	return _c
}

func (_c *MockFetchersRegistry_RegisterFetchers_Call) Return(_a0 error) *MockFetchersRegistry_RegisterFetchers_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFetchersRegistry_RegisterFetchers_Call) RunAndReturn(run func([]*ParsedFetcher, uniqueness.Manager) error) *MockFetchersRegistry_RegisterFetchers_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields: ctx, key, metadata
func (_m *MockFetchersRegistry) Run(ctx context.Context, key string, metadata fetching.CycleMetadata) error {
	ret := _m.Called(ctx, key, metadata)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, fetching.CycleMetadata) error); ok {
		r0 = rf(ctx, key, metadata)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFetchersRegistry_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockFetchersRegistry_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - metadata fetching.CycleMetadata
func (_e *MockFetchersRegistry_Expecter) Run(ctx interface{}, key interface{}, metadata interface{}) *MockFetchersRegistry_Run_Call {
	return &MockFetchersRegistry_Run_Call{Call: _e.mock.On("Run", ctx, key, metadata)}
}

func (_c *MockFetchersRegistry_Run_Call) Run(run func(ctx context.Context, key string, metadata fetching.CycleMetadata)) *MockFetchersRegistry_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(fetching.CycleMetadata))
	})
	return _c
}

func (_c *MockFetchersRegistry_Run_Call) Return(_a0 error) *MockFetchersRegistry_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFetchersRegistry_Run_Call) RunAndReturn(run func(context.Context, string, fetching.CycleMetadata) error) *MockFetchersRegistry_Run_Call {
	_c.Call.Return(run)
	return _c
}

// ShouldRun provides a mock function with given fields: key
func (_m *MockFetchersRegistry) ShouldRun(key string) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockFetchersRegistry_ShouldRun_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ShouldRun'
type MockFetchersRegistry_ShouldRun_Call struct {
	*mock.Call
}

// ShouldRun is a helper method to define mock.On call
//   - key string
func (_e *MockFetchersRegistry_Expecter) ShouldRun(key interface{}) *MockFetchersRegistry_ShouldRun_Call {
	return &MockFetchersRegistry_ShouldRun_Call{Call: _e.mock.On("ShouldRun", key)}
}

func (_c *MockFetchersRegistry_ShouldRun_Call) Run(run func(key string)) *MockFetchersRegistry_ShouldRun_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockFetchersRegistry_ShouldRun_Call) Return(_a0 bool) *MockFetchersRegistry_ShouldRun_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFetchersRegistry_ShouldRun_Call) RunAndReturn(run func(string) bool) *MockFetchersRegistry_ShouldRun_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MockFetchersRegistry) Stop() {
	_m.Called()
}

// MockFetchersRegistry_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockFetchersRegistry_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockFetchersRegistry_Expecter) Stop() *MockFetchersRegistry_Stop_Call {
	return &MockFetchersRegistry_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockFetchersRegistry_Stop_Call) Run(run func()) *MockFetchersRegistry_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFetchersRegistry_Stop_Call) Return() *MockFetchersRegistry_Stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockFetchersRegistry_Stop_Call) RunAndReturn(run func()) *MockFetchersRegistry_Stop_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockFetchersRegistry interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockFetchersRegistry creates a new instance of MockFetchersRegistry. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockFetchersRegistry(t mockConstructorTestingTNewMockFetchersRegistry) *MockFetchersRegistry {
	mock := &MockFetchersRegistry{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
