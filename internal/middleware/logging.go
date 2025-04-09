package middleware

import (
	"go-boilerplate/pkg/logger"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	// Simpan request body untuk logging
	var requestBody string
	if c.Request().Body() != nil {
		requestBody = string(c.Request().Body())
	}

	// Jalankan handler berikutnya
	err := c.Next()

	// Hitung durasi request
	duration := time.Since(start)

	// Simpan response body untuk logging (opsional)
	var responseBody string
	if c.Response().Body() != nil {
		responseBody = string(c.Response().Body())
	}

	// Log informasi request dan response
	logger.Info("Request Log", map[string]interface{}{
		"timestamp":     time.Now().Format(time.RFC3339),
		"method":        c.Method(),
		"path":          c.Path(),
		"query_params":  string(c.Request().URI().QueryString()),
		"request_body":  requestBody,
		"status":        c.Response().StatusCode(),
		"response_body": responseBody,
		"duration":      duration.String(),
	})

	// Log error jika ada
	if err != nil {
		logger.Error("Middleware error", err)
	}

	return err
}
