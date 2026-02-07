package lap

import (
	"performance_tracker_v2_be/core"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Controller struct {
	Pool    *pgxpool.Pool
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
//	@Success		200			{object}	core.SwaggerSuccessResponse[GetLapsResponse]
//	@Router			/lap [get]
func (controller *Controller) GetList(extraction *core.ExtractorResult[any, GetLapsFilter, any]) *core.ActionFuncResponse {
	laps, lapsError := controller.Service.GetAllLaps(extraction.Context, extraction.Pagination, extraction.Sort, extraction.QueryParams)
	if lapsError != nil {
		return core.DbErrorResponse(lapsError)
	}

	totalCount, countError := controller.Service.GetTotalLapsCount(extraction.Context, extraction.QueryParams)
	if countError != nil {
		return core.DbErrorResponse(countError)
	}

	return core.SuccessResponse(GetLapsResponse{
		Laps:       laps,
		TotalCount: int(totalCount),
	})
}
