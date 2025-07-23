package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	_ = godotenv.Load() // Загружает переменные из .env

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSL")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", user, password, dbname, ssl)
	return sql.Open("postgres", dsn)
}
