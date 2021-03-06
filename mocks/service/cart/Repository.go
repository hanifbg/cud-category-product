// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	cart "github.com/hanifbg/cud-category-product/service/cart"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateCart provides a mock function with given fields: _a0
func (_m *Repository) CreateCart(_a0 cart.Cart) (*cart.Cart, error) {
	ret := _m.Called(_a0)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(cart.Cart) *cart.Cart); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cart.Cart) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCartProduct provides a mock function with given fields: _a0
func (_m *Repository) CreateCartProduct(_a0 cart.CartProduct) (*cart.CartProduct, error) {
	ret := _m.Called(_a0)

	var r0 *cart.CartProduct
	if rf, ok := ret.Get(0).(func(cart.CartProduct) *cart.CartProduct); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.CartProduct)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cart.CartProduct) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindCartByUserID provides a mock function with given fields: _a0
func (_m *Repository) FindCartByUserID(_a0 int) (*cart.Cart, error) {
	ret := _m.Called(_a0)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(int) *cart.Cart); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindCartProduct provides a mock function with given fields: _a0, _a1
func (_m *Repository) FindCartProduct(_a0 int, _a1 int) (*cart.CartProduct, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *cart.CartProduct
	if rf, ok := ret.Get(0).(func(int, int) *cart.CartProduct); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.CartProduct)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindCartProductByCart provides a mock function with given fields: _a0
func (_m *Repository) FindCartProductByCart(_a0 int) ([]cart.CartProduct, error) {
	ret := _m.Called(_a0)

	var r0 []cart.CartProduct
	if rf, ok := ret.Get(0).(func(int) []cart.CartProduct); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cart.CartProduct)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SumPrice provides a mock function with given fields: _a0
func (_m *Repository) SumPrice(_a0 int) int {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// UpdateCart provides a mock function with given fields: _a0
func (_m *Repository) UpdateCart(_a0 cart.Cart) (*cart.Cart, error) {
	ret := _m.Called(_a0)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(cart.Cart) *cart.Cart); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cart.Cart) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCartProduct provides a mock function with given fields: _a0
func (_m *Repository) UpdateCartProduct(_a0 cart.CartProduct) (*cart.CartProduct, error) {
	ret := _m.Called(_a0)

	var r0 *cart.CartProduct
	if rf, ok := ret.Get(0).(func(cart.CartProduct) *cart.CartProduct); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.CartProduct)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cart.CartProduct) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
