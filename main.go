package main

import (
	"database/sql"
	"fmt"

	"github.com/Valkiria/pkg/movie"
	"github.com/Valkiria/pkg/server"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password01"
	dbname   = "valkiria"
)

func main() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	movie.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer movie.DB.Close()

	err = movie.DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	server.StartServer()
}
