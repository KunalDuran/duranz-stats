package utils

import "fmt"

func ESPNPlayerURL(playerID string) string {
	// var url = fmt.Sprintf("https://www.espncricinfo.com/player/player-name-%s", playerID)
	var json_url = fmt.Sprintf("http://core.espnuk.org/v2/sports/cricket/athletes/%s", playerID)
	return json_url
}
