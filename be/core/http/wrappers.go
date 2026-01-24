package http

import "fmt"

func WrapForOneParameter[T any](fn func(*T) *ActionFuncResponse) ActionFunc {
	return func(params ...interface{}) *ActionFuncResponse {
		if len(params) != 1 {
			return &ActionFuncResponse{
				Status: 400,
				Error:  fmt.Errorf("expected 1 parameter of type *%T, got %d", new(T), len(params)),
			}
		}

		param, ok := params[0].(*T)
		if !ok || param == nil {
			return &ActionFuncResponse{
				Status: 400,
				Error:  fmt.Errorf("parameter is not of type *%T", new(T)),
			}
		}

		return fn(param)
	}
}
