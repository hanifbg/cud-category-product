package product

import (
	"github.com/hanifbg/cud-category-product/common"
	"github.com/hanifbg/cud-category-product/handler/product/request"
	"github.com/hanifbg/cud-category-product/service/product"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service product.Service
}

func NewHandler(service product.Service) *Handler {
	return &Handler{
		service,
	}
}

func (handler *Handler) AddProductHandler(c echo.Context) error {
	addProductReq := new(request.AddCategoryRequest)

	if err := c.Bind(addProductReq); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := handler.service.AddProduct(*addProductReq.ConvertToProductData())
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
