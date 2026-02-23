package game

import "performance_tracker_v2_be/core"

type Controller struct {
	Service *Service
}

// GetList @Summary Get games
//
//	@Description	Get all games
//	@Tags			Game
//	@Produce		json
//	@Param			page		query		int		false	"Page"
//	@Param			limit		query		int		false	"Limit"
//	@Param			sortBy		query		string	false	"Sort by"
//	@Param			sortOrder	query		string	false	"Sort order"
//	@Param			name		query		string	false	"Name filter (fuzzy)"
//	@Success		200			{object}	core.SwaggerSuccessResponse[GetListResponse]
//	@Router			/game [get]
func (controller *Controller) GetList(extraction *core.ExtractorResult[core.Empty, GetFilters, core.Empty]) *core.ActionFuncResponse {
	games, err := controller.Service.GetList(extraction.Context, extraction.Pagination, extraction.Sort, extraction.QueryParams)
	if err != nil {
		return core.DbErrorResponse(err)
	}

	totalCount, err := controller.Service.GetTotalCount(extraction.Context, extraction.QueryParams)
	if err != nil {
		return core.DbErrorResponse(err)
	}

	return core.SuccessResponse(GetListResponse{
		Games:      games,
		TotalCount: totalCount,
	})
}

// Create @Summary Create game
//
//	@Description	Create a new game
//	@Tags			Game
//	@Accept			json
//	@Produce		json
//	@Param			game	body		CreateRequest	true	"Game to create"
//	@Success		200	{object}	core.SwaggerSuccessResponse[models.Game]
//	@Router			/game [post]
func (controller *Controller) Create(extraction *core.ExtractorResult[CreateRequest, core.Empty, core.Empty]) *core.ActionFuncResponse {
	gameID, err := controller.Service.Create(extraction.Context, extraction.Body)

	if err != nil {
		return core.DbErrorResponse(err)
	}

	game, err := controller.Service.GetByID(extraction.Context, gameID)
	if err != nil {
		return core.DbErrorResponse(err)
	}

	return core.SuccessResponse(game)
}

// GetByID @Summary Get game by ID
//
//	@Description	Get a game by its ID
//	@Tags			Game
//	@Produce		json
//	@Param			id	path		int	true	"Game ID"
//	@Success		200	{object}	core.SwaggerSuccessResponse[models.Game]
//	@Router			/game/{id} [get]
func (controller *Controller) GetByID(extraction *core.ExtractorResult[core.Empty, core.Empty, core.GetByIdParams]) *core.ActionFuncResponse {
	game, err := controller.Service.GetByID(extraction.Context, extraction.Params.ID)
	if err != nil {
		return core.DbErrorResponse(err)
	}

	return core.SuccessResponse(game)
}

// UpdateByID @Summary Update game by ID
//
//	@Description	Update a game by its ID
//	@Tags			Game
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"Game ID"
//	@Param			game	body		UpdateRequestInput	true	"Game fields to update"
//	@Success		200	{object}	core.SwaggerSuccessResponse[models.Game]
//	@Router			/game/{id} [put]
func (controller *Controller) UpdateByID(extraction *core.ExtractorResult[UpdateRequestParsed, core.Empty, core.GetByIdParams]) *core.ActionFuncResponse {
	err := controller.Service.UpdateByID(extraction.Context, extraction.Params.ID, extraction.Body)
	if err != nil {
		return core.DbErrorResponse(err)
	}

	return core.DefaultSuccessResponse()
}

// DeleteByID @Summary Delete game by ID
//
//	@Description	Delete a game by its ID
//	@Tags			Game
//	@Produce		json
//	@Param			id	path		int	true	"Game ID"
//	@Success		200	{object}	core.SwaggerSuccessResponse[core.Empty]
//	@Router			/game/{id} [delete]
func (controller *Controller) DeleteByID(extraction *core.ExtractorResult[core.Empty, core.Empty, core.GetByIdParams]) *core.ActionFuncResponse {
	game, err := controller.Service.GetByID(extraction.Context, extraction.Params.ID)
	if err != nil {
		return core.DbErrorResponse(err)
	}

	if game == nil {
		return core.NotFoundErrorResponse("game")
	}

	err = controller.Service.DeleteByID(extraction.Context, extraction.Params.ID)
	if err != nil {
		return core.DbErrorResponse(err)
	}

	return core.DefaultSuccessResponse()
}
