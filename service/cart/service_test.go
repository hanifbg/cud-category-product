package cart_test

import (
	"os"
	"testing"
	"time"

	serv "github.com/hanifbg/cud-category-product/service"
	"github.com/hanifbg/cud-category-product/service/cart"
	productServ "github.com/hanifbg/cud-category-product/service/product"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	cartMock "github.com/hanifbg/cud-category-product/mocks/service/cart"
	productMock "github.com/hanifbg/cud-category-product/mocks/service/product"
)

const (
	id        = 1
	productId = 1
	count     = 2
	price     = 10
	name      = "name"
	stock     = 1
	image     = "image"
	detail    = "detail"
)

var (
	cartService cart.Service
	cartRepo    cartMock.Repository
	productRepo productMock.Repository

	cartData         cart.CreateCartData
	sliceCartProduct []cart.CartProduct
	emptySliceCart   []cart.CartProduct
	invalidCart      cart.CreateCartData
	cartProductData  cart.CartProduct

	pointCartData    *cart.Cart
	pointCartProduct *cart.CartProduct
	pointProduct     *productServ.Product
	emptyStock       *productServ.Product
)

func setup() {
	emptySliceCart = []cart.CartProduct{}
	sliceCartProduct = []cart.CartProduct{
		{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
			CartID:    id,
			ProductID: productId,
			Count:     count,
			Price:     price,
		},
		{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
			CartID:    id,
			ProductID: productId,
			Count:     count,
			Price:     price,
		},
	}

	pointProduct = &productServ.Product{
		ID:         id,
		CategoryID: id,
		Name:       name,
		Price:      price,
		Stock:      stock,
		Image:      image,
		Detail:     detail,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		DeletedAt:  nil,
	}

	emptyStock = &productServ.Product{
		ID:         id,
		CategoryID: id,
		Name:       name,
		Price:      price,
		Stock:      0,
		Image:      image,
		Detail:     detail,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		DeletedAt:  nil,
	}

	pointCartData = &cart.Cart{
		ID:         id,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		DeletedAt:  nil,
		UserID:     id,
		TotalPrice: price,
	}

	pointCartProduct = &cart.CartProduct{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
		CartID:    id,
		ProductID: productId,
		Count:     count,
		Price:     price,
	}

	cartData = cart.CreateCartData{
		ProductID: productId,
		Count:     count,
	}

	invalidCart = cart.CreateCartData{}

	cartProductData = cart.CartProduct{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
		CartID:    id,
		ProductID: productId,
		Count:     count,
		Price:     price,
	}

	cartService = cart.NewService(&cartRepo, &productRepo)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestCreateCheckout(t *testing.T) {
	t.Run("Expect Create Checkout Success", func(t *testing.T) {
		cartRepo.On("FindCartByUserID", mock.AnythingOfType("int")).Return(nil, nil).Once()
		cartRepo.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(pointCartData, nil).Once()
		cartRepo.On("FindCartProduct", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(pointCartProduct, nil).Once()
		productRepo.On("FindProductById", mock.AnythingOfType("int")).Return(pointProduct, nil).Once()
		cartRepo.On("CreateCartProduct", mock.AnythingOfType("cart.CartProduct")).Return(pointCartProduct, nil).Once()
		cartRepo.On("SumPrice", mock.AnythingOfType("int")).Return(price).Once()
		cartRepo.On("UpdateCart", mock.AnythingOfType("cart.Cart")).Return(pointCartData, nil).Once()
		cartRepo.On("UpdateCartProduct", mock.AnythingOfType("cart.CartProduct")).Return(pointCartProduct, nil).Once()

		err := cartService.CreateCart(id, cartData)
		assert.Nil(t, err)
	})

	t.Run("Expect Create Checkout Success With Nil CartProduct", func(t *testing.T) {
		cartRepo.On("FindCartByUserID", mock.AnythingOfType("int")).Return(nil, nil).Once()
		cartRepo.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(pointCartData, nil).Once()
		cartRepo.On("FindCartProduct", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil, nil).Once()
		productRepo.On("FindProductById", mock.AnythingOfType("int")).Return(pointProduct, nil).Once()
		cartRepo.On("CreateCartProduct", mock.AnythingOfType("cart.CartProduct")).Return(pointCartProduct, nil).Once()
		cartRepo.On("SumPrice", mock.AnythingOfType("int")).Return(price).Once()
		cartRepo.On("UpdateCart", mock.AnythingOfType("cart.Cart")).Return(pointCartData, nil).Once()
		cartRepo.On("UpdateCartProduct", mock.AnythingOfType("cart.CartProduct")).Return(pointCartProduct, nil).Once()

		err := cartService.CreateCart(id, cartData)
		assert.Nil(t, err)
	})

	t.Run("Expect Create Cart Failed Empty Stock", func(t *testing.T) {
		cartRepo.On("FindCartByUserID", mock.AnythingOfType("int")).Return(nil, nil).Once()
		cartRepo.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(pointCartData, nil).Once()
		cartRepo.On("FindCartProduct", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil, nil).Once()
		productRepo.On("FindProductById", mock.AnythingOfType("int")).Return(emptyStock, nil).Once()

		err := cartService.CreateCart(id, cartData)
		assert.NotNil(t, err)
		assert.Equal(t, err, serv.ErrEmptyStock)
	})

	t.Run("Expect Create Cart Failed Invalid Data", func(t *testing.T) {
		cartRepo.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(nil, serv.ErrInvalidData).Once()

		err := cartService.CreateCart(id, invalidCart)
		assert.NotNil(t, err)
		assert.Equal(t, err, serv.ErrInvalidData)
	})
}

func TestFindCart(t *testing.T) {
	t.Run("Expect Find Cart Success", func(t *testing.T) {
		cartRepo.On("FindCartByUserID", mock.AnythingOfType("int")).Return(pointCartData, nil).Once()
		cartRepo.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(pointCartData, nil).Once()
		cartRepo.On("FindCartProductByCart", mock.AnythingOfType("int")).Return(sliceCartProduct, nil).Once()

		result, err := cartService.FindCartByUserID(id)
		assert.Nil(t, err)
		assert.EqualValues(t, sliceCartProduct, result)
	})

	t.Run("Expect Find Empty Cart", func(t *testing.T) {
		cartRepo.On("FindCartByUserID", mock.AnythingOfType("int")).Return(nil, nil).Once()
		cartRepo.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(pointCartData, serv.ErrInvalidData).Once()

		result, err := cartService.FindCartByUserID(id)
		assert.NotNil(t, err)
		assert.Equal(t, err, serv.ErrInvalidData)
		assert.EqualValues(t, emptySliceCart, result)
	})
}
