// Code generated by mockery v2.15.0. DO NOT EDIT.

package scada

import (
	net "net"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// MockProvider is an autogenerated mock type for the Provider type
type MockProvider struct {
	mock.Mock
}

type MockProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockProvider) EXPECT() *MockProvider_Expecter {
	return &MockProvider_Expecter{mock: &_m.Mock}
}

// GetMeta provides a mock function with given fields:
func (_m *MockProvider) GetMeta() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// MockProvider_GetMeta_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMeta'
type MockProvider_GetMeta_Call struct {
	*mock.Call
}

// GetMeta is a helper method to define mock.On call
func (_e *MockProvider_Expecter) GetMeta() *MockProvider_GetMeta_Call {
	return &MockProvider_GetMeta_Call{Call: _e.mock.On("GetMeta")}
}

func (_c *MockProvider_GetMeta_Call) Run(run func()) *MockProvider_GetMeta_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProvider_GetMeta_Call) Return(_a0 map[string]string) *MockProvider_GetMeta_Call {
	_c.Call.Return(_a0)
	return _c
}

// LastError provides a mock function with given fields:
func (_m *MockProvider) LastError() (time.Time, error) {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProvider_LastError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LastError'
type MockProvider_LastError_Call struct {
	*mock.Call
}

// LastError is a helper method to define mock.On call
func (_e *MockProvider_Expecter) LastError() *MockProvider_LastError_Call {
	return &MockProvider_LastError_Call{Call: _e.mock.On("LastError")}
}

func (_c *MockProvider_LastError_Call) Run(run func()) *MockProvider_LastError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProvider_LastError_Call) Return(_a0 time.Time, _a1 error) *MockProvider_LastError_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Listen provides a mock function with given fields: capability
func (_m *MockProvider) Listen(capability string) (net.Listener, error) {
	ret := _m.Called(capability)

	var r0 net.Listener
	if rf, ok := ret.Get(0).(func(string) net.Listener); ok {
		r0 = rf(capability)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.Listener)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(capability)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProvider_Listen_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Listen'
type MockProvider_Listen_Call struct {
	*mock.Call
}

// Listen is a helper method to define mock.On call
//   - capability string
func (_e *MockProvider_Expecter) Listen(capability interface{}) *MockProvider_Listen_Call {
	return &MockProvider_Listen_Call{Call: _e.mock.On("Listen", capability)}
}

func (_c *MockProvider_Listen_Call) Run(run func(capability string)) *MockProvider_Listen_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockProvider_Listen_Call) Return(_a0 net.Listener, _a1 error) *MockProvider_Listen_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// SessionStatus provides a mock function with given fields:
func (_m *MockProvider) SessionStatus() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockProvider_SessionStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SessionStatus'
type MockProvider_SessionStatus_Call struct {
	*mock.Call
}

// SessionStatus is a helper method to define mock.On call
func (_e *MockProvider_Expecter) SessionStatus() *MockProvider_SessionStatus_Call {
	return &MockProvider_SessionStatus_Call{Call: _e.mock.On("SessionStatus")}
}

func (_c *MockProvider_SessionStatus_Call) Run(run func()) *MockProvider_SessionStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProvider_SessionStatus_Call) Return(_a0 string) *MockProvider_SessionStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

// Start provides a mock function with given fields:
func (_m *MockProvider) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockProvider_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockProvider_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *MockProvider_Expecter) Start() *MockProvider_Start_Call {
	return &MockProvider_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *MockProvider_Start_Call) Run(run func()) *MockProvider_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProvider_Start_Call) Return(_a0 error) *MockProvider_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MockProvider) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockProvider_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockProvider_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockProvider_Expecter) Stop() *MockProvider_Stop_Call {
	return &MockProvider_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockProvider_Stop_Call) Run(run func()) *MockProvider_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProvider_Stop_Call) Return(_a0 error) *MockProvider_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

// UpdateMeta provides a mock function with given fields: _a0
func (_m *MockProvider) UpdateMeta(_a0 map[string]string) {
	_m.Called(_a0)
}

// MockProvider_UpdateMeta_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateMeta'
type MockProvider_UpdateMeta_Call struct {
	*mock.Call
}

// UpdateMeta is a helper method to define mock.On call
//   - _a0 map[string]string
func (_e *MockProvider_Expecter) UpdateMeta(_a0 interface{}) *MockProvider_UpdateMeta_Call {
	return &MockProvider_UpdateMeta_Call{Call: _e.mock.On("UpdateMeta", _a0)}
}

func (_c *MockProvider_UpdateMeta_Call) Run(run func(_a0 map[string]string)) *MockProvider_UpdateMeta_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]string))
	})
	return _c
}

func (_c *MockProvider_UpdateMeta_Call) Return() *MockProvider_UpdateMeta_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewMockProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockProvider creates a new instance of MockProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockProvider(t mockConstructorTestingTNewMockProvider) *MockProvider {
	mock := &MockProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
