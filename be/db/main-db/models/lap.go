package models

import "time"

type Lap struct {
	CreatedAt time.Time `json:"created_at"`

	ID        string `json:"id"`
	CarId     string `json:"carId"`
	TrackId   string `json:"trackId"`
	GameId    string `json:"gameId"`
	WheelId   string `json:"wheelId"`
	CockpitId string `json:"cockpitId"`
	PedalsId  string `json:"pedalsId"`
	GearboxId string `json:"gearboxId"`

	Time float32 `json:"time"`

	IsClear              bool `json:"isClear"`
	HasSignificantErrors bool `json:"hasSignificantErrors"`
}
