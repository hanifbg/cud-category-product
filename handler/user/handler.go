package user

import (
	"net/http"

	"GO_MOD_MODULE_NAME/service/user"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
)

type Handler struct {
	service user.Service
}

func NewHandler(service user.Service) *Handler {
	return &Handler{
		service,
	}
}

func (handler *Handler) UserHandler(c echo.Context) error {
	err := handler.service.ServiceFuncForUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "ok")
}

func (handler *Handler) AuthUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims) //conver to jwt.MapClaims

	userID, ok := claims["id"]
	if !ok {
		return c.JSON(http.StatusForbidden, ok)
	}

	return c.JSON(http.StatusOK, userID)
}
