package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/KunalDuran/duranz-stats/internal/data"
	"github.com/KunalDuran/duranz-stats/internal/models"
	"github.com/KunalDuran/duranz-stats/internal/utils"
	"github.com/julienschmidt/httprouter"
)

func DisplayPageCMS(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	param := r.URL.Query().Get("page")
	offset := r.URL.Query().Get("offset")
	var offsetInt int64
	offsetInt, _ = strconv.ParseInt(offset, 10, 64)

	if param != "players" && param != "venues" && param != "teams" {
		w.Write([]byte("This is CMS Page, Pass players, venues or teams in query param 'page'"))
		return
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	body, _, err := data.RequestAPIData("http://localhost:5000/cms/missing/"+param, headers)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.ParseFiles("templates/" + param + ".html")
	if err != nil {
		panic(err)
	}

	switch param {
	case "players":

		finalPlayers := struct {
			Content []models.PlayerDetailsExt `json:"content"`
		}{}
		err = json.Unmarshal(body, &finalPlayers)
		fmt.Println(len(finalPlayers.Content))
		if err != nil {
			panic(err)
		}
		if len(finalPlayers.Content) == 0 {
			w.Write([]byte("This is CMS Page, No missing player details"))
			return
		}
		if len(finalPlayers.Content) < int(offsetInt)+1 {
			offsetInt = 0
		}
		tmpl.Execute(w, finalPlayers.Content[offsetInt])

	case "teams":

		finalTeams := struct {
			Content []models.PlayerDetailsExt `json:"content"`
		}{}
		err = json.Unmarshal(body, &finalTeams)
		fmt.Println(len(finalTeams.Content))
		if err != nil {
			panic(err)
		}
		if len(finalTeams.Content) == 0 {
			w.Write([]byte("This is CMS Page, No missing team details"))
			return
		}
		if len(finalTeams.Content) < int(offsetInt)+1 {
			offsetInt = 0
		}
		tmpl.Execute(w, finalTeams.Content[offsetInt])

	case "venues":

		finalVenues := struct {
			Content []models.VenueDetailsExt `json:"content"`
		}{}
		err = json.Unmarshal(body, &finalVenues)
		fmt.Println(len(finalVenues.Content))
		if err != nil {
			panic(err)
		}
		if len(finalVenues.Content) == 0 {
			w.Write([]byte("This is CMS Page, No missing Venue details"))
			return
		}
		if len(finalVenues.Content) < int(offsetInt)+1 {
			offsetInt = 0
		}
		tmpl.Execute(w, finalVenues.Content[offsetInt])
	}
}

func GetMissingPlayerDetails(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	objPlayers := data.GetPlayerDetails()

	var finalPlayers []models.PlayerDetailsExt
	for _, objPlayer := range objPlayers {

		var finalPlayer models.PlayerDetailsExt

		finalPlayer.PlayerID = objPlayer.PlayerID.Int64
		finalPlayer.PlayerName = objPlayer.PlayerName.String
		finalPlayer.DisplayName = objPlayer.DisplayName.String
		finalPlayer.FirstName = objPlayer.FirstName.String
		finalPlayer.LastName = objPlayer.LastName.String
		finalPlayer.ShortName = objPlayer.ShortName.String
		finalPlayer.UniqueShortName = objPlayer.UniqueShortName.String
		finalPlayer.DOB = objPlayer.DOB.String
		finalPlayer.BattingStyle = objPlayer.BattingStyle.Int64
		finalPlayer.BowlingStyle = objPlayer.BowlingStyle.Int64
		finalPlayer.IsOverseas = objPlayer.IsOverseas.Int64
		finalPlayer.CricSheetID = objPlayer.CricSheetID.String
		finalPlayer.DateAdded = objPlayer.DateAdded.String
		finalPlayer.Status = objPlayer.Status.Int64

		finalPlayers = append(finalPlayers, finalPlayer)
	}

	final := utils.JSONMessageWrappedObj(http.StatusOK, finalPlayers)
	utils.WebResponseJSONObject(w, r, http.StatusOK, final)
}

func GetMissingTeamDetails(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	objTeams := data.GetTeamDetails()

	var finalTeams []models.TeamDetailsExt
	for _, objTeam := range objTeams {

		var finalTeam models.TeamDetailsExt

		finalTeam.TeamID = objTeam.TeamID.Int64
		finalTeam.TeamName = objTeam.TeamName.String
		finalTeam.TeamType = objTeam.TeamType.String
		finalTeam.FilterName = objTeam.FilterName.String
		finalTeam.ABBR = objTeam.ABBR.String
		finalTeam.TeamColor = objTeam.TeamColor.String
		finalTeam.Icon = objTeam.Icon.String
		finalTeam.URL = objTeam.URL.String
		finalTeam.Jersey = objTeam.Jersey.String
		finalTeam.Flag = objTeam.Flag.String
		finalTeam.Status = objTeam.Status.Int64
		finalTeam.DateAdded = objTeam.DateAdded.String

		finalTeams = append(finalTeams, finalTeam)
	}

	final := utils.JSONMessageWrappedObj(http.StatusOK, finalTeams)
	utils.WebResponseJSONObject(w, r, http.StatusOK, final)
}

func GetMissingVenueDetails(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	objVenues := data.GetVenueDetails()

	var finalVenues []models.VenueDetailsExt
	for _, objVenue := range objVenues {

		var finalVenue models.VenueDetailsExt
		finalVenue.VenueID = objVenue.VenueID.Int64
		finalVenue.Venue = objVenue.Venue.String
		finalVenue.FilterName = objVenue.FilterName.String
		finalVenue.FriendlyName = objVenue.FriendlyName.String
		finalVenue.City = objVenue.City.String
		finalVenue.Country = objVenue.Country.String
		finalVenue.State = objVenue.State.String
		finalVenue.StateABBR = objVenue.StateABBR.String
		finalVenue.OfficialTeam = objVenue.OfficialTeam.String
		finalVenue.Capacity = objVenue.Capacity.String
		finalVenue.Dimension = objVenue.Dimension.String
		finalVenue.Opened = objVenue.Opened.String
		finalVenue.Description = objVenue.Description.String
		finalVenue.ShortName = objVenue.ShortName.String
		finalVenue.TimeZone = objVenue.TimeZone.String
		finalVenue.Weather = objVenue.Weather.String
		finalVenue.PitchType = objVenue.PitchType.String
		finalVenue.DateAdded = objVenue.DateAdded.String
		finalVenue.Status = objVenue.Status.Int64
		finalVenues = append(finalVenues, finalVenue)
	}

	final := utils.JSONMessageWrappedObj(http.StatusOK, finalVenues)
	utils.WebResponseJSONObject(w, r, http.StatusOK, final)
}

func MapVenueDetails(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var finalVenue models.VenueDetailsExt
	venueIDInt, _ := strconv.ParseInt(r.FormValue("venue_id"), 10, 64)
	finalVenue.VenueID = venueIDInt
	finalVenue.Venue = r.FormValue("venue_name")
	finalVenue.FilterName = r.FormValue("filter_name")
	finalVenue.FriendlyName = r.FormValue("friendly_name")
	finalVenue.City = r.FormValue("city")
	finalVenue.Country = r.FormValue("country")
	finalVenue.State = r.FormValue("state")
	finalVenue.StateABBR = r.FormValue("state_abbr")
	finalVenue.OfficialTeam = r.FormValue("official_team")
	finalVenue.Capacity = r.FormValue("capacity")
	finalVenue.Dimension = r.FormValue("dimension")
	finalVenue.Opened = r.FormValue("opened")
	finalVenue.Description = r.FormValue("description")
	finalVenue.ShortName = r.FormValue("shortname")
	finalVenue.TimeZone = r.FormValue("timezone")
	finalVenue.Weather = r.FormValue("weather")
	finalVenue.PitchType = r.FormValue("pitch_type")
	finalVenue.DateAdded = time.Now().Format("2006-01-02 15:04:05")
	venueStatus, _ := strconv.ParseInt(r.FormValue("status"), 10, 64)
	finalVenue.Status = venueStatus

	data.InsertVenueDetails(finalVenue)
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}

func MapPlayerDetails(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var finalPlayer models.PlayerDetailsExt

	playerIDInt, _ := strconv.ParseInt(r.FormValue("player_id"), 10, 64)
	finalPlayer.PlayerID = playerIDInt
	finalPlayer.PlayerName = r.FormValue("player_name")
	finalPlayer.DisplayName = r.FormValue("display_name")
	finalPlayer.FirstName = r.FormValue("first_name")
	finalPlayer.LastName = r.FormValue("last_name")
	finalPlayer.ShortName = r.FormValue("short_name")
	finalPlayer.UniqueShortName = r.FormValue("unique_short_name")
	dobFormat, _ := time.Parse("2006-01-02", r.FormValue("dob"))
	finalPlayer.DOB = dobFormat.Format("2006-01-02")
	battingStyleInt, _ := strconv.ParseInt(r.FormValue("batting_style"), 10, 64)
	bowlingStyleInt, _ := strconv.ParseInt(r.FormValue("bowling_style"), 10, 64)
	finalPlayer.BattingStyle = battingStyleInt
	finalPlayer.BowlingStyle = bowlingStyleInt
	overseasInt, _ := strconv.ParseInt(r.FormValue("is_overseas"), 10, 64)
	finalPlayer.IsOverseas = overseasInt
	finalPlayer.CricSheetID = r.FormValue("cricsheet_id")
	finalPlayer.DateAdded = time.Now().Format("2006-01-02 15:04:05")
	statusInt, _ := strconv.ParseInt(r.FormValue("status"), 10, 64)
	finalPlayer.Status = statusInt

	data.InsertPlayerDetails(finalPlayer)
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}

func MapTeamDetails(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var finalTeam models.TeamDetailsExt

	teamIDInt, _ := strconv.ParseInt(r.FormValue(""), 10, 64)
	finalTeam.TeamID = teamIDInt
	finalTeam.TeamName = r.FormValue("")
	finalTeam.TeamType = r.FormValue("")
	finalTeam.FilterName = r.FormValue("")
	finalTeam.ABBR = r.FormValue("")
	finalTeam.TeamColor = r.FormValue("")
	finalTeam.Icon = r.FormValue("")
	finalTeam.URL = r.FormValue("")
	finalTeam.Jersey = r.FormValue("")
	finalTeam.Flag = r.FormValue("")
	statusInt, _ := strconv.ParseInt(r.FormValue(""), 10, 64)
	finalTeam.Status = statusInt
	finalTeam.DateAdded = r.FormValue("")

	data.InsertTeamDetails(finalTeam)
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)

}
