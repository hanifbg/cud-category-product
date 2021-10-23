package request

import "github.com/hanifbg/cud-category-product/service/category"

type AddCategoryRequest struct {
	Name string `json:"name"`
}

func (req *AddCategoryRequest) ConvertToCategoryData() *category.CreateCategoryData {
	var data category.CreateCategoryData

	data.Name = req.Name

	return &data
}
