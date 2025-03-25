package repo

import (
	"context"
	"go-boilerplate/internal/models"
	"time"

	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(ctx context.Context, role *models.Roles) error
	GetRoleByID(ctx context.Context, id int64) (models.Roles, error)
	GetListRole(ctx context.Context) ([]models.Roles, error)
	UpdateRoleByID(ctx context.Context, id int64, updatedAt time.Time, role models.Roles) (models.Roles, error)
	DeleteRoleByID(ctx context.Context, id int64, updatedAt time.Time) error
}

type roleRepository struct {
	AbstractRepo
}


func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		AbstractRepo: AbstractRepo{
			db: db,
		},
	}
}

func (r *roleRepository) Create(ctx context.Context, role *models.Roles) error {
	return r.getDB(ctx).WithContext(ctx).Create(role).Error
}

func (r *roleRepository) GetRoleByID(ctx context.Context, id int64) (models.Roles, error) {
	var role models.Roles
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&role).Error
	if err != nil {
		return models.Roles{}, err
	}
	return role, nil
}

func (r *roleRepository) GetListRole(ctx context.Context) ([]models.Roles, error) {
	var roles []models.Roles
	err := r.db.WithContext(ctx).Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *roleRepository) UpdateRoleByID(ctx context.Context, id int64, updatedAt time.Time, role models.Roles) (models.Roles, error) {
    db := r.getDB(ctx) // Gunakan DB dari context jika ada transaksi

    err := db.WithContext(ctx).
        Model(&role).
        Where("id = ? AND updated_at = ?", id, updatedAt).
        Updates(role).Error
    if err != nil {
        return models.Roles{}, err
    }
    return role, nil
}

func (r *roleRepository) DeleteRoleByID(ctx context.Context, id int64, updatedAt time.Time) error {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).
		Where("id = ? AND updated_at = ?", id, updatedAt).
		Delete(&models.Roles{}).Error
	if err != nil {
		return err
	}
	return nil
}
