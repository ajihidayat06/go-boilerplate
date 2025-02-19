package main

import (
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/config"
	"go-boilerplate/internal/middleware"
	"go-boilerplate/internal/router"
	"go-boilerplate/pkg/logger"
	"log"
	"os"
)

func main() {
	logger.InitLogger()
	logger.Info("Starting API server...", nil)

	config.LoadEnv()
	db := config.InitDatabase()
	app := fiber.New()
	//config.RunMigrations()

	app.Use(middleware.LoggingMiddleware)

	router.SetupRoutes(app, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server started on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
