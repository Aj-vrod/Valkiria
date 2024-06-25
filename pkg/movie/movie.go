package movie

import (
	"github.com/Valkiria/pkg/database"
)

type Movie struct {
	ID    int    `json:"movie_id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
}

func (m *Movie) getProduct() error {
	return database.DB.QueryRow(
		"SELECT name, genre FROM movies WHERE movie_id = $1", m.ID,
	).Scan(&m.Name, &m.Genre)
}

func (m *Movie) createProduct() error {
	return database.DB.QueryRow(
		"INSERT INTO movies(name, genre) VALUES($1, $2) RETURNING movie_id", m.Name, m.Genre,
	).Scan(&m.ID)
}

func (m *Movie) deleteProduct() error {
	_, err := database.DB.Exec(
		"DELETE FROM movies WHERE movie_id = $1", m.ID,
	)

	return err
}
