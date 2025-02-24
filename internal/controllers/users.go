package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/internal/usecase"
	"go-boilerplate/internal/utils"
	"go-boilerplate/pkg/logger"
)

type UserController struct {
	UserUC usecase.UserUseCase
}

func NewUserController(userUC usecase.UserUseCase) *UserController {
	return &UserController{UserUC: userUC}
}

func (h *UserController) Profile(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		logger.Error("Failed to retrieve user ID from context", nil)
		return utils.SetResponseInternalServerError(c, "User ID not found", fmt.Errorf("User ID not found"))
	}

	return c.JSON(fiber.Map{"message": "Welcome, user", "user_id": userID})
}
