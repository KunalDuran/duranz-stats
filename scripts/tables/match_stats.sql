
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
