package category

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

type CreateCategoryData struct {
	Name string `validate:"required"`
}

func (s *service) AddCategory(data CreateCategoryData) error {
	err := validator.GetValidator().Struct(data)
	if err != nil {
		return serv.ErrInvalidData
	}

	category := NewCategory(
		data.Name,
		time.Now(),
		time.Now(),
	)

	err = s.repository.AddCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateCategory(id int, name string, isActive bool) error {

	category, err := s.repository.FindCategoryById(id)
	if err != nil {
		return serv.ErrInvalidData
	} else if category == nil {
		return serv.ErrNotFound
	}

	modifCategory := category.ModifyCategory(name, isActive, time.Now())

	return s.repository.UpdateCategory(modifCategory)
}
