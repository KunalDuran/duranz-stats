
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
