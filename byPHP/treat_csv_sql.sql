CREATE TABLE `treat_csv`.`customers` 
	( `id` INT NULL AUTO_INCREMENT ,
	 `name` VARCHAR(50) NOT NULL , 
	 `sex` SMALLINT NOT NULL , 
	 `tel` VARCHAR(50) NOT NULL , 
	 PRIMARY KEY (`id`)) 
ENGINE = InnoDB CHARACTER SET utf8 COLLATE utf8_unicode_ci;


INSERT INTO `treat_csv`.`customers` (`id`, `name`, `sex`, `tel`) 
VALUES 
	(NULL, 'john', '0', '000-0000-0000'), 
	(NULL, 'micheal', '1', '000-0000-0000');