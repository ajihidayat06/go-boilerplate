package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/internal/utils"
	"go-boilerplate/pkg/logger"
	"runtime"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	logger.Error("Unhandled error", err)
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return utils.SetResponseAPI(c, code, "unhandled error", err.Error(), nil)
}

func RecoverMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				// Menggunakan runtime.Caller untuk mendapatkan file dan baris.
				if _, file, line, ok := runtime.Caller(6); ok {
					logger.Error("Panic", fmt.Errorf("Panic terjadi di %s:%d - %v", file, line, r))
				} else {
					logger.Error("Panic", fmt.Errorf("Panic terjadi: %v", r))
				}

				utils.SetResponseInternalServerError(c, "Internal Server Error", fmt.Errorf("panic: %v", r))
			}
		}()
		return c.Next()
	}
}
