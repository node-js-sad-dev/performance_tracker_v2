package handler

import (
	"mime/multipart"
	"performance_tracker_v2_be/config"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
)

// Global decoder caches struct metadata for much better performance across requests
var formDecoder = schema.NewDecoder()

func Extract[
	Body any,
	Query any,
	Params any,
](config *config.Config, c *gin.Context) (*ExtractorResult[Body, Query, Params], error) {
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

	headers := make(map[string]string)
	for k, v := range c.Request.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}

	params := new(Params)
	if err := c.ShouldBindUri(params); err != nil {
		return nil, err
	}

	queryParams := new(Query)
	if err := c.ShouldBindQuery(queryParams); err != nil {
		return nil, err
	}

	body := new(Body)
	var files map[string][]*multipart.FileHeader

	if c.Request.ContentLength > 0 {
		contentType := c.ContentType()

		switch contentType {
		case "application/json":
			if err := c.ShouldBindJSON(body); err != nil {
				return nil, err
			}

		case "multipart/form-data", "application/x-www-form-urlencoded":

			if contentType == "multipart/form-data" {
				form, err := c.MultipartForm()
				if err != nil {
					return nil, err
				}

				if err := formDecoder.Decode(body, form.Value); err != nil {
					return nil, err
				}

				files = form.File
			} else {
				if err := c.Request.ParseForm(); err != nil {
					return nil, err
				}

				if err := formDecoder.Decode(body, c.Request.PostForm); err != nil {
					return nil, err
				}
			}
		}
	}

	return &ExtractorResult[Body, Query, Params]{
		Params:      params,
		Pagination:  pagination,
		Sort:        sort,
		QueryParams: queryParams,
		Body:        body,
		Headers:     &headers,
		Context:     c.Request.Context(),
		Config:      config,
		Files:       files,
	}, nil
}
