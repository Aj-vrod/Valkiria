package movie

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var DB *sql.DB

type movie struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
}

func respondeWithJSON(w http.ResponseWriter, code int, m interface{}) {
	resp, _ := json.Marshal(m)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}

func (m *movie) getProduct() error {
	return DB.QueryRow("SELECT name, genre FROM movie WHERE genre = $1", m.Genre).Scan(&m.Name, &m.Genre)
}

func (m *movie) createProduct() error {
	err := DB.QueryRow("INSERT INTO movie(name, genre) VALUES($1, $2) RETURNING id", m.Name, m.Genre).Scan(&m.ID)

	return err
}

func (m *movie) deleteProduct() error {
	_, err := DB.Exec("DELETE FROM movie WHERE genre = $1", m.Genre)

	return err
}

func GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	genre := vars["genre"]

	var m movie
	m.Genre = genre
	if err := m.getProduct(); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondeWithJSON(w, http.StatusNotFound, "Movie not found")
		default:
			respondeWithJSON(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondeWithJSON(w, http.StatusOK, m)
}

func CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	var m movie
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&m); err != nil {
		respondeWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := m.createProduct(); err != nil {
		respondeWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondeWithJSON(w, http.StatusCreated, m)

}

func DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	genre := vars["genre"]

	var m movie
	m.Genre = genre

	if err := m.deleteProduct(); err != nil {
		respondeWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondeWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
