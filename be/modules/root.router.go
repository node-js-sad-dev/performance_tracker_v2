package modules

import (
	"performance_tracker_v2_be/modules/car"

	"github.com/gin-gonic/gin"
)

func RootRouter(app *gin.Engine) {
	rootRouterGroup := app.Group("/api/v1")

	car.Router(rootRouterGroup)

	//rootRouterGroup.Use(middlewares.AuthMiddleware())
	//
	//task.Router(rootRouterGroup)
	//sse.Router(rootRouterGroup)
	//auth.Router(rootRouterGroup)
	//file.Router(rootRouterGroup)
	//
	//rootRouterGroup.GET("/socket", socket.Handler)
}
