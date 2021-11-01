package checkout

import (
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/hanifbg/cud-category-product/common"
	"github.com/hanifbg/cud-category-product/handler/checkout/request"
	"github.com/hanifbg/cud-category-product/service/checkout"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service checkout.Service
}

func NewHandler(service checkout.Service) *Handler {
	return &Handler{
		service,
	}
}

func (handler *Handler) CreateCheckoutHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims) //conver to jwt.MapClaims
	userId, ok := claims["id"]
	if !ok {
		return c.JSON(common.NewForbiddenResponse())
	}

	addCheckoutReq := new(request.AddCheckoutRequest)

	if err := c.Bind(addCheckoutReq); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := handler.service.CreateCheckout(int(userId.(float64)), *addCheckoutReq.ConvertToCheckoutData())
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (handler *Handler) UpdateCheckoutHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	updateCheckoutReq := new(request.UpdateCheckoutRequest)

	if err := c.Bind(updateCheckoutReq); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := handler.service.UpdateCheckout(id, *updateCheckoutReq.ConvertToUpdateCheckoutData())
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
