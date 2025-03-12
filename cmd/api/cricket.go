package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/KunalDuran/duranz-stats/internal/data"
	"github.com/KunalDuran/duranz-stats/internal/models"
	"github.com/KunalDuran/duranz-stats/internal/utils"
	"github.com/go-chi/chi/v5"
)

// GetCricsheetData : Reads the match json file
func GetCricsheetData(f_path string) (models.Match, error) {
	var matchData models.Match
	body, err := os.ReadFile(f_path)
	if err != nil {
		return matchData, err
	}

	err = json.Unmarshal(body, &matchData)
	if err != nil {
		return matchData, err
	}
	return matchData, nil
}

func MatchStats(w http.ResponseWriter, r *http.Request) {
	cricsheetID := chi.URLParam(r, "file")
	objScoreCard := data.GetMatchScoreCard(cricsheetID)
	// objScoreCard := mapper.ProcessScoreCard(match)
	final := utils.JSONMessageWrappedObj(http.StatusOK, objScoreCard.Data)
	utils.WebResponseJSONObject(w, r, http.StatusOK, final)
}

func PlayerStats(w http.ResponseWriter, r *http.Request) {
	playerName := chi.URLParam(r, "player")
	bio := r.URL.Query().Get("bio")
	format := utils.CleanText(r.URL.Query().Get("format"), true)
	year := utils.CleanText(r.URL.Query().Get("year"), true)
	vsteam := utils.CleanText(r.URL.Query().Get("vsteam"), true)

	if format == "" {
		utils.WebResponseJSONObject(w, r, http.StatusOK, []byte(`{"message":"missing required query param format"}`))
		return
	}

	if playerName == "all" {
		// do something
	}

	if bio != "true" {
		bio = ""
	}

	vsTeamID := 0
	if vsteam != "" {
		vsTeamID = data.GetTeamIDByTeamName(vsteam)
	}
	objAllPlayerStats := data.GetPlayerStats(playerName, format, year, vsTeamID)

	var playerFinal models.PlayerStatsExt
	if len(objAllPlayerStats) == 0 {
		final := utils.JSONMessageWrappedObj(http.StatusOK, playerFinal)
		utils.WebResponseJSONObject(w, r, http.StatusOK, final)
		return
	}

	// General Stats
	playerFinal.TeamID = int64(objAllPlayerStats[0].TeamID)
	playerFinal.PlayerID = int64(objAllPlayerStats[0].PlayerID)
	playerFinal.PlayerName = playerName

	// playerFinal.InningsID += pstat.InningsID
	// playerFinal.OutBowler += pstat.OutBowler
	// playerFinal.OutFielder += pstat.OutFielder
	// playerFinal.OutType += pstat.OutType
	playerFinal.SeasonID = year
	// playerFinal.SeasonType += pstat.SeasonType
	outTypeCounter := make(map[string]int)
	for _, pstat := range objAllPlayerStats {
		playerFinal.MatchesPlayed++
		// bind batting stats
		playerFinal.Batting.BallsFaced += pstat.BallsFaced
		// playerFinal.Batting.BattingOrder += pstat.BattingOrder
		playerFinal.Batting.DotBallsPlayed += pstat.DotBallsPlayed
		playerFinal.Batting.Doubles += pstat.Doubles
		playerFinal.Batting.FoursHit += pstat.FoursHit
		playerFinal.Batting.RunsScored += pstat.RunsScored
		playerFinal.Batting.Singles += pstat.Singles
		playerFinal.Batting.SixesHit += pstat.SixesHit
		playerFinal.Batting.Triples += pstat.Triples
		if pstat.IsBatted {
			playerFinal.Batting.IsBatted++
		}
		if pstat.RunsScored >= 100 {
			playerFinal.Batting.Hundreds++
		} else if pstat.RunsScored >= 50 {
			playerFinal.Batting.Fifties++
		}
		if pstat.OutType == "not out" {
			playerFinal.Batting.NotOuts++
		}

		if pstat.OutType != "" && pstat.OutType != "not out" {
			outTypeCounter[pstat.OutType]++
		}
		if pstat.RunsScored == 0 && pstat.OutType != "not out" && pstat.OutType != "" {
			playerFinal.Batting.Ducks++
		}

		if pstat.RunsScored > playerFinal.Batting.HighestScore {
			playerFinal.Batting.HighestScore = pstat.RunsScored
		}

		// bind bowling stats
		// playerFinal.Bowling.BowlingOrder += pstat.BowlingOrder
		playerFinal.Bowling.DotsBowled += pstat.DotsBowled
		playerFinal.Bowling.MaidenOver += pstat.MaidenOver
		playerFinal.Bowling.BallsBowled += pstat.BallsBowled
		playerFinal.Bowling.ExtrasConceded += pstat.ExtrasConceded
		playerFinal.Bowling.FoursConceded += pstat.FoursConceded
		playerFinal.Bowling.RunsConceded += pstat.RunsConceded
		playerFinal.Bowling.SixesConceded += pstat.SixesConceded
		playerFinal.Bowling.WicketsTaken += pstat.WicketsTaken
		if pstat.WicketsTaken >= 5 {
			playerFinal.Bowling.Fifers++
		}
		if playerFinal.Bowling.BestBowling != "" {
			bowlingFigures := strings.Split(playerFinal.Bowling.BestBowling, "/")
			wickets, runs := bowlingFigures[0], bowlingFigures[1]
			wicketsInt, _ := strconv.ParseInt(wickets, 10, 64)
			runsInt, _ := strconv.ParseInt(runs, 10, 64)
			if wicketsInt < int64(pstat.WicketsTaken) {
				playerFinal.Bowling.BestBowling = fmt.Sprint(pstat.WicketsTaken) + "/" + fmt.Sprint(pstat.RunsConceded)
			} else if wicketsInt == int64(pstat.WicketsTaken) && runsInt > int64(pstat.RunsConceded) {
				playerFinal.Bowling.BestBowling = fmt.Sprint(pstat.WicketsTaken) + "/" + fmt.Sprint(pstat.RunsConceded)
			}
		} else {
			playerFinal.Bowling.BestBowling = fmt.Sprint(pstat.WicketsTaken) + "/" + fmt.Sprint(pstat.RunsConceded)
		}

		// bind fieling stats
		playerFinal.Fielding.Catches += pstat.Catches
		playerFinal.Fielding.Stumpings += pstat.Stumpings
		playerFinal.Fielding.RunOut += pstat.RunOut
	}
	playerFinal.Batting.OutType = outTypeCounter
	if playerFinal.Bowling.BallsBowled > 0 {
		playerFinal.Bowling.OversBowled = fmt.Sprint(playerFinal.Bowling.BallsBowled/6) + "." + fmt.Sprint(playerFinal.Bowling.BallsBowled%6)
	}

	if playerFinal.Batting.IsBatted-playerFinal.Batting.NotOuts > 0 {
		playerFinal.Batting.Average = utils.Round((float64(playerFinal.Batting.RunsScored))/float64(playerFinal.Batting.IsBatted-playerFinal.Batting.NotOuts), 0.01, 2)
	}

	if playerFinal.Batting.BallsFaced > 0 {
		playerFinal.Batting.StrikeRate = utils.Round((float64(playerFinal.Batting.RunsScored)*100)/float64(playerFinal.Batting.BallsFaced), 0.01, 2)
	}

	if playerFinal.Bowling.WicketsTaken > 0 {
		playerFinal.Bowling.Average = utils.Round((float64(playerFinal.Bowling.RunsConceded))/float64(playerFinal.Bowling.WicketsTaken), 0.01, 2)
	}

	if playerFinal.Bowling.BallsBowled > 0 {
		playerFinal.Bowling.Economy = utils.Round((float64(playerFinal.Bowling.RunsConceded))/(float64(playerFinal.Bowling.BallsBowled)/6), 0.01, 2)
	}

	final := utils.JSONMessageWrappedObj(http.StatusOK, playerFinal)
	utils.WebResponseJSONObject(w, r, http.StatusOK, final)
}

