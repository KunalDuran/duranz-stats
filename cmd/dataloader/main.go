package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/KunalDuran/duranz-stats/internal/data"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Time taken:", time.Since(start))
	}()

	dbHost := "localhost"
	dbPort := 5432
	dbUser := "postgres"
	dbName := "duranz"
	dbPass := "password"

	err := data.InitDB(dbHost, dbUser, dbPass, dbName, "postgres", dbPort)
	if err != nil {
		log.Fatal(err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Expected a command such as 'setup', 'delete', or 'process'.")
		return
	}

	command := os.Args[1]

	switch command {
	case "setup":
		data.CreateTables()
	case "delete":
		data.TruncateTables()
	case "venue", "team", "player", "match", "matchstats", "playerstats", "scorecard", "all":
		process(command, os.Args[2:])
	default:
		fmt.Println("Unknown command:", command)
	}
}

func process(cmd string, args []string) {
	var leagueFormat string
	var fileName string

	flags := flag.NewFlagSet("process", flag.ExitOnError)
	flags.StringVar(&leagueFormat, "league", "", "Specify the league format (e.g., odi, test, t20).")
	flags.StringVar(&fileName, "file", "", "Specify a specific file to process.")

	flags.Parse(args)

	if leagueFormat == "" {
		fmt.Println("Please specify the league format using the -league option.")
		return
	}

	fmt.Println("Processing for league:", leagueFormat)

	if fileName != "" {
		fmt.Println("Processing file:", fileName)
	}

	DATASET_BASE = DATASET_BASE + data.GamePath[leagueFormat] + "/"

	RunAllProcess(cmd, fileName)

}
