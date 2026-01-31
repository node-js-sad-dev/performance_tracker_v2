package car

import "performance_tracker_v2_be/db/main-db/models"

type GetCarsFilter struct {
	Name []string
}

type GetCarsResponse struct {
	Cars       []models.Car `json:"cars"`
	TotalCount int          `json:"total_count"`
}
