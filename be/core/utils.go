package core

import (
	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func setCookies(context *gin.Context, accessToken string, refreshToken string) {
	//Context.SetCookie("access", accessToken, config.AccessTokenLife, "/", config.Domain, config.CookieIsSecure, true)
	//Context.SetCookie("refresh", refreshToken, config.RefreshTokenLife, "/", config.Domain, config.CookieIsSecure, true)
}

func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	o.IsSet = true

	if bytes.Equal(data, []byte("null")) {
		o.IsNull = true
		return nil
	}

	return json.Unmarshal(data, &o.Value)
}
