package database

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password01"
	dbname   = "valkiria"
)

func InitDB() error {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	return DB.Ping()
}
