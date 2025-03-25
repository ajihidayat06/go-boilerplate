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
    var reqRole request.ReqRoles
    if err := c.BodyParser(&reqRole); err != nil {
        logger.Error("Failed to parse request body", err)
        return utils.SetResponseBadRequest(c, "Invalid request", err)
    }

    ok, errMsg := utils.ValidateRequest(reqRole, request.ReqRolesErrorMessage)
    if !ok {
        err := fmt.Errorf("%s", errMsg)
        logger.Error("error validate request ", err)
        return utils.SetResponseBadRequest(c, "Invalid request", err)
    }

    if err := ctrl.RoleUseCase.CreateRole(c.Context(), &reqRole); err != nil {
        logger.Error("Failed to create role", err)
        return utils.SetResponseBadRequest(c, "Failed to create role", err)
    }

    return utils.SetResponseOK(c, "success create role", nil)
}

func (ctrl *RoleController) GetRoleByID(c *fiber.Ctx) error {
    id, err := utils.ReadRequestParamID(c)
    if err != nil {
        logger.Error("Failed get param id", err)
        return utils.SetResponseBadRequest(c, "Invalid request", err)
    }

    response, err := ctrl.RoleUseCase.GetRoleByID(c.Context(), id)
    if err != nil {
        logger.Error("Failed get role", err)
        return utils.SetResponseBadRequest(c, "Failed get role", err)
    }

    return utils.SetResponseOK(c, "success get role", response)
}

func (ctrl *RoleController) GetListRole(c *fiber.Ctx) error {
    response, err := ctrl.RoleUseCase.GetListRole(c.Context())
    if err != nil {
        logger.Error("Failed get list role", err)
        return utils.SetResponseBadRequest(c, "Failed get list role", err)
    }

    return utils.SetResponseOK(c, "success get list role", response)
}

func (ctrl *RoleController) UpdateRoleByID(c *fiber.Ctx) error {
    id, err := utils.ReadRequestParamID(c)
    if err != nil {
        logger.Error("Failed get param id", err)
        return utils.SetResponseBadRequest(c, "Invalid request", err)
    }

    reqUpdate := request.ReqRoleUpdate{}
    reqUpdate.ID = id
    if err := c.BodyParser(&reqUpdate); err != nil {
        logger.Error("Failed to parse request body", err)
        return utils.SetResponseBadRequest(c, "Invalid request", err)
    }

    ok, errMsg := utils.ValidateRequest(reqUpdate, request.ReqRoleUpdateErrorMessage)
    if !ok {
        err := fmt.Errorf("%s", errMsg)
        logger.Error("error validate request ", err)
        return utils.SetResponseBadRequest(c, "Invalid request", err)
    }

    response, err := ctrl.RoleUseCase.UpdateRoleByID(c.Context(), &reqUpdate)
    if err != nil {
        logger.Error("Failed update role", err)
        return utils.SetResponseBadRequest(c, "Failed update role", err)
    }

    return utils.SetResponseOK(c, "success update role", response)
}

func (ctrl *RoleController) DeleteRoleByID(c *fiber.Ctx) error {
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

    err = ctrl.RoleUseCase.DeleteRoleByID(c.Context(), id, reqData.UpdatedAt)
    if err != nil {
        logger.Error("Failed delete role", err)
        return utils.SetResponseBadRequest(c, "Failed delete role", err)
    }

    return utils.SetResponseOK(c, "success delete role", nil)
}