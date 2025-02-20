package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/internal/middleware"
	"go-boilerplate/internal/models"
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

func (h *UserController) Register(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		logger.Error("Failed to parse request body", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	if err := h.UserUC.RegisterUser(user); err != nil {
		logger.Error("Failed to register user", err)
		return utils.SetResponseBadRequest(c, "Failed to register user", err)
	}

	return utils.SetResponseOK(c, "success register user", nil)
}

func (h *UserController) Login(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		logger.Error("Failed to parse login request", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		logger.Error("Failed to generate token", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"token": token})
}

func (h *UserController) Profile(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		logger.Error("Failed to retrieve user ID from context", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User ID not found"})
	}

	return c.JSON(fiber.Map{"message": "Welcome, user", "user_id": userID})
}
