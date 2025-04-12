package usecase

import (
	"context"
	"go-boilerplate/internal/dto/request"
	"go-boilerplate/internal/dto/response"
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/repo"
	"go-boilerplate/internal/utils"
	"go-boilerplate/internal/utils/errorutils"
	"go-boilerplate/pkg/logger"

	"gorm.io/gorm"
)

type UserUseCase interface {
	Register(ctx context.Context, reqUser *request.ReqUser) error
	Login(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error)
	GetUserByID(ctx context.Context, user int64) (models.User, error)
	CreateUserDashboard(ctx context.Context, user *request.ReqUser) error
	GetListUser(ctx context.Context, listStruct *models.GetListStruct) (response.ListResponse[response.UserResponse], error)
	UpdateUserByID(ctx context.Context, user *request.ReqUserUpdate) (models.User, error)
	DeleteUserByID(ctx context.Context, id int64, reqData request.AbstractRequest) error
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
	return processWithTx(ctx, u.db, func(ctx context.Context) error {
		err := u.UserRepo.Create(ctx, &user)
		if err != nil {
			logger.Error("Failed to create user", err)
			return err
		}
		return nil
	})
}

func (u userUseCase) Login(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error) {
	// TODO: get user by (username or email) and password
	user := models.UserLogin{}
	return user, nil
}

func (u *userUseCase) GetUserByID(ctx context.Context, id int64) (models.User, error) {
	return u.UserRepo.GetUserByID(ctx, id)
}

func (u *userUseCase) CreateUserDashboard(ctx context.Context, reqUser *request.ReqUser) error {
	// Map request user to model user
	err := reqUser.ValidateRequestCreate()
	if err != nil {
		return err
	}

	user := models.User{
		Name:     reqUser.Name,
		Email:    reqUser.Email,
		Username: reqUser.Username,
		Password: reqUser.Password,
	}

	return processWithTx(ctx, u.db, func(ctx context.Context) error {
		err := u.UserRepo.Create(ctx, &user)
		if err != nil {
			logger.Error("Failed to create user dashboard", err)
			return errorutils.HandleRepoError(err)
		}
		return nil
	})
}

func (u *userUseCase) GetListUser(ctx context.Context, listStruct *models.GetListStruct) (response.ListResponse[response.UserResponse], error) {
	userDb, count, err := u.UserRepo.GetListUser(ctx, listStruct)
	if err != nil {
		logger.Error("Failed to get list user", err)
		return response.ListResponse[response.UserResponse]{}, err
	}

	listResponse := utils.MapToListResponse(response.SetResponseListUser(userDb), count, listStruct.Page, listStruct.Limit)
	return listResponse, nil
}

func (u *userUseCase) UpdateUserByID(ctx context.Context, reqData *request.ReqUserUpdate) (models.User, error) {
	err := reqData.ValidateRequestUpdate()
	if err != nil {
		return models.User{}, err
	}

	// TODO: Mapping request user ke model user
	var (
		res models.User
	)
	err = processWithTx(ctx, u.db, func(ctx context.Context) error {
		res, err = u.UserRepo.UpdateUserByID(ctx, *reqData, models.User{})
		if err != nil {
			logger.Error("Failed to update user", err)
			return err
		}

		return nil
	})
	if err != nil {
		return models.User{}, err
	}

	return res, nil
}

func (u *userUseCase) DeleteUserByID(ctx context.Context, id int64, reqData request.AbstractRequest) error {
	err := reqData.ValidateUpdatedAt()
	if err != nil {
		return err
	}

	return processWithTx(ctx, u.db, func(ctx context.Context) error {
		err := u.UserRepo.DeleteUserByID(ctx, id, reqData.UpdatedAt)
		if err != nil {
			logger.Error("Failed to delete user", err)
			return err
		}
		return nil
	})
}
