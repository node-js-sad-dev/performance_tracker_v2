package car

import (
	"log/slog"
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
//	@Param			name		query		string	false	"Name filter (fuzzy)"
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
//	@Description	Create a new car with an image and details
//	@Tags			Car
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			file		formData	file	true	"Car image file"
//	@Param			name		formData	string	true	"Name of the car"
//	@Param			description	formData	string	false	"Detailed car description"
//	@Success		200			{object}	swagger.SuccessResponse[models.Car]
//	@Router			/car [post]
func (c *Controller) Create(extraction *handler.ExtractorResult[CreateRequest, handler.Empty, handler.Empty]) *handler.ActionFuncResponse {
	carWithSameName, err := c.Service.GetByName(extraction.Context, extraction.Body.Name, nil)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	if carWithSameName != nil {
		return responses.CommonErrorResponse(409, "car with this name already exists")
	}

	filePath, err := files.UploadFileAndGetLink("file", extraction.Files, false)
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
//	@Produce		json
//	@Param			id			path		number	true	"Car ID"
//	@Param			file		formData	file	false	"Car image file"
//	@Param			name		formData	string	false	"Name of the car"
//	@Param			description	formData	string	false	"Detailed car description"
//	@Param			image		formData	string	false	"Image to update if file not passed"
//	@Success		200			{object}	swagger.SuccessResponse[handler.Empty]
//	@Router			/car/{id} [patch]
func (c *Controller) UpdateByID(extraction *handler.ExtractorResult[UpdateRequestParsed, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	carFromDB, err := c.Service.GetByID(extraction.Context, extraction.Params.ID)
	if err != nil {
		return responses.DbErrorResponse(err)
	}
	if carFromDB == nil {
		return responses.NotFoundErrorResponse("car")
	}

	if !extraction.Body.Name.GetIsNull() && extraction.Body.Name.GetIsSet() {
		// cast to string is needed cause GetValue returns any for now and i dont know how to fix it(((
		carWithSameName, err := c.Service.GetByName(extraction.Context, extraction.Body.Name.GetValue().(string), &extraction.Params.ID)
		if err != nil {
			return responses.DbErrorResponse(err)
		}
		if carWithSameName != nil {
			return responses.CommonErrorResponse(409, "car with this name already exists")
		}
	}

	filePath, err := files.UploadFileAndGetLink("file", extraction.Files, false)
	if err != nil {
		return responses.CommonErrorResponse(500, err.Error())
	}

	clientRequestedImageRemoval := extraction.Body.Image.GetIsNull() && extraction.Body.Image.GetIsSet()
	shouldRemoveOldImage := carFromDB.Image != nil && (filePath != nil || clientRequestedImageRemoval)

	updateBody := &UpdateRequestParsed{
		Name:        extraction.Body.Name,
		Description: extraction.Body.Description,
	}

	if filePath != nil {
		updateBody.Image = handler.OptionalBodyField[string]{
			Value:  *filePath,
			IsSet:  true,
			IsNull: false,
		}
	} else {
		updateBody.Image = extraction.Body.Image
	}

	err = c.Service.UpdateByID(extraction.Context, extraction.Params.ID, updateBody)
	if err != nil {
		if filePath != nil {
			_ = files.RemoveFile(*filePath)
		}
		return responses.DbErrorResponse(err)
	}

	if shouldRemoveOldImage {
		err := files.RemoveFile(*carFromDB.Image)
		if err != nil {
			slog.Error("Failed to remove old image on car update")
		}
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
