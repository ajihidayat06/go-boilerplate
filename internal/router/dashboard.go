package router

import (
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/internal/constanta"
	"go-boilerplate/internal/controllers/dashboard"
	"go-boilerplate/internal/middleware"
)

func UserDashboardRoutes(api fiber.Router, handler *dashboard.AuthController) {
	//protected routes
	userDashboard := api.Group("/user")
	userDashboard.Post("/", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionWrite), handler.CreateUserDashboard)
	userDashboard.Get("/", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionRead), handler.CreateUserDashboard)        //todo: diganti jadi get list user
	userDashboard.Get("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionRead), handler.CreateUserDashboard)     //todo: diganti jadi get user by ID
	userDashboard.Put("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionWrite), handler.CreateUserDashboard)    //todo: diganti jadi update user by ID
	userDashboard.Delete("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionWrite), handler.CreateUserDashboard) //todo: diganti jadi delete user by ID
}
