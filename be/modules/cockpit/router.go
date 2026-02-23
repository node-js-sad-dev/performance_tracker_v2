package cockpit

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

	group := router.Group("/cockpit")
	{
		group.GET("/:id", handler.Handler(config, pool, controller.GetByID))
		group.PATCH("/:id", handler.Handler(config, pool, controller.UpdateByID))
		group.DELETE("/:id", handler.Handler(config, pool, controller.DeleteByID))
		group.GET("/", handler.Handler(config, pool, controller.GetList))
		group.POST("/", handler.Handler(config, pool, controller.Create))
	}
}
