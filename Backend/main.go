package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// dsn := "postgres://DeSakke:fakePassword@localhost:5432/UrlShortener?sslmode=disable"

	dsn := "host=localhost port=5432 user=DeSakke password=fakePassword dbname=UrlShortener sslmode=disable"
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	pingError := db.Ping()
	if pingError != nil {
		log.Fatal(pingError)
	}
}
