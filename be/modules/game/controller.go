package game

import (
	"performance_tracker_v2_be/core/files"
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
func (c *Controller) GetList(extraction *handler.ExtractorResult[handler.Empty, GetFilters, handler.Empty]) *handler.ActionFuncResponse {
	games, err := c.Service.GetList(extraction.Context, extraction.Pagination, extraction.Sort, extraction.QueryParams)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	totalCount, err := c.Service.GetTotalCount(extraction.Context, extraction.QueryParams)
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
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			file		formData	file	true	"Game image file"
//	@Param			name		formData	string	true	"Name of the game"
//	@Success		200		{object}	swagger.SuccessResponse[models.Game]
//	@Router			/game [post]
func (c *Controller) Create(extraction *handler.ExtractorResult[CreateRequest, handler.Empty, handler.Empty]) *handler.ActionFuncResponse {
	gameWithSameName, err := c.Service.GetByName(extraction.Context, extraction.Body.Name, nil)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	if gameWithSameName != nil {
		return responses.CommonErrorResponse(409, "game with this name already exists")
	}

	gameImage, err := files.GetFileInfoFromExtraction("file", extraction.Files)

	if err != nil {
		return responses.CommonErrorResponse(500, err.Error())
	}

	if gameImage == nil {
		return responses.CommonErrorResponse(400, "file is required")
	}

	filePath, err := files.SaveFile(gameImage)
	if err != nil {
		return responses.CommonErrorResponse(500, err.Error())
	}

	gameId, err := c.Service.Create(extraction.Context, &CreateRequestParsed{
		Name:  extraction.Body.Name,
		Image: filePath,
	})

	if err != nil {
		return responses.DbErrorResponse(err)
	}

	game, err := c.Service.GetByID(extraction.Context, gameId)
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
func (c *Controller) GetByID(extraction *handler.ExtractorResult[handler.Empty, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	game, err := c.Service.GetByID(extraction.Context, extraction.Params.ID)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(game)
}

// UpdateByID @Summary Update game by ID
//
//	@Description	Update a game by its ID
//	@Tags			Game
//	@Produce		json
//	@Param			id			path		number	true	"Car ID"
//	@Param			file		formData	file	false	"Car image file"
//	@Param			name		formData	string	false	"Name of the car"
//	@Param			description	formData	string	false	"Detailed car description"
//	@Param			image		formData	string	false	"Image to update if file not passed"
//	@Success		200		{object}	swagger.SuccessResponse[models.Game]
//	@Router			/game/{id} [patch]
func (c *Controller) UpdateByID(extraction *handler.ExtractorResult[UpdateRequestParsed, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	err := c.Service.UpdateByID(extraction.Context, extraction.Params.ID, extraction.Body)
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
func (c *Controller) DeleteByID(extraction *handler.ExtractorResult[handler.Empty, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	game, err := c.Service.GetByID(extraction.Context, extraction.Params.ID)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	if game == nil {
		return responses.NotFoundErrorResponse("game")
	}

	err = c.Service.DeleteByID(extraction.Context, extraction.Params.ID)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.DefaultSuccessResponse()
}
