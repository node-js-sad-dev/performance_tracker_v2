package middlewares

import (
	"net/http"
	"performance_tracker_v2_be/config"
	"strings"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware(config *config.Config) gin.HandlerFunc {
	allowedOrigins := make(map[string]struct{})
	for _, origin := range config.GetAllowedOrigins() {
		allowedOrigins[origin] = struct{}{}
	}

	allowedMethods := strings.Join([]string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}, ", ")
	allowedHeaders := strings.Join([]string{"Accept", "Authorization", "Content-Type", "Origin"}, ", ")
	exposedHeaders := strings.Join([]string{"Content-Length", "Content-Type"}, ", ")

	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		_, isAllowedOrigin := allowedOrigins[origin]

		if origin != "" && isAllowedOrigin {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Methods", allowedMethods)
			c.Header("Access-Control-Allow-Headers", allowedHeaders)
			c.Header("Access-Control-Expose-Headers", exposedHeaders)
			c.Header("Vary", "Origin")
		}

		if c.Request.Method == http.MethodOptions {
			if origin != "" && !isAllowedOrigin {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"success": false,
					"error":   "origin not allowed",
				})
				return
			}

			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
