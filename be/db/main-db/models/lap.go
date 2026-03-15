package models

import "time"

type Lap struct {
	CreatedAt time.Time `json:"created_at"`

	ID        int64 `json:"id"`
	CarID     int64 `json:"car_id"`
	TrackID   int64 `json:"track_id"`
	GameID    int64 `json:"game_id"`
	WheelID   int64 `json:"wheel_id"`
	CockpitID int64 `json:"cockpit_id"`
	PedalsID  int64 `json:"pedals_id"`

	Time float64 `json:"time"`

	IsClear              bool `json:"is_clear"`
	HasSignificantErrors bool `json:"has_significant_errors"`
}
