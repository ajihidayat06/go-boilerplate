package main

import (
	"go-boilerplate/config"
	"go-boilerplate/internal/middleware"
	"go-boilerplate/internal/router"
	"go-boilerplate/internal/seeder"
	"go-boilerplate/internal/utils"
	"go-boilerplate/pkg/logger"

	// "go-boilerplate/pkg/redis"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
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

	// err := redis.InitRedis()
    // if err != nil {
    //     log.Fatalf("Failed to initialize Redis: %v", err)
    // }

	if err := seeder.SeedSuperAdmin(db); err != nil {
        logger.Error("Failed to seed superadmin", err)
    }

	// Security middleware: Helmet untuk secure headers
	app.Use(helmet.New())
	// CORS middleware
	app.Use(config.CorsConfig())

	app.Use(middleware.LoggingMiddleware)
	app.Use(middleware.RecoverMiddleware())

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
