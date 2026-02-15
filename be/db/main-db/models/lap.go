package models

import "time"

type Lap struct {
	CreatedAt time.Time `json:"created_at"`

	ID        int64 `json:"id"`
	CarId     int64 `json:"carId"`
	TrackId   int64 `json:"trackId"`
	GameId    int64 `json:"gameId"`
	WheelId   int64 `json:"wheelId"`
	CockpitId int64 `json:"cockpitId"`
	PedalsId  int64 `json:"pedalsId"`
	GearboxId int64 `json:"gearboxId"`

	Time float32 `json:"time"`

	IsClear              bool `json:"isClear"`
	HasSignificantErrors bool `json:"hasSignificantErrors"`
}
