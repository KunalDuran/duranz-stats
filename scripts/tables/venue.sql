
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