func TeamStats(w http.ResponseWriter, r *http.Request) {
	teamName := chi.URLParam(r, "team")
	gender := utils.CleanText(r.URL.Query().Get("gender"), true)
	if gender == "" {
		gender = "male"
	}
	year := utils.CleanText(r.URL.Query().Get("year"), true)
	venue := utils.CleanText(r.URL.Query().Get("venue"), true)
	vsTeam := utils.CleanText(r.URL.Query().Get("vsteam"), true)

	teamID := data.GetTeamIDByTeamName(teamName)
	log.Printf("fetching team stats for TeamName: %s, TeamID: %d, Gender: %s, Year: %s, Venue: %s, VsTeam: %s", teamName, teamID, gender, year, venue, vsTeam)
	objAllTeamStats := data.GetTeamStats(teamID, gender, year, venue, vsTeam)
	teamStatistics := calculateTeamStats(objAllTeamStats, teamID)
	final := utils.JSONMessageWrappedObj(http.StatusOK, teamStatistics)
	utils.WebResponseJSONObject(w, r, http.StatusOK, final)
}

func TeamHeadToHead(w http.ResponseWriter, r *http.Request) {
	teamName := r.URL.Query().Get("team")

	log.Println("data requested for teams ", teamName)
	gender := utils.CleanText(r.URL.Query().Get("gender"), true)
	if gender == "" {
		gender = "male"
	}
	year := utils.CleanText(r.URL.Query().Get("year"), true)
	venue := utils.CleanText(r.URL.Query().Get("venue"), true)

	requestedTeams := strings.Split(teamName, ",")
	teamStatistics := make(map[string]models.DuranzTeamStats)
	for idx, team := range requestedTeams {
		if team == "" || idx > 5 {
			continue
		}
		log.Printf("fetching team stats for TeamName: %s, Gender: %s, Year: %s, Venue: %s", teamName, gender, year, venue)
		teamID := data.GetTeamIDByTeamName(utils.CleanText(team, false))
		objAllTeamStats := data.GetTeamStats(teamID, gender, year, venue, "")
		teamStatistics[utils.CleanText(team, false)] = calculateTeamStats(objAllTeamStats, teamID)
	}
	final := utils.JSONMessageWrappedObj(http.StatusOK, teamStatistics)
	utils.WebResponseJSONObject(w, r, http.StatusOK, final)
}

