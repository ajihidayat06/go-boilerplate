package usecase

import (
	"context"
	"go-boilerplate/internal/dto/request"
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/repo"
	"go-boilerplate/pkg/logger"
	"time"

	"gorm.io/gorm"
)

type RoleUseCase interface {
	CreateRole(ctx context.Context, req *request.ReqRoles) error
	GetRoleByID(ctx context.Context, id int64) (models.Roles, error)
	GetListRole(ctx context.Context) ([]models.Roles, error)
	UpdateRoleByID(ctx context.Context, req *request.ReqRoleUpdate) (models.Roles, error)
	DeleteRoleByID(ctx context.Context, id int64, updatedAt time.Time) error
}

type roleUseCase struct {
	db       *gorm.DB
	roleRepo repo.RoleRepository
}

func NewRoleUseCase(
	db *gorm.DB,
	roleRepo repo.RoleRepository,
) RoleUseCase {
	return &roleUseCase{
		db:       db,
		roleRepo: roleRepo,
	}
}

func (uc *roleUseCase) CreateRole(ctx context.Context, req *request.ReqRoles) error {
	role := models.Roles{
		Code: req.Code,
		Name: req.Name,
	}

	return processWithTx(ctx,  uc.db, func(ctx context.Context) error {
        err := uc.roleRepo.Create(ctx, &role)
        if err != nil {
			logger.Error(ctx, "Failed to create role", err)
            return err
        }
        return nil
    })
}

func (uc *roleUseCase) GetRoleByID(ctx context.Context, id int64) (models.Roles, error) {
	return uc.roleRepo.GetRoleByID(ctx, id)
}

func (uc *roleUseCase) GetListRole(ctx context.Context) ([]models.Roles, error) {
	return uc.roleRepo.GetListRole(ctx)
}

func (uc *roleUseCase) UpdateRoleByID(ctx context.Context, req *request.ReqRoleUpdate) (models.Roles, error) {
	err := req.ValidateUpdatedAt()
	if err != nil {
		return models.Roles{}, err
	}

	var updatedRole models.Roles
	err = processWithTx(ctx, uc.db, func(ctx context.Context) error {
		role := models.Roles{
			ID:   req.ID,
			Code: req.Code,
			Name: req.Name,
		}
	
		updatedRole, err = uc.roleRepo.UpdateRoleByID(ctx, req.ID, req.UpdatedAt, role)
		if err != nil {
			logger.Error(ctx, "Failed to update role", err)
			return err
		}
	
		return nil
	})

	return updatedRole, err
}

func (uc *roleUseCase) DeleteRoleByID(ctx context.Context, id int64, updatedAt time.Time) error {
	return processWithTx(ctx,  uc.db, func(ctx context.Context) error {
        err := uc.roleRepo.DeleteRoleByID(ctx, id, updatedAt)
        if err != nil {
			logger.Error(ctx, "Failed to delete role", err)
            return err
        }
        return nil
    })
}
