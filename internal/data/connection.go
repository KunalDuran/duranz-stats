package data

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Setup(dbName string) error {
	if !databaseExists(Db, dbName) {
		if err := createDatabase(Db, dbName); err != nil {
			return err
		}
	}

	if _, err := Db.Exec("USE " + dbName); err != nil {
		return err
	}

	tableNames := []string{"users", "weather_history"}
	for _, tableName := range tableNames {
		if !tableExists(Db, dbName, tableName) {
			if err := createTable(Db, tableName); err != nil {
				return err
			}
		}
	}

	return nil
}

func databaseExists(db *sql.DB, dbName string) bool {
	var exists string
	query := "SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?"
	err := db.QueryRow(query, dbName).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return exists == dbName
}

func createDatabase(db *sql.DB, dbName string) error {
	_, err := db.Exec("CREATE DATABASE " + dbName)
	return err
}

func tableExists(db *sql.DB, dbName, tableName string) bool {
	var exists bool
	query := "SELECT 1 FROM information_schema.tables WHERE table_schema = ? AND table_name = ? LIMIT 1"
	err := db.QueryRow(query, dbName, tableName).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return exists
}

func createTable(db *sql.DB, tableName string) error {

	var query string
	switch tableName {
	case "users":
		query = `
		CREATE TABLE users (
			id INT PRIMARY KEY AUTO_INCREMENT,
			username VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			date_of_birth DATE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		) ENGINE=InnoDB;
		`
	case "weather_history":
		query = `
		CREATE TABLE weather_history (
		  id INT NOT NULL AUTO_INCREMENT,
		  city_name VARCHAR(255) NOT NULL,
		  user_id INT NOT NULL,
		  coord_lon FLOAT NOT NULL,
		  coord_lat FLOAT NOT NULL,
		  weather_id INT NOT NULL,
		  weather_main VARCHAR(255) NOT NULL,
		  weather_description VARCHAR(255) NOT NULL,
		  weather_icon VARCHAR(255) NOT NULL,
		  base VARCHAR(255) NOT NULL,
		  temp FLOAT NOT NULL,
		  feels_like FLOAT NOT NULL,
		  temp_min FLOAT NOT NULL,
		  temp_max FLOAT NOT NULL,
		  pressure INT NOT NULL,
		  humidity INT NOT NULL,
		  visibility INT NOT NULL,
		  wind_speed FLOAT NOT NULL,
		  wind_deg INT NOT NULL,
		  clouds_all INT NOT NULL,
		  dt INT NOT NULL,
		  sys_type INT NOT NULL,
		  sys_id INT NOT NULL,
		  sys_country VARCHAR(255) NOT NULL,
		  sys_sunrise INT NOT NULL,
		  sys_sunset INT NOT NULL,
		  timezone INT NOT NULL,
		  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		  PRIMARY KEY (id),
		  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
		) ENGINE=InnoDB;`

	}

	_, err := db.Exec(query)
	return err
}

func Delete() error {
	_, err := Db.Exec("DROP DATABASE IF EXISTS duranz")
	return err
}

