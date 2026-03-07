package game

import (
	"performance_tracker_v2_be/core/handler"
	"performance_tracker_v2_be/db/main-db/models"
)

type GetFilters struct {
	Name []string `form:"name"`
}

type GetListResponse struct {
	Games      []models.Game `json:"games"`
	TotalCount int64         `json:"total_count"`
}

type CreateRequest struct {
	Name string `form:"name" binding:"required"`
}

type CreateRequestParsed struct {
	Name  string  `json:"name" binding:"required"`
	Image *string `json:"image"`
}

type UpdateRequestParsed struct {
	Name  handler.OptionalBodyField[string] `form:"name" json:"name"`
	Image handler.OptionalBodyField[string] `form:"image" json:"image"`
}
