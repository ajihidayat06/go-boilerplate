package config

import (
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
