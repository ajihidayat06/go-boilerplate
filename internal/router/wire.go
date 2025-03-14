package router

import (
	"go-boilerplate/internal/controllers"
	"go-boilerplate/internal/controllers/dashboard"
	"go-boilerplate/internal/repo"
	"go-boilerplate/internal/usecase"
	dashboard2 "go-boilerplate/internal/usecase/dashboard"
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
	authUC := dashboard2.NewAuthUseCase(db, userRepo)
	authController := dashboard.NewAuthController(authUC)

	return authController
}
