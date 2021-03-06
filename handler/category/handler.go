package category

import (
	"strconv"

	"github.com/hanifbg/cud-category-product/common"
	"github.com/hanifbg/cud-category-product/handler/category/request"
	"github.com/hanifbg/cud-category-product/service/category"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
)

type Handler struct {
	service category.Service
}

func NewHandler(service category.Service) *Handler {
	return &Handler{
		service,
	}
}

func (handler *Handler) AddCategoryHandler(c echo.Context) error {
	addCategoryReq := new(request.AddCategoryRequest)

	if err := c.Bind(addCategoryReq); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := handler.service.AddCategory(*addCategoryReq.ConvertToCategoryData())
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (handler *Handler) UpdateCategory(c echo.Context) error {
	var newActive bool = true
	id, _ := strconv.Atoi(c.Param("id"))

	updateCategoryRequest := new(request.UpdateCategoryRequest)

	if err := c.Bind(updateCategoryRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	if updateCategoryRequest.IsActive == 0 {
		newActive = false
	}

	err := handler.service.UpdateCategory(id, updateCategoryRequest.Name, newActive)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
func (handler *Handler) AuthUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims) //conver to jwt.MapClaims

	userID, ok := claims["id"]
	if !ok {
		return c.JSON(common.NewForbiddenResponse())
	}

	return c.JSON(common.NewSuccessResponse(userID))
}
