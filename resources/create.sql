CREATE TABLE `ecommerce`.`broad_category` (
  `bcId` INT NOT NULL  AUTO_INCREMENT,
  `bcName` VARCHAR(45) NOT NULL,
  `bcDescription` VARCHAR(45) NULL,
  `bcImageUrl` VARCHAR(200) NULL,
  `bcIsActive` TINYINT NULL DEFAULT 1,
  PRIMARY KEY (`bcId`)
) AUTO_INCREMENT=1000;



INSERT INTO `ecommerce`.`broad_category`
(`bcName`,
 `bcDescription`)
VALUES ("Fruits",
        "These are juicy fruits");


CREATE TABLE `ecommerce`.`sub_category`
(
    `scId`          INT          NOT NULL AUTO_INCREMENT,
    `scName`        VARCHAR(45)  NULL,
    `scDescription` VARCHAR(45)  NULL,
    `scImageUrl`    VARCHAR(200) NULL,
    `scIsActive`    TINYINT      NULL DEFAULT 1,
    `bcId`          INT          NOT NULL,
    PRIMARY KEY (`scId`),
    CONSTRAINT `bcId`
        FOREIGN KEY (`bcId`)
            REFERENCES `ecommerce`.`broad_category` (`bcId`)
) AUTO_INCREMENT = 2000;

INSERT INTO `ecommerce`.`sub_category`
    (`scName`, `scDescription`, `bcId`)
VALUES ("Apple", "This is an apple", 1000);


INSERT INTO `ecommerce`.`sub_category`
    (`scName`, `scDescription`, `bcId`)
VALUES ("Oranges", "These are oranges", 1000);


INSERT INTO `ecommerce`.`sub_category`
    (`scName`, `scDescription`, `bcId`)
VALUES ("Bananas", "These are Bananas", 1000);


INSERT INTO `ecommerce`.`sub_category`
(`scName`,
 `scDescription`,
 `bcId`)
VALUES ("Pears",
        "These are Pears",
        1000);

-- PRODUCT TABLE
CREATE TABLE `ecommerce`.`product`
(
    `productId`          INT          NOT NULL AUTO_INCREMENT,
    `name`               VARCHAR(50)  NOT NULL,
    `sku`                VARCHAR(45)  NULL,
    `description`        VARCHAR(500) NULL,
    `bcId`               INT          NOT NULL,
    `currency`           VARCHAR(45)  NOT NULL,
    `unitsInStock`       INT          NOT NULL,
    `imageUrl`           VARCHAR(200) NULL,
    `discountPercentage` DECIMAL(2)   NULL DEFAULT 0,
    `unitPrice`          DECIMAL(2)   NOT NULL,
    `scId`               INT          NOT NULL,
    PRIMARY KEY (`productId`),
    CONSTRAINT `scId_p`
        FOREIGN KEY (`scId`)
            REFERENCES `ecommerce`.`sub_category` (`scId`),
    CONSTRAINT `bcId_p`
        FOREIGN KEY (`bcId`)
            REFERENCES `ecommerce`.`broad_category` (`bcId`)
);

INSERT INTO `ecommerce`.`product` (`name`, `description`, `bcId`, `currency`, `unitsInStock`, `imageUrl`, `unitPrice`,
                                   `scId`)
VALUES ('Red Apple', 'This is a red apply', '1000', 'Rs', '16',
        'https://raw.githubusercontent.com/shashankvivek/e-food-client/master/e-food/src/assets/apple-1.png', '10',
        '2000');

INSERT INTO `ecommerce`.`product` (`name`, `description`, `bcId`, `currency`, `unitsInStock`, `imageUrl`, `unitPrice`,
                                   `scId`)
VALUES ('Green Apple', 'This is a Green Apple', '1000', 'Rs', '4',
        'https://raw.githubusercontent.com/shashankvivek/e-food-client/master/e-food/src/assets/apple-3.jpg', '15',
        '2000');

INSERT INTO `ecommerce`.`product` (`name`, `description`, `bcId`, `currency`, `unitsInStock`, `imageUrl`, `unitPrice`, `scId`) VALUES ('Banana', 'This is a Banana', '1000', 'Rs', '20', 'https://raw.githubusercontent.com/shashankvivek/e-food-client/master/e-food/src/assets/banana.jpg', '10', '2002');

