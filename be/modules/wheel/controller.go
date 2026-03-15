package wheel

import (
	"performance_tracker_v2_be/core/handler"
	"performance_tracker_v2_be/core/responses"

	_ "performance_tracker_v2_be/core/swagger"
	_ "performance_tracker_v2_be/db/main-db/models"
)

type Controller struct {
	Service *Service
}

// GetList @Summary Get wheels
//
//	@Description	Get all wheels
//	@Tags			Wheel
//	@Produce		json
//	@Param			page		query		int		false	"Page"
//	@Param			limit		query		int		false	"Limit"
//	@Param			sortBy		query		string	false	"Sort by"
//	@Param			sortOrder	query		string	false	"Sort order"
//	@Param			name		query		string	false	"Name filter (fuzzy)"
//	@Success		200			{object}	swagger.SuccessResponse[GetListResponse]
//	@Router			/wheel [get]
func (c *Controller) GetList(extraction *handler.ExtractorResult[handler.Empty, GetFilters, handler.Empty]) *handler.ActionFuncResponse {
	wheels, err := c.Service.GetList(extraction.Context, extraction.Pagination, extraction.Sort, extraction.QueryParams)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	totalCount, err := c.Service.GetTotalCount(extraction.Context, extraction.QueryParams)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(GetListResponse{
		Wheels:     wheels,
		TotalCount: totalCount,
	})
}

// Create @Summary Create wheel
//
//	@Description	Create a new wheel
//	@Tags			Wheel
//	@Accept			json
//	@Produce		json
//	@Param			wheel	body		CreateRequest	true	"Wheel to create"
//	@Success		200			{object}	swagger.SuccessResponse[models.Wheel]
//	@Router			/wheel [post]
func (c *Controller) Create(extraction *handler.ExtractorResult[CreateRequest, handler.Empty, handler.Empty]) *handler.ActionFuncResponse {
	wheelID, err := c.Service.Create(extraction.Context, extraction.Body)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	wheel, err := c.Service.GetByID(extraction.Context, wheelID)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(wheel)
}

// GetByID @Summary Get wheel by ID
//
//	@Description	Get wheel by ID
//	@Tags			Wheel
//	@Produce		json
//	@Param			id	path		int64	true	"Wheel ID"
//	@Success		200	{object}	swagger.SuccessResponse[models.Wheel]
//	@Router			/wheel/{id} [get]
func (c *Controller) GetByID(extraction *handler.ExtractorResult[handler.Empty, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	wheel, err := c.Service.GetByID(extraction.Context, extraction.Params.ID)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(wheel)
}

// UpdateByID @Summary Update wheel by ID
//
//	@Description	Update wheel by ID
//	@Tags			Wheel
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int64				true	"Wheel ID"
//	@Param			wheel	body		UpdateRequestParsed	true	"Wheel data to update"
//	@Success		200		{object}	swagger.SuccessResponse[any]
//	@Router			/wheel/{id} [patch]
func (c *Controller) UpdateByID(extraction *handler.ExtractorResult[UpdateRequestParsed, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	if err := c.Service.UpdateByID(extraction.Context, extraction.Params.ID, extraction.Body); err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.DefaultSuccessResponse()
}

// DeleteByID @Summary Delete wheel by ID
//
//	@Description	Delete wheel by ID
//	@Tags			Wheel
//	@Produce		json
//	@Param			id	path		int64	true	"Wheel ID"
//	@Success		200	{object}	swagger.SuccessResponse[any]
//	@Router			/wheel/{id} [delete]
func (c *Controller) DeleteByID(extraction *handler.ExtractorResult[handler.Empty, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	wheel, err := c.Service.GetByID(extraction.Context, extraction.Params.ID)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	if wheel == nil {
		return responses.NotFoundErrorResponse("wheel")
	}

	if err := c.Service.DeleteByID(extraction.Context, extraction.Params.ID); err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.DefaultSuccessResponse()
}
