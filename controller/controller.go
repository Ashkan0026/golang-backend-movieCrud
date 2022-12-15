package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Ashkan0026/go-movie-crud/models"
	"github.com/Ashkan0026/go-movie-crud/utility"
	"github.com/gorilla/mux"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	students := models.GetAllMovies()
	data, err := json.Marshal(students)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	numberId, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}
	student := models.GetMovieById(numberId)
	data, err := json.Marshal(student)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	models.Delete(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Movie with id : " + strconv.Itoa(id) + " Deleted successfully"))
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	movie := &models.Movie{}
	utility.ParseRequestBody(r, movie)
	movie = models.AddMovie(movie)
	data, err := json.Marshal(movie)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func GetMovieBuildYears(w http.ResponseWriter, r *http.Request) {
	years := models.GetMovieBuiltYears()
	data, err := json.Marshal(years)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	movie := models.GetMovieById(ID)
	var newMovie = &models.Movie{}
	utility.ParseRequestBody(r, newMovie)
	values := make(map[string]interface{})
	if newMovie.Name != "" {
		values["name"] = newMovie.Name
	}
	if newMovie.ID != 0 {
		values["id"] = newMovie.ID
	}
	if newMovie.Director != "" {
		values["director"] = newMovie.Director
	}
	if newMovie.Rating != 0 {
		values["rating"] = newMovie.Rating
	}
	if newMovie.Year != 0 {
		values["year"] = newMovie.Year
	}
	movie.UpdateMovie(ID, values)
}
