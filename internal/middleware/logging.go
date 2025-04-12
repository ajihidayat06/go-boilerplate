package middleware

import (
	"encoding/json"
	"fmt"
	"go-boilerplate/pkg/logger"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	// Simpan request body untuk logging
	requestBody := string(c.Body())

	// Hindari mencatat informasi sensitif
	requestBody = sanitizeSensitiveData(requestBody)

	// Jalankan handler berikutnya
	err := c.Next()

	// Hitung durasi request
	duration := time.Since(start)

	// Siapkan data log
	logData := map[string]interface{}{
		"timestamp":    time.Now().Format(time.RFC3339),
		"method":       c.Method(),
		"path":         c.Path(),
		"query_params": string(c.Request().URI().QueryString()),
		"request_body": requestBody,
		"status":       c.Response().StatusCode(),
		"duration":     duration.String(),
	}

	// Log berdasarkan status
	if err != nil {
		// Jika terjadi error, log sebagai ERROR
		logData["error"] = sanitizeError(err).Error()
		logger.Error(fmt.Sprintf("Request Error: %v | Data: %v", err, logData), err)
	} else {
		// Jika tidak ada error, log sebagai INFO
		logData["error"] = sanitizeError(err).Error()
		logger.Info("Request Log", logData)
	}

	return err
}

// sanitizeSensitiveData mengganti informasi sensitif dalam request body dengan [FILTERED]
func sanitizeSensitiveData(body string) string {
	// Coba parsing body sebagai JSON
	var parsedBody map[string]interface{}
	if err := json.Unmarshal([]byte(body), &parsedBody); err != nil {
		// Jika gagal parsing, kembalikan body asli (tidak difilter)
		return body
	}

	// Daftar key sensitif
	sensitiveFields := []string{"password", "token", "api_key", "secret"}

	// Iterasi melalui key-value dan filter value sensitif
	for key, value := range parsedBody {
		for _, sensitiveField := range sensitiveFields {
			if strings.EqualFold(key, sensitiveField) {
				parsedBody[key] = "[FILTERED]" // Ganti value dengan [FILTERED]
			} else if strValue, ok := value.(string); ok {
				// Jika value adalah string, lakukan sanitasi
				parsedBody[key] = sanitizeString(strValue, sensitiveFields)
			}
		}
	}

	// Kembalikan body yang sudah difilter sebagai JSON string
	filteredBody, err := json.Marshal(parsedBody)
	if err != nil {
		// Jika gagal mengubah kembali ke JSON, kembalikan body asli
		return body
	}
	return string(filteredBody)
}

// sanitizeError memastikan error yang dicatat tidak mengandung informasi sensitif
func sanitizeError(err error) error {
	sensitiveKeywords := []string{"password", "token", "api_key", "secret"}
	sanitizedMessage := err.Error()
	for _, keyword := range sensitiveKeywords {
		if strings.Contains(strings.ToLower(sanitizedMessage), keyword) {
			sanitizedMessage = strings.ReplaceAll(sanitizedMessage, keyword, "[FILTERED]")
		}
	}
	return fiber.NewError(fiber.StatusInternalServerError, sanitizedMessage)
}

func sanitizeString(input string, sensitiveKeywords []string) string {
	for _, keyword := range sensitiveKeywords {
		if strings.Contains(strings.ToLower(input), keyword) {
			input = strings.ReplaceAll(input, keyword, "[FILTERED]")
		}
	}
	return input
}
