package usecase

import (
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/repo"
)

type UserUseCase interface {
	RegisterUser(user *models.User) error
}

type userUseCase struct {
	UserRepo repo.UserRepository
}

func NewUserUseCase(userRepo repo.UserRepository) UserUseCase {
	return &userUseCase{UserRepo: userRepo}
}

func (u *userUseCase) RegisterUser(user *models.User) error {
	return u.UserRepo.Create(user)
}
