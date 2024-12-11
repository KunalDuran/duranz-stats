package mapper

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/KunalDuran/duranz-stats/internal/data"
	"github.com/KunalDuran/duranz-stats/internal/models"
)

func VenueMapper(venueName, city string) {

	if val, ok := data.MappedVenues[venueName]; ok && city == val {
		return
	}
	venue := data.Venue{
		VenueName: venueName,
		City:      city,
	}

	result := data.DB.Create(&venue)
	if result.Error != nil {
		panic(result.Error)
	}
	data.MappedVenues[venueName] = city
}

func TeamMapper(teams []string, teamType string) {
	for _, team := range teams {
		if _, ok := data.MappedTeams[team]; ok {
			continue
		}

		teamObj := data.Team{
			TeamName: team,
			TeamType: teamType,
		}

		result := data.DB.Create(&teamObj)
		if result.Error != nil {
			panic(result.Error)
		}
		data.MappedTeams[team] = teamType
	}
}

func PlayerMapper(players map[string]string) {
	var allTeamPlayers []string
	for _, playerID := range players {
		allTeamPlayers = append(allTeamPlayers, playerID)
	}

	// Check existing players
	var existingPlayers []data.CricketPlayer
	result := data.DB.Where("cricsheet_id IN ?", allTeamPlayers).Find(&existingPlayers)
	if result.Error != nil {
		panic(result.Error)
	}

	// Create map of existing player IDs
	existingPlayerMap := make(map[string]string)
	for _, player := range existingPlayers {
		existingPlayerMap[player.CricsheetID] = player.PlayerName
	}

	// Prepare new players for insertion
	var newPlayers []data.CricketPlayer
	for playerName, playerID := range players {
		if _, exists := existingPlayerMap[playerID]; !exists {
			newPlayers = append(newPlayers, data.CricketPlayer{
				PlayerName:  playerName,
				CricsheetID: playerID,
			})
		}
	}

	// Batch insert new players
	if len(newPlayers) > 0 {
		result := data.DB.Create(&newPlayers)
		if result.Error != nil {
			panic(result.Error)
		}
	}
}

func MatchMapper(match models.Match, fileName string) {

	fileName = strings.Replace(fileName, ".json", "", -1)
	leagueID := models.AllDuranzLeagues[strings.ToLower(match.Info.MatchType)]

	if match.Info.Event.Name == "Indian Premier League" {
		leagueID = models.AllDuranzLeagues["ipl"]
	}
	venueID := data.GetVenueID(match.Info.Venue, match.Info.City)

	var startDate string
	var matchDate time.Time
	var seasonID int
	if len(match.Info.Dates) > 0 {
		startDate = match.Info.Dates[0]
		matchDate, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			fmt.Println("Error in parsing match date", err)
		}
		seasonID = matchDate.Year()

	}

	if len(match.Info.Teams) != 2 || venueID == 0 || leagueID == 0 || startDate == "" {
		fmt.Println("Error in match mapper process")
		if len(match.Info.Teams) != 2 {
			data.InsertErrorLog(data.DATABASE_ERROR, `More than 2 teams`, fileName)
		}
		if venueID == 0 {
			data.InsertErrorLog(data.DATABASE_ERROR, `Venue not found `+match.Info.Venue, fileName)
		}
		if leagueID == 0 {
			data.InsertErrorLog(data.LEAGUE_NOT_FOUND, `League not found `+match.Info.MatchType, fileName)
		}
		if startDate == "" {
			data.InsertErrorLog(data.DATETIME_ERROR, `Start date not found`, fileName)
		}
		return
	}

	home := match.Info.Teams[0]
	away := match.Info.Teams[1]
	homeTeamID := data.GetTeamID(home, match.Info.TeamType)
	awayTeamID := data.GetTeamID(away, match.Info.TeamType)

	var winningTeam, tossWinner int
	if home == match.Info.Outcome.Winner {
		winningTeam = homeTeamID
	} else if away == match.Info.Outcome.Winner {
		winningTeam = awayTeamID
	} else if match.Info.Outcome.Eliminator == home {
		winningTeam = homeTeamID
	} else if match.Info.Outcome.Eliminator == away {
		winningTeam = awayTeamID
	}

	if home == match.Info.Toss.Winner {
		tossWinner = homeTeamID
	} else if away == match.Info.Toss.Winner {
		tossWinner = awayTeamID
	}

	tossDecision := match.Info.Toss.Decision

	if tossWinner == 0 {
		fmt.Println("Error in mapping match ", fileName)
		fmt.Println("TossWinner ", tossWinner)
		data.InsertErrorLog(data.DATABASE_ERROR, `Toss winner not found`, fileName)
		return
	}

	var resultStr string
	if match.Info.Outcome.Result == "no result" {
		resultStr = "no result"
	} else if match.Info.Outcome.Result == "tie" {
		resultStr = "tie"
	} else if match.Info.Outcome.Result == "draw" {
		resultStr = "draw"
	}

	if resultStr != "no result" && resultStr != "tie" && resultStr != "draw" {
		resultStr = match.Info.Outcome.Winner + " Won by "
		if match.Info.Outcome.By.Runs > 0 {
			resultStr += strconv.Itoa(match.Info.Outcome.By.Runs) + " Runs"
		} else if match.Info.Outcome.By.Wickets > 0 {
			resultStr += strconv.Itoa(match.Info.Outcome.By.Wickets) + " Wickets"
		}
	}
	if match.Info.Outcome.Method == "D/L" {
		resultStr += " (D/L Method)"
	}

	matchReferees := strings.Join(match.Info.MatchReferees, ";")
	reserveUmpires := strings.Join(match.Info.ReserveUmpires, ";")
	tvUmpires := strings.Join(match.Info.TvUmpires, ";")
	umpires := strings.Join(match.Info.Umpires, ";")

	cricketMatch := data.CricketMatch{
		LeagueID:          &leagueID,
		SeasonID:          &seasonID,
		HomeTeamID:        &homeTeamID,
		AwayTeamID:        &awayTeamID,
		HomeTeamName:      home,
		AwayTeamName:      away,
		VenueID:           &venueID,
		MatchDate:         &matchDate,
		MatchDateMulti:    strings.Join(match.Info.Dates, ";"),
		CricsheetFileName: fileName,
		Result:            resultStr,
		TossWinner:        &tossWinner,
		TossDecision:      tossDecision,
		WinningTeam:       &winningTeam,
		Gender:            match.Info.Gender,
		MatchRefrees:      matchReferees,
		ReserveUmpires:    reserveUmpires,
		TVUmpires:         tvUmpires,
		Umpires:           umpires,
	}

	if len(match.Info.PlayerOfMatch) > 0 {
		peopleRegistry := match.Info.Register.People
		cricketMatch.ManOfTheMatch = data.GetPlayerID(peopleRegistry[match.Info.PlayerOfMatch[0]])
	}

	result := data.DB.Create(&cricketMatch)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}
