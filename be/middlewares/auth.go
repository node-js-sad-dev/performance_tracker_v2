package middlewares

import (
	//jwtAuth "task_management_backend/modules/auth"
	"github.com/gin-gonic/gin"
)

var allowlist = []struct {
	Method string
	Path   string
}{
	{"POST", "/api/v1/auth/login"},
	{"POST", "/api/v1/auth/register"},
	{"GET", "/api/v1/socket"},
}

func isAllowedRoute(method, path string) bool {
	for _, route := range allowlist {
		if route.Method == method && route.Path == path {
			return true
		}
	}
	return false
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//method := c.Request.Method
		//path := c.Request.URL.Path
		//
		//if isAllowedRoute(method, path) {
		//	c.Next()
		//	return
		//}
		//
		//accessToken, accessTokenErr := c.Cookie("access")
		//
		//if accessTokenErr != nil {
		//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, access token not found"})
		//	c.Abort()
		//	return
		//}
		//
		//claims := &jwtAuth.Claims{}
		//token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		//		return nil, fmt.Errorf("unexpected signing method")
		//	}
		//	return config.JwtSecret, nil
		//})
		//
		//if err != nil || !token.Valid {
		//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
		//	c.Abort()
		//	return
		//}
		//
		//c.Set("UserId", claims.UserId)
		//c.Set("LoginStep", claims.LoginStep)
		c.Next()
	}
}
