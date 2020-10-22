CREATE TABLE `accounts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `document_number` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `document_number_UNIQUE` (`document_number`)
) ENGINE=InnoDB;

CREATE TABLE `operation_types` (
  `id` int NOT NULL AUTO_INCREMENT,
  `description` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

CREATE TABLE `transactions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `account_id` int DEFAULT NULL,
  `operation_type_id` int DEFAULT NULL,
  `amount` int DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `account_fk` (`account_id`),
  KEY `operation_type_fk` (`operation_type_id`),
  CONSTRAINT `account_fk` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`),
  CONSTRAINT `operation_type_fk` FOREIGN KEY (`operation_type_id`) REFERENCES `operation_types` (`id`)
) ENGINE=InnoDB;

INSERT INTO `operation_types` (`id`, `description`) VALUES (1, 'COMPRA A VISTA');
INSERT INTO `operation_types` (`id`, `description`) VALUES (2, 'COMPRA PARCELADA');
INSERT INTO `operation_types` (`id`, `description`) VALUES (3, 'SAQUE');
INSERT INTO `operation_types` (`id`, `description`) VALUES (4, 'PAGAMENTO');