

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

