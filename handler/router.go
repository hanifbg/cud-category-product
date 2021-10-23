package handler

import (
	user "github.com/hanifbg/cud-category-product/handler/category"
	"github.com/hanifbg/cud-category-product/middleware"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, categoryHandler *user.Handler) {

	userV1 := e.Group("v1")
	userV1.POST("/add_category", categoryHandler.UserHandler)

	cobaV1 := e.Group("v1/auth")
	cobaV1.Use(middleware.JWTMiddleware())
	cobaV1.GET("/", categoryHandler.AuthUser)
}
