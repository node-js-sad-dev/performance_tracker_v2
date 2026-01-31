package modules

import (
	"performance_tracker_v2_be/modules/car"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RootRouter(app *gin.Engine, pool *pgxpool.Pool) {
	rootRouterGroup := app.Group("/api/v1")

	car.Router(rootRouterGroup, pool)

	//rootRouterGroup.Use(middlewares.AuthMiddleware())
	//
	//task.Router(rootRouterGroup)
	//sse.Router(rootRouterGroup)
	//auth.Router(rootRouterGroup)
	//file.Router(rootRouterGroup)
	//
	//rootRouterGroup.GET("/socket", socket.Handler)
}
