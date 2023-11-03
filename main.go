package main

import (
	"fmt"
	"log"

	"github.com/Valkiria/pkg/database"
	"github.com/Valkiria/pkg/server"
	_ "github.com/lib/pq"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")
	server.StartServer()
}
