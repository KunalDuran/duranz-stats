package data

import (
	"fmt"
	"log"
	"strconv"

	"github.com/KunalDuran/duranz-stats/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// InsertErrorLog : Insert All Error log record according to alert ID, file name
func InsertErrorLog(alertid, errormsg, fileName, err string) {
	errorLog := ErrorLog{
		AlertID:  alertid,
		Error:    err,
		ErrorMsg: errormsg,
		FileName: fileName,
	}

	log.Println(err)

	if err := DB.Create(&errorLog).Error; err != nil {
		panic(err)
	}
}

var GamePath = map[string]string{
	"odi":  "odis_json",
	"test": "tests_json",
	"t20":  "t20s_json",
	"ipl":  "ipl_json",
}

// cache simulation
var MappedTeams = map[string]string{}
var MappedVenues = map[string]string{}

func InsertMatchStats(matchID int, objMatchStats models.MatchStats) {
	matchStats := MatchStats{
		MatchID:       matchID,
		TeamID:        objMatchStats.TeamID,
		FallOfWickets: objMatchStats.FOW,
		Extras:        objMatchStats.Extras,
		Score:         objMatchStats.Score,
		SuperOver:     objMatchStats.SuperOver,
		Wickets:       objMatchStats.Wickets,
		OversPlayed:   objMatchStats.OversPlayed,
		Innings:       objMatchStats.InningsID,
	}

	if err := DB.Create(&matchStats).Error; err != nil {
		panic(err)
	}
}

func GetVenueID(venueName, city string) int {
	var venue Venue
	if err := DB.Select("venue_id").Where("venue_name = ? AND city = ?", venueName, city).First(&venue).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			panic(err)
		}
		return 0
	}
	return venue.VenueID
}

func GetVenueIDbyName(venueName string) int {
	var venue Venue
	if err := DB.Select("venue_id").Where("venue_name = ?", venueName).First(&venue).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			panic(err)
		}
		return 0
	}
	return venue.VenueID
}
func GetTeamID(teamName, teamType string) int {
	var team Team
	if err := DB.Select("team_id").Where("team_name = ? AND team_type = ?", teamName, teamType).First(&team).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			panic(err)
		}
		return 0
	}
	return team.TeamID
}

func GetPlayerID(cricsheetID string) int {
	var player CricketPlayer
	if err := DB.Select("player_id").Where("cricsheet_id = ?", cricsheetID).First(&player).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			panic(err)
		}
		return 0
	}
	return player.PlayerID
}

func GetMatchID(cricsheetID string) int {
	var match CricketMatch
	if err := DB.Select("match_id").Where("cricsheet_file_name = ?", cricsheetID).First(&match).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			panic(err)
		}
		return 0
	}
	return match.MatchID
}

func InsertMappingInfo(fileName string, mappingInfo models.MappingInfo) {
	fileMapping := FileMapping{
		FileName:    fileName,
		LeagueID:    mappingInfo.LeagueID,
		Teams:       mappingInfo.Teams,
		Players:     mappingInfo.Players,
		Venue:       mappingInfo.Venue,
		Matches:     mappingInfo.Match,
		MatchStats:  mappingInfo.MatchStats,
		PlayerStats: mappingInfo.PlayerStats,
	}

	if err := DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "file_name"}},
		DoUpdates: clause.AssignmentColumns([]string{"teams", "players", "venue", "matches", "match_stats", "player_stats"}),
	}).Create(&fileMapping).Error; err != nil {
		panic(err)
	}
}

func GetMappingDetails() map[string]models.MappingInfo {
	var fileMappings []FileMapping
	objMappingInfo := map[string]models.MappingInfo{}

	if err := DB.Find(&fileMappings).Error; err != nil {
		panic(err)
	}

	for _, mapping := range fileMappings {
		objMappingInfo[mapping.FileName] = models.MappingInfo{
			LeagueID:    mapping.LeagueID,
			Teams:       mapping.Teams,
			Players:     mapping.Players,
			Venue:       mapping.Venue,
			Match:       mapping.Matches,
			MatchStats:  mapping.MatchStats,
			PlayerStats: mapping.PlayerStats,
		}
	}

	return objMappingInfo
}

func InsertPlayerStats(matchID, seasonID int, teamInningPlayerStats map[string]map[string]models.PlayerStats, allPlayerID map[string]int, innBatTeamMap map[int]int) {
	var playerStats []PlayerMatchStats

	for teamID, allPlayerStats := range teamInningPlayerStats {
		for _, stats := range allPlayerStats {
			tempPlayerID := allPlayerID[stats.Name]
			if tempPlayerID == 0 {
				fmt.Println(stats.Name)
				continue
			}

			teamIDint, _ := strconv.Atoi(teamID)
			innID := innBatTeamMap[teamIDint]

			playerStats = append(playerStats, PlayerMatchStats{
				MatchID:        matchID,
				SeasonID:       strconv.Itoa(seasonID),
				InningsID:      strconv.Itoa(innID),
				TeamID:         teamIDint,
				PlayerID:       tempPlayerID,
				BattingOrder:   stats.BattingOrder,
				RunsScored:     stats.RunsScored,
				BallsFaced:     stats.BallsPlayed,
				DotBallsPlayed: stats.DotsPlayed,
				Singles:        stats.Singles,
				Doubles:        stats.Doubles,
				Triples:        stats.Triples,
				FoursHit:       stats.FoursHit,
				SixesHit:       stats.SixesHit,
				OutType:        stats.OutType,
				OutBowler:      stats.OutBowler,
				OutFielder:     stats.OutFielder,
				IsBatted:       stats.IsBatted,
				OversBowled:    stats.OversBowled,
				RunsConceded:   stats.RunsConceded,
				BallsBowled:    stats.BallsBowled,
				DotsBowled:     stats.DotsBowled,
				WicketsTaken:   stats.WicketsTaken,
				FoursConceded:  stats.FoursConceded,
				SixesConceded:  stats.SixesConceded,
				ExtrasConceded: stats.ExtrasConceded,
				MaidenOver:     stats.MaidenOvers,
				BowlingOrder:   stats.BowlingOrder,
				RunOut:         stats.RunOuts,
				Catches:        stats.Catches,
				Stumpings:      stats.Stumpings,
			})
		}
	}

	// Batch insert all player stats
	if len(playerStats) > 0 {
		if err := DB.CreateInBatches(playerStats, 100).Error; err != nil {
			panic(err)
		}
	}
}

func PseudoCacheLayer(teamType string) {
	// Clear existing cache
	MappedTeams = make(map[string]string)
	MappedVenues = make(map[string]string)

	// Load Teams
	if teamType != "ipl" {
		teamType = "international"
	}

	var teams []Team
	if err := DB.Select("team_name").Where("team_type = ?", teamType).Find(&teams).Error; err != nil {
		panic(err)
	}

	for _, team := range teams {
		MappedTeams[team.TeamName] = teamType
	}

	// Load Venues
	var venues []Venue
	if err := DB.Select("venue_name, city").Find(&venues).Error; err != nil {
		panic(err)
	}

	for _, venue := range venues {
		MappedVenues[venue.VenueName] = venue.City
	}
}
