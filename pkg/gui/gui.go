package gui

import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
)

func OnReady() {
	systray.SetIcon(getIcon("./assets/cinema.png"))
	systray.AddMenuItem("Shoot", "pick a random movie")
	systray.AddMenuItem("Mood", "pick a movie with genre")
	systray.AddMenuItem("Quit", "quit the app")
}

func OnExit() {

}

func getIcon(path string) []byte {
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Failed to get icon", err)
	}
	return b
}
