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
	MatchID                int64  `json:"match_id,omitempty"`
	InningsID              string `json:"innings_id,omitempty"`
	SeasonID               string `json:"season_id,omitempty"`
	SeasonType             string `json:"season_type,omitempty"`
	PlayerID               int64  `json:"player_id,omitempty"`
	PlayerName             string `json:"player_name,omitempty"`
	TeamID                 int64  `json:"team_id,omitempty"`
	LastUpdate             string `json:"last_update,omitempty"`
	PlayedAbandonedMatches int64  `json:"played_abandoned_matches,omitempty"`

	Batting  PlayerBattingStats  `json:"batting,omitempty"`
	Bowling  PlayerBowlingStats  `json:"bowling,omitempty"`
	Fielding PlayerFieldingStats `json:"fielding,omitempty"`
}

type PlayerBattingStats struct {
	BallsFaced     int     `json:"balls_faced,omitempty"`
	BattingOrder   int     `json:"batting_order,omitempty"`
	DotBallsPlayed int     `json:"dot_balls_played,omitempty"`
	Doubles        int     `json:"doubles,omitempty"`
	FoursHit       int     `json:"fours_hit,omitempty"`
	IsBatted       int     `json:"is_batted,omitempty"`
	OutBowler      string  `json:"out_bowler,omitempty"`
	OutFielder     string  `json:"out_fielder,omitempty"`
	OutType        string  `json:"out_type,omitempty"`
	RunsScored     int     `json:"runs_scored,omitempty"`
	Singles        int     `json:"singles,omitempty"`
	SixesHit       int     `json:"sixes_hit,omitempty"`
	Triples        int     `json:"triples,omitempty"`
	Fifties        int     `json:"fifties,omitempty"`
	Hundreds       int     `json:"hundreds,omitempty"`
	Average        float64 `json:"average,omitempty"`
	HighestScore   int     `json:"highest_score,omitempty"`
	StrikeRate     float64 `json:"strike_rate,omitempty"`
	NotOuts        int     `json:"not_outs,omitempty"`
	Ducks          int     `json:"ducks,omitempty"`
}

type PlayerBowlingStats struct {
	BowlingOrder   int     `json:"bowling_order,omitempty"`
	BallsBowled    int     `json:"balls_bowled,omitempty"`
	DotsBowled     int     `json:"dots_bowled,omitempty"`
	ExtrasConceded int     `json:"extras_conceded,omitempty"`
	FoursConceded  int     `json:"fours_conceded,omitempty"`
	MaidenOver     int     `json:"maiden_over,omitempty"`
	OversBowled    string  `json:"overs_bowled,omitempty"`
	RunsConceded   int     `json:"runs_conceded,omitempty"`
	SixesConceded  int     `json:"sixes_conceded,omitempty"`
	WicketsTaken   int     `json:"wickets_taken,omitempty"`
	Economy        float64 `json:"economy,omitempty"`
	Average        float64 `json:"average,omitempty"`
	Fifers         float64 `json:"fifers,omitempty"`
	BestBowling    string  `json:"best_bowling,omitempty"`
}

type PlayerFieldingStats struct {
	RunOut    int `json:"run_out,omitempty"`
	Stumpings int `json:"stumpings,omitempty"`
	Catches   int `json:"catches,omitempty"`
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
	TotalMatches    int     `json:"total_matches,omitempty"`
	MatchWin        int     `json:"match_win,omitempty"`
	MatchWinPercent float64 `json:"win_percent,omitempty"`
	BatFirstWin     int     `json:"bat_first_win,omitempty"`
	ChasingWin      int     `json:"chasing_win,omitempty"`
	BatFirstWinPer  float64 `json:"bat_first_win_per,omitempty"`
	ChasingWinPer   float64 `json:"chasing_win_per,omitempty"`
	AvgScoreInn     float64 `json:"avg_score_inn,omitempty"`
	HighestScore    int     `json:"highest_score,omitempty"`
	LowestScore     int     `json:"lowest_score,omitempty"`
	TossWin         float64 `json:"toss_win,omitempty"`
	TossWinPercent  float64 `json:"toss_win_percent,omitempty"`
}
