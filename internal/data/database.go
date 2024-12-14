package data

import (
	"fmt"
	"strconv"

	"github.com/KunalDuran/duranz-stats/internal/models"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func GetPlayerStats(playerName, league, season string, vsTeam int) map[string][]PlayerMatchStats {
	objAllPlayerStats := map[string][]PlayerMatchStats{}

	query := DB.Table("cricket_players as player").
		Select("player.player_name, pms.*").
		Joins("LEFT JOIN player_match_stats AS pms ON pms.player_id = player.player_id").
		Joins("LEFT JOIN cricket_matches matches ON matches.match_id = pms.match_id").
		Where("player_name = ? AND league_id = ?", playerName, models.AllDuranzLeagues[league])

	if season != "" {
		if seasonID, err := strconv.ParseInt(season, 10, 64); err == nil && seasonID > 1950 {
			query = query.Where("pms.season_id = ?", season)
		}
	}

	if vsTeam != 0 {
		query = query.Where("pms.team_id != ? AND (matches.away_team_id = ? OR matches.home_team_id = ?)",
			vsTeam, vsTeam, vsTeam)
	}

	var results []struct {
		PlayerName string
		PlayerMatchStats
	}

	if err := query.Find(&results).Error; err != nil {
		panic(err)
	}

	for _, result := range results {
		objAllPlayerStats[result.PlayerName] = append(
			objAllPlayerStats[result.PlayerName],
			result.PlayerMatchStats,
		)
	}

	return objAllPlayerStats
}

func GetTeamStats(teamID int, gender, season, venue, vsTeam string) []models.DuranzMatchStats {
	var objAllTeamStats []models.DuranzMatchStats

	queryString := "(home_team_id = ? OR away_team_id = ?) AND gender = ?"
	if venue != "" {
		venueID := GetVenueIDbyName(venue)
		queryString += fmt.Sprintf(" AND venue_id = %d", venueID)
	}

	if vsTeam != "" {
		vsTeamID := GetTeamIDByTeamName(vsTeam)
		queryString += fmt.Sprintf(" AND away_team_id = %d", vsTeamID)
	}

	query := DB.Model(&CricketMatch{}).
		Where(queryString,
			teamID, teamID, gender)

	if err := query.Find(&objAllTeamStats).Error; err != nil {
		panic(err)
	}

	return objAllTeamStats
}

func GetTeamIDByTeamName(teamName string) int {
	var team Team
	if err := DB.Select("team_id").
		Where("team_name = ?", teamName).
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
