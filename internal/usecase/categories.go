package usecase

import (
	"context"
	"go-boilerplate/internal/dto/request"
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/repo"
	"go-boilerplate/pkg/logger"
	"time"

	"gorm.io/gorm"
)

type CategoryUseCase interface {
    CreateCategory(ctx context.Context, req *request.ReqCategory) error
    GetCategoryByID(ctx context.Context, id int64) (models.Category, error)
    GetListCategory(ctx context.Context) ([]models.Category, error)
    UpdateCategoryByID(ctx context.Context, req *request.ReqCategoryUpdate) (models.Category, error)
    DeleteCategoryByID(ctx context.Context, id int64, updatedAt time.Time) error
}

type categoryUseCase struct {
    db       *gorm.DB
    categoryRepo repo.CategoryRepository
}

func NewCategoryUseCase(categoryRepo repo.CategoryRepository) CategoryUseCase {
    return &categoryUseCase{
        categoryRepo: categoryRepo,
    }
}

func (uc *categoryUseCase) CreateCategory(ctx context.Context, req *request.ReqCategory) error {
    category := models.Category{
        Name: req.Name,
    }

    return processWithTx(ctx,  uc.db, func(ctx context.Context) error {
        err := uc.categoryRepo.Create(ctx, &category)
        if err != nil {
            logger.Error(ctx, "Failed to create category", err)
            return err
        }
        return nil
    })
}

func (uc *categoryUseCase) GetCategoryByID(ctx context.Context, id int64) (models.Category, error) {
    return uc.categoryRepo.GetCategoryByID(ctx, id)
}

func (uc *categoryUseCase) GetListCategory(ctx context.Context) ([]models.Category, error) {
    return uc.categoryRepo.GetListCategory(ctx)
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
            logger.Error(ctx, "Failed to update category", err)
			return err
		}

        return nil
    })

    if err != nil {
        return models.Category{}, err   
    }

    return res, nil
}

func (uc *categoryUseCase) DeleteCategoryByID(ctx context.Context, id int64, updatedAt time.Time) error {
    return processWithTx(ctx,  uc.db, func(ctx context.Context) error {
        err := uc.categoryRepo.DeleteCategoryByID(ctx, id, updatedAt)
        if err != nil {
            logger.Error(ctx, "Failed to delete category", err)
            return err
        }
        return nil
    })
}