package dashboard

import (
	"fmt"
	"go-boilerplate/internal/dto/request"
	"go-boilerplate/internal/usecase"
	"go-boilerplate/internal/utils"
	"go-boilerplate/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type RoleController struct {
	RoleUseCase usecase.RoleUseCase
}

func NewRoleController(roleUC usecase.RoleUseCase) *RoleController {
	return &RoleController{RoleUseCase: roleUC}
}

func (ctrl *RoleController) CreateRole(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	var reqRole request.ReqRoles
	if err := c.BodyParser(&reqRole); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	ok, errMsg := utils.ValidateRequest(reqRole, request.ReqRolesErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error(ctx, "error validate request ", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	if err := ctrl.RoleUseCase.CreateRole(ctx, &reqRole); err != nil {
		logger.Error(ctx, "Failed to create role", err)
		return utils.SetResponseBadRequest(c, "Failed to create role", err)
	}

	return utils.SetResponseOK(c, "success create role", nil)
}

func (ctrl *RoleController) GetRoleByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	response, err := ctrl.RoleUseCase.GetRoleByID(ctx, id)
	if err != nil {
		logger.Error(ctx, "Failed get role", err)
		return utils.SetResponseBadRequest(c, "Failed get role", err)
	}

	return utils.SetResponseOK(c, "success get role", response)
}

func (ctrl *RoleController) GetListRole(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	response, err := ctrl.RoleUseCase.GetListRole(ctx)
	if err != nil {
		logger.Error(ctx, "Failed get list role", err)
		return utils.SetResponseBadRequest(c, "Failed get list role", err)
	}

	return utils.SetResponseOK(c, "success get list role", response)
}

func (ctrl *RoleController) UpdateRoleByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	reqUpdate := request.ReqRoleUpdate{}
	reqUpdate.ID = id
	if err := c.BodyParser(&reqUpdate); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	ok, errMsg := utils.ValidateRequest(reqUpdate, request.ReqRoleUpdateErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error(ctx, "error validate request ", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	response, err := ctrl.RoleUseCase.UpdateRoleByID(ctx, &reqUpdate)
	if err != nil {
		logger.Error(ctx, "Failed update role", err)
		return utils.SetResponseBadRequest(c, "Failed update role", err)
	}

	return utils.SetResponseOK(c, "success update role", response)
}

func (ctrl *RoleController) DeleteRoleByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	reqData := request.AbstractRequest{}
	if err := c.BodyParser(&reqData); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	err = ctrl.RoleUseCase.DeleteRoleByID(ctx, id, reqData.UpdatedAt)
	if err != nil {
		logger.Error(ctx, "Failed delete role", err)
		return utils.SetResponseBadRequest(c, "Failed delete role", err)
	}

	return utils.SetResponseOK(c, "success delete role", nil)
}
