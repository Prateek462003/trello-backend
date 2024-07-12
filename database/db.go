package database

import (
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB *sql.DB

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	DB, err = sql.Open("postgres", os.Getenv("URI"))
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
