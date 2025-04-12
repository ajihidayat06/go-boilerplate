package utils

import (
	"go-boilerplate/internal/dto/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
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

func SetResponseUnauthorized(c *fiber.Ctx, message string, err string) error {
	return SetResponseAPI(c, http.StatusUnauthorized, message, err, nil)
}

func SetResponseForbiden(c *fiber.Ctx, message string) error {
	return SetResponseAPI(c, http.StatusForbidden, message, "", nil)
}

func SetResponseNotFound(c *fiber.Ctx, message string, err error) error {
	return SetResponseAPI(c, http.StatusNotFound, message, err.Error(), nil)
}

func MapToListResponse[T any](list []T, totalCount int64, page, pageSize int) response.ListResponse[T] {
	return response.ListResponse[T]{
		List:       list,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
	}
}
