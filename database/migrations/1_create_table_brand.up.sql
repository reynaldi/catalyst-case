CREATE TABLE IF NOT EXISTS `brands` (
    `brand_id` INT NOT NULL AUTO_INCREMENT,
    `brand_name` VARCHAR(200) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`brand_id`)    
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin;