package models

import "time"

type Track struct {
	CreatedAt time.Time `json:"created_at"`

	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`

	Image *string `json:"image"`
}
