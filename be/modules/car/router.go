package car

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

	carGroup := router.Group("/car")
	{
		carGroup.GET("/:id", core.Handler(config, pool, controller.GetByID))
		carGroup.PATCH("/:id", core.Handler(config, pool, controller.UpdateByID))
		carGroup.DELETE("/:id", core.Handler(config, pool, controller.DeleteByID))
		carGroup.GET("/", core.Handler(config, pool, controller.GetList))
		carGroup.POST("/", core.Handler(config, pool, controller.Create))
	}
}
