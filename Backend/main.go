package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	dsn := "host=localhost port=5432 user=DeSakke password=fakePassword dbname=UrlShortener sslmode=disable"
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	pingError := db.Ping()
	if pingError != nil {
		log.Fatal(pingError)
	}

	router := gin.Default()
	router.Use(attachDbMiddleWare(db))
	router.POST("/Url", UrlPostHandler)
	router.GET("/Url/:key", UrlGetHandler)

	router.Run("localhost:8080")
}

func attachDbMiddleWare(db *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("dbConnection", db)
		context.Next()
	}
}

func UrlPostHandler(context *gin.Context) {
	var url Url
	err := context.Bind(&url)

	if err != nil {
		log.Fatal(err)
	}

	var stringProvider defaultRandomStringProvider

	key, _ := stringProvider.RandomString(7)
	db, ok := context.MustGet("dbConnection").(*sql.DB)

	if !ok {
		log.Fatal("Failed to get DB connection")
	}

	PostUrl(db, key, url.Url)

	context.JSON(http.StatusOK, KeyUrl{key, url.Url})
}

func UrlGetHandler(context *gin.Context) {

	key := context.Param("key")
	if key == "" {
		context.JSON(http.StatusBadRequest, "Please use key parameter")
		return
	}

	db, ok := context.MustGet("dbConnection").(*sql.DB)

	if !ok {
		log.Fatal("Failed to get DB connection")
	}

	url, ok := GetUrl(db, key)

	if !ok {
		context.JSON(http.StatusNotFound, "Failed to find key")
		return
	}

	context.JSON(http.StatusOK, KeyUrl{key, url})
}
