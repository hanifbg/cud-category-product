package cart

import (
	"fmt"
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

type CreateCartData struct {
	ProductID int `validate:"required"`
	Count     int `validate:"required"`
	Price     int `validate:"required"`
}

func (s *service) CreateCart(userID int, data CreateCartData) error {
	err := validator.GetValidator().Struct(data)
	if err != nil {
		return serv.ErrInvalidData
	}

	totalPrice := data.Count * data.Price

	//cart
	cartData, _ := s.repository.FindCartByUserID(userID)

	if cartData == nil {
		newCartData := NewCart(userID, totalPrice, time.Now(), time.Now()) //create cart if notfound
		cartData, err = s.repository.CreateCart(newCartData)
		if err != nil {
			return err
		}
	}
	//cartProduct
	cartProductData, _ := s.repository.FindCartProduct(int(cartData.ID), data.ProductID)
	if cartProductData == nil {
		newCartProductData := NewCartProduct(int(cartData.ID), data.ProductID, data.Count, data.Count*data.Price, time.Now(), time.Now())
		cartProductData, err = s.repository.CreateCartProduct(newCartProductData)
		if err != nil {
			return err
		}
	} else {
		updateCartProductData := cartProductData.ModifyCartProduct(data.Count*data.Price, data.Count, time.Now())
		cartProductData, err = s.repository.UpdateCartProduct(updateCartProductData)
		if err != nil {
			return err
		}
	}
	// update total price
	newTotalPrice := s.repository.SumPrice(int(cartData.ID), data.ProductID)
	updateCartData := cartData.ModifyCart(newTotalPrice, time.Now())
	cartData, err = s.repository.UpdateCart(updateCartData)
	if err != nil {
		return err
	}

	fmt.Println(newTotalPrice)
	fmt.Println(cartProductData)
	fmt.Println(cartData)

	return nil
}
