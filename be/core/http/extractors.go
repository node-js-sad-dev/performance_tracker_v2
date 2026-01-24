package http

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBody[T any]() ParamFunc {
	return func(c *gin.Context) (interface{}, error) {
		var body T
		if err := c.ShouldBindJSON(&body); err != nil {
			return nil, fmt.Errorf("invalid JSON body: %w", err)
		}
		return &body, nil
	}
}

func GetQuery[T any]() ParamFunc {
	return func(c *gin.Context) (interface{}, error) {
		query := make(map[string]string)

		for key, values := range c.Request.URL.Query() {
			if len(values) > 0 {
				query[key] = values[0]
			}
		}

		jsonData, err := json.Marshal(query)
		if err != nil {
			return nil, err
		}

		var p T
		if err := json.Unmarshal(jsonData, &p); err != nil {
			return nil, err
		}

		return p, nil
	}
}

func GetParams[T any]() ParamFunc {
	return func(c *gin.Context) (interface{}, error) {
		params := make(map[string]string)

		for _, param := range c.Params {
			params[param.Key] = param.Value
		}

		jsonData, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}

		var result T
		if err := json.Unmarshal(jsonData, &result); err != nil {
			return nil, err
		}

		return result, nil
	}
}

func GetIdFromRequest(context *gin.Context) (int, error) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetRequestPagination(context *gin.Context) *Pagination {
	page, err := strconv.Atoi(context.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(context.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}

	return &Pagination{
		Page:  page,
		Limit: limit,
	}
}

func GetRequestSort(context *gin.Context) *Sort {
	sortBy := context.DefaultQuery("sortBy", "id")
	sortOrder := context.DefaultQuery("sortOrder", "asc")

	return &Sort{
		SortBy:    sortBy,
		SortOrder: sortOrder,
	}
}
