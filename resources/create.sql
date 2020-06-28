CREATE TABLE `ecommerce`.`broad_category` (
  `bcId` MEDIUMINT NOT NULL  AUTO_INCREMENT,
  `bcName` VARCHAR(45) NOT NULL,
  `bcDescription` VARCHAR(45) NULL,
  `bcImageUrl` VARCHAR(200) NULL,
  `bcIsActive` TINYINT NULL DEFAULT 1,
  PRIMARY KEY (`bcId`)
);



INSERT INTO `ecommerce`.`broad_category`
(
`bcName`,
`bcDescription`)
VALUES
("Fruits",
"These are juicy fruits");


CREATE TABLE `ecommerce`.`sub_category` (
  `scId` MEDIUMINT NOT NULL  AUTO_INCREMENT,
  `scName` VARCHAR(45) NULL,
  `scDescription` VARCHAR(45) NULL,
  `scImageUrl` VARCHAR(200) NULL,
  `scIsActive` TINYINT NULL DEFAULT 1,
  `bcId` MEDIUMINT NOT NULL,
  FOREIGN KEY (`bcId`) REFERENCES `ecommerce`.`broad_category`(`bcId`),
  PRIMARY KEY (`scId`)
);

INSERT INTO `ecommerce`.`sub_category`
(
`scName`,
`scDescription`,
`bcId`)
VALUES
("Apple",
"This is an apple",
1);


INSERT INTO `ecommerce`.`sub_category`
(
`scName`,
`scDescription`,
`bcId`)
VALUES
("Oranges",
"These are oranges",
1);


INSERT INTO `ecommerce`.`sub_category`
(
`scName`,
`scDescription`,
`bcId`)
VALUES
("Bananas",
"These are Bananas",
1);


INSERT INTO `ecommerce`.`sub_category`
(
`scName`,
`scDescription`,
`bcId`)
VALUES
("Pears",
"These are Pears",
1);

