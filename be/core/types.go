package core

import (
	"github.com/gin-gonic/gin"
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

type ExtractorResult[
	Body any,
	Query any,
	Params any,
	Headers map[string]string,
] struct {
	Params      *Params
	Pagination  *Pagination
	Sort        *Sort
	QueryParams *Query
	Body        *Body
	Headers     *Headers
	// @todo -> for projects with auth module add user here too
}

type GetEntityListPayload struct {
	Pool         *pgxpool.Pool
	Context      *gin.Context
	TableName    string
	Pagination   *Pagination
	Sort         *Sort
	Filters      map[string][]string
	FilterRules  map[string]FilterRule
	SelectFields []string
}
