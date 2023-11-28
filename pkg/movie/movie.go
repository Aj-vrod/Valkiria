package movie

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Valkiria/pkg/database"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID    int    `json:"movie_id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
}

func respondeWithJSON(w http.ResponseWriter, code int, m interface{}) {
	resp, _ := json.Marshal(m)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}

func (m *Movie) getProduct() error {
	return database.DB.QueryRow("SELECT name, genre FROM movies WHERE movie_id = $1", m.ID).Scan(&m.Name, &m.Genre)
}

func (m *Movie) createProduct() error {
	err := database.DB.QueryRow("INSERT INTO movies(name, genre) VALUES($1, $2) RETURNING movie_id", m.Name, m.Genre).Scan(&m.ID)

	return err
}

func (m *Movie) deleteProduct() error {
	_, err := database.DB.Exec("DELETE FROM movies WHERE movie_id = $1", m.ID)

	return err
}

func GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondeWithJSON(w, http.StatusBadRequest, "Invalid product ID")
	}

	var m Movie
	m.ID = id
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
	var m Movie
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
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondeWithJSON(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var m Movie
	m.ID = id

	if err := m.deleteProduct(); err != nil {
		respondeWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondeWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
