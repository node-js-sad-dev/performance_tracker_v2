package modules

import (
	"performance_tracker_v2_be/config"
	"performance_tracker_v2_be/modules/car"
	"performance_tracker_v2_be/modules/cockpit"
	"performance_tracker_v2_be/modules/game"
	"performance_tracker_v2_be/modules/lap"
	"performance_tracker_v2_be/modules/pedals"
	"performance_tracker_v2_be/modules/track"
	"performance_tracker_v2_be/modules/wheel"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RootRouter(config *config.Config, pool *pgxpool.Pool, app *gin.Engine) {
	rootRouterGroup := app.Group("/api/v1")

	// call global middlewares here before routers

	// routers

	car.Router(config, pool, rootRouterGroup)
	cockpit.Router(config, pool, rootRouterGroup)
	lap.Router(config, pool, rootRouterGroup)
	game.Router(config, pool, rootRouterGroup)
	pedals.Router(config, pool, rootRouterGroup)
	track.Router(config, pool, rootRouterGroup)
	wheel.Router(config, pool, rootRouterGroup)
}
