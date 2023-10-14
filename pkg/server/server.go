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
	fmt.Println("Starting server. Listening on port 3000...")
	http.HandleFunc("/", root.GetRoot)
	http.HandleFunc("/hello", hello.GetHello)
	http.HandleFunc("/movie", movie.GetMovie)

	err := http.ListenAndServe(":3000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed :(\n)")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
