package lap

import (
	"performance_tracker_v2_be/config"
	"performance_tracker_v2_be/core"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Router(config *config.Config, pool *pgxpool.Pool, router *gin.RouterGroup) {
	controller := &Controller{
		Pool:    pool,
		Service: &Service{Pool: pool},
	}

	lapGroup := router.Group("/lap")
	{
		lapGroup.GET("/", core.Handler[any, GetLapsFilter, any](config, pool, controller.GetList))
	}
}
