package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/KunalDuran/duranz-stats/internal/data"
	"github.com/KunalDuran/duranz-stats/internal/mapper"
	"github.com/KunalDuran/duranz-stats/internal/models"
)

var PWD, _ = os.Getwd()

var DATASET_BASE = PWD + `/datasets/`

func RunAllProcess(process, fileName string) {
	var newFiles []string
	if fileName != "" {
		newFiles = append(newFiles, fileName)
	} else {
		allFiles := ListFiles(DATASET_BASE)
		newFiles = GetNewFiles(allFiles)
	}

	var namesMap = map[string]string{}
	if process == "all" || process == "player" {
		in, err := os.ReadFile("docs/names.csv")
		if err != nil {
			log.Fatal(err)
		}
		r := csv.NewReader(strings.NewReader(string(in)))

		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			namesMap[record[0]] = record[1]
		}
	}

	for _, file := range newFiles {
		var mappingInfo models.MappingInfo
		match, err := mapper.GetCricsheetData(DATASET_BASE + file)
		if err != nil {
			data.InsertErrorLog(data.CRICSHEET_FILE_ERROR, `Error in fetching file `, file, err.Error())
			return
		}

		mappingInfo.LeagueID = models.AllDuranzLeagues[match.Info.MatchType]

		if match.Info.TeamType == "club" && (match.Info.Event.Name == "Indian Premier League" || strings.ToLower(match.Info.Event.Name) == "ipl") {
			match.Info.TeamType = "ipl"
		}

		if process == "venue" || process == "all" {
			mapper.VenueMapper(match.Info.Venue, match.Info.City)
			mappingInfo.Venue = true
		}

		if process == "player" || process == "all" {
			mapper.PlayerMapper(match.Info.Register.People, namesMap)
			mappingInfo.Players = true
		}

		if process == "team" || process == "all" {
			mapper.TeamMapper(match.Info.Teams, match.Info.TeamType)
			mappingInfo.Teams = true
		}

		if process == "match" || process == "all" {
			mapper.MatchMapper(match, file)
			mappingInfo.Match = true
		}

		if process == "matchstats" || process == "all" {
			mapper.ProcessMatchStats(match, file)
			mappingInfo.MatchStats = true
		}

		if process == "playerstats" || process == "all" {
			mapper.ProcessPlayerStats(match, file)
			mappingInfo.PlayerStats = true
		}

		if process == "scorecard" || process == "all" {
			mapper.ScorecardMapper(match, file)
			mappingInfo.ScoreCard = true
		}

		data.InsertMappingInfo(file, mappingInfo)
	}
}

func ListFiles(path string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		data.InsertErrorLog(data.CRICSHEET_FILE_ERROR, `Error in ListFiles : `, path, err.Error())
		return []string{}
	}

	var fileList []string
	for _, f := range files {
		if strings.Contains(f.Name(), "json") {
			fileList = append(fileList, f.Name())
		}
	}
	return fileList
}

func GetNewFiles(allFiles []string) []string {
	objMappingDetails := data.GetMappingDetails()
	var newFiles []string

	for _, file := range allFiles {
		if row, exist := objMappingDetails[file]; !exist {
			if row.Match && row.MatchStats && row.PlayerStats && row.ScoreCard && row.Teams && row.Venue && row.Players {
				continue
			}
			newFiles = append(newFiles, file)
		}
	}

	return newFiles
}
