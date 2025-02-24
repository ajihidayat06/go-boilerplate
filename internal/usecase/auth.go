package usecase

import (
	"go-boilerplate/internal/constanta"
	"go-boilerplate/internal/dto/request"
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/repo"
)

type AuthUseCase interface {
	Login(req *request.ReqLogin) (models.User, error)
	RegisterUser(user *request.ReqUser) error
}

type authUseCase struct {
	UserRepo repo.UserRepository
}

func NewAuthUseCase(userRepo repo.UserRepository) AuthUseCase {
	return &authUseCase{UserRepo: userRepo}
}

func (u *authUseCase) RegisterUser(reqUser *request.ReqUser) error {
	//Mapping request user ke model user

	user := models.User{}
	return u.UserRepo.Create(&user)
}

func (a authUseCase) Login(req *request.ReqLogin) (models.User, error) {
	// get user by (username or email) and password
	var permissions []models.Permissions
	permissions = append(permissions, models.Permissions{
		ID:        1,
		Code:      "user_read",
		Name:      "user read",
		Action:    constanta.ActionRead,
		GroupMenu: constanta.MenuGroupUser,
	})

	user := models.User{
		ID:          1,
		Username:    "ajihidayat",
		Password:    "ajihdiayat6",
		Email:       "ajihidayat@gmail.com",
		RoleID:      1,
		Permissions: permissions,
	}

	return user, nil
}
