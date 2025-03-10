package data

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/KunalDuran/duranz-stats/internal/models"
	"github.com/KunalDuran/duranz-stats/internal/utils"
)

func GetPlayerStats(playerName, league, season string, vsTeam int) []PlayerMatchStats {
	playerName = strings.ToLower(playerName)
	query := DB.Table("cricket_players AS player").
		Select("player.player_name, pms.*").
		Joins("LEFT JOIN player_match_stats AS pms ON pms.player_id = player.player_id").
		Joins("LEFT JOIN cricket_matches AS matches ON matches.match_id = pms.match_id").
		Where("(LOWER(player.player_name) = ? OR LOWER(player.display_name) = ?) AND matches.league_id = ?", playerName, playerName, models.AllDuranzLeagues[league])

	if season != "" {
		if seasonID, err := strconv.ParseInt(season, 10, 64); err == nil && seasonID > 1950 {
			query = query.Where("pms.season_id = ?", season)
		}
	}

	if vsTeam != 0 {
		query = query.Where("pms.team_id != ? AND (matches.away_team_id = ? OR matches.home_team_id = ?)",
			vsTeam, vsTeam, vsTeam)
	}

	var results []PlayerMatchStats
	if err := query.Find(&results).Error; err != nil {
		fmt.Println("Query Execution Error:", err)
		return nil
	}

	return results
}

func GetTeamStats(teamID int, gender, season, venue, vsTeam string) []MatchStatsExt {
	var objAllTeamStats []MatchStatsExt

	// Base query string
	query := DB.Table("cricket_matches AS matches").
		Joins("LEFT JOIN match_stats AS ms ON ms.match_id = matches.match_id").
		Where("(home_team_id = ? OR away_team_id = ?) AND ms.team_id = ? AND LOWER(gender) = ? AND matches.result != 'no result' AND matches.season_id = ?", teamID, teamID, teamID, utils.CleanText(gender, true), utils.CleanText(season, true))

	// Add venue filter if provided
	if venue != "" {
		venueID := GetVenueIDbyName(venue)
		query = query.Where("venue_id = ?", venueID)
	}

	// Add vsTeam filter if provided
	if vsTeam != "" {
		vsTeamID := GetTeamIDByTeamName(vsTeam)
		query = query.Where("away_team_id = ?", vsTeamID)
	}

	// Execute query
	if err := query.Select("matches.*, ms.*").Find(&objAllTeamStats).Error; err != nil {
		panic(err)
	}

	return objAllTeamStats
}

func GetTeamIDByTeamName(teamName string) int {
	var team Team
	if err := DB.Select("team_id").
		Where("LOWER(team_name) = ?", utils.CleanText(teamName, true)).
		First(&team).Error; err != nil {
		panic(err)
	}
	return team.TeamID
}

func GetPlayerIDByPlayerName(playerName string) int {
	var player CricketPlayer
	if err := DB.Select("player_id").
		Where("player_name = ?", playerName).
		First(&player).Error; err != nil {
		panic(err)
	}
	return player.PlayerID
}

func GetPlayerList(matchCount string) []string {
	var playerList []string
	// cachekey := "player-list"

	// Try to get from cache first
	// if cacheData, err := CacheGet(cachekey); err == nil {
	// 	if err := json.Unmarshal(cacheData, &playerList); err == nil {
	// 		return playerList
	// 	}
	// }

	// If not in cache, query database
	subQuery := DB.Model(&PlayerMatchStats{}).
		Select("player_id").
		Group("player_id").
		Having("COUNT(player_id) > ?", matchCount)

	if err := DB.Model(&CricketPlayer{}).
		Select("player_name").
		Where("player_id IN (?)", subQuery).
		Find(&playerList).Error; err != nil {
		panic(err)
	}

	// Cache the results
	// if result, err := json.Marshal(playerList); err == nil {
	// 	if err := CacheSetExp(cachekey, result, 43200); err != nil {
	// 		panic(err)
	// 	}
	// }

	return playerList
}

func GetTeamList() []string {
	var teamList []string
	// cachekey := "team-list"

	// // Try to get from cache first
	// if cacheData, err := CacheGet(cachekey); err == nil {
	// 	if err := json.Unmarshal(cacheData, &teamList); err == nil {
	// 		return teamList
	// 	}
	// }

	// If not in cache, query database
	if err := DB.Model(&Team{}).
		Select("team_name").
		Find(&teamList).Error; err != nil {
		panic(err)
	}

	// Cache the results
	// if result, err := json.Marshal(teamList); err == nil {
	// 	if err := CacheSetExp(cachekey, result, 43200); err != nil {
	// 		panic(err)
	// 	}
	// }

	return teamList
}

func GetMatchList(team, vsTeam int) []CricketMatch {
	var matchList []CricketMatch
	if err := DB.Model(&CricketMatch{}).
		Where("(home_team_id = ? AND away_team_id = ?) OR (home_team_id = ? AND away_team_id = ?)", team, vsTeam, vsTeam, team).
		Find(&matchList).Error; err != nil {
		panic(err)
	}
	return matchList
}

func GetMatchScoreCard(cricsheetID string) MatchScoreCard {
	var matchScoreCard MatchScoreCard
	if err := DB.Model(&MatchScoreCard{}).
		Where("cricsheet_id = ?", cricsheetID).
		Find(&matchScoreCard).Error; err != nil {
		panic(err)
	}
	return matchScoreCard
}
