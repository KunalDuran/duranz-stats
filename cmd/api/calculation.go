package main

import (
	"math"

	"github.com/KunalDuran/duranz-stats/internal/data"
	"github.com/KunalDuran/duranz-stats/internal/models"
	"github.com/KunalDuran/duranz-stats/internal/utils"
)

func calculateTeamStats(objAllTeamStats []data.MatchStatsExt, teamID int) models.DuranzTeamStats {

	var teamStatistics models.DuranzTeamStats

	var totalScore int
	for _, objTeam := range objAllTeamStats {
		totalScore += objTeam.Score

		if objTeam.Score > teamStatistics.HighestScore.Runs {
			teamStatistics.HighestScore.Runs = objTeam.Score
			teamStatistics.HighestScore.Match = objTeam.Match.MatchID
		}

		if objTeam.Score < teamStatistics.LowestScore.Runs || teamStatistics.LowestScore.Runs == 0 {
			teamStatistics.LowestScore.Runs = objTeam.Score
			teamStatistics.LowestScore.Match = objTeam.Match.MatchID
		}

		if objTeam.TossWinner == teamID {
			teamStatistics.TossWin++
		}

		if objTeam.WinningTeam == teamID {
			teamStatistics.MatchWin++

			// count if team won while batting first or chasing first
			if objTeam.TossWinner == teamID { // team won the toss
				if objTeam.TossDecision == "bat" {
					teamStatistics.BatFirstWin++
				} else {
					teamStatistics.ChasingWin++
				}
			} else {
				if objTeam.TossDecision == "bat" { // other team won the toss
					teamStatistics.ChasingWin++
				} else {
					teamStatistics.BatFirstWin++
				}
			}
		}
	}

	teamStatistics.TotalMatches = len(objAllTeamStats)
	if teamStatistics.TotalMatches > 0 {
		teamStatistics.MatchWinPercent = utils.Round(float64(teamStatistics.MatchWin)/float64(teamStatistics.TotalMatches), 0.01, 2) * 100
		teamStatistics.TossWinPercent = utils.Round(float64(teamStatistics.TossWin)/float64(teamStatistics.TotalMatches), 0.01, 2) * 100

		teamStatistics.AvgScoreInn = utils.Round(float64(totalScore)/float64(teamStatistics.TotalMatches), 0.01, 2)
	}

	if teamStatistics.MatchWin > 0 {
		teamStatistics.ChasingWinPer = math.Round((float64(teamStatistics.ChasingWin)/float64(teamStatistics.MatchWin))*10000) / 100
		teamStatistics.BatFirstWinPer = math.Round((float64(teamStatistics.BatFirstWin)/float64(teamStatistics.MatchWin))*10000) / 100
	}
	return teamStatistics
}
