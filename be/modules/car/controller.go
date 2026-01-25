package car

import (
	"performance_tracker_v2_be/core"

	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context) {
	core.SuccessResponse(c, "Car list retrieved successfully")
}
