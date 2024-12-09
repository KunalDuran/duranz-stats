package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/KunalDuran/duranz-stats/internal/models"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// InitDB initialises the database pools with
func InitDB(host, port, user, password string) (*sql.DB, error) {
	SportsDb, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/duranz")

	if err != nil {
		return nil, err
	}
	if err = Db.Ping(); err != nil {
		return nil, err
	}
	Db = SportsDb
	return SportsDb, nil
}

// RequestAPIData - Calls a (API) URL and return the data from the request.
func RequestAPIData(url string, headers map[string]string) ([]byte, int, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 500, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := http.Client{}
	client.Timeout = time.Duration(120 * time.Second)
	resp, err := client.Do(req)
	if err != nil {
		return nil, 500, err
	}
	statusCode := resp.StatusCode

	// read body
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, statusCode, err
	}
	return body, statusCode, nil
}

func GetPlayerDetails() []models.PlayerDetailsInt {
	var objAllPlayer = []models.PlayerDetailsInt{}

	sqlStr := `SELECT * FROM duranz_cricket_players WHERE 
		(display_name IS NULL or display_name='') OR 
		(first_name IS NULL or first_name='') OR
		(last_name IS NULL or last_name='') OR
		(dob IS NULL)`

	rows, err := Db.Query(sqlStr)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var objPlayer models.PlayerDetailsInt
		err = rows.Scan(
			&objPlayer.PlayerID,
			&objPlayer.PlayerName,
			&objPlayer.DisplayName,
			&objPlayer.FirstName,
			&objPlayer.LastName,
			&objPlayer.ShortName,
			&objPlayer.UniqueShortName,
			&objPlayer.DOB,
			&objPlayer.BattingStyle,
			&objPlayer.BowlingStyle,
			&objPlayer.IsOverseas,
			&objPlayer.CricSheetID,
			&objPlayer.DateAdded,
			&objPlayer.Status,
		)
		if err != nil {
			panic(err)
		}
		objAllPlayer = append(objAllPlayer, objPlayer)
	}
	return objAllPlayer
}

func GetTeamDetails() []models.TeamDetailsInt {
	var objAllTeams = []models.TeamDetailsInt{}

	sqlStr := `SELECT * FROM duranz_teams`

	rows, err := Db.Query(sqlStr)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var objTeam models.TeamDetailsInt
		err = rows.Scan(
			&objTeam.TeamID,
			&objTeam.TeamName,
			&objTeam.TeamType,
			&objTeam.FilterName,
			&objTeam.ABBR,
			&objTeam.TeamColor,
			&objTeam.Icon,
			&objTeam.URL,
			&objTeam.Jersey,
			&objTeam.Flag,
			&objTeam.Status,
			&objTeam.DateAdded,
		)
		if err != nil {
			panic(err)
		}
		objAllTeams = append(objAllTeams, objTeam)
	}
	return objAllTeams
}

func GetVenueDetails() []models.VenueDetailsInt {
	var objAllVenue = []models.VenueDetailsInt{}

	sqlStr := `SELECT * FROM duranz_venue WHERE 
				(city IS NULL OR city='') OR (state IS NULL OR state='') OR (country IS NULL OR country='')`

	rows, err := Db.Query(sqlStr)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var objVenue models.VenueDetailsInt
		err = rows.Scan(
			&objVenue.VenueID,
			&objVenue.Venue,
			&objVenue.FilterName,
			&objVenue.FriendlyName,
			&objVenue.City,
			&objVenue.Country,
			&objVenue.State,
			&objVenue.StateABBR,
			&objVenue.OfficialTeam,
			&objVenue.Capacity,
			&objVenue.Dimension,
			&objVenue.Opened,
			&objVenue.Description,
			&objVenue.ShortName,
			&objVenue.TimeZone,
			&objVenue.Weather,
			&objVenue.PitchType,
			&objVenue.DateAdded,
			&objVenue.Status,
		)
		if err != nil {
			panic(err)
		}
		objAllVenue = append(objAllVenue, objVenue)
	}
	return objAllVenue
}

