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
