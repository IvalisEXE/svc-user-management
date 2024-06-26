// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	auth "svc-user-management/lib/auth"

	mock "github.com/stretchr/testify/mock"
)

// JWTManager is an autogenerated mock type for the JWTManager type
type JWTManager struct {
	mock.Mock
}

// Generate provides a mock function with given fields: sessionUser
func (_m *JWTManager) Generate(sessionUser auth.SessionUser) (string, error) {
	ret := _m.Called(sessionUser)

	if len(ret) == 0 {
		panic("no return value specified for Generate")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(auth.SessionUser) (string, error)); ok {
		return rf(sessionUser)
	}
	if rf, ok := ret.Get(0).(func(auth.SessionUser) string); ok {
		r0 = rf(sessionUser)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(auth.SessionUser) error); ok {
		r1 = rf(sessionUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetClaims provides a mock function with given fields: token
func (_m *JWTManager) GetClaims(token string) (interface{}, error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for GetClaims")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (interface{}, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Parse provides a mock function with given fields: _a0
func (_m *JWTManager) Parse(_a0 string) (auth.SessionUser, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Parse")
	}

	var r0 auth.SessionUser
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (auth.SessionUser, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) auth.SessionUser); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(auth.SessionUser)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewJWTManager creates a new instance of JWTManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJWTManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *JWTManager {
	mock := &JWTManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
