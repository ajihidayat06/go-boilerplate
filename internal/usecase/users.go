package usecase

import (
	"context"
	"go-boilerplate/internal/dto/request"
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/repo"
	"gorm.io/gorm"
)

type UserUseCase interface {
	Register(ctx context.Context, reqUser *request.ReqUser) error
	Login(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error)
	GetUserByID(ctx context.Context, user int64) (models.User, error)
	CreateUserDashboard(ctx context.Context, user *request.ReqUser) error
	GetListUser(ctx context.Context, listStruct *models.GetListStruct) ([]models.User, error)
	UpdateUserByID(ctx context.Context, id int64) (models.User, error)
}

type userUseCase struct {
	db       *gorm.DB
	UserRepo repo.UserRepository
}

func NewUserUseCase(db *gorm.DB, userRepo repo.UserRepository) UserUseCase {
	return &userUseCase{
		db:       db,
		UserRepo: userRepo,
	}
}

func (u *userUseCase) Register(ctx context.Context, reqUser *request.ReqUser) error {
	user := models.User{}
	return u.UserRepo.Create(ctx, &user)
}

func (u userUseCase) Login(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error) {
	// get user by (username or email) and password
	user := models.UserLogin{}
	return user, nil
}

func (u *userUseCase) GetUserByID(ctx context.Context, id int64) (models.User, error) {
	return u.UserRepo.GetUserByID(ctx, id)
}

func (u *userUseCase) CreateUserDashboard(ctx context.Context, reqUser *request.ReqUser) error {
	//Mapping request user ke model user

	user := models.User{}
	return u.UserRepo.Create(ctx, &user)
}

func (u *userUseCase) GetListUser(ctx context.Context, listStruct *models.GetListStruct) ([]models.User, error) {
	return u.UserRepo.GetListUser(ctx, listStruct)
}

func (u *userUseCase) UpdateUserByID(ctx context.Context, id int64) (models.User, error) {
	return u.UserRepo.UpdateUserByID(ctx, id)
}
