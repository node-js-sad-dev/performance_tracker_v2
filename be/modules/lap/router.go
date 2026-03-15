package lap

import (
	"performance_tracker_v2_be/config"
	"performance_tracker_v2_be/core/handler"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Router(config *config.Config, pool *pgxpool.Pool, router *gin.RouterGroup) {
	controller := &Controller{
		Service: &Service{Pool: pool},
	}

	lapGroup := router.Group("/lap")
	{
		lapGroup.GET("/:id", handler.Handler(config, controller.GetByID))
		lapGroup.PATCH("/:id", handler.Handler(config, controller.UpdateByID))
		lapGroup.DELETE("/:id", handler.Handler(config, controller.DeleteByID))
		lapGroup.GET("", handler.Handler(config, controller.GetList))
		lapGroup.POST("", handler.Handler(config, controller.Create))
	}
}
