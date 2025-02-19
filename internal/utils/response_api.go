package utils

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type APIResponse[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
	Error   string `json:"error"`
}

func SetResponseAPI[T any](c *fiber.Ctx, status int, message, err string, data T) error {
	response := APIResponse[T]{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   err,
	}
	return c.Status(status).JSON(response)
}

func SetResponseOK[T any](c *fiber.Ctx, message string, data T) error {
	return SetResponseAPI[T](c, http.StatusOK, message, "", data)
}

func SetResponseBadRequest(c *fiber.Ctx, message string, err error) error {
	return SetResponseAPI[any](c, http.StatusBadRequest, message, err.Error(), nil)
}

func SetResponseInternalServerError(c *fiber.Ctx, message string, err error) error {
	return SetResponseAPI[any](c, http.StatusInternalServerError, message, err.Error(), nil)
}
