package movie

import (
	"fmt"
	"io"
	"net/http"
)

func pickMovie(genre string) string {
	var movie string = "Mamma Mia"
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
	return movie
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /movie request\n")
	genre := r.URL.Query().Get("genre")
	resp := "Movie of today is "
	w.WriteHeader(http.StatusOK)
	pickedMovie := pickMovie(genre)

	io.WriteString(w, resp+pickedMovie+"\n")
}
