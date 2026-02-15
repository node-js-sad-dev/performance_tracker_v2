package models

import "time"

type Car struct {
	CreatedAt time.Time `json:"created_at"`

	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Image       *string `json:"image"`
	Description *string `json:"description"`
}
