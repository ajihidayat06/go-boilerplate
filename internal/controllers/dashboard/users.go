package dashboard

import (
	"fmt"
	"go-boilerplate/internal/dto/request"
	"go-boilerplate/internal/usecase"
	"go-boilerplate/internal/utils"
	"go-boilerplate/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type UserDahboardController struct {
	UserDashboardUsecase usecase.UserUseCase
}

func NewUserDashboardController(
	userUC usecase.UserUseCase,
) *UserDahboardController {
	return &UserDahboardController{UserDashboardUsecase: userUC}
}

func (ctrl *UserDahboardController) CreateUserDashboard(c *fiber.Ctx) error {
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

	if err := ctrl.UserDashboardUsecase.CreateUserDashboard(c.Context(), &reqUser); err != nil {
		logger.Error("Failed to register user", err)
		return utils.SetResponseBadRequest(c, "Failed to register user", err)
	}

	return utils.SetResponseOK(c, "success register user", nil)
}

func (ctrl *UserDahboardController) GetUserDashboard(c *fiber.Ctx) error {

	return utils.SetResponseOK(c, "success get user", nil)
}
