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
	userUC := usecase.NewUserUseCase(db, userRepo)
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
	userDashboardUC := usecase.NewUserUseCase(db, userRepo)
	userDashboardController := dashboard.NewUserDashboardController(userDashboardUC)

	return userDashboardController
}

func InitCategoryDashboard(db *gorm.DB) *dashboard.CategoryDashboardController {
    categoryRepo := repo.NewCategoryRepository(db)
    categoryUC := usecase.NewCategoryUseCase(categoryRepo)
    categoryController := dashboard.NewCategoryController(categoryUC)

    return categoryController
}