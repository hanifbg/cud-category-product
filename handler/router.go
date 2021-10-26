package handler

import (
	"github.com/hanifbg/cud-category-product/handler/cart"
	"github.com/hanifbg/cud-category-product/handler/category"
	"github.com/hanifbg/cud-category-product/handler/checkout"
	"github.com/hanifbg/cud-category-product/handler/product"
	"github.com/hanifbg/cud-category-product/middleware"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, categoryHandler *category.Handler, productHandler *product.Handler, cartHandler *cart.Handler, checkoutHandler *checkout.Handler) {

	userV1 := e.Group("v1")
	userV1.POST("/category", categoryHandler.AddCategoryHandler)
	userV1.PUT("/category/:id", categoryHandler.UpdateCategory)
	userV1.POST("/product", productHandler.AddProductHandler)
	userV1.PUT("/product/:id", productHandler.UpdateProductHandler)

	cobaV1 := e.Group("v1")
	cobaV1.Use(middleware.JWTMiddleware())
	cobaV1.GET("/", categoryHandler.AuthUser)
	cobaV1.POST("/cart", cartHandler.CartHandler)
	cobaV1.POST("/checkout", checkoutHandler.CreateCheckoutHandler)
}
