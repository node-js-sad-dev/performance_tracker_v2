package http

import (
	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type Sort struct {
	SortBy    string `json:"sortBy"`
	SortOrder string `json:"sortOrder"`
}

type SuccessOperationResponse struct {
	Message string `json:"message" default:"Operation completed successfully"`
}

type Cookies struct {
	Access  string
	Refresh string
}

type ActionFuncResponse struct {
	Status  int
	Data    interface{}
	Error   error
	Cookies *Cookies
}

type ActionFunc func(params ...interface{}) *ActionFuncResponse

type ParamFunc func(*gin.Context) (interface{}, error)
