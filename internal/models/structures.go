package models

import "database/sql"

var AllDuranzLeagues = map[string]int{
	"odi":                   1,
	"test":                  2,
	"t20":                   3,
	"ipl":                   4,
	"Indian Premier League": 4,
}

type ScoreCard struct {
	Innings []Innings `json:"innings"`
	Result  string    `json:"result"`
}

type PlayerDetailsInt struct {
	PlayerID        sql.NullInt64
	PlayerName      sql.NullString
	DisplayName     sql.NullString
	FirstName       sql.NullString
	LastName        sql.NullString
	ShortName       sql.NullString
	UniqueShortName sql.NullString
	DOB             sql.NullString
	BattingStyle    sql.NullInt64
	BowlingStyle    sql.NullInt64
	IsOverseas      sql.NullInt64
	CricSheetID     sql.NullString
	DateAdded       sql.NullString
	Status          sql.NullInt64
}

type PlayerDetailsExt struct {
	PlayerID        int64  `json:"player_id"`
	PlayerName      string `json:"player_name"`
	DisplayName     string `json:"display_name"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	ShortName       string `json:"short_name"`
	UniqueShortName string `json:"unique_short_name"`
	DOB             string `json:"dob"`
	BattingStyle    int64  `json:"batting_style"`
	BowlingStyle    int64  `json:"bowling_style"`
	IsOverseas      int64  `json:"is_overseas"`
	CricSheetID     string `json:"cricsheet_id"`
	DateAdded       string `json:"date_added"`
	Status          int64  `json:"status"`
}

type TeamDetailsInt struct {
	TeamID     sql.NullInt64
	TeamName   sql.NullString
	TeamType   sql.NullString
	FilterName sql.NullString
	ABBR       sql.NullString
	TeamColor  sql.NullString
	Icon       sql.NullString
	URL        sql.NullString
	Jersey     sql.NullString
	Flag       sql.NullString
	Status     sql.NullInt64
	DateAdded  sql.NullString
}

type TeamDetailsExt struct {
	TeamID     int64  `json:"team_id"`
	TeamName   string `json:"team_name"`
	TeamType   string `json:"team_type"`
	FilterName string `json:"filter_name"`
	ABBR       string `json:"abbr"`
	TeamColor  string `json:"team_color"`
	Icon       string `json:"icon"`
	URL        string `json:"url"`
	Jersey     string `json:"jersey"`
	Flag       string `json:"flag"`
	Status     int64  `json:"status"`
	DateAdded  string `json:"date_added"`
}

type VenueDetailsInt struct {
	VenueID      sql.NullInt64
	Venue        sql.NullString
	FilterName   sql.NullString
	FriendlyName sql.NullString
	City         sql.NullString
	Country      sql.NullString
	State        sql.NullString
	StateABBR    sql.NullString
	OfficialTeam sql.NullString
	Capacity     sql.NullString
	Dimension    sql.NullString
	Opened       sql.NullString
	Description  sql.NullString
	ShortName    sql.NullString
	TimeZone     sql.NullString
	Weather      sql.NullString
	PitchType    sql.NullString
	DateAdded    sql.NullString
	Status       sql.NullInt64
}

type VenueDetailsExt struct {
	VenueID      int64  `json:"venue_id"`
	Venue        string `json:"venue_name"`
	FilterName   string `json:"filter_name"`
	FriendlyName string `json:"friendly_name"`
	City         string `json:"city"`
	Country      string `json:"country"`
	State        string `json:"state"`
	StateABBR    string `json:"state_abbr"`
	OfficialTeam string `json:"official_team"`
	Capacity     string `json:"capacity"`
	Dimension    string `json:"dimension"`
	Opened       string `json:"opened"`
	Description  string `json:"description"`
	ShortName    string `json:"short_name"`
	TimeZone     string `json:"time_zone"`
	Weather      string `json:"weather"`
	PitchType    string `json:"pitch_type"`
	DateAdded    string `json:"date_added"`
	Status       int64  `json:"status"`
}

type PlayerStatsInt struct {
	MatchID                sql.NullInt64
	InningsID              sql.NullString
	BallsBowled            sql.NullInt64
	BallsFaced             sql.NullInt64
	BattingOrder           sql.NullInt64
	BowlingOrder           sql.NullInt64
	Catches                sql.NullInt64
	DotBallsPlayed         sql.NullInt64
	DotsBowled             sql.NullInt64
	Doubles                sql.NullInt64
	ExtrasConceded         sql.NullInt64
	FoursConceded          sql.NullInt64
	FoursHit               sql.NullInt64
	IsBatted               sql.NullInt64
	LastUpdate             sql.NullString
	MaidenOver             sql.NullInt64
	PlayedAbandonedMatches sql.NullInt64
	PlayerID               sql.NullInt64
	RunOut                 sql.NullInt64
	RunsConceded           sql.NullInt64
	RunsScored             sql.NullInt64
	Singles                sql.NullInt64
	SixesConceded          sql.NullInt64
	SixesHit               sql.NullInt64
	Stumpings              sql.NullInt64
	TeamID                 sql.NullInt64
	Triples                sql.NullInt64
	WicketsTaken           sql.NullInt64
	OutBowler              sql.NullString
	OutFielder             sql.NullString
	OutType                sql.NullString
	OversBowled            sql.NullString
	SeasonID               sql.NullString
	SeasonType             sql.NullString
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
	BallsFaced     int64   `json:"balls_faced,omitempty"`
	BattingOrder   int64   `json:"batting_order,omitempty"`
	DotBallsPlayed int64   `json:"dot_balls_played,omitempty"`
	Doubles        int64   `json:"doubles,omitempty"`
	FoursHit       int64   `json:"fours_hit,omitempty"`
	IsBatted       int64   `json:"is_batted,omitempty"`
	OutBowler      string  `json:"out_bowler,omitempty"`
	OutFielder     string  `json:"out_fielder,omitempty"`
	OutType        string  `json:"out_type,omitempty"`
	RunsScored     int64   `json:"runs_scored,omitempty"`
	Singles        int64   `json:"singles,omitempty"`
	SixesHit       int64   `json:"sixes_hit,omitempty"`
	Triples        int64   `json:"triples,omitempty"`
	Fifties        int64   `json:"fifties,omitempty"`
	Hundreds       int64   `json:"hundreds,omitempty"`
	Average        float64 `json:"average,omitempty"`
	HighestScore   int64   `json:"highest_score,omitempty"`
	StrikeRate     float64 `json:"strike_rate,omitempty"`
	NotOuts        int64   `json:"not_outs,omitempty"`
	Ducks          int64   `json:"ducks,omitempty"`
}
type PlayerBowlingStats struct {
	BowlingOrder   int64   `json:"bowling_order,omitempty"`
	BallsBowled    int64   `json:"balls_bowled,omitempty"`
	DotsBowled     int64   `json:"dots_bowled,omitempty"`
	ExtrasConceded int64   `json:"extras_conceded,omitempty"`
	FoursConceded  int64   `json:"fours_conceded,omitempty"`
	MaidenOver     int64   `json:"maiden_over,omitempty"`
	OversBowled    string  `json:"overs_bowled,omitempty"`
	RunsConceded   int64   `json:"runs_conceded,omitempty"`
	SixesConceded  int64   `json:"sixes_conceded,omitempty"`
	WicketsTaken   int64   `json:"wickets_taken,omitempty"`
	Economy        float64 `json:"economy,omitempty"`
	Average        float64 `json:"average,omitempty"`
	Fifers         float64 `json:"fifers,omitempty"`
	BestBowling    string  `json:"best_bowling,omitempty"`
}
type PlayerFieldingStats struct {
	RunOut    int64 `json:"run_out,omitempty"`
	Stumpings int64 `json:"stumpings,omitempty"`
	Catches   int64 `json:"catches,omitempty"`
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
