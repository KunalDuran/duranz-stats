

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

