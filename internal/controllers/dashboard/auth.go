package dashboard

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/internal/dto/request"
	"go-boilerplate/internal/dto/response"
	"go-boilerplate/internal/middleware"
	"go-boilerplate/internal/usecase/dashboard"
	"go-boilerplate/internal/utils"
	"go-boilerplate/pkg/logger"
)

type AuthController struct {
	AuthUsecase dashboard.AuthUseCase
}

func NewAuthController(
	authUC dashboard.AuthUseCase,
) *AuthController {
	return &AuthController{AuthUsecase: authUC}
}

func (ctrl *AuthController) CreateUserDashboard(c *fiber.Ctx) error {
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

	if err := ctrl.AuthUsecase.CreateUserDashboard(c.Context(), &reqUser); err != nil {
		logger.Error("Failed to register user", err)
		return utils.SetResponseBadRequest(c, "Failed to register user", err)
	}

	return utils.SetResponseOK(c, "success register user", nil)
}

func (ctrl *AuthController) Login(c *fiber.Ctx) error {
	var reqLogin request.ReqLogin
	if err := c.BodyParser(&reqLogin); err != nil {
		logger.Error("Failed to parse login request", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	// get user by (email or username) and password
	user, err := ctrl.AuthUsecase.LoginDashboard(c.Context(), &reqLogin)
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
