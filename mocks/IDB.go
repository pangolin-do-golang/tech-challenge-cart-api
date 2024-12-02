// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	sql "database/sql"

	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// IDB is an autogenerated mock type for the IDB type
type IDB struct {
	mock.Mock
}

// Count provides a mock function with given fields: count
func (_m *IDB) Count(count *int64) *gorm.DB {
	ret := _m.Called(count)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(*int64) *gorm.DB); ok {
		r0 = rf(count)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Create provides a mock function with given fields: value
func (_m *IDB) Create(value interface{}) *gorm.DB {
	ret := _m.Called(value)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}) *gorm.DB); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// CreateInBatches provides a mock function with given fields: value, batchSize
func (_m *IDB) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	ret := _m.Called(value, batchSize)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, int) *gorm.DB); ok {
		r0 = rf(value, batchSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: value, conds
func (_m *IDB) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(value, conds...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Exec provides a mock function with given fields: _a0, values
func (_m *IDB) Exec(_a0 string, values ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(string, ...interface{}) *gorm.DB); ok {
		r0 = rf(_a0, values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Find provides a mock function with given fields: dest, conds
func (_m *IDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, dest)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(dest, conds...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// First provides a mock function with given fields: dest, conds
func (_m *IDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, dest)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(dest, conds...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Last provides a mock function with given fields: dest, conds
func (_m *IDB) Last(dest interface{}, conds ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, dest)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(dest, conds...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Limit provides a mock function with given fields: limit
func (_m *IDB) Limit(limit int) *gorm.DB {
	ret := _m.Called(limit)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(int) *gorm.DB); ok {
		r0 = rf(limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Model provides a mock function with given fields: value
func (_m *IDB) Model(value interface{}) *gorm.DB {
	ret := _m.Called(value)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}) *gorm.DB); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Order provides a mock function with given fields: value
func (_m *IDB) Order(value interface{}) *gorm.DB {
	ret := _m.Called(value)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}) *gorm.DB); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Row provides a mock function with given fields:
func (_m *IDB) Row() *sql.Row {
	ret := _m.Called()

	var r0 *sql.Row
	if rf, ok := ret.Get(0).(func() *sql.Row); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Row)
		}
	}

	return r0
}

// Rows provides a mock function with given fields:
func (_m *IDB) Rows() (*sql.Rows, error) {
	ret := _m.Called()

	var r0 *sql.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func() (*sql.Rows, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *sql.Rows); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: value
func (_m *IDB) Save(value interface{}) *gorm.DB {
	ret := _m.Called(value)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}) *gorm.DB); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Scan provides a mock function with given fields: dest
func (_m *IDB) Scan(dest interface{}) *gorm.DB {
	ret := _m.Called(dest)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}) *gorm.DB); ok {
		r0 = rf(dest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Select provides a mock function with given fields: query, args
func (_m *IDB) Select(query interface{}, args ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Take provides a mock function with given fields: dest, conds
func (_m *IDB) Take(dest interface{}, conds ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, dest)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(dest, conds...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Update provides a mock function with given fields: column, value
func (_m *IDB) Update(column string, value interface{}) *gorm.DB {
	ret := _m.Called(column, value)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(string, interface{}) *gorm.DB); ok {
		r0 = rf(column, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Where provides a mock function with given fields: query, args
func (_m *IDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

type mockConstructorTestingTNewIDB interface {
	mock.TestingT
	Cleanup(func())
}

// NewIDB creates a new instance of IDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIDB(t mockConstructorTestingTNewIDB) *IDB {
	mock := &IDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}