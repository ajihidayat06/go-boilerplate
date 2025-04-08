package router

import (
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
	role := InitRoleDashboard(db)
	permissions := InitPermissionDashboard(db)
	rolePermissions := InitRolePermissionsDashboard(db)

	api := app.Group("/api/v1/dashboard")
	// Public routes

	AuthRoutes(app, auth)
	
	RoleRoutesDashboard(api, role)
	PermissionRoutesDashboard(api, permissions)
	RolePermissionsRoutesDashboard(api, rolePermissions)

	UserRoutesDashboard(api, user)
	CategoryRoutesdashboard(api, category)
}

func WebRoute(app *fiber.App, db *gorm.DB) {
	user := InitUser(db)

	api := app.Group("/api/v1")

	api.Post("/login", user.Login)
	api.Post("/register", user.Register)
	api.Post("/logout", middleware.AuthMiddleware(), user.Logout)

	// Protected routes
	UserRoutesWeb(api, user)
}
