package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func GetDB() (*sql.DB, error) {
	open, err := sql.Open("mysql", ConnectionString)

	DB = open
	return open, err
}

func Connect() {
	db, err := GetDB()
	if err != nil {
		log.Printf("Error with database: %s", err.Error())
		return
	}

	err = db.Ping()

	if err != nil {
		log.Printf("Error making connection with database: %s", err.Error())
		return
	}
}
