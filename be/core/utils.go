package core

import "github.com/gin-gonic/gin"

func setCookies(context *gin.Context, accessToken string, refreshToken string) {
	//context.SetCookie("access", accessToken, config.AccessTokenLife, "/", config.Domain, config.CookieIsSecure, true)
	//context.SetCookie("refresh", refreshToken, config.RefreshTokenLife, "/", config.Domain, config.CookieIsSecure, true)
}
