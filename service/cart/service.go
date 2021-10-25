package cart

import (
	"fmt"
	"time"

	serv "github.com/hanifbg/cud-category-product/service"
	productServ "github.com/hanifbg/cud-category-product/service/product"
	"github.com/hanifbg/cud-category-product/util/validator"
)

type service struct {
	repository  Repository
	productRepo productServ.Repository
}

func NewService(repository Repository, productRepo productServ.Repository) Service {
	return &service{
		repository,
		productRepo,
	}
}

type CreateCartData struct {
	ProductID int `validate:"required"`
	Count     int `validate:"required"`
}

func (s *service) CreateCart(userID int, data CreateCartData) error {
	err := validator.GetValidator().Struct(data)
	if err != nil {
		return serv.ErrInvalidData
	}

	//cart
	cartData, _ := s.repository.FindCartByUserID(userID)

	if cartData == nil {
		newCartData := NewCart(userID, 0, time.Now(), time.Now()) //create cart if notfound
		cartData, err = s.repository.CreateCart(newCartData)
		if err != nil {
			return err
		}
	}
	//cartProduct
	cartProductData, _ := s.repository.FindCartProduct(int(cartData.ID), data.ProductID)
	productData, err := s.productRepo.FindProductById(data.ProductID)
	if err != nil {
		return err
	}

	if cartProductData == nil {
		newCartProductData := NewCartProduct(int(cartData.ID), data.ProductID, data.Count, data.Count*productData.Price, time.Now(), time.Now())
		cartProductData, err = s.repository.CreateCartProduct(newCartProductData)
		if err != nil {
			return err
		}
	} else {
		updateCartProductData := cartProductData.ModifyCartProduct(data.Count*productData.Price, data.Count, time.Now())
		cartProductData, err = s.repository.UpdateCartProduct(updateCartProductData)
		if err != nil {
			return err
		}
	}
	// update total price
	newTotalPrice := s.repository.SumPrice(int(cartData.ID))
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
