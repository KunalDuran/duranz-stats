package models

import "encoding/json"

type Match struct {
	Meta struct {
		DataVersion string `json:"data_version"`
		Created     string `json:"created"`
		Revision    int    `json:"revision"`
	} `json:"meta"`
	Info struct {
		BallsPerOver int      `json:"balls_per_over"`
		City         string   `json:"city"`
		Dates        []string `json:"dates"`
		Event        struct {
			Name        string `json:"name"`
			MatchNumber int    `json:"match_number"`
		} `json:"event"`
		Gender          string `json:"gender"`
		MatchType       string `json:"match_type"`
		MatchTypeNumber int    `json:"match_type_number"`
		Officials       `json:"officials"`
		Outcome         struct {
			By struct {
				Runs    int `json:"runs"`
				Wickets int `json:"wickets"`
			} `json:"by"`
			Winner     string `json:"winner"`
			Eliminator string `json:"eliminator"`
			Result     string `json:"result"`
			Method     string `json:"method"`
		} `json:"outcome"`
		Overs         int                 `json:"overs"`
		PlayerOfMatch []string            `json:"player_of_match"`
		Players       map[string][]string `json:"players"`
		Register      Registry            `json:"registry"`
		Season        json.RawMessage     `json:"season"`
		TeamType      string              `json:"team_type"`
		Teams         []string            `json:"teams"`
		Toss          struct {
			Decision string `json:"decision"`
			Winner   string `json:"winner"`
		} `json:"toss"`
		Venue string `json:"venue"`
	} `json:"info"`
	Innings []struct {
		Team       string        `json:"team"`
		Overs      []OverDetails `json:"overs"`
		SuperOvers bool          `json:"super_over"`
		Powerplays []struct {
			From float64 `json:"from"`
			To   float64 `json:"to"`
			Type string  `json:"type"`
		} `json:"powerplays"`
		Target struct {
			Overs json.RawMessage `json:"overs"`
			Runs  int             `json:"runs"`
		} `json:"target,omitempty"`
	} `json:"innings"`
}

type Registry struct {
	People map[string]string `json:"people"`
}

type Innings struct {
	InningID      int            `json:"innings_id"`
	InningDetail  string         `json:"innings_detail"`
	Bowling       []BowlingStats `json:"bowling"`
	Batting       []BattingStats `json:"batting"`
	Extras        `json:"extras"`
	FallOfWickets string `json:"fall_of_wickets"`
}

type Extras struct {
	Wides   int `json:"wides"`
	NoBall  int `json:"noballs"`
	Byes    int `json:"byes"`
	LegByes int `json:"legbyes"`
	Total   int `json:"total"`
}
