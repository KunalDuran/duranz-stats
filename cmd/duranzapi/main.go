package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/KunalDuran/duranz-stats/internal/cache"
	"github.com/KunalDuran/duranz-stats/internal/data"
	_ "github.com/go-sql-driver/mysql"

	"github.com/julienschmidt/httprouter"
)

var PWD, _ = os.Getwd()

var DATASET_BASE = PWD + `/datasets/odis_json/`

func main() {

	dbHost := "localhost"
	dbPort := 3306
	dbUser := "root"
	dbName := "duranz"
	dbPass := ""
	port := ":5000"

	err := data.InitDB(dbHost, dbUser, dbPass, dbName, dbPort)
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

	router := httprouter.New()
	router.RedirectTrailingSlash = true
	addRouteHandlers(router)

	fmt.Println("Duranz API initialized")
	log.Fatal(http.ListenAndServe(port, router))
}
