package dashboard

import (
	"fmt"
	"go-boilerplate/internal/dto/request"
	"go-boilerplate/internal/usecase"
	"go-boilerplate/internal/utils"
	"go-boilerplate/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type RolePermissionsController struct {
	RolePermissionsUseCase usecase.RolePermissionsUsecase
}

func NewRolePermissionsController(rolePermissionsUC usecase.RolePermissionsUsecase) *RolePermissionsController {
	return &RolePermissionsController{RolePermissionsUseCase: rolePermissionsUC}
}

func (ctrl *RolePermissionsController) CreateRolePermission(c *fiber.Ctx) error {
	var reqRolePermission request.ReqRolePermission
	if err := c.BodyParser(&reqRolePermission); err != nil {
		logger.Error("Failed to parse request body", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	ok, errMsg := utils.ValidateRequest(reqRolePermission, request.ReqRolePermissionErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error("error validate request ", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	if err := ctrl.RolePermissionsUseCase.CreateRolePermission(c.Context(), reqRolePermission); err != nil {
		logger.Error("Failed to create role permission", err)
		return utils.SetResponseBadRequest(c, "Failed to create role permission", err)
	}

	return utils.SetResponseOK(c, "success create role permission", nil)
}

func (ctrl *RolePermissionsController) GetRolePermissionByID(c *fiber.Ctx) error {
	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error("Failed get param id", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	response, err := ctrl.RolePermissionsUseCase.GetRolePermissionByID(c.Context(), id)
	if err != nil {
		logger.Error("Failed get role permission", err)
		return utils.SetResponseBadRequest(c, "Failed get role permission", err)
	}

	return utils.SetResponseOK(c, "success get role permission", response)
}

func (ctrl *RolePermissionsController) GetListRolePermissions(c *fiber.Ctx) error {
	response, err := ctrl.RolePermissionsUseCase.GetListRolePermissions(c.Context())
	if err != nil {
		logger.Error("Failed get list role permissions", err)
		return utils.SetResponseBadRequest(c, "Failed get list role permissions", err)
	}

	return utils.SetResponseOK(c, "success get list role permissions", response)
}

func (ctrl *RolePermissionsController) UpdateRolePermissionByID(c *fiber.Ctx) error {
	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error("Failed get param id", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	reqUpdate := request.ReqRolePermission{}
	if err := c.BodyParser(&reqUpdate); err != nil {
		logger.Error("Failed to parse request body", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	ok, errMsg := utils.ValidateRequest(reqUpdate, request.ReqRolePermissionErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error("error validate request ", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	response, err := ctrl.RolePermissionsUseCase.UpdateRolePermissionByID(c.Context(), id, reqUpdate.UpdatedAt, reqUpdate)
	if err != nil {
		logger.Error("Failed update role permission", err)
		return utils.SetResponseBadRequest(c, "Failed update role permission", err)
	}

	return utils.SetResponseOK(c, "success update role permission", response)
}

func (ctrl *RolePermissionsController) DeleteRolePermissionByID(c *fiber.Ctx) error {
	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error("Failed get param id", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	reqData := request.AbstractRequest{}
	if err := c.BodyParser(&reqData); err != nil {
		logger.Error("Failed to parse request body", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	err = ctrl.RolePermissionsUseCase.DeleteRolePermissionByID(c.Context(), id, reqData.UpdatedAt)
	if err != nil {
		logger.Error("Failed delete role permission", err)
		return utils.SetResponseBadRequest(c, "Failed delete role permission", err)
	}

	return utils.SetResponseOK(c, "success delete role permission", nil)
}
