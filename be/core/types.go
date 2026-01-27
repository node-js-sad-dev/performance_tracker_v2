package core

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type Sort struct {
	SortBy    string `json:"sortBy"`
	SortOrder string `json:"sortOrder"`
}
