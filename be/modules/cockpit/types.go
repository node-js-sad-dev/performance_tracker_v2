package cockpit

import (
	"performance_tracker_v2_be/core"
	"performance_tracker_v2_be/db/main-db/models"
)

type GetFilters struct {
	Name []string `form:"name"`
}

type GetListResponse struct {
	Cockpits   []models.Cockpit `json:"cockpits"`
	TotalCount int64            `json:"total_count"`
}

type CreateRequest struct {
	Name      string `json:"name" binding:"required"`
	IsDefault bool   `json:"is_default"`
}

type UpdateRequestInput struct {
	Name      *string `json:"name"`
	IsDefault *bool   `json:"is_default"`
}

type UpdateRequestParsed struct {
	Name      core.OptionalBodyField[string] `json:"name"`
	IsDefault core.OptionalBodyField[bool]   `json:"is_default"`
}
