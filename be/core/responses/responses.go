package responses

import (
	"database/sql"
	"errors"
	"log/slog"
	"performance_tracker_v2_be/core/handler"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

func SuccessResponse(data interface{}) *handler.ActionFuncResponse {
	return &handler.ActionFuncResponse{
		Status: 200,
		Data:   data,
		Error:  nil,
	}
}

func DefaultSuccessResponse() *handler.ActionFuncResponse {
	return &handler.ActionFuncResponse{
		Status: 200,
		Data: gin.H{
			"message": "success",
		},
		Error: nil,
	}
}

func CommonErrorResponse(status int, errorMessage string) *handler.ActionFuncResponse {
	return &handler.ActionFuncResponse{
		Status: status,
		Data:   nil,
		Error:  errors.New(errorMessage),
	}
}

func NotFoundErrorResponse(entityName string) *handler.ActionFuncResponse {
	return &handler.ActionFuncResponse{
		Status: 404,
		Data:   nil,
		Error:  errors.New(entityName + " not found"),
	}
}

func DbErrorResponse(err error) *handler.ActionFuncResponse {
	if errors.Is(err, sql.ErrNoRows) {
		return &handler.ActionFuncResponse{
			Status: 404,
			Error:  errors.New("record not found"),
		}
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return &handler.ActionFuncResponse{
				Status: 400,
				Error:  errors.New("duplicate record"),
			}
		}
	}

	slog.Error("Database error: %v", err)
	return &handler.ActionFuncResponse{
		Status: 500,
		Error:  errors.New("internal server error"),
	}
}
