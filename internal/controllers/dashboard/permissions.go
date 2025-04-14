package dashboard

import (
	"fmt"
	"go-boilerplate/internal/dto/request"
	"go-boilerplate/internal/usecase"
	"go-boilerplate/internal/utils"
	"go-boilerplate/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type PermissionController struct {
	PermissionUseCase usecase.PermissionUseCase
}

func NewPermissionController(permissionUC usecase.PermissionUseCase) *PermissionController {
	return &PermissionController{PermissionUseCase: permissionUC}
}

func (ctrl *PermissionController) CreatePermission(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	var reqPermission request.ReqPermission
	if err := c.BodyParser(&reqPermission); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	ok, errMsg := utils.ValidateRequest(reqPermission, request.ReqPermissionErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error(ctx, "error validate request ", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	if err := ctrl.PermissionUseCase.CreatePermission(ctx, &reqPermission); err != nil {
		logger.Error(ctx, "Failed to create permission", err)
		return utils.SetResponseBadRequest(c, "Failed to create permission", err)
	}

	return utils.SetResponseOK(c, "success create permission", nil)
}

func (ctrl *PermissionController) GetPermissionByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	response, err := ctrl.PermissionUseCase.GetPermissionByID(ctx, id)
	if err != nil {
		logger.Error(ctx, "Failed get permission", err)
		return utils.SetResponseBadRequest(c, "Failed get permission", err)
	}

	return utils.SetResponseOK(c, "success get permission", response)
}

func (ctrl *PermissionController) GetListPermission(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	response, err := ctrl.PermissionUseCase.GetListPermission(ctx)
	if err != nil {
		logger.Error(ctx, "Failed get list permission", err)
		return utils.SetResponseBadRequest(c, "Failed get list permission", err)
	}

	return utils.SetResponseOK(c, "success get list permission", response)
}

func (ctrl *PermissionController) UpdatePermissionByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	reqUpdate := request.ReqPermissionUpdate{}
	reqUpdate.ID = id
	if err := c.BodyParser(&reqUpdate); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	ok, errMsg := utils.ValidateRequest(reqUpdate, request.ReqPermissionUpdateErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error(ctx, "error validate request ", err)
		return utils.SetResponseBadRequest(c, "Invalid request", err)
	}

	response, err := ctrl.PermissionUseCase.UpdatePermissionByID(ctx, &reqUpdate)
	if err != nil {
		logger.Error(ctx, "Failed update permission", err)
		return utils.SetResponseBadRequest(c, "Failed update permission", err)
	}

	return utils.SetResponseOK(c, "success update permission", response)
}

func (ctrl *PermissionController) DeletePermissionByID(c *fiber.Ctx) error {
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

	err = ctrl.PermissionUseCase.DeletePermissionByID(ctx, id, reqData.UpdatedAt)
	if err != nil {
		logger.Error(ctx, "Failed delete permission", err)
		return utils.SetResponseBadRequest(c, "Failed delete permission", err)
	}

	return utils.SetResponseOK(c, "success delete permission", nil)
}
