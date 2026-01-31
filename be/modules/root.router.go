package modules

import (
	"performance_tracker_v2_be/config"
	"performance_tracker_v2_be/modules/car"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RootRouter(config *config.Config, app *gin.Engine, pool *pgxpool.Pool) {
	rootRouterGroup := app.Group("/api/v1")

	// call global middlewares here before routers

	car.Router(config, pool, rootRouterGroup)
}
