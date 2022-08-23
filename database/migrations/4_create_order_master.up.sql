CREATE TABLE IF NOT EXISTS `orders`(
    `order_id` INT NOT NULL AUTO_INCREMENT,
    `grand_total` DECIMAL(18,2) DEFAULT NULL,
    `created_at` DATETIME NOT NULL,
    `created_by` INT NOT NULL,
    PRIMARY KEY (`order_id`), 
    KEY `created_by` (`created_by`),
    CONSTRAINT `fk_orders_created_by_users_user_id` FOREIGN KEY (`created_by`) REFERENCES `customers` (`customer_id`)   
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin;