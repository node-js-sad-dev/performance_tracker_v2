package models

import "time"

type Game struct {
	CreatedAt time.Time `json:"created_at"`

	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Image *string `json:"image"`
}
