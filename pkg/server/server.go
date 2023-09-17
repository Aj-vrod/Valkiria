package server

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Valkiria/pkg/movie"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Hello, user!\n")
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /movie request\n")
	genre := r.URL.Query().Get("genre")
	resp := "Movie of today is "
	w.WriteHeader(http.StatusOK)
	pickedMovie := movie.PickMovie(genre)

	io.WriteString(w, resp+pickedMovie+"\n")
}

func StartServer() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/movie", getMovie)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed :(\n)")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
