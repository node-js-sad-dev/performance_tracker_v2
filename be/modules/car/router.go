package car

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	carGroup := router.Group("/car")
	{
		carGroup.GET("/", GetList)
	}
}
