package modules

import (
	"github.com/gin-gonic/gin"
)

func RootRouter(app *gin.Engine) {
	rootRouterGroup := app.Group("/api/v1")

	//rootRouterGroup.Use(middlewares.AuthMiddleware())
	//
	//task.Router(rootRouterGroup)
	//sse.Router(rootRouterGroup)
	//auth.Router(rootRouterGroup)
	//file.Router(rootRouterGroup)
	//
	//rootRouterGroup.GET("/socket", socket.Handler)
}
