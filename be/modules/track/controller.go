package track

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

// GetList @Summary Get tracks
//
//	@Description	Get all tracks
//	@Tags			Track
//	@Produce		json
//	@Param			page		query		int		false	"Page"
//	@Param			limit		query		int		false	"Limit"
//	@Param			sortBy		query		string	false	"Sort by"
//	@Param			sortOrder	query		string	false	"Sort order"
//	@Param			name		query		string	false	"Name filter (fuzzy)"
//	@Success		200			{object}	swagger.SuccessResponse[GetListResponse]
//	@Router			/track [get]
func (c *Controller) GetList(extraction *handler.ExtractorResult[handler.Empty, GetFilters, handler.Empty]) *handler.ActionFuncResponse {
	tracks, err := c.Service.GetList(extraction.Context, extraction.Pagination, extraction.Sort, extraction.QueryParams)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	totalCount, err := c.Service.GetTotalCount(extraction.Context, extraction.QueryParams)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(GetListResponse{
		Tracks:     tracks,
		TotalCount: totalCount,
	})
}

// Create @Summary Create track
//
//	@Description	Create a new track with an image and details
//	@Tags			Track
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			file		formData	file	false	"Track image file"
//	@Param			name		formData	string	true	"Name of the track"
//	@Param			description	formData	string	true	"Detailed track description"
//	@Success		200			{object}	swagger.SuccessResponse[models.Track]
//	@Router			/track [post]
func (c *Controller) Create(extraction *handler.ExtractorResult[CreateRequest, handler.Empty, handler.Empty]) *handler.ActionFuncResponse {
	trackWithSameName, err := c.Service.GetByName(extraction.Context, extraction.Body.Name, nil)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	if trackWithSameName != nil {
		return responses.CommonErrorResponse(409, "track with this name already exists")
	}

	filePath, err := files.UploadFileAndGetLink("file", extraction.Files, false)
	if err != nil {
		return responses.CommonErrorResponse(500, err.Error())
	}

	trackID, err := c.Service.Create(extraction.Context, &CreateRequestParsed{
		Name:        extraction.Body.Name,
		Description: extraction.Body.Description,
		Image:       filePath,
	})
	if err != nil {
		if filePath != nil {
			_ = files.RemoveFile(*filePath)
		}
		return responses.DbErrorResponse(err)
	}

	track, err := c.Service.GetByID(extraction.Context, trackID)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(track)
}

// GetByID @Summary Get track by ID
//
//	@Description	Get a track by its ID
//	@Tags			Track
//	@Produce		json
//	@Param			id	path		string	true	"Track ID"
//	@Success		200	{object}	swagger.SuccessResponse[models.Track]
//	@Router			/track/{id} [get]
func (c *Controller) GetByID(extraction *handler.ExtractorResult[handler.Empty, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	track, err := c.Service.GetByID(extraction.Context, extraction.Params.ID)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.SuccessResponse(track)
}

// UpdateByID @Summary Update track by ID
//
//	@Description	Update a track by its ID
//	@Tags			Track
//	@Produce		json
//	@Param			id			path		number	true	"Track ID"
//	@Param			file		formData	file	false	"Track image file"
//	@Param			name		formData	string	false	"Name of the track"
//	@Param			description	formData	string	false	"Detailed track description"
//	@Param			image		formData	string	false	"Image to update if file not passed"
//	@Success		200			{object}	swagger.SuccessResponse[handler.Empty]
//	@Router			/track/{id} [patch]
func (c *Controller) UpdateByID(extraction *handler.ExtractorResult[UpdateRequestParsed, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	trackFromDB, err := c.Service.GetByID(extraction.Context, extraction.Params.ID)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	if trackFromDB == nil {
		return responses.NotFoundErrorResponse("track")
	}

	if !extraction.Body.Name.GetIsNull() && extraction.Body.Name.GetIsSet() {
		trackWithSameName, err := c.Service.GetByName(extraction.Context, extraction.Body.Name.GetValue().(string), &extraction.Params.ID)
		if err != nil {
			return responses.DbErrorResponse(err)
		}

		if trackWithSameName != nil {
			return responses.CommonErrorResponse(409, "track with this name already exists")
		}
	}

	filePath, err := files.UploadFileAndGetLink("file", extraction.Files, false)
	if err != nil {
		return responses.CommonErrorResponse(500, err.Error())
	}

	clientRequestedImageRemoval := extraction.Body.Image.GetIsNull() && extraction.Body.Image.GetIsSet()
	shouldRemoveOldImage := trackFromDB.Image != nil && (filePath != nil || clientRequestedImageRemoval)

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
		if err := files.RemoveFile(*trackFromDB.Image); err != nil {
			slog.Error("Failed to remove old image on track update", "error", err)
		}
	}

	return responses.DefaultSuccessResponse()
}

// DeleteByID @Summary Delete track by ID
//
//	@Description	Delete a track by its ID
//	@Tags			Track
//	@Produce		json
//	@Param			id	path		string	true	"Track ID"
//	@Success		200	{object}	swagger.SuccessResponse[handler.Empty]
//	@Router			/track/{id} [delete]
func (c *Controller) DeleteByID(extraction *handler.ExtractorResult[handler.Empty, handler.Empty, handler.GetByIdParams]) *handler.ActionFuncResponse {
	track, err := c.Service.GetByID(extraction.Context, extraction.Params.ID)
	if err != nil {
		return responses.DbErrorResponse(err)
	}

	if track == nil {
		return responses.NotFoundErrorResponse("track")
	}

	if err := c.Service.DeleteByID(extraction.Context, extraction.Params.ID); err != nil {
		return responses.DbErrorResponse(err)
	}

	return responses.DefaultSuccessResponse()
}
