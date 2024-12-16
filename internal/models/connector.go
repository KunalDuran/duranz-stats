package models

type BowlingStats struct {
	BowlingOrder  int     `json:"bowling_order"`
	Name          string  `json:"name"`
	Overs         string  `json:"overs"`
	Maiden        int     `json:"maiden"`
	Runs          int     `json:"runs"`
	Wickets       int     `json:"wickets"`
	Economy       float64 `json:"economy"`
	Balls         int     `json:"-"`
	Dots          int     `json:"-"`
	IsBowled      bool    `json:"-"`
	FoursConceded int     `json:"-"`
	SixesConceded int     `json:"-"`
	Extras        int     `json:"-"`
}

type BattingStats struct {
	BattingOrder int     `json:"batting_order"`
	Name         string  `json:"name"`
	Runs         int     `json:"runs"`
	Balls        int     `json:"balls"`
	Fours        int     `json:"fours"`
	Sixes        int     `json:"sixes"`
	StrikeRate   float64 `json:"strike_rate"`
	Out          string  `json:"out"`
	OutBowler    int     `json:"-"`
	OutFielder   int     `json:"-"`
	Singles      int     `json:"-"`
	Doubles      int     `json:"-"`
	Triples      int     `json:"-"`
	Dots         int     `json:"-"`
	NotOut       bool    `json:"-"`
	IsBatted     bool    `json:"-"`
}

type PlayerStats struct {
	PlayerID int
	Name     string

	// bowling stats
	BowlingOrder   int
	OversBowled    string
	MaidenOvers    int
	RunsConceded   int
	WicketsTaken   int
	Economy        float64
	BallsBowled    int
	DotsBowled     int
	FoursConceded  int
	SixesConceded  int
	ExtrasConceded int

	// batting stats
	BattingOrder int
	RunsScored   int
	BallsPlayed  int
	Singles      int
	Doubles      int
	Triples      int
	FoursHit     int
	SixesHit     int
	StrikeRate   float64
	OutType      string
	OutBowler    int
	OutFielder   int
	DotsPlayed   int
	NotOut       bool
	IsBatted     bool

	// fielding stats
	RunOuts   int
	Catches   int
	Stumpings int
}

type FieldingStats struct {
	Name      string
	RunOuts   int
	Catches   int
	Stumpings int
}

type MatchStats struct {
	MatchID     int
	TeamID      int
	InningsID   int
	Captain     string
	FOW         string
	Extras      int
	Wickets     int
	OversPlayed string
	Score       int
	SuperOver   bool
}

type PlayerStatsBind struct {
	Batting  map[string]BattingStats
	Bowling  map[string]BowlingStats
	Fielding map[string]FieldingStats
}

type MappingInfo struct {
	FileName    string
	LeagueID    int
	Teams       bool
	Players     bool
	Venue       bool
	Match       bool
	MatchStats  bool
	PlayerStats bool
}
