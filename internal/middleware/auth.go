package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/utils"
	"go-boilerplate/pkg/logger"
	"os"
	"strings"
	"time"
)

func GenerateTokenUser(user models.UserLogin) (string, error) {
	claims := jwt.MapClaims{
		"user_id":          user.ID,
		"role_id":          user.RoleID,
		"role_name":        user.RoleName,
		"role_permissions": user.Permissions,
		"exp":              time.Now().Add(time.Hour * 24).Unix(), // Token berlaku 24 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}

func AuthMiddleware(menuAction string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil header Authorization
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			logger.Error("Missing Authorization header", nil)
			return utils.SetResponseUnauthorized(c, "Unauthorized")
		}

		// Hapus prefix "Bearer " jika ada
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		secret := os.Getenv("JWT_SECRET")

		// Parse token JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				logger.Error("Unexpected signing method", nil)
				return nil, fiber.ErrUnauthorized
			}
			return []byte(secret), nil
		})
		if err != nil {
			return utils.SetResponseUnauthorized(c, "Invalid or expired token")
		}

		// Ambil claims dan periksa validitas token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return utils.SetResponseUnauthorized(c, "Invalid token")
		}

		// Cek expired
		exp := int64(claims["exp"].(float64))
		if time.Now().Unix() > exp {
			return utils.SetResponseUnauthorized(c, "Token expired")
		}

		rolePermissionsInf, exists := claims["role_permissions"]
		if !exists {
			return utils.SetResponseForbiden(c, "Permissions deny")
		}

		// Konversi role_permissions ke []models.Permissions
		rawPermissions, ok := rolePermissionsInf.([]interface{})
		if !ok {
			return utils.SetResponseForbiden(c, "Invalid permissions data")
		}

		var permissions []models.Permissions
		for _, item := range rawPermissions {
			// Pastikan item adalah map[string]interface{} sehingga kita bisa meng-cast field-nya
			if permMap, ok := item.(map[string]interface{}); ok {
				permission := models.Permissions{
					GroupMenu: permMap["group_menu"].(string),
					Action:    permMap["action"].(string),
				}
				permissions = append(permissions, permission)
			} else {
				return utils.SetResponseForbiden(c, "Invalid permissions format")
			}
		}

		// Validasi apakah user memiliki permission sesuai menuAction
		isValid := validateUserScopePermission(permissions, menuAction)
		if !isValid {
			return utils.SetResponseForbiden(c, "Permissions deny")
		}

		// Simpan user_id ke context agar bisa digunakan di handler selanjutnya
		c.Locals("user_id", uint(claims["user_id"].(float64)))
		c.Locals("role_id", uint(claims["user_id"].(float64)))
		c.Locals("role_name", claims["role_name"].(string))
		return c.Next()
	}
}

func validateUserScopePermission(userPermissions []models.Permissions, menuAction string) bool {
	if len(userPermissions) == 0 {
		return false
	}

	parts := strings.Split(menuAction, ":")
	if len(parts) != 2 {
		return false
	}
	menu := parts[0]
	action := parts[1]

	for _, permission := range userPermissions {
		if permission.GroupMenu == menu && permission.Action == action {
			return true
		}
	}
	return false
}
