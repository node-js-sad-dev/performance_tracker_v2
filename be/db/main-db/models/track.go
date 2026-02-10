package models

type Track struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Image       *string `json:"image"`
	Description *string `json:"description"`
}
