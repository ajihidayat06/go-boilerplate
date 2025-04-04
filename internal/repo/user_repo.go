package repo

import (
	"context"
	"go-boilerplate/internal/dto/request"
	"go-boilerplate/internal/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	Login(ctx context.Context, emailOrUsername, password string) (*models.User, error)
	GetUserByID(ctx context.Context, id int64) (models.User, error)
	GetListUser(ctx context.Context, listStruct *models.GetListStruct) ([]models.User, error)
	UpdateUserByID(ctx context.Context, reqData request.ReqUserUpdate, user models.User) (models.User, error)
	DeleteUserByID(ctx context.Context, id int64, updatedAt time.Time) error
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
	return r.getDB(ctx).Create(user).Error
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

func (r *userRepository) UpdateUserByID(ctx context.Context, reqData request.ReqUserUpdate, user models.User) (models.User, error) {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).
		Scopes(r.withCheckScope(ctx)).
		Model(user).
		Where("id = ? AND updated_at = ?", reqData.ID, reqData.UpdatedAt).
		Updates(user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// DeleteUserByID implements UserRepository.
func (r *userRepository) DeleteUserByID(ctx context.Context, id int64, updatedAt time.Time) error {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).
	Scopes(r.withCheckScope(ctx)).
	Where("id = ? AND updated_at = ?", id, updatedAt).
	Delete(&models.User{}).Error
	if err != nil {
		return err
	}

	return nil
}