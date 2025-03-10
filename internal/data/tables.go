package data

import (
	"encoding/json"
	"time"
)

type CricketMatch struct {
	MatchID           int       `gorm:"column:match_id;primaryKey;autoIncrement" json:"match_id"`
	LeagueID          int       `gorm:"column:league_id" json:"league_id"`
	Gender            string    `gorm:"column:gender" json:"gender"`
	SeasonID          int       `gorm:"column:season_id" json:"season_id"`
	HomeTeamID        int       `gorm:"column:home_team_id" json:"home_team_id"`
	AwayTeamID        int       `gorm:"column:away_team_id" json:"away_team_id"`
	HomeTeamName      string    `gorm:"column:home_team_name;size:120" json:"home_team_name"`
	AwayTeamName      string    `gorm:"column:away_team_name;size:120" json:"away_team_name"`
	VenueID           int       `gorm:"column:venue_id" json:"venue_id"`
	Result            string    `gorm:"column:result;size:200" json:"result"`
	ManOfTheMatch     int       `gorm:"column:man_of_the_match" json:"man_of_the_match"`
	TossWinner        int       `gorm:"column:toss_winner" json:"toss_winner"`
	TossDecision      string    `gorm:"column:toss_decision" json:"toss_decision"`
	WinningTeam       int       `gorm:"column:winning_team" json:"winning_team"`
	CricsheetFileName string    `gorm:"column:cricsheet_file_name;size:20" json:"cricsheet_file_name"`
	MatchDate         time.Time `gorm:"column:match_date" json:"match_date"`
	MatchDateMulti    string    `gorm:"column:match_date_multi;size:120" json:"match_date_multi"`
	MatchTime         time.Time `gorm:"column:match_time" json:"match_time"`
	IsReschedule      bool      `gorm:"column:is_reschedule;default:false" json:"is_reschedule"`
	IsAbandoned       bool      `gorm:"column:is_abandoned;default:false" json:"is_abandoned"`
	IsNeutral         bool      `gorm:"column:is_neutral;default:false" json:"is_neutral"`
	MatchRefrees      string    `gorm:"column:match_refrees;size:100" json:"match_refrees"`
	ReserveUmpires    string    `gorm:"column:reserve_umpires;size:100" json:"reserve_umpires"`
	TVUmpires         string    `gorm:"column:tv_umpires;size:100" json:"tv_umpires"`
	Umpires           string    `gorm:"column:umpires;size:100" json:"umpires"`
	EventName         string    `gorm:"column:event_name;size:100" json:"event_name"`
	MatchNumber       int       `gorm:"column:match_number" json:"match_number"`
	DateAdded         time.Time `gorm:"column:date_added;autoCreateTime" json:"date_added"`
	LastUpdate        time.Time `gorm:"column:last_update;autoUpdateTime" json:"last_update"`
	MatchEndTime      time.Time `gorm:"column:match_end_time" json:"match_end_time"`
	Status            string    `gorm:"column:status;size:2" json:"status"`

	// metadata
	Scores map[int]string `gorm:"-" json:"scores"`
}

type CricketPlayer struct {
	PlayerID        int       `gorm:"column:player_id;primaryKey;autoIncrement" json:"player_id"`
	PlayerName      string    `gorm:"column:player_name" json:"player_name"`
	DisplayName     string    `gorm:"column:display_name" json:"display_name"`
	FirstName       string    `gorm:"column:first_name" json:"first_name"`
	LastName        string    `gorm:"column:last_name" json:"last_name"`
	ShortName       string    `gorm:"column:short_name" json:"short_name"`
	UniqueShortName string    `gorm:"column:unique_short_name" json:"unique_short_name"`
	DOB             time.Time `gorm:"column:dob" json:"dob"`
	BattingStyle    string    `gorm:"column:batting_style" json:"batting_style"`
	BowlingStyle    string    `gorm:"column:bowling_style" json:"bowling_style"`
	IsOverseas      bool      `gorm:"column:is_overseas;default:false" json:"is_overseas"`
	CricsheetID     string    `gorm:"column:cricsheet_id;unique" json:"cricsheet_id"`
	DateAdded       time.Time `gorm:"column:date_added;autoCreateTime" json:"date_added"`
	Status          bool      `gorm:"column:status;default:true" json:"status"`
}

