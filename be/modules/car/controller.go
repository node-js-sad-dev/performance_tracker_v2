package car

import (
	"performance_tracker_v2_be/core"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Controller struct {
	Pool    *pgxpool.Pool
	Service *Service
}

// GetList @Summary Get cars
//
//	@Description	Get all cars
//	@Tags			Car
//	@Produce		json
//	@Param			page		query		int		false	"Page"
//	@Param			limit		query		int		false	"Limit"
//	@Param			sortBy		query		string	false	"Sort by"
//	@Param			sortOrder	query		string	false	"Sort order"
//	@Success		200			{object}	GetCarsResponse
//	@Router			/car [get]
func (controller *Controller) GetList(extraction *core.ExtractorResult[any, GetCarsFilter, any]) *core.ActionFuncResponse {
	cars, carsError := controller.Service.GetAllCars(extraction.Context, extraction.Pagination, extraction.Sort, extraction.QueryParams)
	if carsError != nil {
		return core.DbErrorResponse(carsError)
	}

	totalCount, countError := controller.Service.GetTotalCarsCount(extraction.Context, extraction.QueryParams)
	if countError != nil {
		return core.DbErrorResponse(countError)
	}

	return core.SuccessResponse(GetCarsResponse{
		Cars:       cars,
		TotalCount: int(totalCount),
	})
}
