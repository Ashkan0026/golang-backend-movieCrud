package models

import (
	data "github.com/Ashkan0026/go-movie-crud/database"
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Name     string  `json:"name"`
	Rating   float32 `json:"rating"`
	Director string  `json:"director"`
	Year     int     `json:"year"`
}

var db *gorm.DB

func Init() {
	data.Init()
	db = data.GetDB()
	db.AutoMigrate(&Movie{})
}

func AddMovie(m *Movie) *Movie {
	db.Create(&m)
	return m
}

func DeleteMovie(ID int) {
	var movie Movie
	db.Where("id=?", ID).Delete(&movie)
}

func GetAllMovies() []Movie {
	var movies []Movie
	db.Find(&movies)
	return movies
}

func GetMovieById(ID int) *Movie {
	var movie Movie
	db.Where("id=?", ID).Find(&movie)
	return &movie
}

func Delete(ID int) {
	var movie Movie
	db.Where("id=?", ID).Delete(&movie)
}

func (m *Movie) UpdateMovie(ID int, values map[string]interface{}) {
	db.Model(m).Where("ID=?", ID).Updates(values)
}

func GetMovieBuiltYears() []int {
	var years []int
	db.Model(&Movie{}).Pluck("year", &years)
	return years
}
