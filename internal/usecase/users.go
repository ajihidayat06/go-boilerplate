package usecase

import (
	"go-boilerplate/internal/dto/request"
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/repo"
)

type UserUseCase interface {
	RegisterUser(user *request.ReqUser) error
}

type userUseCase struct {
	UserRepo repo.UserRepository
}

func NewUserUseCase(userRepo repo.UserRepository) UserUseCase {
	return &userUseCase{UserRepo: userRepo}
}

func (u *userUseCase) RegisterUser(reqUser *request.ReqUser) error {
	//Mapping request user ke model user

	user := models.User{}
	return u.UserRepo.Create(&user)
}
