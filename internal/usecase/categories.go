package usecase

import (
	"context"
	"go-boilerplate/internal/dto/request"
	"go-boilerplate/internal/dto/response"
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/repo"
	"go-boilerplate/internal/utils"
	"go-boilerplate/internal/utils/errorutils"
	"go-boilerplate/pkg/logger"

	"gorm.io/gorm"
)

type CategoryUseCase interface {
	CreateCategory(ctx context.Context, req *request.ReqCategory) error
	GetCategoryByID(ctx context.Context, id int64) (response.CategoryResponse, error)
	GetListCategory(ctx context.Context, listStruct *models.GetListStruct) (response.ListResponse[response.CategoryResponse], error)
	UpdateCategoryByID(ctx context.Context, req *request.ReqCategoryUpdate) (models.Category, error)
	DeleteCategoryByID(ctx context.Context, id int64, ureqData request.AbstractRequest) error
}

type categoryUseCase struct {
	db           *gorm.DB
	categoryRepo repo.CategoryRepository
}

func NewCategoryUseCase(db *gorm.DB, categoryRepo repo.CategoryRepository) CategoryUseCase {
	return &categoryUseCase{
		db:           db,
		categoryRepo: categoryRepo,
	}
}

func (uc *categoryUseCase) CreateCategory(ctx context.Context, req *request.ReqCategory) error {
	err := req.ValidateRequestCreate()
	if err != nil {
		logger.Error(ctx, "Failed to validate request", err)
		return err
	}
	userLogin, err := utils.GetUserIDFromCtx(ctx)
	if err != nil {
		logger.Error(ctx, "Failed to get user id from context", err)
		return errorutils.ErrDataNotFound
	}

	// TODO: get data category by code or name

	category := models.Category{
		Name:      req.Name,
		Code:      req.Code,
		Slug:      req.Slug,
		CreatedBy: userLogin,
		UpdatedBy: userLogin,
	}

	return processWithTx(ctx, uc.db, func(ctx context.Context) error {
		err := uc.categoryRepo.Create(ctx, &category)
		if err != nil {
			return errorutils.HandleRepoErrorWrite(ctx, err, repo.GetContraintErrMessage(uc.categoryRepo))
		}
		return nil
	})
}

func (uc *categoryUseCase) GetCategoryByID(ctx context.Context, id int64) (response.CategoryResponse, error) {
	categoryDb, err := uc.categoryRepo.GetCategoryByID(ctx, id)
	if err != nil {
		return response.CategoryResponse{}, errorutils.HandleRepoError(ctx, err)
	}

	return response.SetCategoryResponse(categoryDb), nil
}

func (uc *categoryUseCase) GetListCategory(ctx context.Context, listStruct *models.GetListStruct) (response.ListResponse[response.CategoryResponse], error) {
	categoryDb, count, err := uc.categoryRepo.GetListCategory(ctx, listStruct)
	if err != nil {
		return response.ListResponse[response.CategoryResponse]{}, errorutils.HandleRepoError(ctx, err)
	}

	listResponse := utils.MapToListResponse(response.SetResponseListCategory(categoryDb), count, listStruct, repo.GetFilterAvailableFromRepo(uc.categoryRepo))
	return listResponse, nil
}

func (uc *categoryUseCase) UpdateCategoryByID(ctx context.Context, req *request.ReqCategoryUpdate) (models.Category, error) {
	category := models.Category{
		ID:   req.ID,
		Name: req.Name,
	}

	var (
		res models.Category
		err error
	)
	err = processWithTx(ctx, uc.db, func(ctx context.Context) error {
		res, err = uc.categoryRepo.UpdateCategoryByID(ctx, req.ID, req.UpdatedAt, category)
		if err != nil {
			return errorutils.HandleRepoErrorWrite(ctx, err, repo.GetContraintErrMessage(uc.categoryRepo))
		}

		return nil
	})

	if err != nil {
		return models.Category{}, err
	}

	return res, nil
}

func (uc *categoryUseCase) DeleteCategoryByID(ctx context.Context, id int64, reqData request.AbstractRequest) error {
	err := reqData.ValidateUpdatedAt()
	if err != nil {
		return err
	}

	category, err := uc.categoryRepo.GetCategoryByID(ctx, id)
	if err != nil {
		return errorutils.HandleRepoError(ctx, err)
	}

	if !utils.ValidateUpdatedAtRequest(reqData.UpdatedAt, category.UpdatedAt) {
		return errorutils.ErrDataDataUpdated
	}

	return processWithTx(ctx, uc.db, func(ctx context.Context) error {
		err := uc.categoryRepo.DeleteCategoryByID(ctx, id, reqData.UpdatedAt)
		if err != nil {
			return errorutils.HandleRepoError(ctx, err)
		}
		return nil
	})
}
