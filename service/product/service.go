package product

import (
	"time"

	serv "github.com/hanifbg/cud-category-product/service"
	"github.com/hanifbg/cud-category-product/util/validator"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

type CreateProductData struct {
	CategoryId int    `validate:"required"`
	Name       string `validate:"required"`
	Price      int    `validate:"required"`
	Stock      int    `validate:"required"`
	Image      string `validate:"required"`
	Detail     string `validate:"required"`
}

type UpdateProduct struct {
	NewCategoryId int
	NewName       string
	NewPrice      int
	NewStock      int
	NewImage      string
	NewDetail     string
	DeletedAt     *time.Time
}

func (s *service) AddProduct(data CreateProductData) error {
	err := validator.GetValidator().Struct(data)
	if err != nil {
		return serv.ErrInvalidData
	}

	product := NewProduct(
		data.CategoryId,
		data.Name,
		data.Price,
		data.Stock,
		data.Image,
		data.Detail,
		time.Now(),
		time.Now(),
	)

	err = s.repository.AddProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateProduct(id int, data UpdateProduct) error {
	product, err := s.repository.FindProductById(id)
	if err != nil {
		return serv.ErrInvalidData
	} else if product == nil {
		return serv.ErrNotFound
	}

	modifProduct := product.ModifyProduct(
		data.NewCategoryId,
		data.NewName,
		data.NewPrice,
		data.NewStock,
		data.NewImage,
		data.NewDetail,
		data.DeletedAt,
	)

	return s.repository.UpdateProduct(modifProduct)
}
