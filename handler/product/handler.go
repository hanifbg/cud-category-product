package product

import (
	"strconv"

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
	addProductReq := new(request.AddProductRequest)

	if err := c.Bind(addProductReq); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := handler.service.AddProduct(*addProductReq.ConvertToProductData())
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (handler *Handler) UpdateProductHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	updateProductReq := new(request.UpdateProductRequest)

	if err := c.Bind(updateProductReq); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := handler.service.UpdateProduct(id, *updateProductReq.ConvertToUpdateProductData())
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
