package data

import (
	"time"
)

type CricketMatch struct {
	MatchID           int       `gorm:"column:match_id;primaryKey;autoIncrement"`
	LeagueID          int       `gorm:"column:league_id"`
	Gender            string    `gorm:"column:gender"`
	SeasonID          int       `gorm:"column:season_id"`
	HomeTeamID        int       `gorm:"column:home_team_id"`
	AwayTeamID        int       `gorm:"column:away_team_id"`
	HomeTeamName      string    `gorm:"column:home_team_name;size:120"`
	AwayTeamName      string    `gorm:"column:away_team_name;size:120"`
	VenueID           int       `gorm:"column:venue_id"`
	Result            string    `gorm:"column:result;size:200"`
	ManOfTheMatch     int       `gorm:"column:man_of_the_match"`
	TossWinner        int       `gorm:"column:toss_winner"`
	TossDecision      string    `gorm:"column:toss_decision"`
	WinningTeam       int       `gorm:"column:winning_team"`
	CricsheetFileName string    `gorm:"column:cricsheet_file_name;size:20"`
	MatchDate         time.Time `gorm:"column:match_date"`
	MatchDateMulti    string    `gorm:"column:match_date_multi;size:120"`
	MatchTime         time.Time `gorm:"column:match_time"`
	IsReschedule      bool      `gorm:"column:is_reschedule;default:false"`
	IsAbandoned       bool      `gorm:"column:is_abandoned;default:false"`
	IsNeutral         bool      `gorm:"column:is_neutral;default:false"`
	MatchRefrees      string    `gorm:"column:match_refrees;size:100"`
	ReserveUmpires    string    `gorm:"column:reserve_umpires;size:100"`
	TVUmpires         string    `gorm:"column:tv_umpires;size:100"`
	Umpires           string    `gorm:"column:umpires;size:100"`
	DateAdded         time.Time `gorm:"column:date_added;autoCreateTime"`
	LastUpdate        time.Time `gorm:"column:last_update;autoUpdateTime"`
	MatchEndTime      time.Time `gorm:"column:match_end_time"`
	Status            string    `gorm:"column:status;size:2"`
}

type CricketPlayer struct {
	PlayerID        int       `gorm:"column:player_id;primaryKey;autoIncrement"`
	PlayerName      string    `gorm:"column:player_name"`
	DisplayName     string    `gorm:"column:display_name"`
	FirstName       string    `gorm:"column:first_name"`
	LastName        string    `gorm:"column:last_name"`
	ShortName       string    `gorm:"column:short_name"`
	UniqueShortName string    `gorm:"column:unique_short_name"`
	DOB             time.Time `gorm:"column:dob"`
	BattingStyle    string    `gorm:"column:batting_style"`
	BowlingStyle    string    `gorm:"column:bowling_style"`
	IsOverseas      bool      `gorm:"column:is_overseas;default:false"`
	CricsheetID     string    `gorm:"column:cricsheet_id;unique"`
	DateAdded       time.Time `gorm:"column:date_added;autoCreateTime"`
	Status          bool      `gorm:"column:status;default:true"`
}

type Team struct {
	TeamID       int       `gorm:"column:team_id;primaryKey;autoIncrement"`
	TeamName     string    `gorm:"column:team_name;not null"`
	TeamType     string    `gorm:"column:team_type;not null"`
	Filtername   string    `gorm:"column:filtername"`
	Abbreviation string    `gorm:"column:abbreviation;size:4"`
	TeamColor    string    `gorm:"column:team_color"`
	Icon         string    `gorm:"column:icon"`
	URL          string    `gorm:"column:url"`
	Jersey       string    `gorm:"column:jersey"`
	Flag         string    `gorm:"column:flag"`
	Status       bool      `gorm:"column:status;default:true"`
	DateAdded    time.Time `gorm:"column:dateadded;autoCreateTime"`
}

type Venue struct {
	VenueID      int       `gorm:"column:venue_id;primaryKey;autoIncrement"`
	VenueName    string    `gorm:"column:venue_name;not null"`
	Filtername   string    `gorm:"column:filtername"`
	Friendlyname string    `gorm:"column:friendlyname"`
	City         string    `gorm:"column:city"`
	Country      string    `gorm:"column:country"`
	State        string    `gorm:"column:state"`
	StateAbbr    string    `gorm:"column:state_abbr;size:5"`
	OfficialTeam string    `gorm:"column:official_team"`
	Capacity     int       `gorm:"column:capacity"`
	Dimensions   string    `gorm:"column:dimensions"`
	Opened       int       `gorm:"column:opened"`
	Description  string    `gorm:"column:description"`
	Shortname    string    `gorm:"column:shortname"`
	Timezone     string    `gorm:"column:timezone"`
	Weather      string    `gorm:"column:weather"`
	PitchType    string    `gorm:"column:pitch_type"`
	DateAdded    time.Time `gorm:"column:dateadded;autoCreateTime"`
	Status       bool      `gorm:"column:status;default:true"`
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
	Teams       bool      `gorm:"column:teams"`
	Players     bool      `gorm:"column:players"`
	Venue       bool      `gorm:"column:venue"`
	Matches     bool      `gorm:"column:matches"`
	MatchStats  bool      `gorm:"column:match_stats"`
	PlayerStats bool      `gorm:"column:player_stats"`
	OverStats   int       `gorm:"column:over_stats"`
	DateAdded   time.Time `gorm:"column:dateadded;autoCreateTime"`
}

type MatchStatsExt struct {
	CricketMatch
	MatchStats
}
