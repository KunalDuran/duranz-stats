package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// addRouteHandlers adds routes for various APIs.
func addRouteHandlers(router *chi.Mux) {
	router.Get("/", index)
	router.Get("/scorecard/{file}", GetScoreCard)

	router.Get("/player-stats/{player}", PlayerStatsAPI)
	router.Get("/team-stats/{team}", TeamStatsAPI)
	router.Get("/player-vs-player/{player}", PlayerVSPlayer)
	router.Get("/batsman-vs-bowler/", BatsmanVSBowlerAPI)

	router.Get("/player-list/", PlayerList)
	router.Get("/team-list/", TeamList)

}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Duranz Statistics API"))
}
