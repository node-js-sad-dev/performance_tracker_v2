package game

import (
	"performance_tracker_v2_be/core/handler"
	"performance_tracker_v2_be/core/responses"

	_ "performance_tracker_v2_be/core/swagger"
	_ "performance_tracker_v2_be/db/main-db/models"
)

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
//	@Success		200			{object}	swagger.SuccessResponse[GetListResponse]
//	@Router			/game [get]
func (controller *Controller) GetList(extraction *handler.ExtractorResult[handler.Empty, GetFilters, handler.Empty]) *handler.ActionFuncResponse {
	games, err := controller.Service.GetList(extraction.Context, extraction.Pagination, extraction.Sort, extraction.QueryParams)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	totalCount, err := controller.Service.GetTotalCount(extraction.Context, extraction.QueryParams)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(GetListResponse{
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
//	@Success		200	{object}	swagger.SuccessResponse[models.Game]
//	@Router			/game [post]
func (controller *Controller) Create(extraction *handler.ExtractorResult[CreateRequest, handler.Empty, handler.Empty]) *handler.ActionFuncResponse {
	gameID, err := controller.Service.Create(extraction.Context, extraction.Body)

	if err != nil {
		return responses.DbErrorResponse(err)
	}

	game, err := controller.Service.GetByID(extraction.Context, gameID)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(game)
}

// GetByID @Summary Get game by ID
//
//	@Description	Get a game by its ID
//	@Tags			Game
//	@Produce		json
//	@Param			id	path		int	true	"Game ID"
//	@Success		200	{object}	swagger.SuccessResponse[models.Game]
//	@Router			/game/{id} [get]
func (controller *Controller) GetByID(extraction *handler.ExtractorResult[handler.Empty, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	game, err := controller.Service.GetByID(extraction.Context, extraction.Params.ID)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(game)
}

// UpdateByID @Summary Update game by ID
//
//	@Description	Update a game by its ID
//	@Tags			Game
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"Game ID"
//	@Param			game	body		UpdateRequestInput	true	"Game fields to update"
//	@Success		200	{object}	swagger.SuccessResponse[models.Game]
//	@Router			/game/{id} [patch]
func (controller *Controller) UpdateByID(extraction *handler.ExtractorResult[UpdateRequestParsed, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	err := controller.Service.UpdateByID(extraction.Context, extraction.Params.ID, extraction.Body)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.DefaultSuccessResponse()
}

// DeleteByID @Summary Delete game by ID
//
//	@Description	Delete a game by its ID
//	@Tags			Game
//	@Produce		json
//	@Param			id	path		int	true	"Game ID"
//	@Success		200	{object}	swagger.SuccessResponse[handler.Empty]
//	@Router			/game/{id} [delete]
func (controller *Controller) DeleteByID(extraction *handler.ExtractorResult[handler.Empty, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	game, err := controller.Service.GetByID(extraction.Context, extraction.Params.ID)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	if game == nil {
		return responses.NotFoundErrorResponse("game")
	}

	err = controller.Service.DeleteByID(extraction.Context, extraction.Params.ID)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.DefaultSuccessResponse()
}
