package cart

import (
	"github.com/golang-jwt/jwt"
	"github.com/hanifbg/cud-category-product/common"
	"github.com/hanifbg/cud-category-product/handler/cart/request"
	"github.com/hanifbg/cud-category-product/service/cart"
	echo "github.com/labstack/echo/v4"
)

type Handler struct {
	service cart.Service
}

func NewHandler(service cart.Service) *Handler {
	return &Handler{
		service,
	}
}

func (handler *Handler) CartHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims) //conver to jwt.MapClaims
	userID, ok := claims["id"]
	if !ok {
		return c.JSON(common.NewForbiddenResponse())
	}

	addCartReq := new(request.AddCartRequest)
	if err := c.Bind(addCartReq); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := handler.service.CreateCart(int(userID.(float64)), *addCartReq.ConvertToCartData())
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}
	return c.JSON(common.NewSuccessResponseWithoutData())
}
