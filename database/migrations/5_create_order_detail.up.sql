CREATE TABLE IF NOT EXISTS `order_details` (
    `order_detail_id` INT NOT NULL AUTO_INCREMENT,
    `order_id` INT NOT NULL,
    `product_id` INT NOT NULL,
    `product_name` VARCHAR(200) NOT NULL,
    `unit_price` DECIMAL(18,2) NOT NULL,
    `qty` INT NOT NULL,
    PRIMARY KEY (`order_detail_id`),
    KEY `order_id` (`order_id`),
    CONSTRAINT `fk_order_details_orders_order_id` FOREIGN KEY (`order_id`) REFERENCES `orders` (`order_id`)    
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin;