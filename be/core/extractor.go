package core

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func Extract[
	Body any,
	Query any,
	Params any,
](c *gin.Context) (*ExtractorResult[Body, Query, Params, map[string]string], error) {

	// 1. Extract Pagination (Standard Logic)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if limit < 1 {
		limit = 10
	}

	pagination := &Pagination{Page: page, Limit: limit}

	// 2. Extract Sort (Standard Logic)
	sort := &Sort{
		SortBy:    c.DefaultQuery("sortBy", "created_at"),
		SortOrder: c.DefaultQuery("sortOrder", "desc"),
	}

	// 3. Extract & Validate URI Params
	// We use 'new' to create a pointer to the generic type
	params := new(Params)
	if err := c.ShouldBindUri(params); err != nil {
		return nil, err
	}

	// 4. Extract & Validate Custom Query Params
	// ShouldBindQuery binds ?key=value to the struct
	queryParams := new(Query)
	if err := c.ShouldBindQuery(queryParams); err != nil {
		return nil, err
	}

	// 5. Extract Headers (Manual extraction is usually still best for headers)
	headers := make(map[string]string)
	for k, v := range c.Request.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}

	// 6. Extract & Validate Body
	body := new(Body)
	// Only attempt to bind body if content is present
	if c.Request.ContentLength > 0 {
		if err := c.ShouldBindJSON(body); err != nil {
			// This returns an error if JSON is malformed OR if validation tags fail
			return nil, err
		}
	}

	return &ExtractorResult[Body, Query, Params, map[string]string]{
		Params:      params,
		Pagination:  pagination,
		Sort:        sort,
		QueryParams: queryParams,
		Body:        body,
		Headers:     &headers,
	}, nil
}
