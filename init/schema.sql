CREATE TABLE `accounts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `document_number` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `document_number_UNIQUE` (`document_number`)
) ENGINE=InnoDB;

CREATE TABLE `transactions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `account_id` int DEFAULT NULL,
  `operation_type_id` enum('1','2','3','4') DEFAULT NULL,
  `amount` int DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `account_fk` (`account_id`),
  CONSTRAINT `account_fk` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`)
) ENGINE=InnoDB;