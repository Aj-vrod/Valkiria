package main

import (
	"fmt"
	"log"

	"github.com/Valkiria/pkg/database"
	"github.com/Valkiria/pkg/gui"
	"github.com/Valkiria/pkg/server"
	"github.com/getlantern/systray"
	_ "github.com/lib/pq"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer database.DB.Close()

	systray.Run(gui.OnReady, gui.OnExit)

	fmt.Println("Connected!")
	server.StartServer()
}
