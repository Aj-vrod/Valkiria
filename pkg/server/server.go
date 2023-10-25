package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/Valkiria/pkg/hello"
	"github.com/Valkiria/pkg/movie"
	"github.com/Valkiria/pkg/root"
)

func StartServer() {
	fmt.Println("Starting server. Listening on port 3000...")

	r := mux.NewRouter()
	r.HandleFunc("/", root.HomeHandler).Methods("GET")
	r.HandleFunc("/hello", hello.HelloHandler).Methods("GET")
	r.HandleFunc("/movie/{genre:[a-z]+}", movie.GetMovieHandler).Methods("GET")
	r.HandleFunc("/movie", movie.CreateMovieHandler).Methods("POST")
	r.HandleFunc("/movie/{genre:[a-z]+}", movie.DeleteMovieHandler).Methods("DELETE")

	err := http.ListenAndServe(":3000", r)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed :(\n)")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
