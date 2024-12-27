package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func addRouteHandlers(router *chi.Mux) {
	router.Get("/", index)
	router.Get("/match-stats/{file}", MatchStats)
	router.Get("/player-stats/{player}", PlayerStats)
	router.Get("/team-stats/{team}", TeamStats)

	router.Get("/player-list", PlayerList)
	router.Get("/team-list", TeamList)
	router.Get("/match-list", MatchList)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Duranz Statistics API"))
}
