package database

import (
	"database/sql"
	"os"

	"github.com/Valgard/godotenv"
)

var DB *sql.DB

func InitDB() error {
	dotenv := godotenv.New()
	if err := dotenv.Load(".env"); err != nil {
		return err
	}

	dns := os.Getenv("DB_DNS")
	if dns == "" {
		panic("Missing env variable DB_DNS")
	}

	var err error
	DB, err = sql.Open("postgres", dns)
	if err != nil {
		return err
	}

	return DB.Ping()
}
