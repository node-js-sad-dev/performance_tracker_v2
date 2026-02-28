package car

import (
	"performance_tracker_v2_be/core/files"
	"performance_tracker_v2_be/core/handler"
	"performance_tracker_v2_be/core/responses"

	_ "performance_tracker_v2_be/core/swagger"
	_ "performance_tracker_v2_be/db/main-db/models"
)

type Controller struct {
	Service *Service
}

// GetList @Summary Get cars
//
//	@Description	Get all cars
//	@Tags			Car
//	@Produce		json
//	@Param			page		query		int		false	"Page"
//	@Param			limit		query		int		false	"Limit"
//	@Param			sortBy		query		string	false	"Sort by"
//	@Param			sortOrder	query		string	false	"Sort order"
//	@Success		200			{object}	swagger.SuccessResponse[GetListResponse]
//	@Router			/car [get]
func (c *Controller) GetList(extraction *handler.ExtractorResult[handler.Empty, GetFilters, handler.Empty]) *handler.ActionFuncResponse {
	cars, err := c.Service.GetList(extraction.Context, extraction.Pagination, extraction.Sort, extraction.QueryParams)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	totalCount, err := c.Service.GetTotalCount(extraction.Context, extraction.QueryParams)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(GetListResponse{
		Cars:       cars,
		TotalCount: totalCount,
	})
}

// Create @Summary      Create car
//
// @Description  Create a new car with an image and details
// @Tags         Car
// @Accept       multipart/form-data
// @Produce      json
// @Param        file         formData  file    true  "Car image file"
// @Param        name         formData  string  true  "Name of the car"
// @Param        description  formData  string  false "Detailed car description"
// @Success      200  {object}  swagger.SuccessResponse[models.Car]
// @Router       /car [post]
func (c *Controller) Create(extraction *handler.ExtractorResult[CreateRequest, handler.Empty, handler.Empty]) *handler.ActionFuncResponse {
	carImage, err := files.GetFileInfoFromExtraction("file", extraction.Files)

	if err != nil {
		return responses.CommonErrorResponse(500, err.Error())
	}

	if carImage == nil {
		return responses.CommonErrorResponse(400, "file is required")
	}

	filePath, err := files.SaveFile(carImage)
	if err != nil {
		return responses.CommonErrorResponse(500, err.Error())
	}

	carId, err := c.Service.Create(extraction.Context, &CreateRequestParsed{
		Name:        extraction.Body.Name,
		Description: extraction.Body.Description,
		Image:       filePath,
	})

	if err != nil {
		return responses.DbErrorResponse(err)
	}

	car, err := c.Service.GetByID(extraction.Context, carId)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(car)
}

// GetByID @Summary Get car by ID
//
//	@Description	Get a car by its ID
//	@Tags			Car
//	@Produce		json
//	@Param			id	path		string	true	"Car ID"
//	@Success		200	{object}	swagger.SuccessResponse[models.Car]
//	@Router			/car/{id} [get]
func (c *Controller) GetByID(extraction *handler.ExtractorResult[handler.Empty, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	car, err := c.Service.GetByID(extraction.Context, extraction.Params.ID)

	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(car)
}

// UpdateByID @Summary Update car by ID
//
//	@Description	Update a car by its ID
//	@Tags			Car
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string				true	"Car ID"
//	@Param			car	body		UpdateRequestInput	true	"Car fields to update"
//	@Success		200	{object}	swagger.SuccessResponse[handler.Empty]
//	@Router			/car/{id} [patch]
func (c *Controller) UpdateByID(extraction *handler.ExtractorResult[UpdateRequestParsed, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	car, err := c.Service.GetByID(extraction.Context, extraction.Params.ID)

	if err != nil {
		return responses.DbErrorResponse(err)
	}

	if car == nil {
		return responses.NotFoundErrorResponse("car")
	}

	if err := c.Service.UpdateByID(extraction.Context, extraction.Params.ID, extraction.Body); err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.DefaultSuccessResponse()
}

// DeleteByID @Summary Delete car by ID
//
//	@Description	Delete a car by its ID
//	@Tags			Car
//	@Produce		json
//	@Param			id	path		string	true	"Car ID"
//	@Success		200	{object}	swagger.SuccessResponse[handler.Empty]
//	@Router			/car/{id} [delete]
func (c *Controller) DeleteByID(extraction *handler.ExtractorResult[handler.Empty, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	car, err := c.Service.GetByID(extraction.Context, extraction.Params.ID)

	if err != nil {
		return responses.DbErrorResponse(err)
	}

	if car == nil {
		return responses.NotFoundErrorResponse("car")
	}

	if err := c.Service.DeleteByID(extraction.Context, extraction.Params.ID); err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.DefaultSuccessResponse()
}
