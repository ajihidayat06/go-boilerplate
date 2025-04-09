package usecase

import (
	"context"
	"go-boilerplate/internal/dto/request"
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/repo"
	"go-boilerplate/internal/utils"
	"go-boilerplate/internal/utils/errors"

	"gorm.io/gorm"
)

type AuthUseCase interface {
	LoginDashboard(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error)
	Login(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error)
	LogoutDashboard(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error)
	Logout(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error)
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

func (u *authUseCase) LoginDashboard(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error) {
	// Ambil user dari repository
	user, err := u.UserRepo.Login(ctx, req.UsernameOrEmail, req.Password)
	if err != nil {
		return models.UserLogin{}, err
	}

	// Validasi password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return models.UserLogin{}, errors.ErrInvalidCredentials // Ensure the error is defined in the errors package
	}

	// Mapping RolePermissions ke UserLogin
	var rolePermissions []models.RolePermissions
	for _, rp := range *user.Roles.RolePermissions {
		rolePermissions = append(rolePermissions, models.RolePermissions{
			ID:            rp.ID,
			RoleID:        rp.RoleID,
			PermissionsID: rp.PermissionsID,
			AccessScope:   rp.AccessScope,
			Permissions: &models.Permissions{
				ID:        rp.Permissions.ID,
				Code:      rp.Permissions.Code,
				Name:      rp.Permissions.Name,
				Action:    rp.Permissions.Action,
				GroupMenu: rp.Permissions.GroupMenu,
			},
		})
	}

	// Buat UserLogin
	userLogin := models.UserLogin{
		ID:              user.ID,
		RoleID:          user.RoleID,
		RoleName:        user.Roles.Name,
		RolePermissions: rolePermissions,
	}

	return userLogin, nil
}

// Login implements AuthUseCase.
func (u *authUseCase) Login(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error) {
	panic("unimplemented")
}

// Logout implements AuthUseCase.
func (u *authUseCase) Logout(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error) {
	panic("unimplemented")
}

// LogoutDashboard implements AuthUseCase.
func (u *authUseCase) LogoutDashboard(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error) {
	panic("unimplemented")
}
