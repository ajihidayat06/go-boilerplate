package dashboard

import (
	"go-boilerplate/internal/dto/request"
	"go-boilerplate/internal/dto/response"
	"go-boilerplate/internal/middleware"
	"go-boilerplate/internal/usecase"
	"go-boilerplate/internal/utils"
	"go-boilerplate/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthUsecase usecase.AuthUseCase
}

func NewAuthController(
	authUC usecase.AuthUseCase,
) *AuthController {
	return &AuthController{AuthUsecase: authUC}
}

func (ctrl *AuthController) LoginDashboard(c *fiber.Ctx) error {
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

// TODO: implemet logout logic
func (ctrl *AuthController) LogoutDashboard(c *fiber.Ctx) error {
	return utils.SetResponseOK(c, "success logout", nil)
}
