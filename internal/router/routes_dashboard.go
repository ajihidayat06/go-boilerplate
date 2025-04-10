package router

import (
	"go-boilerplate/internal/constanta"
	"go-boilerplate/internal/controllers/dashboard"
	"go-boilerplate/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(api fiber.Router, authController *dashboard.AuthController) {
	auth := api.Group("/auth")

	auth.Post("/validate", authController.ValidateCredentials) // Endpoint pertama
	auth.Post("/token", authController.GenerateAccessToken)    // Endpoint kedua

	auth.Post("/logout", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionRead), authController.LogoutDashboard)
}

func UserRoutesDashboard(api fiber.Router, handler *dashboard.UserDahboardController) {
	//protected routes
	userDashboard := api.Group("/user")
	userDashboard.Post("/", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionWrite), handler.CreateUserDashboard)
	userDashboard.Get("/", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionRead), handler.GetListUser)
	userDashboard.Get("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionRead), handler.GetUserByID)
	userDashboard.Put("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionWrite), handler.UpdateUserByID)
	userDashboard.Delete("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionWrite), handler.DeleteUserByID)
}

func CategoryRoutesdashboard(api fiber.Router, handler *dashboard.CategoryDashboardController) {
	// Protected routes
	category := api.Group("/category")
	category.Post("/", middleware.AuthMiddlewareDashboard(constanta.MenuCategoryActionWrite), handler.CreateCategory)
	category.Get("/", middleware.AuthMiddlewareDashboard(constanta.MenuCategoryActionRead), handler.GetListCategory)
	category.Get("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuCategoryActionRead), handler.GetCategoryByID)
	category.Put("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuCategoryActionWrite), handler.UpdateCategoryByID)
	category.Delete("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuCategoryActionWrite), handler.DeleteCategoryByID)
}

func RoleRoutesDashboard(api fiber.Router, handler *dashboard.RoleController) {
	// Protected routes
	role := api.Group("/role")
	role.Post("/", middleware.AuthMiddlewareDashboard(constanta.MenuRoleActionWrite), middleware.CheckAdminRoleMiddleware(), handler.CreateRole)
	role.Get("/", middleware.AuthMiddlewareDashboard(constanta.MenuRoleActionRead), middleware.CheckAdminRoleMiddleware(), handler.GetListRole)
	role.Get("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuRoleActionRead), middleware.CheckAdminRoleMiddleware(), handler.GetRoleByID)
	role.Put("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuRoleActionWrite), middleware.CheckAdminRoleMiddleware(), handler.UpdateRoleByID)
	role.Delete("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuRoleActionWrite), middleware.CheckAdminRoleMiddleware(), handler.DeleteRoleByID)
}

func PermissionRoutesDashboard(api fiber.Router, handler *dashboard.PermissionController) {
	// Protected routes
	permission := api.Group("/permission")
	permission.Post("/", middleware.AuthMiddlewareDashboard(constanta.MenuPermissionsActionWrite), middleware.CheckAdminRoleMiddleware(), handler.CreatePermission)
	permission.Get("/", middleware.AuthMiddlewareDashboard(constanta.MenuPermissionsActionRead), middleware.CheckAdminRoleMiddleware(), handler.GetListPermission)
	permission.Get("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuPermissionsActionRead), middleware.CheckAdminRoleMiddleware(), handler.GetPermissionByID)
	permission.Put("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuPermissionsActionWrite), middleware.CheckAdminRoleMiddleware(), handler.UpdatePermissionByID)
	permission.Delete("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuPermissionsActionWrite), middleware.CheckAdminRoleMiddleware(), handler.DeletePermissionByID)
}

func RolePermissionsRoutesDashboard(api fiber.Router, handler *dashboard.RolePermissionsController) {
	// Protected routes
	rolePermissions := api.Group("/role-permissions")
	rolePermissions.Post("/", middleware.AuthMiddlewareDashboard(constanta.MenuRolePermissionsActionWrite), middleware.CheckAdminRoleMiddleware(), handler.CreateRolePermission)
	rolePermissions.Get("/", middleware.AuthMiddlewareDashboard(constanta.MenuRolePermissionsActionRead), middleware.CheckAdminRoleMiddleware(), handler.GetListRolePermissions)
	rolePermissions.Get("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuRolePermissionsActionRead), middleware.CheckAdminRoleMiddleware(), handler.GetRolePermissionByID)
	rolePermissions.Put("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuRolePermissionsActionWrite), middleware.CheckAdminRoleMiddleware(), handler.UpdateRolePermissionByID)
	rolePermissions.Delete("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuRolePermissionsActionWrite), middleware.CheckAdminRoleMiddleware(), handler.DeleteRolePermissionByID)
}
