package request

import "github.com/hanifbg/cud-category-product/service/product"

type AddProductRequest struct {
	Name       string `json:"name"`
	CategoryId int    `json:"category_id"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	Image      string `json:"image"`
	Detail     string `json:"detail"`
}

func (req *AddProductRequest) ConvertToProductData() *product.CreateProductData {
	var data product.CreateProductData

	data.CategoryId = req.CategoryId
	data.Name = req.Name
	data.Price = req.Price
	data.Stock = req.Stock
	data.Image = req.Image
	data.Detail = req.Detail

	return &data
}
