-- -----------------------------------------------------
-- このファイルはDDLファイルのサンプルです
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Database sample
-- -----------------------------------------------------

CREATE DATABASE IF NOT EXISTS `sample` DEFAULT CHARACTER SET utf8mb4 ;
USE `sample` ;

SET CHARSET utf8mb4;

-- -----------------------------------------------------
-- Table `sample`.`item`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `sample`.`item` (
  `id` VARCHAR(128) NOT NULL COMMENT 'ID',
  `name` VARCHAR(128) NOT NULL COMMENT '名前',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = '商品';


CREATE TABLE IF NOT EXISTS `sample`.`users` (
    `id` char(36) NOT NULL,
    `name` varchar(255) NOT NULL UNIQUE,
    `token` char(36) NOT NULL,
    `high_score` integer NOT NULL DEFAULT 0,
    `coin` integer NOT NULL DEFAULT 0,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`))
ENGINE = InnoDB;