func CreateAllTables() {
	duranzMatches := "CREATE TABLE `duranz_cricket_matches` ( `match_id` INT(10) NOT NULL AUTO_INCREMENT, `league_id` INT(10) NULL DEFAULT NULL, `gender` VARCHAR(20) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `season_id` INT(10) NULL DEFAULT NULL, `home_team_id` INT(10) NULL DEFAULT NULL, `away_team_id` INT(10) NULL DEFAULT NULL, `home_team_name` VARCHAR(120) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `away_team_name` VARCHAR(120) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `venue_id` INT(10) NULL DEFAULT NULL, `result` VARCHAR(200) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `man_of_the_match` VARCHAR(100) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `toss_winner` INT(10) NULL DEFAULT NULL, `toss_decision` VARCHAR(20) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `winning_team` INT(10) NULL DEFAULT NULL, `cricsheet_file_name` VARCHAR(20) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `match_date` DATE NULL DEFAULT NULL, `match_date_multi` VARCHAR(120) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `match_time` TIME NULL DEFAULT NULL, `is_reschedule` SMALLINT(5) NULL DEFAULT '0', `is_abandoned` SMALLINT(5) NULL DEFAULT '0', `is_neutral` SMALLINT(5) NULL DEFAULT '0', `match_refrees` VARCHAR(100) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `reserve_umpires` VARCHAR(100) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `tv_umpires` VARCHAR(100) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `umpires` VARCHAR(100) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `date_added` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP, `last_update` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP, `match_end_time` DATETIME NULL DEFAULT NULL, `status` VARCHAR(2) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', PRIMARY KEY (`match_id`) USING BTREE ) COLLATE='utf8_unicode_ci' ENGINE=InnoDB ;"
	duranzPlayers := "CREATE TABLE `duranz_cricket_players` ( `player_id` INT(10) NOT NULL AUTO_INCREMENT, `player_name` VARCHAR(200) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `display_name` VARCHAR(200) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `first_name` VARCHAR(100) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `last_name` VARCHAR(100) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `short_name` VARCHAR(100) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `unique_short_name` VARCHAR(100) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `dob` DATE NULL DEFAULT NULL, `batting_style_1_id` INT(10) NULL DEFAULT NULL, `bowling_style_1_id` INT(10) NULL DEFAULT NULL, `is_overseas` TINYINT(3) NULL DEFAULT '0', `cricsheet_id` VARCHAR(200) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `date_added` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP, `status` SMALLINT(5) NULL DEFAULT '1', PRIMARY KEY (`player_id`) USING BTREE, UNIQUE INDEX `cricsheet_id` (`cricsheet_id`) USING BTREE ) COLLATE='utf8_unicode_ci' ENGINE=InnoDB ; "
	duranzErrorlog := "CREATE TABLE `duranz_errorlog` ( `alert_id` VARCHAR(11) NOT NULL COLLATE 'utf8_general_ci', `error_msg` VARCHAR(250) NOT NULL COLLATE 'utf8_general_ci', `file_name` VARCHAR(100) NOT NULL COLLATE 'utf8_general_ci', `dateadded` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY (`alert_id`, `error_msg`) USING BTREE ) COLLATE='utf8_general_ci' ENGINE=InnoDB ; "
	duranzMappings := "CREATE TABLE `duranz_file_mappings` ( `file_name` VARCHAR(20) NOT NULL COLLATE 'utf8_general_ci', `league_id` INT(10) NULL DEFAULT NULL, `teams` INT(10) NULL DEFAULT NULL, `players` INT(10) NULL DEFAULT NULL, `venue` INT(10) NULL DEFAULT NULL, `matches` INT(10) NULL DEFAULT NULL, `match_stats` INT(10) NULL DEFAULT NULL, `player_stats` INT(10) NULL DEFAULT NULL, `dateadded` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY (`file_name`) USING BTREE ) COLLATE='utf8_general_ci' ENGINE=InnoDB ; "
	duranzMatchStats := "CREATE TABLE `duranz_match_stats` ( `match_id` INT(10) NOT NULL, `team_id` INT(10) NULL DEFAULT NULL, `innings` INT(10) NULL DEFAULT NULL, `fall_of_wickets` VARCHAR(400) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `extras` INT(10) NULL DEFAULT NULL, `score` INT(10) NULL DEFAULT NULL, `wickets` INT(10) NULL DEFAULT NULL, `overs_played` VARCHAR(11) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `super_over` INT(10) NULL DEFAULT NULL, `last_update` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ) COLLATE='utf8_unicode_ci' ENGINE=InnoDB ; "
	duranzPlayerStats := "CREATE TABLE `duranz_player_match_stats` ( `match_id` INT(10) NOT NULL, `season_id` VARCHAR(40) NOT NULL COLLATE 'utf8_unicode_ci', `innings_id` VARCHAR(40) NOT NULL COLLATE 'utf8_unicode_ci', `season_type` VARCHAR(40) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `team_id` INT(10) NOT NULL, `player_id` INT(10) NOT NULL, `batting_order` INT(10) NULL DEFAULT NULL, `runs_scored` INT(10) NULL DEFAULT NULL, `balls_faced` INT(10) NULL DEFAULT NULL, `dot_balls_played` INT(10) NULL DEFAULT NULL, `singles` INT(10) NULL DEFAULT NULL, `doubles` INT(10) NULL DEFAULT NULL, `triples` INT(10) NULL DEFAULT NULL, `fours_hit` INT(10) NULL DEFAULT NULL, `sixes_hit` INT(10) NULL DEFAULT NULL, `out_type` VARCHAR(40) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `out_fielder` VARCHAR(40) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `out_bowler` VARCHAR(40) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `is_batted` INT(10) NULL DEFAULT NULL, `overs_bowled` VARCHAR(16) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `bowling_order` INT(10) NULL DEFAULT NULL, `runs_conceded` INT(10) NULL DEFAULT NULL, `balls_bowled` INT(10) NULL DEFAULT NULL, `dots_bowled` INT(10) NULL DEFAULT NULL, `wickets_taken` INT(10) NULL DEFAULT NULL, `fours_conceded` INT(10) NULL DEFAULT NULL, `sixes_conceded` INT(10) NULL DEFAULT NULL, `extras_conceded` INT(10) NULL DEFAULT NULL, `maiden_over` INT(10) NULL DEFAULT NULL, `run_out` INT(10) NULL DEFAULT NULL, `catches` INT(10) NULL DEFAULT NULL, `stumpings` INT(10) NULL DEFAULT NULL, `played_abandoned_matches` INT(10) NULL DEFAULT NULL, `last_update` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ) COLLATE='utf8_unicode_ci' ENGINE=InnoDB ; "
	duranzTeams := "CREATE TABLE `duranz_teams` ( `team_id` INT(10) NOT NULL AUTO_INCREMENT, `team_name` VARCHAR(120) NOT NULL COLLATE 'utf8_general_ci', `team_type` VARCHAR(120) NOT NULL DEFAULT '0' COLLATE 'utf8_unicode_ci', `filtername` VARCHAR(120) NULL DEFAULT NULL COLLATE 'utf8_general_ci', `abbreviation` VARCHAR(4) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `team_color` VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `icon` VARCHAR(200) NULL DEFAULT NULL COLLATE 'utf8_general_ci', `url` VARCHAR(100) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `jersey` VARCHAR(100) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `flag` VARCHAR(200) NULL DEFAULT NULL COLLATE 'utf8_general_ci', `status` SMALLINT(5) NOT NULL DEFAULT '1', `dateadded` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY (`team_id`) USING BTREE ) COLLATE='utf8_unicode_ci' ENGINE=InnoDB ; "
	duranzVenues := "CREATE TABLE `duranz_venue` ( `venue_id` INT(10) NOT NULL AUTO_INCREMENT, `venue` VARCHAR(200) NOT NULL COLLATE 'utf8_unicode_ci', `filtername` VARCHAR(200) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `friendlyname` VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `city` VARCHAR(200) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `country` VARCHAR(200) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `state` VARCHAR(200) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `state_abbr` VARCHAR(5) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `official_team` VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `capacity` VARCHAR(200) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `dimensions` VARCHAR(200) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `opened` VARCHAR(200) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `description` VARCHAR(5000) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `shortname` VARCHAR(200) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `timezone` VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8_general_ci', `weather` VARCHAR(100) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `pitch_type` VARCHAR(100) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci', `dateadded` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, `status` SMALLINT(5) NOT NULL DEFAULT '1', PRIMARY KEY (`venue_id`) USING BTREE ) COLLATE='utf8_unicode_ci' ENGINE=InnoDB ; "

	allTables := []string{duranzMatches, duranzPlayers, duranzErrorlog, duranzMappings, duranzMatchStats, duranzPlayerStats, duranzTeams, duranzVenues}

	for _, tableCode := range allTables {
		_, err := Db.Exec(tableCode)
		if err != nil {
			log.Fatal("Error in creating Tables ", err)
		}
	}
}
