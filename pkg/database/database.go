package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Valgard/godotenv"
)

var DB *sql.DB

func InitDB() error {
	dotenv := godotenv.New()
	if err := dotenv.Load(".env"); err != nil {
		return err
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	return DB.Ping()
}
