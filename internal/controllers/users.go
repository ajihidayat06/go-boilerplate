package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/internal/dto/request"
	"go-boilerplate/internal/dto/response"
	"go-boilerplate/internal/middleware"
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

func (ctrl *UserController) Register(c *fiber.Ctx) error {
	var reqUser request.ReqUser
	if err := c.BodyParser(&reqUser); err != nil {
		logger.Error("Failed to parse request body", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	ok, errMsg := utils.ValidateRequest(reqUser, request.ReqUserErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error("error validate request ", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	if err := ctrl.UserUC.Register(c.Context(), &reqUser); err != nil {
		logger.Error("Failed to register user", err)
		return utils.SetResponseBadRequest(c, "Failed to register user", err)
	}

	return utils.SetResponseOK(c, "success register user", nil)
}

func (ctrl *UserController) Login(c *fiber.Ctx) error {
	var reqLogin request.ReqLogin
	if err := c.BodyParser(&reqLogin); err != nil {
		logger.Error("Failed to parse login request", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	// get user by (email or username) and password
	user, err := ctrl.UserUC.Login(c.Context(), &reqLogin)
	if err != nil {
		logger.Error("login failed", err)
		return utils.SetResponseBadRequest(c, "login failed", err)
	}

	token, err := middleware.GenerateTokenUserDashboard(user)
	if err != nil {
		logger.Error("Failed to generate token", err)
		return utils.SetResponseInternalServerError(c, "Failed to generate token", err)
	}

	return utils.SetResponseOK(c, "succes create token", response.ResAuth{Token: token})
}

func (ctrl *UserController) Logout(c *fiber.Ctx) error {
	// TODO: implement logout
	return utils.SetResponseOK(c, "success logout", nil)
}

func (ctrl *UserController) Profile(c *fiber.Ctx) error {
	_, err := ctrl.UserUC.Profile(c.Context(), request.ReqUser{})
	if err != nil {
		logger.Error("Failed to retrieve user database", err)
		return utils.SetResponseInternalServerError(c, "User ID not found", err)
	}

	return utils.SetResponseOK(c, "success get profile", nil)
}
