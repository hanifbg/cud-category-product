package checkout

import (
	"time"

	serv "github.com/hanifbg/cud-category-product/service"
	cartServ "github.com/hanifbg/cud-category-product/service/cart"
	"github.com/hanifbg/cud-category-product/util/validator"
)

type service struct {
	repository Repository
	cartRepo   cartServ.Repository
}

func NewService(repository Repository, cartRepo cartServ.Repository) Service {
	return &service{
		repository,
		cartRepo,
	}
}

type CreateCheckoutData struct {
	Payment string `validate:"required"`
}

type UpdateCheckout struct {
	NewActive     bool
	NewStatusPay  bool
	NewStatusDeli bool
	NewStatusOver bool
	DeletedAt     *time.Time
}

func (s *service) CreateCheckout(userId int, data CreateCheckoutData) error {
	err := validator.GetValidator().Struct(data)
	if err != nil {
		return serv.ErrInvalidData
	}

	cartData, err := s.cartRepo.FindCartByUserID(userId)
	if err != nil {
		return serv.ErrNotFound
	}

	checkout := NewCheckout(
		int(cartData.ID),
		userId,
		cartData.TotalPrice,
		data.Payment,
		true,
		false,
		false,
		false,
		time.Now(),
		time.Now(),
	)

	err = s.repository.CreateCheckout(checkout)
	if err != nil {
		return err
	}
	return nil
}
