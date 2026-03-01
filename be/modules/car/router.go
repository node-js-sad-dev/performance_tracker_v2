package car

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

	group := router.Group("/car")
	{
		group.GET("/:id", handler.Handler(config, controller.GetByID))
		group.PATCH("/:id", handler.Handler(config, controller.UpdateByID))
		group.DELETE("/:id", handler.Handler(config, controller.DeleteByID))
		group.GET("", handler.Handler(config, controller.GetList))
		group.POST("", handler.Handler(config, controller.Create))
		// @todo -> investigate how it should be done correctly
	}
}
