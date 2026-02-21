package cockpit

import (
	"performance_tracker_v2_be/core"

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
//	@Success		200			{object}	core.SwaggerSuccessResponse[GetListResponse]
//	@Router			/cockpit [get]
func (controller *Controller) GetList(extraction *core.ExtractorResult[core.Empty, GetFilters, core.Empty]) *core.ActionFuncResponse {
	cockpits, err := controller.Service.GetList(extraction.Context, extraction.Pagination, extraction.Sort, extraction.QueryParams)
	if err != nil {
		return core.DbErrorResponse(err)
	}

	totalCount, err := controller.Service.GetTotalCount(extraction.Context, extraction.QueryParams)
	if err != nil {
		return core.DbErrorResponse(err)
	}

	return core.SuccessResponse(GetListResponse{
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
//	@Success		200	{object}	core.SwaggerSuccessResponse[models.Cockpit]
//	@Router			/cockpit [post]
func (controller *Controller) Create(extraction *core.ExtractorResult[CreateRequest, core.Empty, core.Empty]) *core.ActionFuncResponse {
	cockpitID, err := controller.Service.Create(extraction.Context, extraction.Body)

	if err != nil {
		return core.DbErrorResponse(err)
	}

	cockpit, err := controller.Service.GetById(extraction.Context, cockpitID)
	if err != nil {
		return core.DbErrorResponse(err)
	}

	return core.SuccessResponse(cockpit)
}

// GetByID @Summary Get cockpit by ID
//
//	@Description	Get cockpit by ID
//	@Tags			Cockpit
//	@Produce		json
//	@Param			id	path	int	true	"Cockpit ID"
//	@Success		200	{object}	core.SwaggerSuccessResponse[models.Cockpit]
//	@Router			/cockpit/{id} [get]
func (controller *Controller) GetByID(extraction *core.ExtractorResult[core.Empty, core.Empty, core.GetByIdParams]) *core.ActionFuncResponse {
	cockpit, err := controller.Service.GetById(extraction.Context, extraction.Params.ID)
	if err != nil {
		return core.DbErrorResponse(err)
	}

	return core.SuccessResponse(cockpit)
}

// UpdateByID @Summary Update cockpit by ID
//
//	@Description	Update cockpit by ID
//	@Tags			Cockpit
//	@Accept			json
//	@Produce		json
//	@Param			id		path	int				true	"Cockpit ID"
//	@Param			cockpit	body	UpdateRequestInput	true	"Cockpit data to update"
//	@Success		200	{object}	core.SwaggerSuccessResponse[models.Cockpit]
//	@Router			/cockpit/{id} [patch]
func (controller *Controller) UpdateByID(extraction *core.ExtractorResult[UpdateRequestParsed, core.Empty, core.GetByIdParams]) *core.ActionFuncResponse {
	err := controller.Service.UpdateById(extraction.Context, extraction.Params.ID, extraction.Body)
	if err != nil {
		return core.DbErrorResponse(err)
	}

	return core.DefaultSuccessResponse()
}

// DeleteByID @Summary Delete cockpit by ID
//
//	@Description	Delete cockpit by ID
//	@Tags			Cockpit
//	@Produce		json
//	@Param			id	path	int	true	"Cockpit ID"
//	@Success		200	{object}	core.SwaggerSuccessResponse[core.Empty]
//	@Router			/cockpit/{id} [delete]
func (controller *Controller) DeleteByID(extraction *core.ExtractorResult[core.Empty, core.Empty, core.GetByIdParams]) *core.ActionFuncResponse {
	cockpit, err := controller.Service.GetById(extraction.Context, extraction.Params.ID)

	if err != nil {
		return core.DbErrorResponse(err)
	}

	if cockpit == nil {
		return core.NotFoundErrorResponse("cockpit")
	}

	if err := controller.Service.DeleteById(extraction.Context, extraction.Params.ID); err != nil {
		return core.DbErrorResponse(err)
	}

	return core.DefaultSuccessResponse()
}
