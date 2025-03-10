package mapper

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/KunalDuran/duranz-stats/internal/data"
	"github.com/KunalDuran/duranz-stats/internal/models"
	"github.com/KunalDuran/duranz-stats/internal/utils"
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
		fmt.Println(err, "in file ", f_path)
		return matchData, err
	}
	return matchData, nil
}

func ScorecardMapper(match models.Match, fileName string) error {
	cricSheetID := strings.Replace(fileName, ".json", "", -1)
	sheetID, err := strconv.Atoi(cricSheetID)
	if err != nil {
		return err
	}
	matchID := data.GetMatchID(cricSheetID)

	scorecard := ProcessScoreCard(match)
	scorecard.MatchID = matchID

	scorecard.CricsheetID = sheetID

	if err := data.InsertScoreCard(scorecard); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func ProcessScoreCard(match models.Match) models.ScoreCard {
	var objScoreCard models.ScoreCard
	var AllInnings []models.Innings

	objScoreCard.Event = match.Info.Event.Name
	objScoreCard.MatchNumber = match.Info.Event.MatchNumber
	objScoreCard.Date = strings.Join(match.Info.Dates, ",")

	for _, inning := range match.Innings {
		// fmt.Println("Scorecard process started innings for : ", inning.Team)

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

	return objScoreCard
}

func ProcessPlayerStats(match models.Match, fileName string) {

	var hasErrors bool
	// extract year
	var seasonID int
	if len(match.Info.Dates) > 0 {
		matchDate, err := time.Parse("2006-01-02", match.Info.Dates[0])
		if err != nil {
			data.InsertErrorLog(data.DATETIME_ERROR, `Error in parsing date`+match.Info.Dates[0], fileName, err.Error())
			panic(err)
		}
		seasonID = matchDate.Year()
	}
	// get match id, else log error and continue
	cricSheetID := strings.Replace(fileName, ".json", "", -1)
	matchID := data.GetMatchID(cricSheetID)

	// get team id
	home := match.Info.Teams[0]
	away := match.Info.Teams[1]
	homeTeamID := data.GetTeamID(home, match.Info.TeamType)
	awayTeamID := data.GetTeamID(away, match.Info.TeamType)

	allPlayerID := map[string]int{}
	for player, cricID := range match.Info.Register.People {
		playerID := data.GetPlayerID(cricID)
		if playerID == 0 {
			// log error
			data.InsertErrorLog(data.PLAYER_NOT_FOUND, `player not found`+player, fileName, "")
			hasErrors = true
			fmt.Println("ID of this player is zero ", player)
			continue
		}
		allPlayerID[player] = playerID
	}

	if hasErrors {
		fmt.Println("STOP!! Process have serious errors to be fixed")
		return
	}

	var teamInningPlayerStats = map[string]map[string]models.PlayerStats{}
	var innBatTeamMap = map[int]int{}

	for inningID, inning := range match.Innings {
		if inning.SuperOvers {
			continue
		}
		inningID++
		// fmt.Println("player Stats process started innings for : ", inning.Team)

		var battingTeamID, bowlingTeamID int
		if inning.Team == home {
			battingTeamID = homeTeamID
			bowlingTeamID = awayTeamID
		} else if inning.Team == away {
			battingTeamID = awayTeamID
			bowlingTeamID = homeTeamID
		}

		// for inning id in test also
		innBatTeamMap[battingTeamID] = inningID

		var objBatsman = map[string]models.BattingStats{}
		var objBowler = map[string]models.BowlingStats{}
		var objFielder = map[string]models.FieldingStats{}

		var battingOrder, bowlingOrder int
		var overRunsBowler = map[int]int{}

		for _, over := range inning.Overs {

			overRunsBowler[over.Over] = 0
			for _, delivery := range over.Deliveries {
				// init batsman
				batsman, existBat := objBatsman[delivery.Batter]
				if !existBat {
					battingOrder++
					batsman.BattingOrder = battingOrder
					batsman.Name = delivery.Batter
					batsman.IsBatted = true
				}

				// init non-striker
				nonBatsman, existBat2 := objBatsman[delivery.NonStriker]

				if !existBat2 {
					battingOrder++
					nonBatsman.BattingOrder = battingOrder
					nonBatsman.Name = delivery.NonStriker
					nonBatsman.IsBatted = true
					objBatsman[delivery.NonStriker] = nonBatsman
				}

				// init bowler
				bowler, existBowl := objBowler[delivery.Bowler]
				if !existBowl {
					bowlingOrder++
					bowler.BowlingOrder = bowlingOrder
					bowler.Name = delivery.Bowler
				}

				// init fielder/WK and calculate
				if len(delivery.Wickets) > 0 {
					for _, wicket := range delivery.Wickets {
						for _, f := range wicket.Fielders {
							if f.Substitute && f.Name == "" {
								continue
							}
							fielder, existField := objFielder[f.Name]
							if !existField {
								fielder.Name = f.Name
							}
							if wicket.Kind == "caught" {
								fielder.Catches++
							} else if wicket.Kind == "run out" {
								fielder.RunOuts++
							} else if wicket.Kind == "stumped" {
								fielder.Stumpings++
							}
							objFielder[f.Name] = fielder
						}
						if wicket.Kind == "caught and bowled" {
							fielder := objFielder[bowler.Name]
							fielder.Name = bowler.Name
							fielder.Catches++
							objFielder[bowler.Name] = fielder
						}
					}
				}

				// ================ bowler calculations
				bowlerRuns := delivery.Runs.Batter

				// Calculate Extras for Bowler
				if delivery.Extras != (models.Extras{}) {
					if delivery.Extras.NoBall > 0 {
						bowler.Extras += delivery.Extras.NoBall
						bowlerRuns += delivery.Runs.Total
					} else if delivery.Extras.Wides > 0 {
						bowler.Extras += delivery.Extras.Wides
						bowlerRuns += delivery.Extras.Wides
					}
				}

				overRunsBowler[over.Over] += bowlerRuns
				bowler.Runs += bowlerRuns

				if delivery.Extras.Wides == 0 && delivery.Extras.NoBall == 0 {
					bowler.Balls += 1
				}

				if delivery.Runs.Batter == 0 {
					bowler.Dots++
				}

				// ================  batter calculations
				batsman.Runs += delivery.Runs.Batter
				if delivery.Extras.Wides == 0 {
					batsman.Balls += 1
					if delivery.Runs.Batter == 0 && delivery.Extras.NoBall == 0 &&
						delivery.Extras.Byes == 0 && delivery.Extras.LegByes == 0 {
						batsman.Dots++
					}
				}

				// ==================== common calculations

				// Check for Wicket
				for _, wicket := range delivery.Wickets {
					if wicket.Kind != "" && wicket.PlayerOut != "" {

						var fielderName string

						if wicket.Kind != "caught and bowled" && wicket.Kind != "bowled" && wicket.Kind != "lbw" && len(wicket.Fielders) == 0 {
							fielderName = "NA"
						}
						if len(wicket.Fielders) > 0 {
							fielderName = wicket.Fielders[0].Name

							if wicket.Fielders[0].Substitute && fielderName == "" {
								fielderName = "substitute"
							}
						}

						// bowler
						if wicket.Kind != "run out" {
							bowler.Wickets++
							batsman.Out = wicket.Kind
							batsman.OutBowler = allPlayerID[bowler.Name]
							batsman.OutFielder = allPlayerID[fielderName]
						} else if wicket.Kind == "run out" {
							if batsman.Name == wicket.PlayerOut {
								batsman.Out = wicket.Kind
								batsman.OutFielder = allPlayerID[fielderName]
							} else if nonBatsman.Name == wicket.PlayerOut {
								nonBatsman.Out = wicket.Kind
								nonBatsman.OutFielder = allPlayerID[fielderName]
							}
						}
					}
				}

				// 4s/6s hit and conceded
				if delivery.Runs.Batter == 4 {
					batsman.Fours++
					bowler.FoursConceded++
				} else if delivery.Runs.Batter == 6 {
					batsman.Sixes++
					bowler.SixesConceded++
				} else if delivery.Runs.Batter == 1 {
					batsman.Singles++
				} else if delivery.Runs.Batter == 2 {
					batsman.Doubles++
				} else if delivery.Runs.Batter == 3 {
					batsman.Triples++
				}

				// ======== BIND ALL STATS
				objBatsman[delivery.Batter] = batsman
				objBatsman[delivery.NonStriker] = nonBatsman
				objBowler[delivery.Bowler] = bowler
			}

			// check maiden over
			if val, ok := overRunsBowler[over.Over]; ok && val == 0 {
				if len(over.Deliveries) > 0 {
					bowler := objBowler[over.Deliveries[0].Bowler]
					bowler.Maiden++
					objBowler[over.Deliveries[0].Bowler] = bowler
				}
			}
		}

		for name, batter := range objBatsman {
			if batter.Out == "" {
				batter.Out = "not out"
			}
			objBatsman[name] = batter
		}

		for name, bowler := range objBowler {
			bowler.Overs = fmt.Sprint(bowler.Balls/6) + "." + fmt.Sprint(bowler.Balls%6)
			objBowler[name] = bowler
		}

		// bind batter stats for the batting team on key teamID**inning
		// fmt.Println(objBatsman)
		battingKey := strconv.Itoa(battingTeamID) // + "**" + strconv.Itoa(inningID)
		bowlingKey := strconv.Itoa(bowlingTeamID) // + "**" + strconv.Itoa(inningID)
		battingStats := teamInningPlayerStats[battingKey]
		bowlingStats := teamInningPlayerStats[bowlingKey]

		// teamInningPlayerStats[bowlingTeamID] =
		battingStats = BindBattingPlayerStats(battingStats, objBatsman)
		bowlingStats = BindBowlingPlayerStats(bowlingStats, objBowler, objFielder)

		teamInningPlayerStats[battingKey] = battingStats
		teamInningPlayerStats[bowlingKey] = bowlingStats

		if inningID == 2 || inningID == 4 || (len(match.Innings) == 3 && inningID == 3) {
			// fmt.Println(teamInningPlayerStats)
			if len(match.Innings) == 3 && inningID == 3 {
				innBatTeamMap[bowlingTeamID] = inningID
			}
			data.InsertPlayerStats(matchID, seasonID, teamInningPlayerStats, allPlayerID, innBatTeamMap)

			// after second inning empty the playerstat for next inning
			teamInningPlayerStats = map[string]map[string]models.PlayerStats{}
		}
	}
}

func BindBattingPlayerStats(objPlayers map[string]models.PlayerStats, objBatsman map[string]models.BattingStats) map[string]models.PlayerStats {

	if len(objPlayers) == 0 {
		objPlayers = map[string]models.PlayerStats{}
	}
	for batsmanName, batsman := range objBatsman {
		player := objPlayers[batsmanName]
		player.Name = batsman.Name
		player.BattingOrder = batsman.BattingOrder
		player.RunsScored = batsman.Runs
		player.BallsPlayed = batsman.Balls
		player.Singles = batsman.Singles
		player.Doubles = batsman.Doubles
		player.Triples = batsman.Triples
		player.FoursHit = batsman.Fours
		player.SixesHit = batsman.Sixes
		player.StrikeRate = batsman.StrikeRate
		player.OutType = batsman.Out
		player.OutBowler = batsman.OutBowler
		player.OutFielder = batsman.OutFielder
		player.DotsPlayed = batsman.Dots
		player.NotOut = batsman.NotOut
		player.IsBatted = batsman.IsBatted
		objPlayers[batsmanName] = player
	}
	return objPlayers
}

func BindBowlingPlayerStats(objPlayers map[string]models.PlayerStats, objBowler map[string]models.BowlingStats, objFielder map[string]models.FieldingStats) map[string]models.PlayerStats {
	if len(objPlayers) == 0 {
		objPlayers = map[string]models.PlayerStats{}
	}
	for bowlerName, bowler := range objBowler {
		player := objPlayers[bowlerName]
		player.Name = bowler.Name
		player.BowlingOrder = bowler.BowlingOrder
		player.OversBowled = bowler.Overs
		player.MaidenOvers = bowler.Maiden
		player.RunsConceded = bowler.Runs
		player.WicketsTaken = bowler.Wickets
		player.Economy = bowler.Economy
		player.BallsBowled = bowler.Balls
		player.DotsBowled = bowler.Dots
		player.FoursConceded = bowler.FoursConceded
		player.SixesConceded = bowler.SixesConceded
		player.ExtrasConceded = bowler.Extras
		objPlayers[bowlerName] = player
	}

	for fielderName, fielder := range objFielder {
		player := objPlayers[fielderName]
		player.Name = fielder.Name
		player.Catches = fielder.Catches
		player.RunOuts = fielder.RunOuts
		player.Stumpings = fielder.Stumpings
		objPlayers[fielderName] = player

	}

	return objPlayers
}

func ProcessMatchStats(match models.Match, fileName string) {

	var objMatchStats models.MatchStats
	cricSheetID := strings.Replace(fileName, ".json", "", -1)
	matchID := data.GetMatchID(cricSheetID)
	if matchID == 0 {
		data.InsertErrorLog(data.MATCH_NOT_FOUND, `matchID not found `+cricSheetID, fileName, "")
		return
	}

	// check for super over in a match
	for _, inning := range match.Innings {
		if inning.SuperOvers {
			objMatchStats.SuperOver = true
		}
	}

	for innID, inning := range match.Innings {
		innID++
		var fowArr []string
		var wicketCnt, runningScore, extras int

		if inning.SuperOvers {
			continue
		}
		for _, over := range inning.Overs {

			for _, delivery := range over.Deliveries {

				runningScore += delivery.Runs.Total

				// Check for Wicket
				for _, wicket := range delivery.Wickets {
					if wicket.Kind != "" && wicket.PlayerOut != "" {
						wicketCnt++
						fowStr := fmt.Sprint(wicketCnt, "-", runningScore, "(", wicket.PlayerOut, ")")
						fowArr = append(fowArr, fowStr)
					}
				}

				// Calculate Extras
				if delivery.Extras != (models.Extras{}) {
					if delivery.Extras.Byes > 0 {
						extras += delivery.Runs.Extras
					} else if delivery.Extras.LegByes > 0 {
						extras += delivery.Runs.Extras
					} else if delivery.Extras.NoBall > 0 {
						extras += delivery.Runs.Extras
					} else if delivery.Extras.Wides > 0 {
						extras += delivery.Runs.Extras
					}
				}
			}
		}
		objMatchStats.FOW = strings.Join(fowArr, " , ")
		objMatchStats.Score = runningScore
		objMatchStats.Extras = extras
		objMatchStats.Wickets = wicketCnt
		objMatchStats.InningsID = innID

		tempOvers := len(inning.Overs) - 1
		var lastoverBalls int
		for _, delivery := range inning.Overs[tempOvers].Deliveries {
			if delivery.Extras.Wides == 0 && delivery.Extras.NoBall == 0 {
				lastoverBalls++
			}
		}
		if lastoverBalls == 6 {
			tempOvers++
			lastoverBalls = 0
		}
		totalOvers := strconv.Itoa(tempOvers)
		if lastoverBalls > 0 {
			totalOvers = strconv.Itoa(tempOvers) + "." + strconv.Itoa(lastoverBalls)
		}
		objMatchStats.OversPlayed = totalOvers
		tempTeamID := data.GetTeamID(inning.Team, match.Info.TeamType)
		if tempTeamID == 0 {
			fmt.Println("team id not found")
		}
		objMatchStats.TeamID = tempTeamID

		data.InsertMatchStats(matchID, objMatchStats)
	}
}

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

func PlayerMapper(players, alternateNames map[string]string) {
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
				DisplayName: alternateNames[playerID],
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
			data.InsertErrorLog(data.DATABASE_ERROR, `More than 2 teams`, fileName, "")
		}
		if venueID == 0 {
			data.InsertErrorLog(data.DATABASE_ERROR, `Venue not found `+match.Info.Venue, fileName, "")
		}
		if leagueID == 0 {
			data.InsertErrorLog(data.LEAGUE_NOT_FOUND, `League not found `+match.Info.MatchType, fileName, "")
		}
		if startDate == "" {
			data.InsertErrorLog(data.DATETIME_ERROR, `Start date not found`, fileName, "")
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
		data.InsertErrorLog(data.DATABASE_ERROR, `Toss winner not found`, fileName, "")
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
		LeagueID:          leagueID,
		SeasonID:          seasonID,
		HomeTeamID:        homeTeamID,
		AwayTeamID:        awayTeamID,
		HomeTeamName:      home,
		AwayTeamName:      away,
		VenueID:           venueID,
		MatchDate:         matchDate,
		MatchDateMulti:    strings.Join(match.Info.Dates, ";"),
		CricsheetFileName: fileName,
		Result:            resultStr,
		TossWinner:        tossWinner,
		TossDecision:      tossDecision,
		WinningTeam:       winningTeam,
		Gender:            match.Info.Gender,
		MatchRefrees:      matchReferees,
		ReserveUmpires:    reserveUmpires,
		TVUmpires:         tvUmpires,
		Umpires:           umpires,
		EventName:         match.Info.Event.Name,
		MatchNumber:       match.Info.Event.MatchNumber,
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
