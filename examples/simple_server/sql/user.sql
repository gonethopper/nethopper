CREATE TABLE IF NOT EXISTS `user`(
   `uid` INT UNSIGNED AUTO_INCREMENT,
   `name` VARCHAR(255) NOT NULL,
   `password` VARCHAR(255) NOT NULL,
   `coin` INT NOT NULL DEFAULT 0,
   `avator` VARCHAR(40) NOT NULL,
   `login_date` DATE,
   PRIMARY KEY ( `uid` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;