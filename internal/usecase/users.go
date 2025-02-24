package usecase

import (
	"context"
	"fmt"
	"go-boilerplate/internal/dto/request"
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/repo"
	"gorm.io/gorm"
)

type UserUseCase interface {
	RegisterUser(ctx context.Context, user *request.ReqUser) error
	Profile(ctx context.Context, user request.ReqUser) (models.User, error)
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

func (u *userUseCase) RegisterUser(ctx context.Context, reqUser *request.ReqUser) error {
	//Mapping request user ke model user
	namaSaya := "aji"
	fmt.Println(namaSaya)
	err := processWithTx(u.db, func(tx *gorm.DB) (err error) {

		namaSaya = "Santoso"
		return
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(namaSaya)

	user := models.User{}
	return u.UserRepo.Create(&user)
}

func (u *userUseCase) Profile(ctx context.Context, user request.ReqUser) (models.User, error) {
	return models.User{}, nil
}
