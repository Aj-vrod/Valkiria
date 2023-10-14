package movie

import "testing"

var genres = []string{"horror", "comedy", "drama", "fantasy", "sci-fi", "thriller", "action", "adventure", "western", "crime"}
var movies = []string{"Saw", "What happens in Vegas", "The green mile", "The Lord of the Rings", "The Matrix", "The Silence of the Lambs", "Die Hard", "The Hobbit", "The Good, the Bad and the Ugly", "The Godfather"}

func TestPickMovie(t *testing.T) {
	for i, v := range genres {
		got := pickMovie(v)
		exp := movies[i]

		if got != exp {
			t.Errorf("got %s, expected %s", got, exp)
		}
	}
}
