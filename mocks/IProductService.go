// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	uuid "github.com/google/uuid"
	product "github.com/pangolin-do-golang/tech-challenge-cart-api/internal/core/product"
	mock "github.com/stretchr/testify/mock"
)

// IProductService is an autogenerated mock type for the IProductService type
type IProductService struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *IProductService) Delete(id uuid.UUID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *IProductService) GetByID(id uuid.UUID) (*product.Product, error) {
	ret := _m.Called(id)

	var r0 *product.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*product.Product, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *product.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*product.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: search, category
func (_m *IProductService) Search(search string, category string) (*[]product.Product, error) {
	ret := _m.Called(search, category)

	var r0 *[]product.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*[]product.Product, error)); ok {
		return rf(search, category)
	}
	if rf, ok := ret.Get(0).(func(string, string) *[]product.Product); ok {
		r0 = rf(search, category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]product.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(search, category)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIProductService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIProductService creates a new instance of IProductService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIProductService(t mockConstructorTestingTNewIProductService) *IProductService {
	mock := &IProductService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
