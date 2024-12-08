package main

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// addRouteHandlers adds routes for various APIs.
func addRouteHandlers(router *httprouter.Router) {

	// General links
	router.GET("/", index)
	router.GET("/scorecard/:file", GetScoreCard)

	// CMS
	router.GET("/cms/display", DisplayPageCMS)

	router.GET("/cms/missing/players", GetMissingPlayerDetails)
	router.GET("/cms/missing/teams", GetMissingTeamDetails)
	router.GET("/cms/missing/venues", GetMissingVenueDetails)

	router.POST("/cms/mapping/players", MapPlayerDetails)
	router.POST("/cms/mapping/teams", MapTeamDetails)
	router.POST("/cms/mapping/venues", MapVenueDetails)

	// API
	router.GET("/player-stats/:player", PlayerStatsAPI)
	router.GET("/team-stats/:team", TeamStatsAPI)
	router.GET("/player-vs-team/:player", PlayerStatsAPI)
	router.GET("/batsman-vs-bowler/", BatsmanVSBowlerAPI)

	router.GET("/player-list/", PlayerList)
	router.GET("/team-list/", TeamList)

}

// Plain index page content
func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, sampleJSON)
}

var sampleJSON = `
{
    "license": "Duranz API",
    "statusCode": 200,
    "content": [
        {
            "season_id": "2021",
            "player_id": 149,
            "player_name": "Mohammed Shami",
            "team_id": 12,
            "batting": {
                "balls_faced": 17,
                "dot_balls_played": 7,
                "doubles": 3,
                "is_batted": 4,
                "runs_scored": 13,
                "singles": 7,
                "average": 13,
                "highest_score": 9,
                "strike_rate": 76.48,
                "not_outs": 3
            },
            "bowling": {
                "balls_bowled": 316,
                "dots_bowled": 151,
                "extras_conceded": 8,
                "fours_conceded": 44,
                "maiden_over": 1,
                "overs_bowled": "52.4",
                "runs_conceded": 395,
                "sixes_conceded": 14,
                "wickets_taken": 19,
                "economy": 7.5,
                "average": 20.79,
                "best_bowling": "3/21"
            },
            "fielding": {
                "catches": 1
            }
        }
    ]
}`
