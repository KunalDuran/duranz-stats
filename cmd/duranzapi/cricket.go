package main

import (
	"encoding/json"
	"fmt"
	"math"
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

func GetScoreCard(w http.ResponseWriter, r *http.Request) {
	jsonID := chi.URLParam(r, "file")
	match, err := GetCricsheetData(DATASET_BASE + jsonID + `.json`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var objScoreCard models.ScoreCard

	var AllInnings []models.Innings
	for _, inning := range match.Innings {
		fmt.Println("Scorecard process started innings for : ", inning.Team)

		var objInning models.Innings
		objInning.InningDetail = inning.Team

		partnerships := make(map[int]models.Partnership)

		var objExtra models.Extras
		var objBatsman = map[string]models.BattingStats{}
		var objBowler = map[string]models.BowlingStats{}
		var batsmanCount, bowlerCount, runningScore, wicketCnt int
		var fowArr []string
		var overRuns = map[int]int{}
		for _, over := range inning.Overs {
			overRuns[over.Over] = 0
			for _, delivery := range over.Deliveries {

				// Score Calculations
				runningScore += delivery.Runs.Total
				overRuns[over.Over] += delivery.Runs.Total

				// Batsman Init
				if _, exist := objBatsman[delivery.Batter]; !exist {
					batsmanCount++
					var tempBat models.BattingStats
					tempBat.BattingOrder = batsmanCount
					tempBat.Name = delivery.Batter
					objBatsman[delivery.Batter] = tempBat

				}
				batsman := objBatsman[delivery.Batter]
				batsman.Runs += delivery.Runs.Batter
				batsman.Balls += 1

				//partnership
				if _, ok := partnerships[wicketCnt]; !ok {
					partnerships[wicketCnt] = models.Partnership{
						Batsman1: delivery.Batter,
						Batsman2: delivery.NonStriker,
					}
				}
				partner := partnerships[wicketCnt]
				partner.Runs += delivery.Runs.Total
				partnerships[wicketCnt] = partner

				// Bowler Init
				if _, exist := objBowler[delivery.Bowler]; !exist {
					bowlerCount++
					var tempBowler models.BowlingStats
					tempBowler.BowlingOrder = bowlerCount
					tempBowler.Name = delivery.Bowler
					objBowler[delivery.Bowler] = tempBowler
				}
				bowler := objBowler[delivery.Bowler]
				bowler.Runs += delivery.Runs.Batter
				bowler.Balls += 1

				if delivery.Runs.Batter == 4 {
					batsman.Fours++
				} else if delivery.Runs.Batter == 6 {
					batsman.Sixes++
				}

				// Calculate Extras
				if delivery.Extras != (models.Extras{}) {
					if delivery.Extras.Byes > 0 {
						objExtra.Byes += delivery.Extras.Byes
						overRuns[over.Over] -= delivery.Extras.Byes
					} else if delivery.Extras.LegByes > 0 {
						objExtra.LegByes += delivery.Extras.LegByes
						overRuns[over.Over] -= delivery.Extras.LegByes
					} else if delivery.Extras.NoBall > 0 {
						// remove ball count if No Ball
						batsman.Balls -= 1
						bowler.Balls -= 1
						bowler.Runs += delivery.Extras.NoBall
						objExtra.NoBall += delivery.Extras.NoBall
					} else if delivery.Extras.Wides > 0 {
						// remove ball count if Wide Ball
						batsman.Balls -= 1
						bowler.Balls -= 1
						bowler.Runs += delivery.Extras.Wides
						objExtra.Wides += delivery.Extras.Wides
					}
				}

				// Check for Wicket
				for _, wicket := range delivery.Wickets {
					if wicket.Kind != "" && wicket.PlayerOut != "" {
						batsman.Out = wicket.Kind
						wicketCnt++
						fowStr := fmt.Sprint(wicketCnt, "-", runningScore, "(", wicket.PlayerOut, ")")
						fowArr = append(fowArr, fowStr)

						// bowler
						if wicket.Kind != "run out" {
							bowler.Wickets++
						}
					}
				}

				// bind all info and calculations
				objBatsman[delivery.Batter] = batsman
				objBowler[delivery.Bowler] = bowler
			}

			// check maiden over
			if val, ok := overRuns[over.Over]; ok && val == 0 {
				if len(over.Deliveries) > 0 {
					bowler := objBowler[over.Deliveries[0].Bowler]
					bowler.Maiden++
					objBowler[over.Deliveries[0].Bowler] = bowler
				}
			}
		}

		var allBatsman []models.BattingStats
		for _, batter := range objBatsman {
			if batter.Balls > 0 {
				batter.StrikeRate = utils.Round((float64(batter.Runs)*100)/float64(batter.Balls), 0.01, 2)
			}
			if batter.Out == "" {
				batter.Out = "not out"
			}
			allBatsman = append(allBatsman, batter)
		}

		var allBowler []models.BowlingStats
		for _, bowler := range objBowler {
			if bowler.Balls > 0 {
				bowler.Economy = utils.Round(float64(bowler.Runs)/(float64(bowler.Balls)/float64(6)), 0.01, 2)
			}
			bowler.Overs = fmt.Sprint(bowler.Balls/6) + "." + fmt.Sprint(bowler.Balls%6)
			allBowler = append(allBowler, bowler)
		}

		objExtra.Total = objExtra.Byes + objExtra.LegByes + objExtra.Wides + objExtra.NoBall
		objInning.Extras = objExtra
		objInning.Batting = allBatsman
		objInning.FallOfWickets = strings.Join(fowArr, " , ")
		objInning.Bowling = allBowler

		objInning.Partnerships = partnerships
		fmt.Printf("%v+", partnerships)
		objInning.OverByOver = overRuns
		AllInnings = append(AllInnings, objInning)
	}

	resultStr := match.Info.Outcome.Winner + " Won by "
	if match.Info.Outcome.By.Runs > 0 {
		resultStr += strconv.Itoa(match.Info.Outcome.By.Runs) + " Runs"
	} else if match.Info.Outcome.By.Wickets > 0 {
		resultStr += strconv.Itoa(match.Info.Outcome.By.Wickets) + " Wickets"
	}
	objScoreCard.Result = resultStr
	objScoreCard.Innings = AllInnings

	final := utils.JSONMessageWrappedObj(http.StatusOK, objScoreCard)
	utils.WebResponseJSONObject(w, r, http.StatusOK, final)
}

func PlayerStatsAPI(w http.ResponseWriter, r *http.Request) {
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
		// playerFinal.Batting.IsBatted += pstat.IsBatted
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

func TeamStatsAPI(w http.ResponseWriter, r *http.Request) {
	teamName := chi.URLParam(r, "team")
	gender := utils.CleanText(r.URL.Query().Get("gender"), true)
	if gender == "" {
		gender = "male"
	}
	year := utils.CleanText(r.URL.Query().Get("year"), true)
	venue := utils.CleanText(r.URL.Query().Get("venue"), true)
	vsTeam := utils.CleanText(r.URL.Query().Get("vsteam"), true)

	teamID := data.GetTeamIDByTeamName(teamName)
	objAllTeamStats := data.GetTeamStats(teamID, gender, year, venue, vsTeam)
	var teamStatistics models.DuranzTeamStats

	for _, objTeam := range objAllTeamStats {

		if objTeam.TossWinner == teamID {
			teamStatistics.TossWin++
		}

		if objTeam.WinningTeam == teamID {
			teamStatistics.MatchWin++

			// count if team won while batting first or chasing first
			if objTeam.TossWinner == teamID { // team won the toss
				if objTeam.TossDecision == "bat" {
					teamStatistics.BatFirstWin++
				} else {
					teamStatistics.ChasingWin++
				}
			} else {
				if objTeam.TossDecision == "bat" { // other team won the toss
					teamStatistics.ChasingWin++
				} else {
					teamStatistics.BatFirstWin++
				}
			}
		}
	}

	teamStatistics.TotalMatches = len(objAllTeamStats)
	if teamStatistics.TotalMatches > 0 {
		teamStatistics.MatchWinPercent = utils.Round(float64(teamStatistics.MatchWin)/float64(teamStatistics.TotalMatches), 0.01, 2) * 100
		teamStatistics.TossWinPercent = utils.Round(float64(teamStatistics.TossWin)/float64(teamStatistics.TotalMatches), 0.01, 2) * 100
	}

	if teamStatistics.MatchWin > 0 {
		teamStatistics.ChasingWinPer = math.Round((float64(teamStatistics.ChasingWin)/float64(teamStatistics.MatchWin))*10000) / 100
		teamStatistics.BatFirstWinPer = math.Round((float64(teamStatistics.BatFirstWin)/float64(teamStatistics.MatchWin))*10000) / 100
	}
	//AvgScore /Inn
	//Highest Score
	//Lowest Score

	final := utils.JSONMessageWrappedObj(http.StatusOK, teamStatistics)
	utils.WebResponseJSONObject(w, r, http.StatusOK, final)
}

func BatsmanVSBowlerAPI(w http.ResponseWriter, r *http.Request) {
	//
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
	// year := utils.CleanText(r.URL.Query().Get("year"), true)
	// format := utils.CleanText(r.URL.Query().Get("format"), true)
	// team := utils.CleanText(r.URL.Query().Get("team"), true)
	// matchList := data.GetMatchList(year, format, team)
	// final := utils.JSONMessageWrappedObj(http.StatusOK, matchList)
	// utils.WebResponseJSONObject(w, r, http.StatusOK, final)
}

func PlayerVSPlayer(w http.ResponseWriter, r *http.Request) {
	// playerName := chi.URLParam(r, "player")

}
