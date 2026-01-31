package car

import (
	"performance_tracker_v2_be/core"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Router(router *gin.RouterGroup, pool *pgxpool.Pool) {
	controller := &Controller{
		Pool:    pool,
		Service: &Service{Pool: pool},
	}

	carGroup := router.Group("/car")
	{
		carGroup.GET("/", core.Handler[any, any, any](pool, controller.GetList))
	}
}
