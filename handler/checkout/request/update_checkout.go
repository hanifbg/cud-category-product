package request

import (
	"github.com/hanifbg/cud-category-product/service/checkout"
)

type UpdateCheckoutRequest struct {
	IsActive       bool `json:"is_active"`
	StatusPayment  bool `json:"status_payment"`
	StatusDelivery bool `json:"status_delivery"`
	StatusOverall  bool `json:"status_overall"`
	DeletedAt      bool `json:"deleted_at"`
}

func (req *UpdateCheckoutRequest) ConvertToUpdateCheckoutData() *checkout.UpdateCheckout {
	var data checkout.UpdateCheckout

	data.NewActive = req.IsActive
	data.NewStatusPay = req.StatusPayment
	data.NewStatusDeli = req.StatusDelivery
	data.NewStatusOver = req.StatusOverall

	return &data
}
