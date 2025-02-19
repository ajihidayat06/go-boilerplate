package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/pkg/logger"
	"time"
)

func LoggingMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	duration := time.Since(start)

	logger.Info("Request", map[string]interface{}{
		"timestamp": time.Now().Format(time.RFC3339),
		"method":    c.Method(),
		"path":      c.Path(),
		"status":    c.Response().StatusCode(),
		"duration":  duration.String(),
	})

	if err != nil {
		logger.Error("Middleware error", err)
	}

	return err
}
