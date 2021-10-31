package checkout_test

import (
	"os"
	"testing"
	"time"

	cartMock "github.com/hanifbg/cud-category-product/mocks/service/cart"
	checkoutMock "github.com/hanifbg/cud-category-product/mocks/service/checkout"
	serv "github.com/hanifbg/cud-category-product/service"
	"github.com/hanifbg/cud-category-product/service/cart"
	"github.com/hanifbg/cud-category-product/service/checkout"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	id      = 1
	payment = "payment"
	price   = 100
)

var (
	checkoutService checkout.Service
	checkoutRepo    checkoutMock.Repository
	cartRepo        cartMock.Repository

	checkoutData    checkout.CreateCheckoutData
	invalidCheckout checkout.CreateCheckoutData
	cartData        *cart.Cart

	pointData *checkout.Checkout
)

func setup() {
	pointData = &checkout.Checkout{
		Payment: payment,
	}

	checkoutData = checkout.CreateCheckoutData{
		Payment: payment,
	}

	cartData = &cart.Cart{
		ID:         id,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		DeletedAt:  nil,
		UserID:     id,
		TotalPrice: price,
	}

	invalidCheckout = checkout.CreateCheckoutData{}

	checkoutService = checkout.NewService(&checkoutRepo, &cartRepo)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestCreateCheckout(t *testing.T) {
	t.Run("Expect Create Checkout Success", func(t *testing.T) {
		cartRepo.On("FindCartByUserID", mock.AnythingOfType("int")).Return(cartData, nil).Once()
		checkoutRepo.On("CreateCheckout", mock.AnythingOfType("checkout.Checkout")).Return(nil).Once()

		err := checkoutService.CreateCheckout(id, checkoutData)
		assert.Nil(t, err)
	})

	t.Run("Expect Failed Create Checkout Invalid Data", func(t *testing.T) {
		checkoutRepo.On("CreateCheckout", mock.AnythingOfType("checkout.Checkout")).Return(serv.ErrInvalidData).Once()

		err := checkoutService.CreateCheckout(id, invalidCheckout)
		assert.NotNil(t, err)
		assert.Equal(t, err, serv.ErrInvalidData)
	})

	t.Run("Expect Failed Create Checkout Cart Not Found", func(t *testing.T) {
		cartRepo.On("FindCartByUserID", mock.AnythingOfType("int")).Return(nil, serv.ErrNotFound).Once()

		err := checkoutService.CreateCheckout(id, checkoutData)
		assert.NotNil(t, err)
		assert.Equal(t, err, serv.ErrNotFound)
	})
}
