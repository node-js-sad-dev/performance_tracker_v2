package car

import (
	"performance_tracker_v2_be/core"

	_ "performance_tracker_v2_be/db/main-db/models"
)

type Controller struct {
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
//	@Success		200			{object}	core.SwaggerSuccessResponse[GetListResponse]
//	@Router			/car [get]
func (controller *Controller) GetList(extraction *core.ExtractorResult[core.Empty, GetFilters, core.Empty]) *core.ActionFuncResponse {
	cars, err := controller.Service.GetList(extraction.Context, extraction.Pagination, extraction.Sort, extraction.QueryParams)
	if err != nil {
		return core.DbErrorResponse(err)
	}

	totalCount, err := controller.Service.GetTotalCount(extraction.Context, extraction.QueryParams)
	if err != nil {
		return core.DbErrorResponse(err)
	}

	return core.SuccessResponse(GetListResponse{
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
//	@Param			car	body		CreateRequest	true	"Car to create"
//	@Success		200	{object}	core.SwaggerSuccessResponse[models.Car]
//	@Router			/car [post]
func (controller *Controller) Create(extraction *core.ExtractorResult[CreateRequest, core.Empty, core.Empty]) *core.ActionFuncResponse {
	carID, err := controller.Service.Create(extraction.Context, extraction.Body)

	if err != nil {
		return core.DbErrorResponse(err)
	}

	car, err := controller.Service.GetById(extraction.Context, carID)
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
	car, err := controller.Service.GetById(extraction.Context, extraction.Params.ID)

	if err != nil {
		return core.DbErrorResponse(err)
	}

	return core.SuccessResponse(car)
}

// UpdateByID @Summary Update car by ID
//
//	@Description	Update a car by its ID
//	@Tags			Car
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string				true	"Car ID"
//	@Param			car	body		UpdateRequestInput	true	"Car fields to update"
//	@Success		200	{object}	core.SwaggerSuccessResponse[core.Empty]
//	@Router			/car/{id} [patch]
func (controller *Controller) UpdateByID(extraction *core.ExtractorResult[UpdateRequestParsed, core.Empty, core.GetByIdParams]) *core.ActionFuncResponse {
	car, err := controller.Service.GetById(extraction.Context, extraction.Params.ID)

	if err != nil {
		return core.DbErrorResponse(err)
	}

	if car == nil {
		return core.NotFoundErrorResponse("car")
	}

	if err := controller.Service.UpdateById(extraction.Context, extraction.Params.ID, extraction.Body); err != nil {
		return core.DbErrorResponse(err)
	}

	return core.DefaultSuccessResponse()
}

// DeleteByID @Summary Delete car by ID
//
//	@Description	Delete a car by its ID
//	@Tags			Car
//	@Produce		json
//	@Param			id	path		string	true	"Car ID"
//	@Success		200	{object}	core.SwaggerSuccessResponse[core.Empty]
//	@Router			/car/{id} [delete]
func (controller *Controller) DeleteByID(extraction *core.ExtractorResult[core.Empty, core.Empty, core.GetByIdParams]) *core.ActionFuncResponse {
	car, err := controller.Service.GetById(extraction.Context, extraction.Params.ID)

	if err != nil {
		return core.DbErrorResponse(err)
	}

	if car == nil {
		return core.NotFoundErrorResponse("car")
	}

	if err := controller.Service.DeleteById(extraction.Context, extraction.Params.ID); err != nil {
		return core.DbErrorResponse(err)
	}

	return core.DefaultSuccessResponse()
}
