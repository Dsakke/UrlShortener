package main

import (
	"database/sql"
	"log"
)

func PostUrl(db *sql.DB, key string, url string) sql.Result {
	result, err := db.Exec("INSERT INTO urls (urlKey, url) VALUES ($1, $2)", key, url)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func GetUrl(db *sql.DB, key string) (string, bool) {
	url := ""
	err := db.QueryRow("SELECT url FROM urls WHERE urlKey = $1", key).Scan(&url)

	return url, err == nil
}
