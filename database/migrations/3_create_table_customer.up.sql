CREATE TABLE IF NOT EXISTS `customers` (
    `customer_id` INT NOT NULL AUTO_INCREMENT,
    `email` VARCHAR(200) NOT NULL,
    `name` VARCHAR(200) DEFAULT NULL,
    `created_at` DATETIME NOT NULL,
    PRIMARY KEY (`customer_id`),
    UNIQUE KEY `email` (`email`)

)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin;