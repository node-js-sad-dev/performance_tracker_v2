package models

type Cockpit struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	IsDefault bool `json:"isDefault"`
}
