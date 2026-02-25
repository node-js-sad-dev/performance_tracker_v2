package pedals

import (
	"performance_tracker_v2_be/core/handler"
	"performance_tracker_v2_be/db/main-db/models"
)

type GetFilters struct {
	Name []string `form:"name"`
}

type GetListResponse struct {
	Pedals     []models.Pedals `json:"pedals"`
	TotalCount int64           `json:"total_count"`
}

type CreateRequest struct {
	Name      string `json:"name" binding:"required"`
	IsDefault bool   `json:"is_default" binding:"required"`
}

type UpdateRequestInput struct {
	Name      *string `json:"name"`
	IsDefault *bool   `json:"is_default"`
}

type UpdateRequestParsed struct {
	Name      handler.OptionalBodyField[string] `json:"name"`
	IsDefault handler.OptionalBodyField[bool]   `json:"is_default"`
}
