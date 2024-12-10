package data

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/KunalDuran/duranz-stats/internal/models"
)

// InsertErrorLog : Insert All Error log record according to alert ID, file name
func InsertErrorLog(alertid, errormsg, fileName string) {
	sqlStr := `INSERT INTO errorlog (alert_id, error_msg, file_name) VALUES (?, ?, ?)`

	_, err := Db.Exec(sqlStr, alertid, errormsg, fileName)
	if err != nil {
		panic(err)
	}
}

var AllDuranzLeagues = map[string]int{
	"ODI":                   1,
	"Test":                  2,
	"T20":                   3,
	"ipl":                   4,
	"Indian Premier League": 4,
}

var GamePath = map[string]string{
	"odi":  "odis_json",
	"test": "tests_json",
	"t20":  "t20s_json",
	"ipl":  "ipl_json",
}

// cache simulation
var MappedTeams = map[string]string{}
var MappedVenues = map[string]string{}

func InsertMatchStats(matchID int, objMatchStats models.MatchStats) {

	sqlStr := `INSERT INTO match_stats(
	match_id ,	
	team_id, 
	fall_of_wickets, 
	extras, 
	score, 
	super_over,  
	wickets, 
	overs_played, 
	innings
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := Db.Exec(sqlStr,
		matchID,
		objMatchStats.TeamID,
		objMatchStats.FOW,
		objMatchStats.Extras,
		objMatchStats.Score,
		objMatchStats.SuperOver,
		objMatchStats.Wickets,
		objMatchStats.OversPlayed,
		objMatchStats.InningsID,
	)
	if err != nil {
		panic(err)
	}
}

func GetVenueID(venueName, city string) int {
	var venueID int
	sqlStr := `SELECT venue_id FROM venue WHERE venue = ? AND city = ?`

	err := Db.QueryRow(sqlStr, venueName, city).Scan(&venueID)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return venueID
}

func GetTeamID(teamName, team_type string) int {
	var teamID int
	sqlStr := `SELECT team_id FROM teams WHERE team_name = ? AND team_type = ?`

	err := Db.QueryRow(sqlStr, teamName, team_type).Scan(&teamID)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return teamID
}

func GetPlayerID(cricsheetID string) int {
	var playerID int
	sqlStr := `SELECT player_id FROM cricket_players WHERE cricsheet_id = ?`

	err := Db.QueryRow(sqlStr, cricsheetID).Scan(&playerID)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return playerID
}

func CheckTable(tableName, whereClause string) (int, error) {
	var countVal int
	// sqlCheck := `SELECT COUNT(1) FROM ? WHERE ?`
	sqlCheck := `SELECT COUNT(1) FROM ` + tableName + ` WHERE ` + whereClause
	// err := data.Db.QueryRow(sqlCheck, tableName, whereClause).Scan(&countVal)
	// fmt.Println(sqlCheck)
	err := Db.QueryRow(sqlCheck).Scan(&countVal)
	if err != nil {
		return countVal, err
	}

	return countVal, nil
}

func GetMatchID(cricsheetID string) int {
	var matchID int
	sqlStr := `SELECT match_id FROM cricket_matches WHERE cricsheet_file_name = ?`

	err := Db.QueryRow(sqlStr, cricsheetID).Scan(&matchID)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return matchID
}

func InsertMappingInfo(fileName string, mappingInfo models.MappingInfo) {
	sqlStr := `INSERT INTO file_mappings (file_name, league_id, teams, players, venue, matches, match_stats, player_stats) 
				VALUES (?, ?, ?, ?, ?, ?, ?, ?) 
				ON DUPLICATE KEY UPDATE file_name=values(file_name), teams=values(teams), players=values(players), venue=values(venue), matches=values(matches), match_stats=values(match_stats), player_stats = values(player_stats)`

	_, err := Db.Exec(sqlStr, fileName, mappingInfo.LeagueID, mappingInfo.Teams, mappingInfo.Players, mappingInfo.Venue, mappingInfo.Match, mappingInfo.MatchStats, mappingInfo.PlayerStats)
	if err != nil {
		panic(err)
	}
}

func DeleteAllTableData() {

	tables := []string{
		"errorlog",
		"teams",
		"venue",
		"cricket_players",
		"cricket_matches",
		"match_stats",
		"player_match_stats",
		"file_mappings",
	}

	for _, table := range tables {
		sqlStr := `DELETE FROM ` + table
		_, err := Db.Exec(sqlStr)
		if err != nil {
			panic(err)
		}

		sqlStr2 := `ALTER TABLE ` + table + ` AUTO_INCREMENT = 1`
		_, err = Db.Exec(sqlStr2)
		if err != nil {
			panic(err)
		}
	}
}

func GetMappingDetails() map[string]models.MappingInfo {

	var objMappingInfo = map[string]models.MappingInfo{}

	sqlStr := `SELECT file_name, league_id, teams, players, venue, matches, match_stats, player_stats FROM file_mappings`
	rows, err := Db.Query(sqlStr)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var mappingInfo models.MappingInfo
		var fileName string
		err := rows.Scan(
			&fileName,
			&mappingInfo.LeagueID,
			&mappingInfo.Teams,
			&mappingInfo.Players,
			&mappingInfo.Venue,
			&mappingInfo.Match,
			&mappingInfo.MatchStats,
			&mappingInfo.PlayerStats,
		)

		if err != nil {
			panic(err)
		}

		objMappingInfo[fileName] = mappingInfo
	}

	return objMappingInfo
}

func InsertPlayerStats(matchID, seasonID int, teamInningPlayerStats map[string]map[string]models.PlayerStats, allPlayerID map[string]int, innBatTeamMap map[int]int) {

	var valueStr []string
	valArgs := []interface{}{}
	sqlStr := `INSERT INTO player_match_stats(
	match_id ,
	season_id ,
	innings_id ,
	team_id ,
	player_id ,
	batting_order ,
	runs_scored ,
	balls_faced ,
	dot_balls_played ,
	singles ,
	doubles ,
	triples ,
	fours_hit ,
	sixes_hit ,
	out_type ,
	out_bowler ,
	out_fielder ,
	is_batted ,
	overs_bowled ,
	runs_conceded ,
	balls_bowled ,
	dots_bowled ,
	wickets_taken ,
	fours_conceded ,
	sixes_conceded ,
	extras_conceded ,
	maiden_over ,
	bowling_order,
	run_out ,
	catches ,
	stumpings
) VALUES `
	for teamID, allPlayerStats := range teamInningPlayerStats {
		for _, playerStats := range allPlayerStats {
			tempPlayerID := allPlayerID[playerStats.Name]
			if tempPlayerID == 0 {
				fmt.Println(playerStats.Name)
				// continue
			}
			teamIDint, _ := strconv.Atoi(teamID)
			innID := innBatTeamMap[teamIDint]
			playerStats.PlayerID = tempPlayerID
			valueStr = append(valueStr, `(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)

			valArgs = append(valArgs,
				matchID,
				seasonID,
				innID,
				teamID,
				playerStats.PlayerID,
				playerStats.BattingOrder,
				playerStats.RunsScored,
				playerStats.BallsPlayed,
				playerStats.DotsPlayed,
				playerStats.Singles,
				playerStats.Doubles,
				playerStats.Triples,
				playerStats.FoursHit,
				playerStats.SixesHit,
				playerStats.OutType,
				playerStats.OutBowler,
				playerStats.OutFielder,
				playerStats.IsBatted,

				playerStats.OversBowled,
				playerStats.RunsConceded,
				playerStats.BallsBowled,
				playerStats.DotsBowled,
				playerStats.WicketsTaken,
				playerStats.FoursConceded,
				playerStats.SixesConceded,
				playerStats.ExtrasConceded,
				playerStats.MaidenOvers,
				playerStats.BowlingOrder,

				playerStats.RunOuts,
				playerStats.Catches,
				playerStats.Stumpings,
			)

		}
	}

	sqlStr = sqlStr + strings.Join(valueStr, ",")
	_, err := Db.Exec(sqlStr, valArgs...)
	if err != nil {
		panic(err)
	}
}

func PseudoCacheLayer(teamType string) {
	// load Teams
	if teamType != "ipl" {
		teamType = "international"
	}

	sqlStr := `SELECT team_name FROM teams WHERE team_type = ?`

	rows, err := Db.Query(sqlStr, teamType)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	for rows.Next() {
		var teamName string
		err = rows.Scan(&teamName)
		if err != nil {
			panic(err)
		}
		MappedTeams[teamName] = teamType
	}

	// load Venues
	sqlStr = `SELECT venue, city FROM venue`

	rows, err = Db.Query(sqlStr)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	for rows.Next() {
		var venue, city string
		err = rows.Scan(&venue, &city)
		if err != nil {
			panic(err)
		}
		MappedVenues[venue] = city
	}
}
