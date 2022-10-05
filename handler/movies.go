package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"native-restful-api/helper"
	"native-restful-api/models"
	"net/http"
	"strconv"
)

type BaseHandler struct {
	db *sql.DB
}

func NewBaseHandler(db *sql.DB) *BaseHandler {
	return &BaseHandler{db: db}
}

func (h *BaseHandler) TestHandler(res http.ResponseWriter, req *http.Request) {
	HandlerMessage := []byte(`{"success": true, "message": "Hello World"}`)

	helper.Json(res, http.StatusOK, HandlerMessage)
}

func (h *BaseHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.JsonError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	rows, err := h.db.Query("SELECT * FROM movies")
	if err != nil {
		helper.JsonError(w, http.StatusInternalServerError, "Failed to fetch movies")
		return
	}

	var movies []models.Movie

	for rows.Next() {
		var movie models.Movie
		err = rows.Scan(&movie.Id, &movie.Title, &movie.Description, &movie.Year)
		if err != nil {
			helper.JsonError(w, http.StatusInternalServerError, "Failed to fetch movies")
			return
		}
		movies = append(movies, movie)
	}

	data, err := json.Marshal(&movies)
	if err != nil {
		helper.JsonError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	helper.Json(w, http.StatusOK, data)
	return
}

func (h *BaseHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var movie models.Movie

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		helper.JsonError(w, http.StatusInternalServerError, "Failed to decode request body")
		return
	}

	query := fmt.Sprintf("INSERT INTO \"movies\" (\"title\", \"description\", \"release_year\") VALUES ('%s', '%s', %d)", movie.Title, movie.Description, movie.Year)

	result, err := h.db.Exec(query)
	_ = result

	if err != nil {
		log.Print(err)
		helper.JsonError(w, http.StatusInternalServerError, "Failed to insert new movie")
		return
	}

	helper.Json(w, http.StatusOK, []byte(`{"success": true, "data": "Movie successfully created"}`))
	return
}

func (h *BaseHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if _, ok := r.URL.Query()["id"]; !ok {
		helper.JsonError(w, http.StatusBadRequest, "id is required")
		return
	}

	var exists bool
	isExists := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM \"movies\" WHERE \"id\" = %s)", r.URL.Query()["id"][0])
	checkErr := h.db.QueryRow(isExists).Scan(&exists)

	if checkErr != nil || !exists {
		helper.JsonError(w, http.StatusInternalServerError, "Movie not found")
		return
	}

	var movie models.Movie

	err := json.NewDecoder(r.Body).Decode(&movie)

	if err != nil {
		helper.JsonError(w, http.StatusInternalServerError, "Failed to decode request body")
		return
	}

	id, _ := strconv.Atoi(r.URL.Query()["id"][0])
	query := fmt.Sprintf("UPDATE \"movies\" SET \"title\" = '%s', \"description\" = '%s', \"release_year\" = %d WHERE \"id\" = %d", movie.Title, movie.Description, movie.Year, id)

	result, err := h.db.Exec(query)
	_ = result

	if err != nil {
		log.Print(err)
		helper.JsonError(w, http.StatusInternalServerError, "Failed to update movie")
		return
	}

	helper.Json(w, http.StatusOK, []byte(`{"success": true, "data": "Movie successfully updated"}`))
	return
}

func (h *BaseHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	if _, ok := r.URL.Query()["id"]; !ok {
		helper.JsonError(w, http.StatusBadRequest, "id is required")
		return
	}

	var exists bool
	isExists := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM \"movies\" WHERE \"id\" = %s)", r.URL.Query()["id"][0])
	err := h.db.QueryRow(isExists).Scan(&exists)

	if err != nil || !exists {
		helper.JsonError(w, http.StatusInternalServerError, "Movie not found")
		return
	}

	query := fmt.Sprintf("DELETE FROM \"movies\" WHERE \"id\" = %s", r.URL.Query()["id"][0])

	result, err := h.db.Exec(query)
	_ = result

	if err != nil {
		log.Print(err)
		helper.JsonError(w, http.StatusInternalServerError, "Failed to delete movie")
		return
	}

	helper.Json(w, http.StatusOK, []byte(`{"success": true, "data": "Movie successfully deleted"}`))
	return
}
