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
	OutBowler    string  `json:"-"`
	OutFielder   string  `json:"-"`
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
	OutBowler    string
	OutFielder   string
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

type TeamStats struct {
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

type Venue struct {
}

type PlayerStatsBind struct {
	Batting  map[string]BattingStats
	Bowling  map[string]BowlingStats
	Fielding map[string]FieldingStats
}

type Officials struct {
	MatchReferees  []string `json:"match_referees"`
	ReserveUmpires []string `json:"reserve_umpires"`
	TvUmpires      []string `json:"tv_umpires"`
	Umpires        []string `json:"umpires"`
}

type OverDetails struct {
	Over       int `json:"over"`
	Deliveries []struct {
		Batter     string `json:"batter"`
		Bowler     string `json:"bowler"`
		Extras     `json:"extras"`
		NonStriker string `json:"non_striker"`
		Runs       struct {
			Batter int `json:"batter"`
			Extras int `json:"extras"`
			Total  int `json:"total"`
		} `json:"runs"`
		Wickets []struct {
			Kind      string `json:"kind"`
			PlayerOut string `json:"player_out"`
			Fielders  []struct {
				Name       string `json:"name"`
				Substitute bool   `json:"substitute"`
			} `json:"fielders"`
		}
	} `json:"deliveries"`
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
