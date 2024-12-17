package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// addRouteHandlers adds routes for various APIs.
func addRouteHandlers(router *httprouter.Router) {

	router.GET("/", index)
	router.GET("/scorecard/:file", GetScoreCard)

	router.GET("/player-stats/:player", PlayerStatsAPI)
	router.GET("/team-stats/:team", TeamStatsAPI)
	router.GET("/player-vs-player/:player", PlayerVSPlayer)
	router.GET("/batsman-vs-bowler/", BatsmanVSBowlerAPI)

	router.GET("/player-list/", PlayerList)
	router.GET("/team-list/", TeamList)

}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("Duranz Statistics API"))
}
