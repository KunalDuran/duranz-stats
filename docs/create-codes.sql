CREATE TABLE `cricket_matches` (
    `match_id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `league_id` INT,
    `gender` ENUM('Male', 'Female'),
    `season_id` INT,
    `home_team_id` INT,
    `away_team_id` INT,
    `home_team_name` VARCHAR(120),
    `away_team_name` VARCHAR(120),
    `venue_id` INT,
    `result` VARCHAR(200),
    `man_of_the_match` VARCHAR(100),
    `toss_winner` INT,
    `toss_decision` ENUM('Bat', 'Field'),
    `winning_team` INT,
    `cricsheet_file_name` VARCHAR(20),
    `match_date` DATE,
    `match_date_multi` VARCHAR(120),
    `match_time` TIME,
    `is_reschedule` BOOLEAN DEFAULT FALSE,
    `is_abandoned` BOOLEAN DEFAULT FALSE,
    `is_neutral` BOOLEAN DEFAULT FALSE,
    `match_refrees` VARCHAR(100),
    `reserve_umpires` VARCHAR(100),
    `tv_umpires` VARCHAR(100),
    `umpires` VARCHAR(100),
    `date_added` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `last_update` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `match_end_time` DATETIME,
    `status` CHAR(2)
) ENGINE = InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;




CREATE TABLE `cricket_players` (
    `player_id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `player_name` VARCHAR(200) COLLATE utf8_unicode_ci,
    `display_name` VARCHAR(200) COLLATE utf8_unicode_ci,
    `first_name` VARCHAR(100) COLLATE utf8_unicode_ci,
    `last_name` VARCHAR(100) COLLATE utf8_unicode_ci,
    `short_name` VARCHAR(100) COLLATE utf8_unicode_ci,
    `unique_short_name` VARCHAR(100) COLLATE utf8_unicode_ci,
    `dob` DATE,
    `batting_style_1_id` INT,
    `bowling_style_1_id` INT,
    `is_overseas` BOOLEAN DEFAULT FALSE,
    `cricsheet_id` VARCHAR(200) COLLATE utf8_unicode_ci UNIQUE,
    `date_added` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `status` BOOLEAN DEFAULT TRUE
) ENGINE = InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


CREATE TABLE `teams` (
    `team_id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `team_name` VARCHAR(120) NOT NULL COLLATE 'utf8_unicode_ci',
    `team_type` VARCHAR(50) NOT NULL COLLATE 'utf8_unicode_ci',
    `filtername` VARCHAR(120) COLLATE 'utf8_unicode_ci',
    `abbreviation` VARCHAR(4) COLLATE 'utf8_unicode_ci',
    `team_color` VARCHAR(50) COLLATE 'utf8_unicode_ci',
    `icon` VARCHAR(200) COLLATE 'utf8_unicode_ci',
    `url` VARCHAR(100) COLLATE 'utf8_unicode_ci',
    `jersey` VARCHAR(100) COLLATE 'utf8_unicode_ci',
    `flag` VARCHAR(200) COLLATE 'utf8_unicode_ci',
    `status` BOOLEAN NOT NULL DEFAULT TRUE,
    `dateadded` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE = InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


CREATE TABLE `venue` (
    `venue_id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `venue_name` VARCHAR(200) NOT NULL COLLATE 'utf8_unicode_ci',
    `filtername` VARCHAR(200) COLLATE 'utf8_unicode_ci',
    `friendlyname` VARCHAR(50) COLLATE 'utf8_unicode_ci',
    `city` VARCHAR(200) COLLATE 'utf8_unicode_ci',
    `country` VARCHAR(200) COLLATE 'utf8_unicode_ci',
    `state` VARCHAR(200) COLLATE 'utf8_unicode_ci',
    `state_abbr` VARCHAR(5) COLLATE 'utf8_unicode_ci',
    `official_team` VARCHAR(50) COLLATE 'utf8_unicode_ci',
    `capacity` INT,
    `dimensions` VARCHAR(200) COLLATE 'utf8_unicode_ci',
    `opened` YEAR,
    `description` VARCHAR(5000) COLLATE 'utf8_unicode_ci',
    `shortname` VARCHAR(200) COLLATE 'utf8_unicode_ci',
    `timezone` VARCHAR(50) COLLATE 'utf8_unicode_ci',
    `weather` VARCHAR(100) COLLATE 'utf8_unicode_ci',
    `pitch_type` VARCHAR(100) COLLATE 'utf8_unicode_ci',
    `dateadded` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `status` BOOLEAN NOT NULL DEFAULT TRUE
) ENGINE = InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


