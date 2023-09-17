package movie

func PickMovie(genre string) string {
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
