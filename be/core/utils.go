package core

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

func setCookies(context *gin.Context, accessToken string, refreshToken string) {
	//Context.SetCookie("access", accessToken, config.AccessTokenLife, "/", config.Domain, config.CookieIsSecure, true)
	//Context.SetCookie("refresh", refreshToken, config.RefreshTokenLife, "/", config.Domain, config.CookieIsSecure, true)
}

// StructToMap have a strong feeling it is not the best approach but good idea to test
func StructToMap[Result any](input any) map[string]Result {
	result := make(map[string]Result)

	if input == nil {
		return result
	}

	val := reflect.ValueOf(input)

	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return result
		}
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return result
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldVal := val.Field(i)

		if !field.IsExported() {
			continue
		}

		key := field.Tag.Get("json")

		if key == "-" {
			continue
		}

		if key == "" {
			key = field.Name
		}

		result[key] = fieldVal.Interface()
	}

	return result
}
