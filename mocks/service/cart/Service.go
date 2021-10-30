// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	cart "github.com/hanifbg/cud-category-product/service/cart"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateCart provides a mock function with given fields: _a0, _a1
func (_m *Service) CreateCart(_a0 int, _a1 cart.CreateCartData) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, cart.CreateCartData) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindCartByUserID provides a mock function with given fields: userID
func (_m *Service) FindCartByUserID(userID int) ([]cart.CartProduct, error) {
	ret := _m.Called(userID)

	var r0 []cart.CartProduct
	if rf, ok := ret.Get(0).(func(int) []cart.CartProduct); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cart.CartProduct)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}