package router

import (
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/internal/middleware"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	user := InitUser(db)

	api := app.Group("/api")

	// Public routes
	api.Post("/login", user.Login)
	api.Post("/register", user.Register)

	// Protected routes
	auth := api.Group("/user", middleware.AuthMiddleware)
	auth.Get("/profile", user.Profile)
}
