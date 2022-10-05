package main

import (
	"log"
	"native-restful-api/config"
	"native-restful-api/handler"
	"net/http"
	"os"
)

func main() {
	db := config.DbConnect()

	movie := handler.NewBaseHandler(db)

	http.HandleFunc("/test", movie.TestHandler)
	http.HandleFunc("/movies", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			movie.GetMovies(w, r)

		case http.MethodPost:
			movie.CreateMovie(w, r)

		case http.MethodPut:
			movie.UpdateMovie(w, r)

		case http.MethodDelete:
			movie.DeleteMovie(w, r)
		}
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
