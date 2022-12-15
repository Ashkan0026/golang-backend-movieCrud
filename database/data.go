package data

import (
	"fmt"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "@shk1382"
	dbname   = "postgres"
)

var psqlConn string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
var db *gorm.DB

func Init() {
	dbC, err := gorm.Open(postgres.Open(psqlConn))
	if err != nil {
		panic(err)
	}
	db = dbC
}

func GetDB() *gorm.DB {
	return db
}
