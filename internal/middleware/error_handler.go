package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/internal/utils"
	"go-boilerplate/pkg/logger"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	logger.Error("Unhandled error", err)
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return utils.SetResponseAPI(c, code, "unhandled error", err.Error(), nil)
}