type Team struct {
	TeamID       int       `gorm:"column:team_id;primaryKey;autoIncrement" json:"team_id"`
	TeamName     string    `gorm:"column:team_name;not null" json:"team_name"`
	TeamType     string    `gorm:"column:team_type;not null" json:"team_type"`
	Filtername   string    `gorm:"column:filtername" json:"filtername"`
	Abbreviation string    `gorm:"column:abbreviation;size:4" json:"abbreviation"`
	TeamColor    string    `gorm:"column:team_color" json:"team_color"`
	Icon         string    `gorm:"column:icon" json:"icon"`
	URL          string    `gorm:"column:url" json:"url"`
	Jersey       string    `gorm:"column:jersey" json:"jersey"`
	Flag         string    `gorm:"column:flag" json:"flag"`
	Status       bool      `gorm:"column:status;default:true" json:"status"`
	DateAdded    time.Time `gorm:"column:dateadded;autoCreateTime" json:"date_added"`
}

type Venue struct {
	VenueID      int       `gorm:"column:venue_id;primaryKey;autoIncrement" json:"venue_id"`
	VenueName    string    `gorm:"column:venue_name;not null" json:"venue_name"`
	Filtername   string    `gorm:"column:filtername" json:"filtername"`
	Friendlyname string    `gorm:"column:friendlyname" json:"friendlyname"`
	City         string    `gorm:"column:city" json:"city"`
	Country      string    `gorm:"column:country" json:"country"`
	State        string    `gorm:"column:state" json:"state"`
	StateAbbr    string    `gorm:"column:state_abbr;size:5" json:"state_abbr"`
	OfficialTeam string    `gorm:"column:official_team" json:"official_team"`
	Capacity     int       `gorm:"column:capacity" json:"capacity"`
	Dimensions   string    `gorm:"column:dimensions" json:"dimensions"`
	Opened       int       `gorm:"column:opened" json:"opened"`
	Description  string    `gorm:"column:description" json:"description"`
	Shortname    string    `gorm:"column:shortname" json:"shortname"`
	Timezone     string    `gorm:"column:timezone" json:"timezone"`
	Weather      string    `gorm:"column:weather" json:"weather"`
	PitchType    string    `gorm:"column:pitch_type" json:"pitch_type"`
	DateAdded    time.Time `gorm:"column:dateadded;autoCreateTime" json:"date_added"`
	Status       bool      `gorm:"column:status;default:true" json:"status"`
}

type MatchStats struct {
	MatchID       int          `gorm:"column:match_id;primaryKey"`
	TeamID        int          `gorm:"column:team_id;primaryKey"`
	Innings       int          `gorm:"column:innings;primaryKey"`
	FallOfWickets string       `gorm:"column:fall_of_wickets"`
	Extras        int          `gorm:"column:extras"`
	Score         int          `gorm:"column:score"`
	Wickets       int          `gorm:"column:wickets"`
	OversPlayed   string       `gorm:"column:overs_played"`
	SuperOver     bool         `gorm:"column:super_over"`
	LastUpdate    time.Time    `gorm:"column:last_update;autoUpdateTime"`
	Match         CricketMatch `gorm:"foreignKey:match_id"`
	Team          Team         `gorm:"foreignKey:team_id"`
}

