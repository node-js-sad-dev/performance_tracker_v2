package core

import (
	"performance_tracker_v2_be/config"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Extract[
	Body any,
	Query any,
	Params any,
](config *config.Config, pool *pgxpool.Pool, c *gin.Context) (*ExtractorResult[Body, Query, Params], error) {
	/*
		TODO
		need special type for body to handle null values,
		need distinction between empty body and body with empty fields,
		for now we will use pointer to struct and check if it's nil or not

	*/

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
		SortBy:    c.DefaultQuery("sortBy", "id"),
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

	return &ExtractorResult[Body, Query, Params]{
		Params:      params,
		Pagination:  pagination,
		Sort:        sort,
		QueryParams: queryParams,
		Body:        body,
		Headers:     &headers,
		Config:      config,
		Pool:        pool,
		Context:     c.Request.Context(),
	}, nil
}
