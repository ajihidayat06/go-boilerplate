package repo

import (
	"context"
	"go-boilerplate/internal/models"
	"time"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(ctx context.Context, category *models.Category) error
	GetCategoryByID(ctx context.Context, id int64) (models.Category, error)
	GetListCategory(ctx context.Context, listStruct *models.GetListStruct) ([]models.Category, int64, error)
	UpdateCategoryByID(ctx context.Context, id int64, updatedAt time.Time, category models.Category) (models.Category, error)
	DeleteCategoryByID(ctx context.Context, id int64, updatedAt time.Time) error
}

type categoryRepository struct {
	AbstractRepo
}

var (
	FilterCategory = map[string]string{
		"name": "name",
	}
	JoinsCategory = map[string]string{}
)

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		AbstractRepo: AbstractRepo{
			db:          db,
			FilterAlias: FilterCategory,
			Joins:       JoinsCategory,
		},
	}
}

func (r *categoryRepository) Create(ctx context.Context, category *models.Category) error {
	return r.getDB(ctx).WithContext(ctx).Create(category).Error
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int64) (models.Category, error) {
	var category models.Category
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&category).Error
	if err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func (r *categoryRepository) GetListCategory(ctx context.Context, listStruct *models.GetListStruct) ([]models.Category, int64, error) {
	var users []models.Category
	var total int64

	err := r.db.WithContext(ctx).
		Model(&models.Category{}).
		Scopes(r.applyFilters(listStruct.Filters)).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).
		Model(&models.User{}).Preload("Roles").
		Scopes(r.applyFiltersAndPaginationAndOrder(listStruct)).
		Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *categoryRepository) UpdateCategoryByID(ctx context.Context, id int64, updatedAt time.Time, category models.Category) (models.Category, error) {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).
		Model(&category).
		Where("id = ? AND updated_at = ?", id, updatedAt).
		Updates(category).Error
	if err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func (r *categoryRepository) DeleteCategoryByID(ctx context.Context, id int64, updatedAt time.Time) error {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).
		Where("id = ? AND updated_at = ?", id, updatedAt).
		Delete(&models.Category{}).Error
	if err != nil {
		return err
	}
	return nil
}
