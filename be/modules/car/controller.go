package car

import (
	"performance_tracker_v2_be/core"

	"github.com/jackc/pgx/v5/pgxpool"

	_ "performance_tracker_v2_be/db/main-db/models"
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
//	@Success		200			{object}	core.SwaggerSuccessResponse[GetCarsResponse]
//	@Router			/car [get]
func (controller *Controller) GetList(extraction *core.ExtractorResult[core.Empty, GetCarsFilter, core.Empty]) *core.ActionFuncResponse {
	cars, err := controller.Service.GetAllCars(extraction.Context, extraction.Pagination, extraction.Sort, extraction.QueryParams)
	if err != nil {
		return core.DbErrorResponse(err)
	}

	totalCount, err := controller.Service.GetTotalCarsCount(extraction.Context, extraction.QueryParams)
	if err != nil {
		return core.DbErrorResponse(err)
	}

	return core.SuccessResponse(GetCarsResponse{
		Cars:       cars,
		TotalCount: totalCount,
	})
}

// Create @Summary Create car
//
//	@Description	Create a new car
//	@Tags			Car
//	@Accept			json
//	@Produce		json
//	@Param			car	body		CreateCarRequest	true	"Car to create"
//	@Success		200	{object}	core.SwaggerSuccessResponse[models.Car]
//	@Router			/car [post]
func (controller *Controller) Create(extraction *core.ExtractorResult[CreateCarRequest, core.Empty, core.Empty]) *core.ActionFuncResponse {
	carID, err := controller.Service.CreateCar(extraction.Context, extraction.Body)

	if err != nil {
		return core.DbErrorResponse(err)
	}

	car, err := controller.Service.GetCarByID(extraction.Context, carID)
	if err != nil {
		return core.DbErrorResponse(err)
	}

	return core.SuccessResponse(car)
}

// GetByID @Summary Get car by ID
//
//	@Description	Get a car by its ID
//	@Tags			Car
//	@Produce		json
//	@Param			id	path		string	true	"Car ID"
//	@Success		200	{object}	core.SwaggerSuccessResponse[models.Car]
//	@Router			/car/{id} [get]
func (controller *Controller) GetByID(extraction *core.ExtractorResult[core.Empty, core.Empty, core.GetByIdParams]) *core.ActionFuncResponse {
	car, err := controller.Service.GetCarByID(extraction.Context, extraction.Params.ID)

	if err != nil {
		return core.DbErrorResponse(err)
	}

	return core.SuccessResponse(car)
}
