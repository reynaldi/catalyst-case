CREATE TABLE IF NOT EXISTS `products` (
    `product_id` INT NOT NULL AUTO_INCREMENT,
    `brand_id` INT NOT NULL,
    `product_name` VARCHAR(200) NOT NULL,
    `price` DECIMAL(18,2) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `deleted_at` DATETIME DEFAULT NULL,    
    PRIMARY KEY (`product_id`),
    KEY `brand_id` (`brand_id`),
    CONSTRAINT `fk_products_brand_id_brands_id` FOREIGN KEY (`brand_id`) REFERENCES `brands` (`brand_id`)
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin;