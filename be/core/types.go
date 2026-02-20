package core

import (
	"context"
	"performance_tracker_v2_be/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type FilterRule struct {
	DBColumn string
	Operator string
	IsFuzzy  bool
}

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type Sort struct {
	SortBy    string `json:"sortBy"`
	SortOrder string `json:"sortOrder"`
}

type ExtractorResult[Body any, Query any, Params any] struct {
	Params      *Params
	Pagination  *Pagination
	Sort        *Sort
	QueryParams *Query
	Body        *Body
	Headers     *map[string]string
	Config      *config.Config
	Pool        *pgxpool.Pool
	Context     context.Context
}

type GetEntityListPayload struct {
	Pool         *pgxpool.Pool
	Context      context.Context
	TableName    string
	Pagination   *Pagination
	Sort         *Sort
	Filters      map[string][]string
	FilterRules  map[string]FilterRule
	SelectFields []string
}

type GetEntityCountPayload struct {
	Pool        *pgxpool.Pool
	Context     context.Context
	TableName   string
	Filters     map[string][]string
	FilterRules map[string]FilterRule
}

type ActionFuncResponse struct {
	Status int
	Data   interface{}
	Error  error
}

type SwaggerSuccessResponse[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}

type GetByIdParams struct {
	ID int `uri:"id" binding:"required"`
}

type UpdateEntityPayload struct {
	ID      int
	Pool    *pgxpool.Pool
	Context context.Context
	Updates map[string]interface{}
	Table   string
}

type Empty struct{}
