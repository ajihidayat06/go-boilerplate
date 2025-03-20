package router

import (
	"go-boilerplate/internal/constanta"
	"go-boilerplate/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	DashboardRoute(app, db)

	WebRoute(app, db)
}

func DashboardRoute(app *fiber.App, db *gorm.DB) {
	auth := InitAuth(db)
	user := InitUserDahboard(db)
	category := InitCategoryDashboard(db)

	api := app.Group("/api/dashboard")
	// Public routes
	api.Post("/login", auth.LoginDashboard)
	api.Post("/logout", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionRead), auth.LogoutDashboard)

	UserRoutesDashboard(api, user)
	CategoryRoutesdashboard(api, category)
}

func WebRoute(app *fiber.App, db *gorm.DB) {
	user := InitUser(db)

	api := app.Group("/api")

	api.Post("/login", user.Login)
	api.Post("/register", user.Register)
	api.Post("/logout", middleware.AuthMiddleware(), user.Logout)

	// Protected routes
	UserRoutesWeb(api, user)
}
