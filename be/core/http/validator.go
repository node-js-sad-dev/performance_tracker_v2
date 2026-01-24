package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate = validator.New()

type SchemaProvider func(c *gin.Context) (interface{}, error)

func ValidateSchema(schemaProvider SchemaProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the value to validate
		payload, err := schemaProvider(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Validate the struct
		if err := validate.Struct(payload); err != nil {
			var validationErrors validator.ValidationErrors
			errors.As(err, &validationErrors)

			var messages []string

			for _, fieldErr := range validationErrors {
				messages = append(messages, fieldErr.Error())
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"error": messages,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
