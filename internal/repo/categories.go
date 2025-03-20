package repo

import (
    "context"
    "go-boilerplate/internal/models"
    "gorm.io/gorm"
    "time"
)

type CategoryRepository interface {
    Create(ctx context.Context, category *models.Category) error
    GetCategoryByID(ctx context.Context, id int64) (models.Category, error)
    GetListCategory(ctx context.Context) ([]models.Category, error)
    UpdateCategoryByID(ctx context.Context, id int64, updatedAt time.Time, category models.Category) (models.Category, error)
    DeleteCategoryByID(ctx context.Context, id int64, updatedAt time.Time) error
}

type categoryRepository struct {
    db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
    return &categoryRepository{
        db: db,
    }
}

func (r *categoryRepository) Create(ctx context.Context, category *models.Category) error {
    return r.db.WithContext(ctx).Create(category).Error
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int64) (models.Category, error) {
    var category models.Category
    err := r.db.WithContext(ctx).Where("id = ?", id).First(&category).Error
    if err != nil {
        return models.Category{}, err
    }
    return category, nil
}

func (r *categoryRepository) GetListCategory(ctx context.Context) ([]models.Category, error) {
    var categories []models.Category
    err := r.db.WithContext(ctx).Find(&categories).Error
    if err != nil {
        return nil, err
    }
    return categories, nil
}

func (r *categoryRepository) UpdateCategoryByID(ctx context.Context, id int64, updatedAt time.Time, category models.Category) (models.Category, error) {
    err := r.db.WithContext(ctx).
        Model(&category).
        Where("id = ? AND updated_at = ?", id, updatedAt).
        Updates(category).Error
    if err != nil {
        return models.Category{}, err
    }
    return category, nil
}

func (r *categoryRepository) DeleteCategoryByID(ctx context.Context, id int64, updatedAt time.Time) error {
    err := r.db.WithContext(ctx).
        Where("id = ? AND updated_at = ?", id, updatedAt).
        Delete(&models.Category{}).Error
    if err != nil {
        return err
    }
    return nil
}