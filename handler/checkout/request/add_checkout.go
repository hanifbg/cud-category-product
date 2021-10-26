package request

import "github.com/hanifbg/cud-category-product/service/checkout"

type AddCheckoutRequest struct {
	Payment string `json:"payment"`
}

func (req *AddCheckoutRequest) ConvertToCheckoutData() *checkout.CreateCheckoutData {
	var data checkout.CreateCheckoutData

	data.Payment = req.Payment

	return &data
}
