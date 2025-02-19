package router

import (
	"go-boilerplate/internal/controllers"
	"go-boilerplate/internal/repo"
	"go-boilerplate/internal/usecase"
	"gorm.io/gorm"
)

func InitUser(db *gorm.DB) *controllers.UserController {
	userRepo := repo.NewUserRepository(db)
	userUC := usecase.NewUserUseCase(userRepo)
	userController := controllers.NewUserController(userUC)

	return userController
}