func InsertPlayerDetails(objPlayer models.PlayerDetailsExt) {

	sqlStr := `INSERT INTO duranz_cricket_players(player_id,player_name,display_name,first_name,last_name,
		short_name,unique_short_name,dob,batting_style_1_id,bowling_style_1_id,is_overseas,cricsheet_id,
		date_added,status)
		VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		player_id=VALUES(player_id),player_name=VALUES(player_name),display_name=VALUES(display_name),
		first_name=VALUES(first_name),last_name=VALUES(last_name),short_name=VALUES(short_name),
		unique_short_name=VALUES(unique_short_name),dob=VALUES(dob),batting_style_1_id=VALUES(batting_style_1_id),
		bowling_style_1_id=VALUES(bowling_style_1_id),is_overseas=VALUES(is_overseas),
		cricsheet_id=VALUES(cricsheet_id),date_added=VALUES(date_added),status=VALUES(status)`

	_, err := Db.Exec(
		sqlStr,
		objPlayer.PlayerID,
		objPlayer.PlayerName,
		objPlayer.DisplayName,
		objPlayer.FirstName,
		objPlayer.LastName,
		objPlayer.ShortName,
		objPlayer.UniqueShortName,
		objPlayer.DOB,
		objPlayer.BattingStyle,
		objPlayer.BowlingStyle,
		objPlayer.IsOverseas,
		objPlayer.CricSheetID,
		objPlayer.DateAdded,
		objPlayer.Status,
	)
	if err != nil {
		panic(err)
	}
}

func InsertTeamDetails(objTeam models.TeamDetailsExt) {

	sqlStr := `INSERT INTO duranz_teams(
		team_id,team_name,team_type,filtername,abbreviation,team_color,icon,url,jersey,flag,status,dateadded) 
		VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) 
		ON DUPLICATE KEY UPDATE 
		team_id=VALUES(team_id),team_name=VALUES(team_name),team_type=VALUES(team_type),
		filtername=VALUES(filtername),abbreviation=VALUES(abbreviation),team_color=VALUES(team_color),
		icon=VALUES(icon),url=VALUES(url),jersey=VALUES(jersey),flag=VALUES(flag),
		status=VALUES(status),dateadded=VALUES(dateadded)`

	_, err := Db.Exec(
		sqlStr,
		objTeam.TeamID,
		objTeam.TeamName,
		objTeam.TeamType,
		objTeam.FilterName,
		objTeam.ABBR,
		objTeam.TeamColor,
		objTeam.Icon,
		objTeam.URL,
		objTeam.Jersey,
		objTeam.Flag,
		objTeam.Status,
		objTeam.DateAdded,
	)
	if err != nil {
		panic(err)
	}
}

func InsertVenueDetails(objVenue models.VenueDetailsExt) {

	sqlStr := `INSERT INTO duranz_venue (
		venue_id,venue,filtername,friendlyname,city,country,state,state_abbr,official_team,capacity,
		dimensions,opened,description,shortname,timezone,weather,pitch_type,dateadded,status)
	 	VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE 	
		venue_id=VALUES(venue_id),venue=VALUES(venue),filtername=VALUES(filtername),
		friendlyname=VALUES(friendlyname),city=VALUES(city),country=VALUES(country),
		state=VALUES(state),state_abbr=VALUES(state_abbr),official_team=VALUES(official_team),
		capacity=VALUES(capacity),dimensions=VALUES(dimensions),opened=VALUES(opened),
		description=VALUES(description),shortname=VALUES(shortname),timezone=VALUES(timezone),
		weather=VALUES(weather),pitch_type=VALUES(pitch_type),dateadded=VALUES(dateadded),
		status=VALUES(status)
	`

	_, err := Db.Exec(
		sqlStr,
		objVenue.VenueID,
		objVenue.Venue,
		objVenue.FilterName,
		objVenue.FriendlyName,
		objVenue.City,
		objVenue.Country,
		objVenue.State,
		objVenue.StateABBR,
		objVenue.OfficialTeam,
		objVenue.Capacity,
		objVenue.Dimension,
		objVenue.Opened,
		objVenue.Description,
		objVenue.ShortName,
		objVenue.TimeZone,
		objVenue.Weather,
		objVenue.PitchType,
		objVenue.DateAdded,
		objVenue.Status,
	)
	if err != nil {
		panic(err)
	}
}

