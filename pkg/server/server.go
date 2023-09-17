package server

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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
	var movie = "Matrix"
	switch genre {
	case "horror":
		movie = "Saw"
	case "comedy":
		movie = "What happens in Vegas"
	case "action":
		movie = "Die Hard"
	case "thriller":
		movie = "The Silence of the Lambs"
	case "drama":
		movie = "The green mile"
	case "fantasy":
		movie = "The Lord of the Rings"
	case "sci-fi":
		movie = "The Matrix"
	case "adventure":
		movie = "The Hobbit"
	case "western":
		movie = "The Good, the Bad and the Ugly"
	case "crime":
		movie = "The Godfather"
	}
	io.WriteString(w, resp+movie+"\n")
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
