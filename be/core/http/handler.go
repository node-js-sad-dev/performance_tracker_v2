package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(action ActionFunc, mapParams ...ParamFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var params []interface{}

		for _, fn := range mapParams {
			value, err := fn(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"error":   err.Error(),
				})
				return
			}
			params = append(params, value)
		}

		result := action(params...)
		if result.Error != nil {
			c.JSON(result.Status, gin.H{
				"success": false,
				"error":   result.Error.Error(),
			})
			return
		}

		if result.Cookies != nil {
			setCookies(c, result.Cookies.Access, result.Cookies.Refresh)
		}

		c.JSON(result.Status, gin.H{
			"success": true,
			"data":    result.Data,
		})
	}
}
