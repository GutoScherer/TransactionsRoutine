CREATE TABLE `Accounts` (
  `Account_ID` int NOT NULL AUTO_INCREMENT,
  `Document_Number` int DEFAULT NULL,
  PRIMARY KEY (`Account_ID`),
  UNIQUE KEY `document_number_UNIQUE` (`Document_Number`)
) ENGINE=InnoDB DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci;

CREATE TABLE `OperationTypes` (
  `OperationType_ID` int NOT NULL AUTO_INCREMENT,
  `Description` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`OperationType_ID`)
) ENGINE=InnoDB DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci;

CREATE TABLE `Transactions` (
  `Transaction_ID` int NOT NULL,
  `Account_ID` int DEFAULT NULL,
  `OperationType_ID` int DEFAULT NULL,
  `Amount` int DEFAULT NULL,
  `EventDate` datetime DEFAULT NULL,
  PRIMARY KEY (`Transaction_ID`),
  KEY `OperationType_FK_idx` (`OperationType_ID`),
  KEY `Account_FK` (`Account_ID`),
  CONSTRAINT `Account_FK` FOREIGN KEY (`Account_ID`) REFERENCES `Accounts` (`Account_ID`),
  CONSTRAINT `OperationType_FK` FOREIGN KEY (`OperationType_ID`) REFERENCES `OperationTypes` (`OperationType_ID`)
) ENGINE=InnoDB DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci;

INSERT INTO `OperationTypes` (`OperationType_ID`, `Description`) VALUES (1, 'COMPRA A VISTA');
INSERT INTO `OperationTypes` (`OperationType_ID`, `Description`) VALUES (2, 'COMPRA PARCELADA');
INSERT INTO `OperationTypes` (`OperationType_ID`, `Description`) VALUES (3, 'SAQUE');
INSERT INTO `OperationTypes` (`OperationType_ID`, `Description`) VALUES (4, 'PAGAMENTO');