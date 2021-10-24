package request

import "github.com/hanifbg/cud-category-product/service/product"

type UpdateProductRequest struct {
	CategoryID int    `json:"category_id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	Image      string `json:"image"`
	Detail     string `json:"detail"`
}

func (req *UpdateProductRequest) ConvertToUpdateProductData() *product.UpdateProduct {
	var data product.UpdateProduct

	data.NewCategoryId = req.CategoryID
	data.NewName = req.Name
	data.NewPrice = req.Price
	data.NewStock = req.Stock
	data.NewImage = req.Image
	data.NewDetail = req.Detail

	return &data
}
