package repo

import (
	"context"
	"go-boilerplate/internal/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	GetRoleByID(ctx context.Context, id int64) (models.Roles, error)
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

func (r roleRepository) GetRoleByID(ctx context.Context, id int64) (models.Roles, error) {
	var role models.Roles
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&role).Error; err != nil {
		return models.Roles{}, err
	}

	return role, nil
}
