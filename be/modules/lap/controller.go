package lap

import (
	"performance_tracker_v2_be/core/handler"
	"performance_tracker_v2_be/core/responses"

	_ "performance_tracker_v2_be/core/swagger"
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
