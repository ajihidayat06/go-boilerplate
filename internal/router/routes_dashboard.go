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
	userDashboard.Get("/", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionRead), handler.CreateUserDashboard) //TODO: diganti jadi get list user
	userDashboard.Get("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionRead), handler.GetUserByID)
	userDashboard.Put("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionWrite), handler.CreateUserDashboard)    //TODO: diganti jadi update user by ID
	userDashboard.Delete("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionWrite), handler.CreateUserDashboard) //TODO: diganti jadi delete user by ID
}
