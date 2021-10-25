package common

import (
	"net/http"

	"github.com/hanifbg/cud-category-product/service"
)

type errorServiceResponse string

const (
	errInternalServerError errorServiceResponse = "internal_server_error"
	errHasBeenModified     errorServiceResponse = "data_has_been modified"
	errNotFound            errorServiceResponse = "data_not_found"
	errInvalidSpec         errorServiceResponse = "invalid_spec"
	errEmptyStock          errorServiceResponse = "empty_stock"
)

//BusinessResponse default payload response
type ServiceResponse struct {
	Code    errorServiceResponse `json:"code"`
	Message string               `json:"message"`
	Data    interface{}          `json:"data"`
}

//NewErrorBusinessResponse Response return choosen http status like 400 bad request 422 unprocessable entity, ETC, based on responseCode
func NewErrorBusinessResponse(err error) (int, ServiceResponse) {
	return errorMapping(err)
}

//errorMapping error for missing header key with given value
func errorMapping(err error) (int, ServiceResponse) {
	switch err {
	default:
		return newInternalServerErrorResponse()
	case service.ErrNotFound:
		return newNotFoundResponse()
	case service.ErrInvalidData:
		return newValidationResponse(err.Error())
	case service.ErrHasBeenModified:
		return newHasBeedModifiedResponse()
	case service.ErrEmptyStock:
		return newEmptyStockResponse(err.Error())
	}
}

func newEmptyStockResponse(message string) (int, ServiceResponse) {
	return http.StatusBadRequest, ServiceResponse{
		errEmptyStock,
		message,
		map[string]interface{}{},
	}
}

//newInternalServerErrorResponse default internal server error response
func newInternalServerErrorResponse() (int, ServiceResponse) {
	return http.StatusInternalServerError, ServiceResponse{
		errInternalServerError,
		"Internal server error",
		map[string]interface{}{},
	}
}

//newHasBeedModifiedResponse failed to validate request payload
func newHasBeedModifiedResponse() (int, ServiceResponse) {
	return http.StatusBadRequest, ServiceResponse{
		errHasBeenModified,
		"Data has been modified",
		map[string]interface{}{},
	}
}

//newNotFoundResponse default not found error response
func newNotFoundResponse() (int, ServiceResponse) {
	return http.StatusNotFound, ServiceResponse{
		errNotFound,
		"Data Not found",
		map[string]interface{}{},
	}
}

//newValidationResponse failed to validate request payload
func newValidationResponse(message string) (int, ServiceResponse) {
	return http.StatusBadRequest, ServiceResponse{
		errInvalidSpec,
		"Validation failed " + message,
		map[string]interface{}{},
	}
}
