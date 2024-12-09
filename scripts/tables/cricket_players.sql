
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
