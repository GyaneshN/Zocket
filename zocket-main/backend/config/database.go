package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database is unreachable:", err)
	}

	fmt.Println("Connected to the database successfully!")
	DB = db
}