func PlayerList(w http.ResponseWriter, r *http.Request) {
	matchCount := utils.CleanText(r.URL.Query().Get("cnt"), true)
	if matchCount == "" {
		matchCount = "10"
	}

	playerList := data.GetPlayerList(matchCount)
	final := utils.JSONMessageWrappedObj(http.StatusOK, playerList)
	utils.WebResponseJSONObject(w, r, http.StatusOK, final)
}

func TeamList(w http.ResponseWriter, r *http.Request) {
	teamList := data.GetTeamList()
	final := utils.JSONMessageWrappedObj(http.StatusOK, teamList)
	utils.WebResponseJSONObject(w, r, http.StatusOK, final)
}

func MatchList(w http.ResponseWriter, r *http.Request) {
	team := utils.CleanText(r.URL.Query().Get("team"), true)
	vsTeam := utils.CleanText(r.URL.Query().Get("vsteam"), true)

	if team == "" && vsTeam == "" {
		utils.WebResponseJSONObject(w, r, http.StatusBadRequest, []byte(`{"message":"missing required query param team or vsteam"}`))
		return
	}

	teamID := data.GetTeamIDByTeamName(team)
	vsTeamID := data.GetTeamIDByTeamName(vsTeam)

	if teamID == 0 || vsTeamID == 0 {
		utils.WebResponseJSONObject(w, r, http.StatusBadRequest, []byte(`{"message":"not found team or vsteam"}`))
		return
	}

	matchList := data.GetMatchList(teamID, vsTeamID)
	final := utils.JSONMessageWrappedObj(http.StatusOK, matchList)
	utils.WebResponseJSONObject(w, r, http.StatusOK, final)
}
