package repo

import (
	"context"
	"go-boilerplate/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	Login(ctx context.Context, emailOrUsername, password string) (*models.User, error)
	GetUserByID(ctx context.Context, id int64) (models.User, error)
	GetListUser(ctx context.Context, listStruct *models.GetListStruct) ([]models.User, error)
	UpdateUserByID(ctx context.Context, id int64) (models.User, error)
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

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Login(ctx context.Context, emailOrUsername, password string) (*models.User, error) {
	var user models.User
	err := r.db.
		Preload("Roles").
		Preload("Roles.RolePermissions").
		Preload("Roles.RolePermissions.Permission").
		Where("(email = ? OR username = ?)", emailOrUsername, emailOrUsername).
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id int64) (models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).
		Scopes(r.withCheckScope(ctx)).
		Where(" id = ? ", id).
		First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetListUser(ctx context.Context, listStruct *models.GetListStruct) ([]models.User, error) {
	var users []models.User

	err := r.db.WithContext(ctx).
		Scopes(r.withCheckScope(ctx), r.applyFiltersAndPaginationAndOrder(listStruct)).
		Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) UpdateUserByID(ctx context.Context, id int64) (models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).
		Scopes(r.withCheckScope(ctx)).
		Model(user).
		Where("id = ? AND updated_at = ?", user.ID, user.UpdatedAt).
		Updates(user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
