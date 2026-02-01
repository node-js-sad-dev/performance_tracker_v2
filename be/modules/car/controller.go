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

// Create @Summary Create car
//
//	@Description	Create a new car
//	@Tags			Car
//	@Accept			json
//	@Produce		json
//	@Param			car	body		CreateCarRequest	true	"Car to create"
//	@Success		200	{object}	core.SwaggerSuccessResponse[models.Car]
//	@Router			/car [post]
func (controller *Controller) Create(extraction *core.ExtractorResult[CreateCarRequest, any, any]) *core.ActionFuncResponse {
	carID, createError := controller.Service.CreateCar(extraction.Context, extraction.Body)

	if createError != nil {
		return core.DbErrorResponse(createError)
	}

	car, getError := controller.Service.GetCarByID(extraction.Context, carID)
	if getError != nil {
		return core.DbErrorResponse(getError)
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
func (controller *Controller) GetByID(extraction *core.ExtractorResult[any, any, core.GetByIdParams]) *core.ActionFuncResponse {
	car, getError := controller.Service.GetCarByID(extraction.Context, extraction.Params.ID)

	if getError != nil {
		return core.DbErrorResponse(getError)
	}

	return core.SuccessResponse(car)
}
