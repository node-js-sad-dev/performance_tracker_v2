package core

import (
	"errors"
)

func SuccessResponse(data interface{}) *ActionFuncResponse {
	return &ActionFuncResponse{
		Status: 200,
		Data:   data,
		Error:  nil,
		//Cookies: nil,
	}
}

//func SuccessResponseWithCookies(data interface{}, accessCookie, refreshCookie string) *ActionFuncResponse {
//	return &ActionFuncResponse{
//		Status:  200,
//		Data:    data,
//		Error:   nil,
//		Cookies: &http.Cookies{Access: accessCookie, Refresh: refreshCookie},
//	}
//}

func CommonErrorResponse(status int, errorMessage string) *ActionFuncResponse {
	return &ActionFuncResponse{
		Status: status,
		Data:   nil,
		Error:  errors.New(errorMessage),
	}
}

func DbErrorResponse(err error) *ActionFuncResponse {
	switch err.Error() {
	case "record not found":
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
