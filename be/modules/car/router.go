package car

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	carGroup := router.Group("/car")
	{
		//carGroup.GET("/", http.Handler(GetList))
	}
}
