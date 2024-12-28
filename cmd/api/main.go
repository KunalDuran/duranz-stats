package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/KunalDuran/duranz-stats/internal/cache"
	"github.com/KunalDuran/duranz-stats/internal/data"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

var PWD, _ = os.Getwd()

var DATASET_BASE = PWD + `/datasets/odis_json/`

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	port := os.Getenv("PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")

	err := data.InitDB(dbHost, dbUser, dbPass, dbName, "postgres", dbPort)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Connect the cache server
	cacheDataHost := "localhost"
	cacheDataPort := "6379"
	err = cache.InitRedis(cacheDataHost, cacheDataPort)
	if err != nil {
		log.Panic(err)
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	addRouteHandlers(router)

	fmt.Println("Duranz API initialized")
	log.Fatal(http.ListenAndServe(port, router))
}
