package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-boilerplate/internal/constanta"
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/utils"
	"go-boilerplate/internal/utils/errors"
	"go-boilerplate/pkg/logger"
	"os"
	"strings"
	"time"
)

func GenerateTokenUserDashboard(user models.UserLogin) (string, error) {
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

func AuthMiddlewareDashboard(menuAction string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil header Authorization
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			logger.Error("Missing Authorization header", nil)
			return utils.SetResponseUnauthorized(c, "Missing Authorization header")
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
			return utils.SetResponseUnauthorized(c, errors.ErrMessageInvalidOrExpiredToken)
		}

		// Ambil claims dan periksa validitas token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return utils.SetResponseUnauthorized(c, errors.ErrMessageInvalidToken)
		}

		// Cek expired
		exp := int64(claims["exp"].(float64))
		if time.Now().Unix() > exp {
			return utils.SetResponseUnauthorized(c, errors.ErrMessageExpiredToken)
		}

		rolePermissionsInf, exists := claims["role_permissions"]
		if !exists {
			return utils.SetResponseForbiden(c, errors.ErrMessageForbidden)
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
					GroupMenu:   permMap["group_menu"].(string),
					Action:      permMap["action"].(string),
					AccessScope: permMap["access_scope"].(string),
				}
				permissions = append(permissions, permission)
			} else {
				return utils.SetResponseForbiden(c, "Invalid permissions format")
			}
		}

		// Validasi apakah user memiliki permission sesuai menuAction
		isValid, scope := validateUserScopePermissionDashboard(permissions, menuAction)
		if !isValid {
			return utils.SetResponseForbiden(c, errors.ErrMessageForbidden)
		}

		// Simpan user_id ke context agar bisa digunakan di handler selanjutnya
		c.Locals(constanta.AuthUserID, uint(claims["user_id"].(float64)))
		c.Locals(constanta.AuthRoleID, uint(claims["role_id"].(float64)))
		c.Locals(constanta.AuthRoleName, claims["role_name"].(string))
		c.Locals(constanta.Scope, scope)
		return c.Next()
	}
}

func validateUserScopePermissionDashboard(userPermissions []models.Permissions, menuAction string) (bool, string) {
	if len(userPermissions) == 0 {
		return false, ""
	}

	parts := strings.Split(menuAction, ":")
	if len(parts) != 2 {
		return false, ""
	}
	menu := parts[0]
	action := parts[1]

	for _, permission := range userPermissions {
		if permission.GroupMenu == menu && permission.Action == action {
			return true, permission.AccessScope
		}
	}
	return false, ""
}

func GenerateTokenUser(user models.UserLogin) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token berlaku 24 jam
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
			return utils.SetResponseUnauthorized(c, "Missing Authorization header")
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
			return utils.SetResponseUnauthorized(c, errors.ErrMessageInvalidOrExpiredToken)
		}

		// Ambil claims dan periksa validitas token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return utils.SetResponseUnauthorized(c, errors.ErrMessageInvalidToken)
		}

		// Cek expired
		exp := int64(claims["exp"].(float64))
		if time.Now().Unix() > exp {
			return utils.SetResponseUnauthorized(c, errors.ErrMessageExpiredToken)
		}

		// Simpan user_id ke context agar bisa digunakan di handler selanjutnya
		c.Locals(constanta.AuthUserID, uint(claims["user_id"].(float64)))
		return c.Next()
	}
}
