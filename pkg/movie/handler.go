package movie

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Writes response with status and body
func respondeWithJSON(w http.ResponseWriter, code int, m interface{}) {
	resp, _ := json.Marshal(m)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
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
