package core

import (
	"performance_tracker_v2_be/config"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Handler[Body any, Query any, Params any](config *config.Config, pool *pgxpool.Pool, actionFunc func(extraction *ExtractorResult[Body, Query, Params]) *ActionFuncResponse) gin.HandlerFunc {
	return func(c *gin.Context) {
		extraction, extractionErr := Extract[Body, Query, Params](config, pool, c)
		if extractionErr != nil {
			c.JSON(400, gin.H{
				"success": false,
				"error":   "failed to extract request data: " + extractionErr.Error(),
			})
			return
		}

		result := actionFunc(extraction)

		if result.Error != nil {
			c.JSON(result.Status, gin.H{
				"success": false,
				"error":   result.Error.Error(),
			})
			return
		}

		c.JSON(result.Status, gin.H{
			"success": true,
			"data":    result.Data,
		})
	}
}
