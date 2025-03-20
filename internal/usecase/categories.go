package usecase

import (
    "context"
    "go-boilerplate/internal/dto/request"
    "go-boilerplate/internal/models"
    "go-boilerplate/internal/repo"
    "time"
)

type CategoryUseCase interface {
    CreateCategory(ctx context.Context, req *request.ReqCategory) error
    GetCategoryByID(ctx context.Context, id int64) (models.Category, error)
    GetListCategory(ctx context.Context) ([]models.Category, error)
    UpdateCategoryByID(ctx context.Context, req *request.ReqCategoryUpdate) (models.Category, error)
    DeleteCategoryByID(ctx context.Context, id int64, updatedAt time.Time) error
}

type categoryUseCase struct {
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
    return uc.categoryRepo.Create(ctx, &category)
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
    return uc.categoryRepo.UpdateCategoryByID(ctx, req.ID, req.UpdatedAt, category)
}

func (uc *categoryUseCase) DeleteCategoryByID(ctx context.Context, id int64, updatedAt time.Time) error {
    return uc.categoryRepo.DeleteCategoryByID(ctx, id, updatedAt)
}