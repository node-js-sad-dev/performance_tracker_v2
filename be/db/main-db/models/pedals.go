package models

import "time"

type Pedals struct {
	CreatedAt time.Time `json:"created_at"`

	ID   int64  `json:"id"`
	Name string `json:"name"`

	IsDefault bool `json:"is_default"`
}
