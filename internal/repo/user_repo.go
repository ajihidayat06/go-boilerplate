package repo

import (
	"go-boilerplate/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	Login(emailOrUsername, password string) (*models.User, error)
}

type userRepository struct {
	AbstractRepo
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		AbstractRepo: AbstractRepo{
			db: db,
		},
	}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Login(eemailOrUsername, password string) (*models.User, error) {
	var user models.User
	err := r.db.
		Preload("Roles").
		Preload("Roles.RolePermissions").
		Preload("Roles.RolePermissions.Permission").
		Where("(email = ? OR username = ?)", eemailOrUsername, eemailOrUsername).
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