func GetPlayerStats(playerName, league, season string, vsTeam int) map[string][]models.PlayerStatsInt {
	objAllPlayerStats := map[string][]models.PlayerStatsInt{}
	var seasonCond, vsTeamCond string

	leagueID := models.AllDuranzLeagues[league]
	if season != "" {
		seasonID, err := strconv.ParseInt(season, 10, 64)
		if err == nil && seasonID > 1950 {
			seasonCond = " AND pms.season_id = " + season
		}
	}

	if vsTeam != 0 {
		vsTeamCond = fmt.Sprintf(" AND pms.team_id !=%d AND (matches.away_team_id=%d OR matches.home_team_id=%d)", vsTeam, vsTeam, vsTeam)
	}

	sqlStr := `SELECT player.player_name, 
	pms.match_id               ,
	pms.balls_bowled           ,
    pms.balls_faced            ,
    pms.batting_order          ,
    pms.bowling_order          ,
    pms.catches                ,
    pms.dot_balls_played       ,
    pms.dots_bowled            ,
    pms.doubles                ,
    pms.extras_conceded        ,
    pms.fours_conceded         ,
    pms.fours_hit              ,
    pms.innings_id             ,
    pms.is_batted              ,
    pms.last_update            ,
    pms.maiden_over            ,
    pms.out_bowler             ,
    pms.out_fielder            ,
    pms.out_type               ,
    pms.overs_bowled           ,
    pms.played_abandoned_matches ,
    pms.player_id              ,
    pms.run_out                ,
    pms.runs_conceded          ,
    pms.runs_scored            ,
    pms.season_id              ,
    pms.season_type            ,
    pms.singles                ,
    pms.sixes_conceded         ,
    pms.sixes_hit              ,
    pms.stumpings              ,
    pms.team_id                ,
    pms.triples                ,
    pms.wickets_taken  
	FROM duranz_cricket_players as player
	LEFT JOIN duranz_player_match_stats AS pms ON pms.player_id = player.player_id 
	LEFT JOIN duranz_cricket_matches matches ON matches.match_id = pms.match_id 
	WHERE player_name = ? AND league_id= ? ` + seasonCond + vsTeamCond

	rows, err := Db.Query(sqlStr, playerName, leagueID)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var objPlayerStats models.PlayerStatsInt
		var playerDisplayName string
		err = rows.Scan(
			&playerDisplayName,
			&objPlayerStats.MatchID,
			&objPlayerStats.BallsBowled,
			&objPlayerStats.BallsFaced,
			&objPlayerStats.BattingOrder,
			&objPlayerStats.BowlingOrder,
			&objPlayerStats.Catches,
			&objPlayerStats.DotBallsPlayed,
			&objPlayerStats.DotsBowled,
			&objPlayerStats.Doubles,
			&objPlayerStats.ExtrasConceded,
			&objPlayerStats.FoursConceded,
			&objPlayerStats.FoursHit,
			&objPlayerStats.InningsID,
			&objPlayerStats.IsBatted,
			&objPlayerStats.LastUpdate,
			&objPlayerStats.MaidenOver,
			&objPlayerStats.OutBowler,
			&objPlayerStats.OutFielder,
			&objPlayerStats.OutType,
			&objPlayerStats.OversBowled,
			&objPlayerStats.PlayedAbandonedMatches,
			&objPlayerStats.PlayerID,
			&objPlayerStats.RunOut,
			&objPlayerStats.RunsConceded,
			&objPlayerStats.RunsScored,
			&objPlayerStats.SeasonID,
			&objPlayerStats.SeasonType,
			&objPlayerStats.Singles,
			&objPlayerStats.SixesConceded,
			&objPlayerStats.SixesHit,
			&objPlayerStats.Stumpings,
			&objPlayerStats.TeamID,
			&objPlayerStats.Triples,
			&objPlayerStats.WicketsTaken,
		)
		if err != nil {
			panic(err)
		}
		objAllPlayerStats[playerDisplayName] = append(objAllPlayerStats[playerDisplayName], objPlayerStats)
	}
	return objAllPlayerStats
}

