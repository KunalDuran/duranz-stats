package models

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
