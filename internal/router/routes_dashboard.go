package router

import (
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/internal/constanta"
	"go-boilerplate/internal/controllers/dashboard"
	"go-boilerplate/internal/middleware"
)

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