package router

import (
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/internal/controllers"
)

func UserWebRoutes(api fiber.Router, user *controllers.UserController) {
	userRoute := api.Group("/user")
	userRoute.Get("/profile", user.Profile)
}
