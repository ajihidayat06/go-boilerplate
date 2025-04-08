package seeder

import (
	"errors"
	"go-boilerplate/internal/models"
	"go-boilerplate/pkg/logger"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedSuperAdmin(db *gorm.DB) error {
	if db == nil {
		err := errors.New("database connection is nil")
		logger.Error("Database connection is nil", err)
		return err
	}

	// Periksa apakah role superadmin sudah ada
	var role models.Roles
	if err := db.Where("name = ?", "superadmin").First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Buat role superadmin
			role = models.Roles{
				Name: "superadmin",
				Code: "super-admin",
			}
			if err := db.Create(&role).Error; err != nil {
				logger.Error("Failed to create superadmin role", err)
				return err
			}
		} else {
			logger.Error("Failed to query superadmin role", err)
			return err
		}
	}

	// Ambil email dan password dari environment variables
	superAdminEmail := os.Getenv("SUPERADMIN_EMAIL")
	if superAdminEmail == "" {
		superAdminEmail = "superadmin@example.com" // Default value
	}

	superAdminPassword := os.Getenv("SUPERADMIN_PASSWORD")
	if superAdminPassword == "" {
		superAdminPassword = "superadmin123" // Default value
	}

	// Periksa apakah user superadmin sudah ada
	var user models.User
	if err := db.Where("email = ?", superAdminEmail).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Hash password
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(superAdminPassword), bcrypt.DefaultCost)
			if err != nil {
				logger.Error("Failed to hash password", err)
				return err
			}

			// Buat user superadmin
			user = models.User{
				Username:  "superadmin",
				Email:     superAdminEmail,
				Password:  string(hashedPassword),
				RoleID:    role.ID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := db.Create(&user).Error; err != nil {
				logger.Error("Failed to create superadmin user", err)
				return err
			}
		} else {
			logger.Error("Failed to query superadmin user", err)
			return err
		}
	}

	logger.Info("Superadmin seeding completed successfully", nil)
	return nil
}
