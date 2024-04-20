// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"
	models "svc-user-management/modules/user/models"

	mock "github.com/stretchr/testify/mock"
)

// UserUsecases is an autogenerated mock type for the UserUsecases type
type UserUsecases struct {
	mock.Mock
}

// CheckUserExistByPhoneNo provides a mock function with given fields: ctx, phoneNo
func (_m *UserUsecases) CheckUserExistByPhoneNo(ctx context.Context, phoneNo string) (bool, error) {
	ret := _m.Called(ctx, phoneNo)

	if len(ret) == 0 {
		panic("no return value specified for CheckUserExistByPhoneNo")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, phoneNo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, phoneNo)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, phoneNo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *UserUsecases) CreateUser(ctx context.Context, user models.User) (int64, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.User) (int64, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.User) int64); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: ctx, userID
func (_m *UserUsecases) GetUserByID(ctx context.Context, userID int64) (*models.User, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*models.User, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.User); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserProfileByID provides a mock function with given fields: ctx, userID, phoneNo, fullname, skipPhoneNo
func (_m *UserUsecases) UpdateUserProfileByID(ctx context.Context, userID int64, phoneNo string, fullname string, skipPhoneNo bool) error {
	ret := _m.Called(ctx, userID, phoneNo, fullname, skipPhoneNo)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserProfileByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string, bool) error); ok {
		r0 = rf(ctx, userID, phoneNo, fullname, skipPhoneNo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserUsecases creates a new instance of UserUsecases. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserUsecases(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserUsecases {
	mock := &UserUsecases{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
