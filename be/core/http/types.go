package http

type Pagination struct {
	Page  int `form:"page" json:"page"`
	Limit int `form:"limit" json:"limit"`
}

type Sort struct {
	SortBy    string `form:"sortBy" json:"sortBy"`
	SortOrder string `form:"sortOrder" json:"sortOrder"`
}

type QueryRequest struct {
	Pagination
	Sort
}
