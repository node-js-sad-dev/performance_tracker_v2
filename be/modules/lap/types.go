package lap

import (
	"performance_tracker_v2_be/core/handler"
)

type GetListLap struct {
	ID    int64  `json:"id"`
	Date  string `json:"date"`
	Game  string `json:"game"`
	Car   string `json:"car"`
	Track string `json:"track"`
	Time  string `json:"time"`
	Clear bool   `json:"clear"`
}

type GetLapsFilter struct {
	Date  []string `form:"date"`
	Game  []string `form:"game"`
	Car   []string `form:"car"`
	Track []string `form:"track"`
	Time  []string `form:"time"`
	Clear []string `form:"clear"`
}

type GetLapsResponse struct {
	Laps       []GetListLap `json:"laps"`
	TotalCount int64        `json:"total_count"`
}

type CreateRequest struct {
	CarID                int64   `json:"car_id" binding:"required"`
	TrackID              int64   `json:"track_id" binding:"required"`
	GameID               int64   `json:"game_id" binding:"required"`
	WheelID              int64   `json:"wheel_id" binding:"required"`
	CockpitID            int64   `json:"cockpit_id" binding:"required"`
	PedalsID             int64   `json:"pedals_id" binding:"required"`
	Time                 float64 `json:"time" binding:"required"`
	IsClear              bool    `json:"is_clear"`
	HasSignificantErrors bool    `json:"has_significant_errors"`
}

type UpdateRequestParsed struct {
	CarID                handler.OptionalBodyField[int64]   `json:"car_id"`
	TrackID              handler.OptionalBodyField[int64]   `json:"track_id"`
	GameID               handler.OptionalBodyField[int64]   `json:"game_id"`
	WheelID              handler.OptionalBodyField[int64]   `json:"wheel_id"`
	CockpitID            handler.OptionalBodyField[int64]   `json:"cockpit_id"`
	PedalsID             handler.OptionalBodyField[int64]   `json:"pedals_id"`
	Time                 handler.OptionalBodyField[float64] `json:"time"`
	IsClear              handler.OptionalBodyField[bool]    `json:"is_clear"`
	HasSignificantErrors handler.OptionalBodyField[bool]    `json:"has_significant_errors"`
}
