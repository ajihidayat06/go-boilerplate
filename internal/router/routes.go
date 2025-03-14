package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	DashboardRoute(app, db)

	WebRoute(app, db)
}

func DashboardRoute(app *fiber.App, db *gorm.DB) {
	auth := InitAuth(db)

	api := app.Group("/api/dashboard")
	// Public routes
	api.Post("/login", auth.Login)
	UserDashboardRoutes(api, auth)
}

func WebRoute(app *fiber.App, db *gorm.DB) {
	user := InitUser(db)

	api := app.Group("/api")

	api.Post("/login", user.Login)
	api.Post("/register", user.Register)

	// Protected routes
	UserWebRoutes(api, user)
}
