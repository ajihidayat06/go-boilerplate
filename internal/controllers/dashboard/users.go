package dashboard

import (
	"fmt"
	"go-boilerplate/internal/constanta"
	"go-boilerplate/internal/dto/request"
	"go-boilerplate/internal/usecase"
	"go-boilerplate/internal/utils"
	"go-boilerplate/internal/utils/errorutils"
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
	ctx := utils.GetContext(c)

	fmt.Println("DEBUG user_id:", ctx.Value(constanta.AuthUserID)) // dari context
	fmt.Println("DEBUG user_id locals:", c.Locals(constanta.AuthUserID))

	var reqUser request.ReqUser
	if err := c.BodyParser(&reqUser); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return utils.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	ok, errMsg := utils.ValidateRequest(reqUser, request.ReqUserErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error(ctx, "error validate request ", err)
		return utils.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	if err := ctrl.UserDashboardUsecase.CreateUserDashboard(ctx, &reqUser); err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed create user")
	}

	return utils.SetResponseOK(c, "success register user", nil)
}

func (ctrl *UserDahboardController) GetUserByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return utils.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	response, err := ctrl.UserDashboardUsecase.GetUserByID(ctx, id)
	if err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed get user")
	}

	return utils.SetResponseOK(c, "success get user", response)
}

func (ctrl *UserDahboardController) GetListUser(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	response, err := ctrl.UserDashboardUsecase.GetListUser(ctx, utils.GetFiltersAndPagination(c))
	if err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed get list user")
	}

	return utils.SetResponseOK(c, "success get list user", response)
}

func (ctrl *UserDahboardController) UpdateUserByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return utils.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	reqUpdate := request.ReqUserUpdate{}
	reqUpdate.ID = id
	if err := c.BodyParser(&reqUpdate); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return utils.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	ok, errMsg := utils.ValidateRequest(reqUpdate, request.ReqUserUpdateErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error(ctx, "error validate request ", err)
		return utils.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	response, err := ctrl.UserDashboardUsecase.UpdateUserByID(ctx, &reqUpdate)
	if err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed update user")
	}

	return utils.SetResponseOK(c, "success update user", response)
}

func (ctrl *UserDahboardController) DeleteUserByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return utils.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	reqData := request.AbstractRequest{}
	if err := c.BodyParser(&reqData); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return utils.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	err = ctrl.UserDashboardUsecase.DeleteUserByID(ctx, id, reqData)
	if err != nil {
		logger.Error(ctx, "Failed delete user", err)
		return utils.SetResponseBadRequest(c, "Failed delete user", err)
	}

	return utils.SetResponseOK(c, "success delete user", nil)
}