INSERT INTO `ecommerce`.`product` (`name`, `description`, `bcId`, `currency`, `unitsInStock`, `imageUrl`, `unitPrice`, `scId`) VALUES ('Oranges', 'This is an Orange', '1000', 'Rs', '5', 'https://raw.githubusercontent.com/shashankvivek/e-food-client/master/e-food/src/assets/oranges.jpg', '20', '2001');

INSERT INTO `ecommerce`.`product` (`name`, `description`, `bcId`, `currency`, `unitsInStock`, `imageUrl`, `unitPrice`, `scId`) VALUES ('Pear', 'This is a Pear', '1000', 'Rs', '10', 'https://raw.githubusercontent.com/shashankvivek/e-food-client/master/e-food/src/assets/pears.jpg', '30', '2003');

-- Creating Guest table for session id
CREATE TABLE `ecommerce`.`guest` (
                                     `sessionId` VARCHAR(40) NOT NULL,
                                     `extraInfo` VARCHAR(200) NULL,
                                     PRIMARY KEY (`sessionId`),
                                     UNIQUE INDEX `sessionId_UNIQUE` (`sessionId` ASC) VISIBLE);

-- create Guest cart Item

CREATE TABLE `ecommerce`.`guest_cart_item`
(
    `sessionId` VARCHAR(40) NOT NULL,
    `totalQty`  INT         NOT NULL,
    `productId` INT         NULL,
    INDEX `sessionId_idx` (`sessionId` ASC) VISIBLE,
    INDEX `productId_idx` (`productId` ASC) VISIBLE,
    CONSTRAINT `gsessionId`
        FOREIGN KEY (`sessionId`)
            REFERENCES `ecommerce`.`guest` (`sessionId`),
    CONSTRAINT `productId`
        FOREIGN KEY (`productId`)
            REFERENCES `ecommerce`.`product` (`productId`)
);


-- create user table, work on password field later

CREATE TABLE `ecommerce`.`customer`
(
    `customerId` INT          NOT NULL AUTO_INCREMENT,
    `firstName`  VARCHAR(45)  NOT NULL,
    `lastName`   VARCHAR(45)  NULL,
    `phoneNo`    INT          NULL,
    `password`   VARCHAR(100) NULL,
    `email`      VARCHAR(60)  NULL,
    PRIMARY KEY (`customerId`),
    UNIQUE INDEX `phoneNo_UNIQUE` (`phoneNo` ASC) VISIBLE,
    UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE
);

-- create coupon table
CREATE TABLE `ecommerce`.`coupons`
(
    `couponId`           VARCHAR(10)   NOT NULL,
    `expiryDate`         DATETIME      NULL,
    `RuleSet`            VARCHAR(1000) NULL,
    `discountPercentage` DECIMAL(2)    NULL,
    PRIMARY KEY (`couponId`)
);

-- create cart for customer

CREATE TABLE `ecommerce`.`cart`
(
    `cartId`          INT         NOT NULL AUTO_INCREMENT,
    `customerId`      INT         NOT NULL,
    `couponId`        VARCHAR(10) NULL,
    `totalCartPrice`  DECIMAL(2)  NOT NULL,
    `totalCartSaving` DECIMAL(2)  NOT NULL,
    `createdAt`       DATETIME    NULL,
    PRIMARY KEY (`cartId`),
    INDEX `cUserId_idx` (`customerId` ASC) VISIBLE,
    CONSTRAINT `cUserId`
        FOREIGN KEY (`customerId`)
            REFERENCES `ecommerce`.`customer` (`customerId`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
);


-- create User cart Item

CREATE TABLE `ecommerce`.`customer_cart_item`
(
    `cartId`      INT        NOT NULL,
    `totalQty`    INT        NOT NULL,
    `totalSaving` DECIMAL(2) NULL,
    `totalPrice`  DECIMAL(2) NOT NULL,
    `productId`   INT        NOT NULL,
    INDEX `cartId_idx` (`cartId` ASC) VISIBLE,
    INDEX `prodId_idx` (`productId` ASC) VISIBLE,
    CONSTRAINT `cartId`
        FOREIGN KEY (`cartId`)
            REFERENCES `ecommerce`.`cart` (`cartId`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION,
    CONSTRAINT `prodId`
        FOREIGN KEY (`productId`)
            REFERENCES `ecommerce`.`product` (`productId`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
);
