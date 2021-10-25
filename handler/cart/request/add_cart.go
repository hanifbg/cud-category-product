package request

import "github.com/hanifbg/cud-category-product/service/cart"

type AddCartRequest struct {
	ProductID int `json:"product_id"`
	Count     int `json:"count"`
	Price     int `json:"price"`
}

func (req *AddCartRequest) ConvertToCartData() *cart.CreateCartData {
	var data cart.CreateCartData

	data.ProductID = req.ProductID
	data.Count = req.Count
	data.Price = req.Price

	return &data
}
