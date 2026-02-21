package cockpit

import (
	"performance_tracker_v2_be/config"
	"performance_tracker_v2_be/core"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Router(config *config.Config, pool *pgxpool.Pool, router *gin.RouterGroup) {
	controller := &Controller{
		Service: &Service{Pool: pool},
	}

	group := router.Group("/cockpit")
	{
		group.GET("/:id", core.Handler(config, pool, controller.GetByID))
		group.PATCH("/:id", core.Handler(config, pool, controller.UpdateByID))
		group.DELETE("/:id", core.Handler(config, pool, controller.DeleteByID))
		group.GET("/", core.Handler(config, pool, controller.GetList))
		group.POST("/", core.Handler(config, pool, controller.Create))
	}
}
