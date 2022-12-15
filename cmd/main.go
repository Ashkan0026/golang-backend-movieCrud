package main

import (
	"net/http"

	"github.com/Ashkan0026/go-movie-crud/controller"
	"github.com/Ashkan0026/go-movie-crud/models"
	"github.com/gorilla/mux"
)

func main() {
	models.Init()
	handler := mux.NewRouter()
	handler.HandleFunc("/movies", controller.GetAllMovies).Methods("GET")
	handler.HandleFunc("/movies/{id}", controller.GetMovieById).Methods("GET")
	handler.HandleFunc("/movies", controller.CreateMovie).Methods("POST")
	handler.HandleFunc("/movies/{id}", controller.UpdateMovie).Methods("PUT")
	handler.HandleFunc("/movies/{id}", controller.DeleteMovie)
	handler.HandleFunc("/moviesyears", controller.GetMovieBuildYears).Methods("GET")

	http.ListenAndServe(":7070", handler)
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
