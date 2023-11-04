package cerrors

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

// GeneralResponse 구조체는 3중 참조를 막기 위한 구조체 입니다.
type generalResponse struct {
	Status  uint   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewErrorHandler() func(*fiber.Ctx, error) error {
	return func(c *fiber.Ctx, err error) error {

		if _, isValidationError := err.(validationErr); isValidationError {
			data := err.Error()
			var messages []map[string]interface{}

			errJson := json.Unmarshal([]byte(data), &messages)
			Sniff(errJson)
			return c.Status(fiber.StatusBadRequest).JSON(generalResponse{
				Status:  fiber.StatusBadRequest,
				Message: "Bad Request",
				Data:    messages,
			})
		}

		if _, isParsingErr := err.(parsingErr); isParsingErr {
			return c.Status(fiber.StatusBadRequest).JSON(generalResponse{
				Status:  fiber.StatusBadRequest,
				Message: "Parsing error",
				Data:    err.Error(),
			})
		}

		if _, isAuthorizationErr := err.(authorizationErr); isAuthorizationErr {
			return c.Status(fiber.StatusUnauthorized).JSON(generalResponse{
				Status:  fiber.StatusUnauthorized,
				Message: "UnAuthorized",
				Data:    err.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(generalResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "General Error",
			Data:    err.Error(),
		})
	}
}
