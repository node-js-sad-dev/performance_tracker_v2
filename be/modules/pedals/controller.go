package pedals

import (
	"performance_tracker_v2_be/core/handler"
	"performance_tracker_v2_be/core/responses"

	_ "performance_tracker_v2_be/core/swagger"
	_ "performance_tracker_v2_be/db/main-db/models"
)

type Controller struct {
	Service *Service
}

// GetList @Summary Get pedals
//
//	@Description	Get all pedals
//	@Tags			Pedal
//	@Produce		json
//	@Param			page		query		int		false	"Page"
//	@Param			limit		query		int		false	"Limit"
//	@Param			sortBy		query		string	false	"Sort by"
//	@Param			sortOrder	query		string	false	"Sort order"
//	@Param			name		query		string	false	"Name filter (fuzzy)"
//	@Success		200			{object}	swagger.SuccessResponse[GetListResponse]
//	@Router			/pedals [get]
func (controller *Controller) GetList(extraction *handler.ExtractorResult[handler.Empty, GetFilters, handler.Empty]) *handler.ActionFuncResponse {
	pedals, err := controller.Service.GetList(extraction.Context, extraction.Pagination, extraction.Sort, extraction.QueryParams)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	totalCount, err := controller.Service.GetTotalCount(extraction.Context, extraction.QueryParams)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(GetListResponse{
		Pedals:     pedals,
		TotalCount: totalCount,
	})
}

// Create @Summary Create pedal
//
//	@Description	Create a new pedal
//	@Tags			Pedal
//	@Accept			json
//	@Produce		json
//	@Param			pedal	body		CreateRequest	true	"Pedal to create"
//	@Success		200	{object}	swagger.SuccessResponse[models.Pedals]
//	@Router			/pedals [post]
func (controller *Controller) Create(extraction *handler.ExtractorResult[CreateRequest, handler.Empty, handler.Empty]) *handler.ActionFuncResponse {
	pedalsId, err := controller.Service.Create(extraction.Context, extraction.Body)

	if err != nil {
		return responses.DbErrorResponse(err)
	}

	pedals, err := controller.Service.GetByID(extraction.Context, pedalsId)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(pedals)
}

// GetByID @Summary Get pedals by ID
//
//	@Description	Get pedals by ID
//	@Tags			Pedal
//	@Produce		json
//	@Param			id		path		int64	true	"Pedal ID"
//	@Success		200	{object}	swagger.SuccessResponse[models.Pedals]
//	@Router			/pedals/{id} [get]
func (controller *Controller) GetByID(extraction *handler.ExtractorResult[handler.Empty, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	pedals, err := controller.Service.GetByID(extraction.Context, extraction.Params.ID)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(pedals)
}

// UpdateByID @Summary Update pedals by ID
//
//	@Description	Update pedals by ID
//	@Tags			Pedal
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int64			true	"Pedal ID"
//	@Param			pedal	body		UpdateRequestParsed	true	"Pedal data to update"
//	@Success		200	{object}	swagger.SuccessResponse[any]
//	@Router			/pedals/{id} [patch]
func (controller *Controller) UpdateByID(extraction *handler.ExtractorResult[UpdateRequestParsed, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	err := controller.Service.UpdateByID(extraction.Context, extraction.Params.ID, extraction.Body)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.DefaultSuccessResponse()
}

// DeleteByID @Summary Delete pedals by ID
//
//	@Description	Delete pedals by ID
//	@Tags			Pedal
//	@Produce		json
//	@Param			id		path		int64	true	"Pedal ID"
//	@Success		200	{object}	swagger.SuccessResponse[any]
//	@Router			/pedals/{id} [delete]
func (controller *Controller) DeleteByID(extraction *handler.ExtractorResult[handler.Empty, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	pedals, err := controller.Service.GetByID(extraction.Context, extraction.Params.ID)

	if err != nil {
		return responses.DbErrorResponse(err)
	}

	if pedals == nil {
		return responses.NotFoundErrorResponse("pedals")
	}

	if err := controller.Service.DeleteByID(extraction.Context, extraction.Params.ID); err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.DefaultSuccessResponse()
}
