package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"success": true,
		"data":    data,
	})
}

func SuccessResponseWithCookies(c *gin.Context, data interface{}, accessCookie, refreshCookie string) {
	setCookies(c, accessCookie, refreshCookie)

	c.JSON(200, gin.H{
		"success": true,
		"data":    data,
	})
}

func CommonErrorResponse(c *gin.Context, status int, errorMessage string) {
	c.JSON(status, gin.H{
		"success": false,
		"error":   errorMessage,
	})
}

func DbErrorResponse(c *gin.Context, err error) {
	var status int
	var errorMessage string

	switch err.Error() {
	case "record not found":
		status = http.StatusNotFound
		errorMessage = "record not found"
	default:
		status = http.StatusInternalServerError
		errorMessage = "internal server error"
	}

	CommonErrorResponse(c, status, errorMessage)
}
