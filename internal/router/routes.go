package router

import (
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/internal/constanta"
	"go-boilerplate/internal/middleware"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	user := InitUser(db)

	// Public routes
	AuthRoute(app, db)

	api := app.Group("/api")

	// Protected routes
	auth := api.Group("/user")
	auth.Get("/profile", middleware.AuthMiddleware(constanta.MenuUserActionRead), user.Profile)
}

func AuthRoute(app *fiber.App, db *gorm.DB) {
	auth := InitAuth(db)

	api := app.Group("/api")

	// Public routes
	api.Post("/login", auth.Login)
	api.Post("/register", auth.Register)
}
