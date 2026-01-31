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
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if limit < 1 {
		limit = 10
	}

	pagination := &Pagination{Page: page, Limit: limit}

	sort := &Sort{
		SortBy:    c.DefaultQuery("sortBy", "created_at"),
		SortOrder: c.DefaultQuery("sortOrder", "desc"),
	}

	params := new(Params)
	if err := c.ShouldBindUri(params); err != nil {
		return nil, err
	}

	queryParams := new(Query)
	if err := c.ShouldBindQuery(queryParams); err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	for k, v := range c.Request.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}

	body := new(Body)
	if c.Request.ContentLength > 0 {
		if err := c.ShouldBindJSON(body); err != nil {
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