func GetTeamStats(teamID int, gender, season string) []models.DuranzMatchStats {
	var objAllTeamStats []models.DuranzMatchStats

	sqlStr := `SELECT 
	match_id,
	league_id,
	gender,
	season_id,
	home_team_id,
	away_team_id,
	home_team_name,
	away_team_name,
	venue_id,
	result,
	man_of_the_match,
	toss_winner,
	toss_decision,
	winning_team,
	cricsheet_file_name,
	match_date,
	match_date_multi,
	match_time,
	is_reschedule,
	is_abandoned,
	is_neutral,
	match_refrees,
	reserve_umpires,
	tv_umpires,
	umpires,
	date_added,
	last_update,
	match_end_time,
	status
	FROM duranz_cricket_matches
	WHERE (home_team_id = ? OR away_team_id = ?) AND gender = ?`

	rows, err := Db.Query(sqlStr, teamID, teamID, gender)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var objTeamStats models.DuranzMatchStats
		err = rows.Scan(
			&objTeamStats.MatchID,
			&objTeamStats.LeagueID,
			&objTeamStats.Gender,
			&objTeamStats.SeasonID,
			&objTeamStats.HomeTeamID,
			&objTeamStats.AwayTeamID,
			&objTeamStats.HomeTeamName,
			&objTeamStats.AwayTeamName,
			&objTeamStats.VenueID,
			&objTeamStats.Result,
			&objTeamStats.ManOfTheMatch,
			&objTeamStats.TossWinner,
			&objTeamStats.TossDecision,
			&objTeamStats.WinningTeam,
			&objTeamStats.CricsheetFileName,
			&objTeamStats.MatchDate,
			&objTeamStats.MatchDateMulti,
			&objTeamStats.MatchTime,
			&objTeamStats.IsReschedule,
			&objTeamStats.IsAbandoned,
			&objTeamStats.IsNeutral,
			&objTeamStats.MatchRefrees,
			&objTeamStats.ReserveUmpires,
			&objTeamStats.TvUmpires,
			&objTeamStats.Umpires,
			&objTeamStats.DateAdded,
			&objTeamStats.LastUpdate,
			&objTeamStats.MatchEndTime,
			&objTeamStats.Status)
		if err != nil {
			panic(err)
		}
		objAllTeamStats = append(objAllTeamStats, objTeamStats)
	}
	return objAllTeamStats
}

func GetTeamIDByTeamName(teamName string) int {

	sqlStr := `SELECT team_id FROM duranz_teams WHERE team_name = ?`
	var teamID int
	row := Db.QueryRow(sqlStr, teamName)
	err := row.Scan(&teamID)
	if err != nil {
		panic(err)
	}
	return teamID
}

func GetPlayerIDByPlayerName(playerName string) int {
	sqlStr := `SELECT player_id FROM duranz_cricket_players WHERE player_name = ?`
	var playerID int
	row := Db.QueryRow(sqlStr, playerName)
	err := row.Scan(&playerID)
	if err != nil {
		panic(err)
	}
	return playerID
}

func GetPlayerList(matchCount string) []string {
	var playerList []string
	cachekey := "player-list"

	cacheData, err := CacheGet(cachekey)
	if err != nil {
		cacheData = nil
	}

	if cacheData != nil {
		err := json.Unmarshal(cacheData, &playerList)
		if err != nil {
			cacheData = nil
		}
	}

	if cacheData != nil {
		return playerList
	}

	sqlStr := `SELECT player_name FROM duranz_cricket_players WHERE player_id IN 
	(SELECT player_id FROM duranz_player_match_stats GROUP BY player_id  HAVING COUNT(player_id) > ?)`
	rows, err := Db.Query(sqlStr, matchCount)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var playerName string
		err = rows.Scan(&playerName)
		if err != nil {
			panic(err)
		}
		playerList = append(playerList, playerName)
	}

	result, err := json.Marshal(playerList)
	err = CacheSetExp(cachekey, result, 43200)
	if err != nil {
		panic(err)
	}

	return playerList
}

func GetTeamList() []string {
	var teamList []string
	cachekey := "team-list"

	cacheData, err := CacheGet(cachekey)
	if err != nil {
		cacheData = nil
	}

	if cacheData != nil {
		err := json.Unmarshal(cacheData, &teamList)
		if err != nil {
			cacheData = nil
		}
	}

	if cacheData != nil {
		return teamList
	}

	sqlStr := `SELECT team_name FROM duranz_teams`
	rows, err := Db.Query(sqlStr)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var teamName string
		err = rows.Scan(&teamName)
		if err != nil {
			panic(err)
		}
		teamList = append(teamList, teamName)
	}

	result, err := json.Marshal(teamList)
	err = CacheSetExp(cachekey, result, 43200)
	if err != nil {
		panic(err)
	}

	return teamList
}
