package http

import "github.com/gin-gonic/gin"

func DecomposeRequestForGetListMethod(c *gin.Context) (*Sort, *Pagination, map[string][]string, error) {
	var req QueryRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		return nil, nil, nil, err
	}

	// todo -> need better approach for defaults
	// Default values if empty
	if req.Page == 0 {
		req.Page = 1
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	filters := make(map[string][]string)
	rawQuery := c.Request.URL.Query()

	reserved := map[string]bool{
		"page":      true,
		"limit":     true,
		"sortBy":    true,
		"sortOrder": true,
	}

	for key, values := range rawQuery {
		if !reserved[key] && len(values) > 0 {
			filters[key] = values
		}
	}

	return &req.Sort, &req.Pagination, filters, nil
}

func DecomposeRequestForGetOneMethod(c *gin.Context) string {
	id := c.Param("id")

	return id
}

//func DecomposeRequestForCreateMethod(c *gin.Context, obj interface{}) error {
//	if err := c.ShouldBindJSON(obj); err != nil {
//
//	}
//}
