package handler

import (
	"github.com/hanifbg/cud-category-product/handler/user"
	"github.com/hanifbg/cud-category-product/middleware"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, userHandler *user.Handler) {

	userV1 := e.Group("v1")
	userV1.POST("/", userHandler.UserHandler)

	cobaV1 := e.Group("v1/auth")
	cobaV1.Use(middleware.JWTMiddleware())
	cobaV1.GET("/", userHandler.AuthUser)
}
