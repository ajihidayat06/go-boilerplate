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

func (ctrl *UserDahboardController) GetUserByID(c *fiber.Ctx) error {
	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error("Failed get param id", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	response, err := ctrl.UserDashboardUsecase.GetUserByID(c.Context(), id)
	if err != nil {
		logger.Error("Failed get user", err)
		return utils.SetResponseBadRequest(c, "Failed get user", err)
	}

	return utils.SetResponseOK(c, "success get user", response)
}

func (ctrl *UserDahboardController) GetListUser(c *fiber.Ctx) error {

	response, err := ctrl.UserDashboardUsecase.GetListUser(c.Context(), utils.GetFiltersAndPagination(c))
	if err != nil {
		logger.Error("Failed get list user", err)
		return utils.SetResponseBadRequest(c, "Failed get list user", err)
	}

	return utils.SetResponseOK(c, "success get list user", response)
}

func (ctrl *UserDahboardController) UpdateUserByID(c *fiber.Ctx) error {
	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error("Failed get param id", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	reqUpdate := request.ReqUserUpdate{}
	reqUpdate.ID = id
	if err := c.BodyParser(&reqUpdate); err != nil {
		logger.Error("Failed to parse request body", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	ok, errMsg := utils.ValidateRequest(reqUpdate, request.ReqUserUpdateErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error("error validate request ", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	response, err := ctrl.UserDashboardUsecase.UpdateUserByID(c.Context(), &reqUpdate)
	if err != nil {
		logger.Error("Failed update user", err)
		return utils.SetResponseBadRequest(c, "Failed update user", err)
	}

	return utils.SetResponseOK(c, "success update user", response)
}

func (ctrl *UserDahboardController) DeleteUserByID(c *fiber.Ctx) error {
	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error("Failed get param id", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	err = ctrl.UserDashboardUsecase.DeleteUserByID(c.Context(), id)
	if err != nil {
		logger.Error("Failed delete user", err)
		return utils.SetResponseBadRequest(c, "Failed delete user", err)
	}

	return utils.SetResponseOK(c, "success delete user", nil)
}
