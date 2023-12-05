package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func SetupDB() *sql.DB {
	const connString = "user=unreal dbname=unreal sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
