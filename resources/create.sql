CREATE TABLE `ecommerce`.`broad_category` (
  `bcId` INT NOT NULL  AUTO_INCREMENT,
  `bcName` VARCHAR(45) NOT NULL,
  `bcDescription` VARCHAR(45) NULL,
  `bcImageUrl` VARCHAR(200) NULL,
  `bcIsActive` TINYINT NULL DEFAULT 1,
  PRIMARY KEY (`bcId`)
) AUTO_INCREMENT=1000;



INSERT INTO `ecommerce`.`broad_category`
(
`bcName`,
`bcDescription`)
VALUES
("Fruits",
"These are juicy fruits");


CREATE TABLE `ecommerce`.`sub_category` (
  `scId` INT NOT NULL  AUTO_INCREMENT,
  `scName` VARCHAR(45) NULL,
  `scDescription` VARCHAR(45) NULL,
  `scImageUrl` VARCHAR(200) NULL,
  `scIsActive` TINYINT NULL DEFAULT 1,
  `bcId` INT NOT NULL,
  PRIMARY KEY (`scId`),
  CONSTRAINT `bcId`
    FOREIGN KEY (`bcId`)
    REFERENCES `ecommerce`.`broad_category` (`bcId`)
    ON DELETE CASCADE
    ON UPDATE CASCADE) AUTO_INCREMENT=2000;

INSERT INTO `ecommerce`.`sub_category`
(
`scName`,
`scDescription`,
`bcId`)
VALUES
("Apple",
"This is an apple",
1000);


INSERT INTO `ecommerce`.`sub_category`
(
`scName`,
`scDescription`,
`bcId`)
VALUES
("Oranges",
"These are oranges",
1000);


INSERT INTO `ecommerce`.`sub_category`
(
`scName`,
`scDescription`,
`bcId`)
VALUES
("Bananas",
"These are Bananas",
1000);


INSERT INTO `ecommerce`.`sub_category`
(
`scName`,
`scDescription`,
`bcId`)
VALUES
("Pears",
"These are Pears",
1000);


CREATE TABLE `ecommerce`.`product` (
  `productId` INT NOT NULL  AUTO_INCREMENT,
  `name` VARCHAR(50) NOT NULL,
  `sku` VARCHAR(45) NULL,
  `description` VARCHAR(500) NULL,
  `bcId` INT NOT NULL,
  `currency` VARCHAR(45) NOT NULL,
  `unitsInStock` INT NOT NULL,
  `imageUrl` VARCHAR(200) NULL,
  `discountPercentage` DECIMAL(2) NULL DEFAULT 0,
  `unitPrice` DECIMAL(2) NOT NULL,
  `scId` INT NOT NULL,
  PRIMARY KEY (`productId`),
  CONSTRAINT `scId_p`
    FOREIGN KEY (`scId`)
    REFERENCES `ecommerce`.`sub_category` (`scId`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `bcId_p`
    FOREIGN KEY (`bcId`)
    REFERENCES `ecommerce`.`broad_category` (`bcId`)
    ON DELETE CASCADE
    ON UPDATE CASCADE);

INSERT INTO `ecommerce`.`product` (`name`, `description`, `bcId`, `currency`, `unitsInStock`, `imageUrl`, `unitPrice`, `scId`) VALUES ('Red Apple', 'This is a red apply', '1000', 'Rs', '16', 'https://raw.githubusercontent.com/shashankvivek/e-food-client/master/e-food/src/assets/apple-1.png', '10', '2000');

INSERT INTO `ecommerce`.`product` (`name`, `description`, `bcId`, `currency`, `unitsInStock`, `imageUrl`, `unitPrice`, `scId`) VALUES ('Green Apple', 'This is a Green Apple', '1000', 'Rs', '4', 'https://raw.githubusercontent.com/shashankvivek/e-food-client/master/e-food/src/assets/apple-3.jpg', '15', '2000');

INSERT INTO `ecommerce`.`product` (`name`, `description`, `bcId`, `currency`, `unitsInStock`, `imageUrl`, `unitPrice`, `scId`) VALUES ('Banana', 'This is a Banana', '1000', 'Rs', '20', 'https://raw.githubusercontent.com/shashankvivek/e-food-client/master/e-food/src/assets/banana.jpg', '10', '2002');

INSERT INTO `ecommerce`.`product` (`name`, `description`, `bcId`, `currency`, `unitsInStock`, `imageUrl`, `unitPrice`, `scId`) VALUES ('Oranges', 'This is an Orange', '1000', 'Rs', '5', 'https://raw.githubusercontent.com/shashankvivek/e-food-client/master/e-food/src/assets/oranges.jpg', '20', '2001');

INSERT INTO `ecommerce`.`product` (`name`, `description`, `bcId`, `currency`, `unitsInStock`, `imageUrl`, `unitPrice`, `scId`) VALUES ('Pear', 'This is a Pear', '1000', 'Rs', '10', 'https://raw.githubusercontent.com/shashankvivek/e-food-client/master/e-food/src/assets/pears.jpg', '30', '2003');
