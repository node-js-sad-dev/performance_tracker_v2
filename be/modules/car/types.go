package car

import (
	"performance_tracker_v2_be/core/handler"
	"performance_tracker_v2_be/db/main-db/models"
)

type GetFilters struct {
	Name []string `form:"name"`
}

type GetListResponse struct {
	Cars       []models.Car `json:"cars"`
	TotalCount int64        `json:"total_count"`
}

type CreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Image       string `json:"image"`
}

type UpdateRequestInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`

	Image *string `json:"image"`
}

type UpdateRequestParsed struct {
	Name        handler.OptionalBodyField[string] `json:"name"`
	Description handler.OptionalBodyField[string] `json:"description"`

	Image handler.OptionalBodyField[string] `json:"image"`
}
