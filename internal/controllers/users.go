package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/internal/dto/request"
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

func (ctrl *UserController) Profile(c *fiber.Ctx) error {
	_, err := ctrl.UserUC.Profile(c.Context(), request.ReqUser{})
	if err != nil {
		logger.Error("Failed to retrieve user database", err)
		return utils.SetResponseInternalServerError(c, "User ID not found", err)
	}

	return utils.SetResponseOK(c, "success get profile", nil)
}
