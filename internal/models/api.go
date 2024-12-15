package models

import "database/sql"

var AllDuranzLeagues = map[string]int{
	"odi":                   1,
	"test":                  2,
	"t20":                   3,
	"ipl":                   4,
	"indian premier league": 4,
}

type ScoreCard struct {
	Innings []Innings `json:"innings"`
	Result  string    `json:"result"`
}

type PlayerStatsExt struct {
	MatchID                int64  `json:"match_id"`
	InningsID              string `json:"innings_id"`
	SeasonID               string `json:"season_id"`
	SeasonType             string `json:"season_type"`
	PlayerID               int64  `json:"player_id"`
	PlayerName             string `json:"player_name"`
	TeamID                 int64  `json:"team_id"`
	LastUpdate             string `json:"last_update"`
	PlayedAbandonedMatches int64  `json:"played_abandoned_matches"`

	Batting  PlayerBattingStats  `json:"batting"`
	Bowling  PlayerBowlingStats  `json:"bowling"`
	Fielding PlayerFieldingStats `json:"fielding"`
}

type PlayerBattingStats struct {
	BallsFaced     int     `json:"balls_faced"`
	BattingOrder   int     `json:"batting_order"`
	DotBallsPlayed int     `json:"dot_balls_played"`
	Doubles        int     `json:"doubles"`
	FoursHit       int     `json:"fours_hit"`
	IsBatted       int     `json:"is_batted"`
	OutBowler      string  `json:"out_bowler"`
	OutFielder     string  `json:"out_fielder"`
	OutType        string  `json:"out_type"`
	RunsScored     int     `json:"runs_scored"`
	Singles        int     `json:"singles"`
	SixesHit       int     `json:"sixes_hit"`
	Triples        int     `json:"triples"`
	Fifties        int     `json:"fifties"`
	Hundreds       int     `json:"hundreds"`
	Average        float64 `json:"average"`
	HighestScore   int     `json:"highest_score"`
	StrikeRate     float64 `json:"strike_rate"`
	NotOuts        int     `json:"not_outs"`
	Ducks          int     `json:"ducks"`
}

type PlayerBowlingStats struct {
	BowlingOrder   int     `json:"bowling_order"`
	BallsBowled    int     `json:"balls_bowled"`
	DotsBowled     int     `json:"dots_bowled"`
	ExtrasConceded int     `json:"extras_conceded"`
	FoursConceded  int     `json:"fours_conceded"`
	MaidenOver     int     `json:"maiden_over"`
	OversBowled    string  `json:"overs_bowled"`
	RunsConceded   int     `json:"runs_conceded"`
	SixesConceded  int     `json:"sixes_conceded"`
	WicketsTaken   int     `json:"wickets_taken"`
	Economy        float64 `json:"economy"`
	Average        float64 `json:"average"`
	Fifers         float64 `json:"fifers"`
	BestBowling    string  `json:"best_bowling"`
}

type PlayerFieldingStats struct {
	RunOut    int `json:"run_out"`
	Stumpings int `json:"stumpings"`
	Catches   int `json:"catches"`
}

type DuranzMatchStats struct {
	MatchID           sql.NullString `json:"match_id"`
	LeagueID          sql.NullString `json:"league_id"`
	Gender            sql.NullString `json:"gender"`
	SeasonID          sql.NullString `json:"season_id"`
	HomeTeamID        sql.NullString `json:"home_team_id"`
	AwayTeamID        sql.NullString `json:"away_team_id"`
	HomeTeamName      sql.NullString `json:"home_team_name"`
	AwayTeamName      sql.NullString `json:"away_team_name"`
	VenueID           sql.NullString `json:"venue_id"`
	Result            sql.NullString `json:"result"`
	ManOfTheMatch     sql.NullString `json:"man_of_the_match"`
	TossWinner        sql.NullInt64  `json:"toss_winner"`
	TossDecision      sql.NullString `json:"toss_decision"`
	WinningTeam       sql.NullInt64  `json:"winning_team"`
	CricsheetFileName sql.NullString `json:"cricsheet_file_name"`
	MatchDate         sql.NullString `json:"match_date"`
	MatchDateMulti    sql.NullString `json:"match_date_multi"`
	MatchTime         sql.NullString `json:"match_time"`
	IsReschedule      sql.NullString `json:"is_reschedule"`
	IsAbandoned       sql.NullString `json:"is_abandoned"`
	IsNeutral         sql.NullString `json:"is_neutral"`
	MatchRefrees      sql.NullString `json:"match_refrees"`
	ReserveUmpires    sql.NullString `json:"reserve_umpires"`
	TvUmpires         sql.NullString `json:"tv_umpires"`
	Umpires           sql.NullString `json:"umpires"`
	DateAdded         sql.NullString `json:"date_added"`
	LastUpdate        sql.NullString `json:"last_update"`
	MatchEndTime      sql.NullString `json:"match_end_time"`
	Status            sql.NullString `json:"status"`
}

type DuranzTeamStats struct {
	TotalMatches    int     `json:"total_matches"`
	MatchWin        int     `json:"match_win"`
	MatchWinPercent float64 `json:"win_percent"`
	BatFirstWin     int     `json:"bat_first_win"`
	ChasingWin      int     `json:"chasing_win"`
	BatFirstWinPer  float64 `json:"bat_first_win_per"`
	ChasingWinPer   float64 `json:"chasing_win_per"`
	AvgScoreInn     float64 `json:"avg_score_inn"`
	HighestScore    int     `json:"highest_score"`
	LowestScore     int     `json:"lowest_score"`
	TossWin         float64 `json:"toss_win"`
	TossWinPercent  float64 `json:"toss_win_percent"`
}
