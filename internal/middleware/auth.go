package middleware

import (
	"context"
	"errors"
	"go-boilerplate/internal/constanta"
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/utils"
	errorutils "go-boilerplate/internal/utils/errors"
	"go-boilerplate/pkg/logger"
	"go-boilerplate/pkg/redis"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokenUserDashboard(user models.UserLogin) (string, error) {
	var permissions []map[string]interface{}
	for _, rolePermission := range user.RolePermissions {
		permissions = append(permissions, map[string]interface{}{
			"group_menu":   rolePermission.Permissions.GroupMenu,
			"action":       rolePermission.Permissions.Action,
			"access_scope": rolePermission.AccessScope, // Ambil AccessScope dari RolePermissions
		})
	}

	claims := jwt.MapClaims{
		"user_id":          user.ID,
		"role_id":          user.RoleID,
		"role_name":        user.RoleName,
		"role_permissions": permissions,                           // Simpan permissions dalam bentuk slice dari map
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
			return utils.SetResponseUnauthorized(c, "Missing Authorization header", "")
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
			return utils.SetResponseUnauthorized(c, errorutils.ErrMessageInvalidOrExpiredToken, "")
		}

		// Ambil claims dan periksa validitas token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return utils.SetResponseUnauthorized(c, errorutils.ErrMessageInvalidToken, "")
		}

		// Periksa token di Redis
		isValid, err := IsTokenInRedis(c.Context(), tokenString)
		if err != nil {
			logger.Error("Failed to validate token in Redis", err)
			return utils.SetResponseInternalServerError(c, "Failed to validate token", err)
		}
		if !isValid {
			return utils.SetResponseUnauthorized(c, "Token is not valid", "")
		}

		// Cek expired
		exp := int64(claims["exp"].(float64))
		if time.Now().Unix() > exp {
			return utils.SetResponseUnauthorized(c, errorutils.ErrMessageExpiredToken, "")
		}

		// Ambil role_permissions dari token
		rawPermissions, exists := claims["role_permissions"]
		if !exists {
			return utils.SetResponseForbiden(c, errorutils.ErrMessageForbidden)
		}

		// Konversi role_permissions ke []models.RolePermissions
		permissionsData, ok := rawPermissions.([]interface{})
		if !ok {
			return utils.SetResponseForbiden(c, "Invalid permissions data")
		}

		var rolePermissions []models.RolePermissions
		for _, item := range permissionsData {
			if permMap, ok := item.(map[string]interface{}); ok {
				rolePermission := models.RolePermissions{
					Permissions: models.Permissions{
						GroupMenu: permMap["group_menu"].(string),
						Action:    permMap["action"].(string),
					},
					AccessScope: permMap["access_scope"].(string), // Ambil AccessScope
				}
				rolePermissions = append(rolePermissions, rolePermission)
			} else {
				return utils.SetResponseForbiden(c, "Invalid permissions format")
			}
		}

		// Validasi apakah user memiliki permission sesuai menuAction
		isValid, scope := validateUserScopePermissionDashboard(rolePermissions, menuAction)
		if !isValid {
			return utils.SetResponseForbiden(c, errorutils.ErrMessageForbidden)
		}

		// Simpan user_id dan scope ke context agar bisa digunakan di handler selanjutnya
		c.Locals(constanta.AuthUserID, uint(claims["user_id"].(float64)))
		c.Locals(constanta.AuthRoleID, uint(claims["role_id"].(float64)))
		c.Locals(constanta.AuthRoleName, claims["role_name"].(string))
		c.Locals(constanta.Scope, scope)
		return c.Next()
	}
}

func CheckAdminRoleMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleName := c.Locals(constanta.AuthRoleName)
		if roleName == nil || (roleName != "admin" && roleName != "superadmin") {
			logger.Error("User does not have admin role", nil)
			return utils.SetResponseForbiden(c, errorutils.ErrMessageForbidden)
		}
		return c.Next()
	}
}

func validateUserScopePermissionDashboard(rolePermissions []models.RolePermissions, menuAction string) (bool, string) {
	if len(rolePermissions) == 0 {
		return false, ""
	}

	parts := strings.Split(menuAction, ":")
	if len(parts) != 2 {
		return false, ""
	}
	menu := parts[0]
	action := parts[1]

	for _, rolePermission := range rolePermissions {
		if rolePermission.Permissions.GroupMenu == menu && rolePermission.Permissions.Action == action {
			return true, rolePermission.AccessScope // Kembalikan AccessScope
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

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.SetResponseUnauthorized(c, "Missing Authorization header", "")
		}

		tokenString := utils.ExtractBearerToken(authHeader)

		// Validasi token JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			return utils.SetResponseUnauthorized(c, "Invalid token", err.Error())
		}

		// Periksa token di Redis
		isValid, err := IsTokenInRedis(c.Context(), tokenString)
		if err != nil {
			return utils.SetResponseInternalServerError(c, "Failed to validate token", err)
		}
		if !isValid {
			return utils.SetResponseUnauthorized(c, "Token is not valid", "")
		}

		// Simpan informasi user ke context
		claims := token.Claims.(jwt.MapClaims)
		c.Locals("user_id", claims["user_id"])
		c.Locals("role_id", claims["role_id"])

		return c.Next()
	}
}

func GenerateTemporaryToken(user models.UserLogin) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role_id": user.RoleID,
		"exp":     time.Now().Add(time.Minute * 5).Unix(), // Temporary token berlaku 5 menit
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}

func ValidateTemporaryToken(temporaryToken string) (models.UserLogin, error) {
	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(temporaryToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return models.UserLogin{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return models.UserLogin{}, errors.New("invalid token")
	}

	// Ambil data user dari token
	user := models.UserLogin{
		ID:     int64(claims["user_id"].(float64)),
		RoleID: int64(claims["role_id"].(float64)),
	}

	return user, nil
}

func SaveTokenToRedis(ctx context.Context, token string, exp time.Time) error {
	ttl := time.Until(exp)
	return redis.SetToRedisWithTTL(ctx, token, true, ttl)
}

func IsTokenInRedis(ctx context.Context, token string) (bool, error) {
	result, err := redis.GetFromRedis(ctx, token)
    if err != nil {
        return false, err
    }
    return result == "true", nil
}

func DeleteTokenFromRedis(ctx context.Context, token string) error {
	return redis.DeleteFromRedis(ctx, token)
}
