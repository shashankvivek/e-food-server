CREATE TABLE `ecommerce`.`broad_category` (
  `BC_Id` MEDIUMINT NOT NULL  AUTO_INCREMENT,
  `BC_Name` VARCHAR(45) NOT NULL,
  `BC_Description` VARCHAR(45) NULL,
  `BC_ImageUrl` VARCHAR(200) NULL,
  `BC_IsActive` TINYINT NULL DEFAULT 1,
  PRIMARY KEY (`BC_Id`)
);



INSERT INTO `ecommerce`.`broad_category`
(
`BC_Name`,
`BC_Description`)
VALUES
("Fruits",
"These are juicy fruits");


CREATE TABLE `ecommerce`.`sub_category` (
  `SC_Id` MEDIUMINT NOT NULL  AUTO_INCREMENT,
  `SC_Name` VARCHAR(45) NULL,
  `SC_Description` VARCHAR(45) NULL,
  `SC_ImageUrl` VARCHAR(200) NULL,
  `SC_IsActive` TINYINT NULL DEFAULT 1,
  `BC_Id` MEDIUMINT NOT NULL,
  FOREIGN KEY (`BC_Id`) REFERENCES `ecommerce`.`broad_category`(`BC_Id`),
  PRIMARY KEY (`SC_Id`)
);

INSERT INTO `ecommerce`.`sub_category`
(
`SC_Name`,
`SC_Description`,
`BC_Id`)
VALUES
("Apple",
"This is an apple",
1);


INSERT INTO `ecommerce`.`sub_category`
(
`SC_Name`,
`SC_Description`,
`BC_Id`)
VALUES
("Oranges",
"These are oranges",
1);


INSERT INTO `ecommerce`.`sub_category`
(
`Name`,
`Description`,
`BroadCategoryId`)
VALUES
("Bananas",
"These are Bananas",
1);


INSERT INTO `ecommerce`.`sub_category`
(
`SC_Name`,
`SC_Description`,
`BC_Id`)
VALUES
("Pears",
"These are Pears",
1);

