CREATE TABLE `customers`
	( `id` INT NOT NULL AUTO_INCREMENT ,
	 `name` VARCHAR(50) NOT NULL ,
	 `sex` SMALLINT NOT NULL ,
	 `tel` VARCHAR(50) NOT NULL ,
	 PRIMARY KEY (`id`))
ENGINE = InnoDB CHARACTER SET utf8 COLLATE utf8_unicode_ci;


INSERT INTO `customers` (`id`, `name`, `sex`, `tel`)
VALUES
	(NULL, 'john', '0', '000-0000-0000'),
	(NULL, 'micheal', '1', '000-0000-0000');
