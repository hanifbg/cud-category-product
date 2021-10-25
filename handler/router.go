package handler

import (
	"github.com/hanifbg/cud-category-product/handler/category"
	"github.com/hanifbg/cud-category-product/handler/product"
	"github.com/hanifbg/cud-category-product/middleware"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, categoryHandler *category.Handler, productHandler *product.Handler) {

	userV1 := e.Group("v1")
	userV1.POST("/category", categoryHandler.AddCategoryHandler)
	userV1.PUT("/category/:id", categoryHandler.UpdateCategory)
	userV1.POST("/product", productHandler.AddProductHandler)
	userV1.PUT("/product/:id", productHandler.UpdateProductHandler)

	cobaV1 := e.Group("v1/auth")
	cobaV1.Use(middleware.JWTMiddleware())
	cobaV1.GET("/", categoryHandler.AuthUser)
}
