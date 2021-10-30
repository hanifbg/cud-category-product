package category_test

import (
	"os"
	"testing"

	categoryMock "github.com/hanifbg/cud-category-product/mocks/service/category"
	serv "github.com/hanifbg/cud-category-product/service"
	"github.com/hanifbg/cud-category-product/service/category"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	id   = 1
	name = "category"
)

var (
	categoryService category.Service
	categoryRepo    categoryMock.Repository

	categoryData        category.CreateCategoryData
	invalidCategoryData category.CreateCategoryData

	updateData category.CreateCategoryData
	pointData  *category.Category
)

func setup() {
	pointData = &category.Category{
		Name: name,
	}

	categoryData = category.CreateCategoryData{
		Name: name,
	}

	invalidCategoryData = category.CreateCategoryData{
		Name: "",
	}

	categoryService = category.NewService(&categoryRepo)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestAddCategory(t *testing.T) {
	t.Run("Expect Add Category Success", func(t *testing.T) {
		categoryRepo.On("AddCategory", mock.AnythingOfType("category.Category")).Return(nil).Once()

		err := categoryService.AddCategory(categoryData)
		assert.Nil(t, err)
	})

	t.Run("Expect Failed Add Product Invalid Data", func(t *testing.T) {
		categoryRepo.On("AddCategory", mock.AnythingOfType("category.Category")).Return(serv.ErrInvalidData).Once()

		err := categoryService.AddCategory(invalidCategoryData)
		assert.NotNil(t, err)
		assert.Equal(t, err, serv.ErrInvalidData)
	})
}

func TestUpdateCategory(t *testing.T) {
	t.Run("Expect Update Product Success", func(t *testing.T) {
		categoryRepo.On("FindCategoryById", mock.AnythingOfType("int")).Return(pointData, nil).Once()
		categoryRepo.On("UpdateCategory", mock.AnythingOfType("category.Category")).Return(nil).Once()

		err := categoryService.UpdateCategory(id, name, true)
		assert.Nil(t, err)
	})
}
