package core

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(data interface{}) *ActionFuncResponse {
	return &ActionFuncResponse{
		Status: 200,
		Data:   data,
		Error:  nil,
	}
}

func DefaultSuccessResponse() *ActionFuncResponse {
	return &ActionFuncResponse{
		Status: 200,
		Data: gin.H{
			"message": "success",
		},
		Error: nil,
	}
}

func CommonErrorResponse(status int, errorMessage string) *ActionFuncResponse {
	return &ActionFuncResponse{
		Status: status,
		Data:   nil,
		Error:  errors.New(errorMessage),
	}
}

func NotFoundErrorResponse(entityName string) *ActionFuncResponse {
	return &ActionFuncResponse{
		Status: 404,
		Data:   nil,
		Error:  errors.New(entityName + " not found"),
	}
}

func DbErrorResponse(err error) *ActionFuncResponse {
	switch err.Error() {
	case "no rows in result set":
		return &ActionFuncResponse{
			Status: 404,
			Data:   nil,
			Error:  errors.New("record not found"),
		}
	default:
		println(err.Error())
		return &ActionFuncResponse{
			Status: 500,
			Data:   nil,
			Error:  errors.New("internal server error"),
		}
	}
}
