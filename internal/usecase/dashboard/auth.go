package dashboard

import (
	"context"
	"go-boilerplate/internal/dto/request"
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/repo"
	"gorm.io/gorm"
)

type AuthUseCase interface {
	LoginDashboard(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error)
	CreateUserDashboard(ctx context.Context, user *request.ReqUser) error
}

type authUseCase struct {
	db       *gorm.DB
	UserRepo repo.UserRepository
}

func NewAuthUseCase(db *gorm.DB, userRepo repo.UserRepository) AuthUseCase {
	return &authUseCase{
		db:       db,
		UserRepo: userRepo,
	}
}

func (u *authUseCase) CreateUserDashboard(ctx context.Context, reqUser *request.ReqUser) error {
	//Mapping request user ke model user

	user := models.User{}
	return u.UserRepo.Create(ctx, &user)
}

func (u authUseCase) LoginDashboard(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error) {
	// get user by (username or email) and password
	var (
		rolePermissions []models.RolePermissions
	)

	rolePermissions = append(rolePermissions, models.RolePermissions{
		ID:            1,
		RoleID:        1,
		PermissionsID: 1,
		Permissions: models.Permissions{
			ID:        1,
			Code:      "user_read",
			Name:      "user read",
			Action:    "read",
			GroupMenu: "user",
		},
	})

	roles := models.Roles{
		ID:              1,
		Code:            "admin",
		Name:            "admin",
		RolePermissions: &rolePermissions,
	}

	user := models.User{
		ID:       1,
		Username: "ajihidayat",
		Password: "ajihdiayat6",
		Email:    "ajihidayat@gmail.com",
		RoleID:   1,
		Roles:    &roles,
	}

	//userData, err := u.UserRepo.LoginDashboard(ctx, req.UsenameOrEmail, req.Password)
	//if err != nil {
	//	return models.User{}, err
	//}

	//mapping response ke model user login
	var resPesmissions []models.Permissions
	for _, rp := range *user.Roles.RolePermissions {
		resPesmissions = append(resPesmissions, models.Permissions{
			ID:        rp.Permissions.ID,
			Code:      rp.Permissions.Code,
			Name:      rp.Permissions.Name,
			Action:    rp.Permissions.Action,
			GroupMenu: rp.Permissions.GroupMenu,
		})
	}

	userLogin := models.UserLogin{
		ID:          user.ID,
		RoleID:      user.RoleID,
		RoleName:    user.Roles.Name,
		Permissions: resPesmissions,
	}

	return userLogin, nil
}
