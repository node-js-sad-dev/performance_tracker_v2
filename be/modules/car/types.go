package car

import "performance_tracker_v2_be/db/main-db/models"

type GetCarsFilter struct {
	Name []string `form:"name"`
}

type GetCarsResponse struct {
	Cars       []models.Car `json:"cars"`
	TotalCount int64        `json:"total_count"`
}

type CreateCarRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Image       string `json:"image"`
}

type UpdateCarRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Image       *string `json:"image"`
}
