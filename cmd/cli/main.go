package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/KunalDuran/duranz-stats/internal/data"
	"github.com/KunalDuran/duranz-stats/internal/mapper"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	_, err = data.InitDB(dbHost, dbPort, dbUser, dbPass)
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
		data.Setup("duranz")
	case "delete":
		data.Delete()
	case "venue", "team", "player", "match":
		process(os.Args[2:])
	default:
		fmt.Println("Unknown command:", command)
	}
}

func process(args []string) {
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

	mapper.RunAllProcess("", fileName)

}
