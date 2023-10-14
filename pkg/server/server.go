package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/Valkiria/pkg/hello"
	"github.com/Valkiria/pkg/movie"
	"github.com/Valkiria/pkg/root"
)

func StartServer() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", root.GetRoot)
	http.HandleFunc("/hello", hello.GetHello)
	http.HandleFunc("/movie", movie.GetMovie)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed :(\n)")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
