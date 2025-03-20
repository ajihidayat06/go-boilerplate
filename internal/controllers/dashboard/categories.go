package dashboard

import (
	"fmt"
	"go-boilerplate/internal/dto/request"
	"go-boilerplate/internal/usecase"
	"go-boilerplate/internal/utils"
	"go-boilerplate/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type CategoryDashboardController struct {
    CategoryUseCase usecase.CategoryUseCase
}

func NewCategoryController(categoryUC usecase.CategoryUseCase) *CategoryDashboardController {
    return &CategoryDashboardController{CategoryUseCase: categoryUC}
}

func (ctrl *CategoryDashboardController) CreateCategory(c *fiber.Ctx) error {
    var reqCategory request.ReqCategory
    if err := c.BodyParser(&reqCategory); err != nil {
        logger.Error("Failed to parse request body", err)
        return utils.SetResponseBadRequest(c, "Invalid request", err)
    }

    ok, errMsg := utils.ValidateRequest(reqCategory, request.ReqCategoryErrorMessage)
    if !ok {
        err := fmt.Errorf("%s", errMsg)
        logger.Error("error validate request ", err)
        return utils.SetResponseBadRequest(c, "Invalid request", err)
    }

    if err := ctrl.CategoryUseCase.CreateCategory(c.Context(), &reqCategory); err != nil {
        logger.Error("Failed to create category", err)
        return utils.SetResponseBadRequest(c, "Failed to create category", err)
    }

    return utils.SetResponseOK(c, "success create category", nil)
}

func (ctrl *CategoryDashboardController) GetCategoryByID(c *fiber.Ctx) error {
    id, err := utils.ReadRequestParamID(c)
    if err != nil {
        logger.Error("Failed get param id", err)
        return utils.SetResponseBadRequest(c, "Invalid request", err)
    }

    response, err := ctrl.CategoryUseCase.GetCategoryByID(c.Context(), id)
    if err != nil {
        logger.Error("Failed get category", err)
        return utils.SetResponseBadRequest(c, "Failed get category", err)
    }

    return utils.SetResponseOK(c, "success get category", response)
}

func (ctrl *CategoryDashboardController) GetListCategory(c *fiber.Ctx) error {
    response, err := ctrl.CategoryUseCase.GetListCategory(c.Context())
    if err != nil {
        logger.Error("Failed get list category", err)
        return utils.SetResponseBadRequest(c, "Failed get list category", err)
    }

    return utils.SetResponseOK(c, "success get list category", response)
}

func (ctrl *CategoryDashboardController) UpdateCategoryByID(c *fiber.Ctx) error {
    id, err := utils.ReadRequestParamID(c)
    if err != nil {
        logger.Error("Failed get param id", err)
        return utils.SetResponseBadRequest(c, "Invalid request", err)
    }

    reqUpdate := request.ReqCategoryUpdate{}
    reqUpdate.ID = id
    if err := c.BodyParser(&reqUpdate); err != nil {
        logger.Error("Failed to parse request body", err)
        return utils.SetResponseBadRequest(c, "Invalid request", err)
    }

    ok, errMsg := utils.ValidateRequest(reqUpdate, request.ReqCategoryUpdateErrorMessage)
    if !ok {
        err := fmt.Errorf("%s", errMsg)
        logger.Error("error validate request ", err)
        return utils.SetResponseBadRequest(c, "Invalid request", err)
    }

    response, err := ctrl.CategoryUseCase.UpdateCategoryByID(c.Context(), &reqUpdate)
    if err != nil {
        logger.Error("Failed update category", err)
        return utils.SetResponseBadRequest(c, "Failed update category", err)
    }

    return utils.SetResponseOK(c, "success update category", response)
}

func (ctrl *CategoryDashboardController) DeleteCategoryByID(c *fiber.Ctx) error {
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

    err = ctrl.CategoryUseCase.DeleteCategoryByID(c.Context(), id, reqData.UpdatedAt)
    if err != nil {
        logger.Error("Failed delete category", err)
        return utils.SetResponseBadRequest(c, "Failed delete category", err)
    }

    return utils.SetResponseOK(c, "success delete category", nil)
}