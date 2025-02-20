package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"go-boilerplate/config"
	"go-boilerplate/internal/middleware"
	"go-boilerplate/internal/router"
	"go-boilerplate/internal/utils"
	"go-boilerplate/pkg/logger"
	"log"
	"os"
)

func main() {
	logger.InitLogger()
	logger.Info("Starting API server...", nil)

	config.LoadEnv()
	db := config.InitDatabase()
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})
	//config.RunMigrations()

	// Security middleware: Helmet untuk secure headers
	app.Use(helmet.New())
	// CORS middleware
	app.Use(config.CorsConfig())

	app.Use(middleware.LoggingMiddleware)

	utils.InitValidator()
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
