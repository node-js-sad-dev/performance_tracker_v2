package cockpit

import (
	"performance_tracker_v2_be/core/handler"
	"performance_tracker_v2_be/core/responses"

	_ "performance_tracker_v2_be/core/swagger"
	_ "performance_tracker_v2_be/db/main-db/models"
)

type Controller struct {
	Service *Service
}

// GetList @Summary Get cockpits
//
//	@Description	Get all cockpits
//	@Tags			Cockpit
//	@Produce		json
//	@Param			page		query		int		false	"Page"
//	@Param			limit		query		int		false	"Limit"
//	@Param			sortBy		query		string	false	"Sort by"
//	@Param			sortOrder	query		string	false	"Sort order"
//	@Param			name		query		string	false	"Name filter (fuzzy)"
//	@Success		200			{object}	swagger.SuccessResponse[GetListResponse]
//	@Router			/cockpit [get]
func (controller *Controller) GetList(extraction *handler.ExtractorResult[handler.Empty, GetFilters, handler.Empty]) *handler.ActionFuncResponse {
	cockpits, err := controller.Service.GetList(extraction.Context, extraction.Pagination, extraction.Sort, extraction.QueryParams)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	totalCount, err := controller.Service.GetTotalCount(extraction.Context, extraction.QueryParams)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(GetListResponse{
		Cockpits:   cockpits,
		TotalCount: totalCount,
	})
}

// Create @Summary Create cockpit
//
//	@Description	Create a new cockpit
//	@Tags			Cockpit
//	@Accept			json
//	@Produce		json
//	@Param			cockpit	body		CreateRequest	true	"Cockpit to create"
//	@Success		200		{object}	swagger.SuccessResponse[models.Cockpit]
//	@Router			/cockpit [post]
func (controller *Controller) Create(extraction *handler.ExtractorResult[CreateRequest, handler.Empty, handler.Empty]) *handler.ActionFuncResponse {
	cockpitID, err := controller.Service.Create(extraction.Context, extraction.Body)

	if err != nil {
		return responses.DbErrorResponse(err)
	}

	cockpit, err := controller.Service.GetByID(extraction.Context, cockpitID)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(cockpit)
}

// GetByID @Summary Get cockpit by ID
//
//	@Description	Get cockpit by ID
//	@Tags			Cockpit
//	@Produce		json
//	@Param			id	path		int	true	"Cockpit ID"
//	@Success		200	{object}	swagger.SuccessResponse[models.Cockpit]
//	@Router			/cockpit/{id} [get]
func (controller *Controller) GetByID(extraction *handler.ExtractorResult[handler.Empty, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	cockpit, err := controller.Service.GetByID(extraction.Context, extraction.Params.ID)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(cockpit)
}

// UpdateByID @Summary Update cockpit by ID
//
//	@Description	Update cockpit by ID
//	@Tags			Cockpit
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Cockpit ID"
//	@Param			cockpit	body		UpdateRequestInput	true	"Cockpit data to update"
//	@Success		200		{object}	swagger.SuccessResponse[models.Cockpit]
//	@Router			/cockpit/{id} [patch]
func (controller *Controller) UpdateByID(extraction *handler.ExtractorResult[UpdateRequestParsed, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	err := controller.Service.UpdateByID(extraction.Context, extraction.Params.ID, extraction.Body)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.DefaultSuccessResponse()
}

// DeleteByID @Summary Delete cockpit by ID
//
//	@Description	Delete cockpit by ID
//	@Tags			Cockpit
//	@Produce		json
//	@Param			id	path		int	true	"Cockpit ID"
//	@Success		200	{object}	swagger.SuccessResponse[handler.Empty]
//	@Router			/cockpit/{id} [delete]
func (controller *Controller) DeleteByID(extraction *handler.ExtractorResult[handler.Empty, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	cockpit, err := controller.Service.GetByID(extraction.Context, extraction.Params.ID)

	if err != nil {
		return responses.DbErrorResponse(err)
	}

	if cockpit == nil {
		return responses.NotFoundErrorResponse("cockpit")
	}

	if err := controller.Service.DeleteByID(extraction.Context, extraction.Params.ID); err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.DefaultSuccessResponse()
}
