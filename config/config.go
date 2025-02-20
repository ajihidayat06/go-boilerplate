package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func InitDatabase() *gorm.DB {
	//dsn := os.Getenv("DATABASE_URL")
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	log.Fatal("Failed to connect to database")
	//}
	return &gorm.DB{}
}

func CorsConfig() fiber.Handler {
	return cors.New(cors.Config{
		//AllowOrigins:     "https://example.com, https://www.example.com",
		AllowOrigins:  "*",
		AllowMethods:  "GET,POST,PUT,DELETE",
		AllowHeaders:  "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders: "Content-Length",
		//AllowCredentials: true,
		MaxAge: 3600,
	})
}
