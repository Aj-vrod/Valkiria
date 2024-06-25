package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/Valkiria/pkg/movie"
)

func StartServer() {
	fmt.Println("Starting server. Listening on port 3000...")

	r := mux.NewRouter()
	r.HandleFunc("/movies/{id:[0-9]+}", movie.GetMovieHandler).Methods("GET")
	r.HandleFunc("/movies", movie.CreateMovieHandler).Methods("POST")
	r.HandleFunc("/movies/{id:[0-9]+}", movie.DeleteMovieHandler).Methods("DELETE")

	err := http.ListenAndServe(":3000", r)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed :(\n)")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
