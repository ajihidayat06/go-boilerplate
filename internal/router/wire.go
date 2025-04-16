package router

import (
	"go-boilerplate/internal/controllers"
	"go-boilerplate/internal/controllers/dashboard"
	"go-boilerplate/internal/repo"
	"go-boilerplate/internal/usecase"

	"gorm.io/gorm"
)

func InitUser(db *gorm.DB) *controllers.UserController {
	userRepo := repo.NewUserRepository(db)
	roleRepo := repo.NewRoleRepository(db)
	userUC := usecase.NewUserUseCase(db, userRepo, roleRepo)
	userController := controllers.NewUserController(userUC)

	return userController
}

func InitAuth(db *gorm.DB) *dashboard.AuthController {
	userRepo := repo.NewUserRepository(db)
	authUC := usecase.NewAuthUseCase(db, userRepo)
	authController := dashboard.NewAuthController(authUC)

	return authController
}

func InitUserDahboard(db *gorm.DB) *dashboard.UserDahboardController {
	userRepo := repo.NewUserRepository(db)
	roleRepo := repo.NewRoleRepository(db)
	userDashboardUC := usecase.NewUserUseCase(db, userRepo, roleRepo)
	userDashboardController := dashboard.NewUserDashboardController(userDashboardUC)

	return userDashboardController
}

func InitCategoryDashboard(db *gorm.DB) *dashboard.CategoryDashboardController {
	categoryRepo := repo.NewCategoryRepository(db)
	categoryUC := usecase.NewCategoryUseCase(db, categoryRepo)
	categoryController := dashboard.NewCategoryController(categoryUC)

	return categoryController
}

func InitRoleDashboard(db *gorm.DB) *dashboard.RoleController {
	roleRepo := repo.NewRoleRepository(db)
	permissionsRepo := repo.NewPermissionsRepository(db)
	rolePermissionsRepo := repo.NewRolePermissionsRepository(db)
	roleUC := usecase.NewRoleUseCase(db, roleRepo, permissionsRepo, rolePermissionsRepo)
	roleController := dashboard.NewRoleController(roleUC)

	return roleController
}

func InitPermissionDashboard(db *gorm.DB) *dashboard.PermissionController {
	permissionRepo := repo.NewPermissionsRepository(db)
	permissionUC := usecase.NewPermissionUseCase(db, permissionRepo)
	permissionController := dashboard.NewPermissionController(permissionUC)

	return permissionController
}

func InitRolePermissionsDashboard(db *gorm.DB) *dashboard.RolePermissionsController {
	rolePermissionsRepo := repo.NewRolePermissionsRepository(db)
	rolePermissionsUC := usecase.NewRolePermissionsUsecase(rolePermissionsRepo)
	rolePermissionsController := dashboard.NewRolePermissionsController(rolePermissionsUC)

	return rolePermissionsController
}
