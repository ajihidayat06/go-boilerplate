package utils

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

func SetResponseAPI(c *fiber.Ctx, status int, message, err string, data interface{}) error {
	response := APIResponse{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   err,
	}

	return c.Status(status).JSON(response)
}

func SetResponseOK(c *fiber.Ctx, message string, data interface{}) error {
	return SetResponseAPI(c, http.StatusOK, message, "", data)
}

func SetResponseBadRequest(c *fiber.Ctx, message string, err error) error {
	return SetResponseAPI(c, http.StatusBadRequest, message, err.Error(), nil)
}

func SetResponseInternalServerError(c *fiber.Ctx, message string, err error) error {
	return SetResponseAPI(c, http.StatusInternalServerError, message, err.Error(), nil)
}
