package models

type Wheel struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	IsDefault bool `json:"isDefault"`
}
