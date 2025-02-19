package config

import (
	"database/sql"
	migrate "github.com/rubenv/sql-migrate"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func RunMigrations() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer db.Close()

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}

	log.Printf("Applied %d migrations\n", n)
}
