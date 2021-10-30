package product_test

import (
	"os"
	"testing"
	"time"

	productMock "github.com/hanifbg/cud-category-product/mocks/service/product"
	serv "github.com/hanifbg/cud-category-product/service"
	"github.com/hanifbg/cud-category-product/service/product"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	id         = 1
	categoryId = 1
	name       = "category"
	price      = 100
	stock      = 10
	image      = "images"
	detail     = "detail"
)

var (
	productService product.Service
	productRepo    productMock.Repository

	productData        product.CreateProductData
	invalidProductData product.CreateProductData

	updateData product.UpdateProduct
	pointData  *product.Product
)

func setup() {
	pointData = &product.Product{
		ID:         id,
		CategoryID: categoryId,
		Name:       name,
		Price:      price,
		Stock:      stock,
		Image:      image,
		Detail:     detail,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		DeletedAt:  nil,
	}

	productData = product.CreateProductData{
		CategoryId: categoryId,
		Name:       name,
		Price:      price,
		Stock:      stock,
		Image:      image,
		Detail:     detail,
	}

	invalidProductData = product.CreateProductData{
		CategoryId: categoryId,
		Name:       name,
		Price:      price,
		Image:      image,
		Detail:     detail,
	}

	updateData = product.UpdateProduct{
		NewCategoryId: categoryId,
		NewName:       name,
		NewPrice:      price,
		NewStock:      stock,
		NewImage:      image,
		NewDetail:     detail,
		DeletedAt:     nil,
	}

	productService = product.NewService(&productRepo)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestAddProduct(t *testing.T) {
	t.Run("Expect Add Product Success", func(t *testing.T) {
		productRepo.On("AddProduct", mock.AnythingOfType("product.Product")).Return(nil).Once()

		err := productService.AddProduct(productData)
		assert.Nil(t, err)
	})

	t.Run("Expect Failed Add Product Invalid Data", func(t *testing.T) {
		productRepo.On("AddProduct", mock.AnythingOfType("product.Product")).Return(serv.ErrInvalidData).Once()

		err := productService.AddProduct(invalidProductData)
		assert.NotNil(t, err)
		assert.Equal(t, err, serv.ErrInvalidData)
	})
}

func TestUpdateProduct(t *testing.T) {
	t.Run("Expect Update Product Success", func(t *testing.T) {
		productRepo.On("FindProductById", mock.AnythingOfType("int")).Return(pointData, nil).Once()
		productRepo.On("UpdateProduct", mock.AnythingOfType("product.Product")).Return(nil).Once()
		err := productService.UpdateProduct(1, updateData)
		assert.Nil(t, err)
	})

	t.Run("Expect Update Product Failed Invalid Data", func(t *testing.T) {
		productRepo.On("FindProductById", mock.AnythingOfType("int")).Return(nil, serv.ErrInvalidData).Once()

		err := productService.UpdateProduct(1, updateData)
		assert.NotNil(t, err)
		assert.Equal(t, err, serv.ErrInvalidData)
	})

	t.Run("Expect Update Product Failed Product Not Found", func(t *testing.T) {
		productRepo.On("FindProductById", mock.AnythingOfType("int")).Return(nil, nil).Once()

		err := productService.UpdateProduct(99, updateData)
		assert.NotNil(t, err)
		assert.Equal(t, err, serv.ErrNotFound)
	})
}
