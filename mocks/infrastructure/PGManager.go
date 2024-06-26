// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"
	databases "svc-user-management/infrastructure/databases"

	mock "github.com/stretchr/testify/mock"

	postgresql "svc-user-management/lib/databases"

	sqlx "github.com/jmoiron/sqlx"
)

// PGManager is an autogenerated mock type for the PGManager type
type PGManager struct {
	mock.Mock
}

// Architecture provides a mock function with given fields:
func (_m *PGManager) Architecture() *postgresql.Architecture {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Architecture")
	}

	var r0 *postgresql.Architecture
	if rf, ok := ret.Get(0).(func() *postgresql.Architecture); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*postgresql.Architecture)
		}
	}

	return r0
}

// GetMaster provides a mock function with given fields: ctx, dbTx
func (_m *PGManager) GetMaster(ctx context.Context, dbTx *sqlx.Tx) databases.GetContext {
	ret := _m.Called(ctx, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetMaster")
	}

	var r0 databases.GetContext
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.Tx) databases.GetContext); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(databases.GetContext)
		}
	}

	return r0
}

// GetTransaction provides a mock function with given fields: dbTx
func (_m *PGManager) GetTransaction(dbTx *sqlx.Tx) databases.ExecContext {
	ret := _m.Called(dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetTransaction")
	}

	var r0 databases.ExecContext
	if rf, ok := ret.Get(0).(func(*sqlx.Tx) databases.ExecContext); ok {
		r0 = rf(dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(databases.ExecContext)
		}
	}

	return r0
}

// SelectMaster provides a mock function with given fields: ctx, dbTx
func (_m *PGManager) SelectMaster(ctx context.Context, dbTx *sqlx.Tx) databases.SelectContext {
	ret := _m.Called(ctx, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for SelectMaster")
	}

	var r0 databases.SelectContext
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.Tx) databases.SelectContext); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(databases.SelectContext)
		}
	}

	return r0
}

// StartTransaction provides a mock function with given fields: ctx
func (_m *PGManager) StartTransaction(ctx context.Context) (*sqlx.Tx, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for StartTransaction")
	}

	var r0 *sqlx.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*sqlx.Tx, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *sqlx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPGManager creates a new instance of PGManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPGManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *PGManager {
	mock := &PGManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
