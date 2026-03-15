package lap

import (
	"errors"
	"performance_tracker_v2_be/core/handler"
	"performance_tracker_v2_be/core/responses"

	_ "performance_tracker_v2_be/core/swagger"
	_ "performance_tracker_v2_be/db/main-db/models"

	"github.com/jackc/pgx/v5"
)

type Controller struct {
	Service *Service
}

// GetList @Summary Get laps
//
//	@Description	Get all laps
//	@Tags			Lap
//	@Produce		json
//	@Param			page		query		int		false	"Page"
//	@Param			limit		query		int		false	"Limit"
//	@Param			sortBy		query		string	false	"Sort by"
//	@Param			sortOrder	query		string	false	"Sort order"
//	@Param			date		query		string	false	"Date filter"
//	@Param			game		query		string	false	"Game filter"
//	@Param			car			query		string	false	"Car filter"
//	@Param			track		query		string	false	"Track filter"
//	@Param			time		query		string	false	"Time filter"
//	@Param			clear		query		string	false	"Clear filter"
//	@Success		200			{object}	swagger.SuccessResponse[GetLapsResponse]
//	@Router			/lap [get]
func (controller *Controller) GetList(extraction *handler.ExtractorResult[handler.Empty, GetLapsFilter, handler.Empty]) *handler.ActionFuncResponse {
	laps, err := controller.Service.GetAllLaps(extraction.Context, extraction.Pagination, extraction.Sort, extraction.QueryParams)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	totalCount, err := controller.Service.GetTotalLapsCount(extraction.Context, extraction.QueryParams)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(GetLapsResponse{
		Laps:       laps,
		TotalCount: totalCount,
	})
}

// Create @Summary Create lap
//
//	@Description	Create a new lap
//	@Tags			Lap
//	@Accept			json
//	@Produce		json
//	@Param			lap	body		CreateRequest	true	"Lap to create"
//	@Success		200			{object}	swagger.SuccessResponse[models.Lap]
//	@Router			/lap [post]
func (controller *Controller) Create(extraction *handler.ExtractorResult[CreateRequest, handler.Empty, handler.Empty]) *handler.ActionFuncResponse {
	lapID, err := controller.Service.Create(extraction.Context, extraction.Body)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	lap, err := controller.Service.GetByID(extraction.Context, lapID)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(lap)
}

// GetByID @Summary Get lap by ID
//
//	@Description	Get lap by ID
//	@Tags			Lap
//	@Produce		json
//	@Param			id	path		int64	true	"Lap ID"
//	@Success		200	{object}	swagger.SuccessResponse[models.Lap]
//	@Router			/lap/{id} [get]
func (controller *Controller) GetByID(extraction *handler.ExtractorResult[handler.Empty, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	lap, err := controller.Service.GetByID(extraction.Context, extraction.Params.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return responses.NotFoundErrorResponse("lap")
		}

		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(lap)
}

// UpdateByID @Summary Update lap by ID
//
//	@Description	Update lap by ID
//	@Tags			Lap
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int64				true	"Lap ID"
//	@Param			lap	body		UpdateRequestParsed	true	"Lap data to update"
//	@Success		200		{object}	swagger.SuccessResponse[any]
//	@Router			/lap/{id} [patch]
func (controller *Controller) UpdateByID(extraction *handler.ExtractorResult[UpdateRequestParsed, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	_, err := controller.Service.GetByID(extraction.Context, extraction.Params.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return responses.NotFoundErrorResponse("lap")
		}

		return responses.DbErrorResponse(err)
	}

	if err := controller.Service.UpdateByID(extraction.Context, extraction.Params.ID, extraction.Body); err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.DefaultSuccessResponse()
}

// DeleteByID @Summary Delete lap by ID
//
//	@Description	Delete lap by ID
//	@Tags			Lap
//	@Produce		json
//	@Param			id	path		int64	true	"Lap ID"
//	@Success		200	{object}	swagger.SuccessResponse[any]
//	@Router			/lap/{id} [delete]
func (controller *Controller) DeleteByID(extraction *handler.ExtractorResult[handler.Empty, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	_, err := controller.Service.GetByID(extraction.Context, extraction.Params.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return responses.NotFoundErrorResponse("lap")
		}

		return responses.DbErrorResponse(err)
	}

	if err := controller.Service.DeleteByID(extraction.Context, extraction.Params.ID); err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.DefaultSuccessResponse()
}