CREATE TABLE `match_stats` (
    `match_id` INT NOT NULL,
    `team_id` INT NOT NULL,
    `innings` INT NOT NULL,
    `fall_of_wickets` VARCHAR(400) COLLATE 'utf8_unicode_ci',
    `extras` INT DEFAULT NULL,
    `score` INT DEFAULT NULL,
    `wickets` INT DEFAULT NULL,
    `overs_played` DECIMAL(5,2) DEFAULT NULL,
    `super_over` INT DEFAULT NULL,
    `last_update` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`match_id`, `team_id`, `innings`),
    FOREIGN KEY (`match_id`) REFERENCES `cricket_matches`(`match_id`),
    FOREIGN KEY (`team_id`) REFERENCES `teams`(`team_id`)
) ENGINE = InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



CREATE TABLE `player_match_stats` (
    `match_id` INT NOT NULL,
    `season_id` VARCHAR(40) NOT NULL,
    `innings_id` VARCHAR(40) NOT NULL,
    `season_type` VARCHAR(40),
    `team_id` INT NOT NULL,
    `player_id` INT NOT NULL,
    `batting_order` INT,
    `runs_scored` INT,
    `balls_faced` INT,
    `dot_balls_played` INT,
    `singles` INT NULL DEFAULT NULL,
    `doubles` INT NULL DEFAULT NULL,
    `triples` INT NULL DEFAULT NULL,
    `fours_hit` INT NULL DEFAULT NULL,
    `sixes_hit` INT NULL DEFAULT NULL,
    `out_type` VARCHAR(40) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci',
    `out_fielder` VARCHAR(40) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci',
    `out_bowler` VARCHAR(40) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci',
    `is_batted` INT NULL DEFAULT NULL,
    `overs_bowled` VARCHAR(16) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci',
    `bowling_order` INT NULL DEFAULT NULL,
    `runs_conceded` INT NULL DEFAULT NULL,
    `balls_bowled` INT NULL DEFAULT NULL,
    `dots_bowled` INT NULL DEFAULT NULL,
    `wickets_taken` INT NULL DEFAULT NULL,
    `fours_conceded` INT NULL DEFAULT NULL,
    `sixes_conceded` INT NULL DEFAULT NULL,
    `extras_conceded` INT NULL DEFAULT NULL,
    `maiden_over` INT NULL DEFAULT NULL,
    `run_out` INT NULL DEFAULT NULL,
    `catches` INT NULL DEFAULT NULL,
    `stumpings` INT NULL DEFAULT NULL,
    `played_abandoned_matches` INT,
    `last_update` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`match_id`, `team_id`, `player_id`),
    FOREIGN KEY (`match_id`) REFERENCES `cricket_matches`(`match_id`),
    FOREIGN KEY (`team_id`) REFERENCES `teams`(`team_id`),
    FOREIGN KEY (`player_id`) REFERENCES `cricket_players`(`player_id`)
) ENGINE = InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


CREATE TABLE `errorlog` (
    `id` INT AUTO_INCREMENT NOT NULL,
    `alert_id` VARCHAR(36) NOT NULL COLLATE 'utf8_general_ci', 
    `error_msg` VARCHAR(250) NOT NULL COLLATE 'utf8_general_ci',
    `file_name` VARCHAR(100) NOT NULL COLLATE 'utf8_general_ci',
    `error_type` VARCHAR(100) COLLATE 'utf8_general_ci', 
    `severity` VARCHAR(50) COLLATE 'utf8_general_ci', 
    `dateadded` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE INDEX (`alert_id`)
) COLLATE = 'utf8_general_ci' ENGINE = InnoDB;



CREATE TABLE `file_mappings` (
    `file_id` INT AUTO_INCREMENT NOT NULL,  
    `file_name` VARCHAR(255) NOT NULL COLLATE 'utf8_general_ci',  
    `league_id` INT NULL DEFAULT NULL,
    `teams` INT NULL DEFAULT NULL,
    `players` INT NULL DEFAULT NULL,
    `venue` INT NULL DEFAULT NULL,
    `matches` INT NULL DEFAULT NULL,
    `match_stats` INT NULL DEFAULT NULL,
    `player_stats` INT NULL DEFAULT NULL,
    `over_stats` INT NULL DEFAULT NULL,  
    `dateadded` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`file_id`),
    UNIQUE (`file_name`)
) COLLATE = 'utf8_general_ci' ENGINE = InnoDB;