type PlayerMatchStats struct {
	MatchID                int           `gorm:"column:match_id;primaryKey"`
	SeasonID               string        `gorm:"column:season_id;primaryKey"`
	InningsID              string        `gorm:"column:innings_id;primaryKey"`
	SeasonType             string        `gorm:"column:season_type"`
	TeamID                 int           `gorm:"column:team_id;primaryKey"`
	PlayerID               int           `gorm:"column:player_id;primaryKey"`
	BattingOrder           int           `gorm:"column:batting_order"`
	RunsScored             int           `gorm:"column:runs_scored"`
	BallsFaced             int           `gorm:"column:balls_faced"`
	DotBallsPlayed         int           `gorm:"column:dot_balls_played"`
	Singles                int           `gorm:"column:singles"`
	Doubles                int           `gorm:"column:doubles"`
	Triples                int           `gorm:"column:triples"`
	FoursHit               int           `gorm:"column:fours_hit"`
	SixesHit               int           `gorm:"column:sixes_hit"`
	OutType                string        `gorm:"column:out_type"`
	OutFielder             int           `gorm:"column:out_fielder"`
	OutBowler              int           `gorm:"column:out_bowler"`
	IsBatted               bool          `gorm:"column:is_batted"`
	OversBowled            string        `gorm:"column:overs_bowled"`
	BowlingOrder           int           `gorm:"column:bowling_order"`
	RunsConceded           int           `gorm:"column:runs_conceded"`
	BallsBowled            int           `gorm:"column:balls_bowled"`
	DotsBowled             int           `gorm:"column:dots_bowled"`
	WicketsTaken           int           `gorm:"column:wickets_taken"`
	FoursConceded          int           `gorm:"column:fours_conceded"`
	SixesConceded          int           `gorm:"column:sixes_conceded"`
	ExtrasConceded         int           `gorm:"column:extras_conceded"`
	MaidenOver             int           `gorm:"column:maiden_over"`
	RunOut                 int           `gorm:"column:run_out"`
	Catches                int           `gorm:"column:catches"`
	Stumpings              int           `gorm:"column:stumpings"`
	PlayedAbandonedMatches int           `gorm:"column:played_abandoned_matches"`
	LastUpdate             time.Time     `gorm:"column:last_update;autoUpdateTime"`
	Match                  CricketMatch  `gorm:"foreignKey:match_id"`
	Team                   Team          `gorm:"foreignKey:team_id"`
	Player                 CricketPlayer `gorm:"foreignKey:player_id"`
}

type ErrorLog struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement"`
	AlertID   string    `gorm:"column:alert_id;unique"`
	Error     string    `gorm:"column:error"`
	ErrorMsg  string    `gorm:"column:error_msg"`
	FileName  string    `gorm:"column:file_name"`
	ErrorType string    `gorm:"column:error_type"`
	Severity  string    `gorm:"column:severity"`
	DateAdded time.Time `gorm:"column:dateadded;autoCreateTime"`
}

type FileMapping struct {
	FileID      int       `gorm:"column:file_id;primaryKey;autoIncrement"`
	FileName    string    `gorm:"column:file_name;unique"`
	LeagueID    int       `gorm:"column:league_id"`
	Teams       bool      `gorm:"column:teams;default:false"`
	Players     bool      `gorm:"column:players;default:false"`
	Venue       bool      `gorm:"column:venue;default:false"`
	Matches     bool      `gorm:"column:matches;default:false"`
	MatchStats  bool      `gorm:"column:match_stats;default:false"`
	PlayerStats bool      `gorm:"column:player_stats;default:false"`
	ScoreCard   bool      `gorm:"column:scorecard;default:false"`
	OverStats   int       `gorm:"column:over_stats"`
	DateAdded   time.Time `gorm:"column:dateadded;autoCreateTime"`
}

type MatchScoreCard struct {
	ID          uint            `gorm:"primaryKey" json:"id"`
	CricsheetID int             `gorm:"column:cricsheet_id" json:"cricsheet_id"`
	MatchID     int             `gorm:"column:match_id" json:"match_id"`
	Data        json.RawMessage `gorm:"type:json" json:"data"`
}

type MatchStatsExt struct {
	CricketMatch
	MatchStats
}
