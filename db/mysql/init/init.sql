-- SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
-- SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
-- SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
SET CHARSET UTF8;
DROP SCHEMA IF EXISTS `pr_card`;
CREATE SCHEMA IF NOT EXISTS `pr_card` DEFAULT CHARACTER SET utf8;
USE `pr_card`;

---- drop ----
DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `cards`;
DROP TABLE IF EXISTS `card_matrixs`;
DROP TABLE IF EXISTS `card_my_word`;
DROP TABLE IF EXISTS `owned_cards`;

---- create ----
CREATE TABLE IF NOT EXISTS `pr_card`.`users` (
`id` VARCHAR(64) NOT NULL COMMENT 'ユーザID',
`card_id` VARCHAR(64) COMMENT '名刺ID',
`pass` VARCHAR(64) COMMENT 'パスワード',
`login_id` VARCHAR(64) NOT NULL COMMENT '自動ログインID',
PRIMARY KEY (`id`),
INDEX `idx_auth_token` (`id` ASC)
)
ENGINE = InnoDB
COMMENT = 'アカウント情報を保存';

CREATE TABLE IF NOT EXISTS `pr_card`.`cards` (
`id` VARCHAR(64) NOT NULL COMMENT 'カード識別ID',
`name` VARCHAR(32) NOT NULL COMMENT 'ユーザ名',
`image_path` VARCHAR(64) COMMENT '画像のパス',
`free_text` TEXT COMMENT '画像のパス',
PRIMARY KEY (`id`),
INDEX `idx_auth_token` (`id` ASC)
)
ENGINE = InnoDB
COMMENT = '名刺の要素を保存';

CREATE TABLE IF NOT EXISTS `pr_card`.`card_matrixs` (
`id` VARCHAR(64) NOT NULL COMMENT '識別ID',
`card_id` VARCHAR(32) NOT NULL COMMENT 'カード識別ID',
`item` VARCHAR(64) COMMENT '項目名',
`socre` INT COMMENT 'スコア',
PRIMARY KEY (`id`),
INDEX `idx_auth_token` (`id` ASC)
)
ENGINE = InnoDB
COMMENT = '名刺のグラフ部分の要素を保存';

CREATE TABLE IF NOT EXISTS `pr_card`.`card_my_word` (
`id` VARCHAR(64) NOT NULL COMMENT '識別ID',
`card_id` VARCHAR(32) NOT NULL COMMENT 'カード識別ID',
`word` VARCHAR(64) COMMENT '自分を表す単語',
PRIMARY KEY (`id`),
INDEX `idx_auth_token` (`id` ASC)
)
ENGINE = InnoDB
COMMENT = '名刺の自分を表すワードを保存';

CREATE TABLE IF NOT EXISTS `pr_card`.`owned_cards` (
`id` VARCHAR(64) NOT NULL COMMENT '識別ID',
`user_id` VARCHAR(32) NOT NULL COMMENT 'ユーザ識別ID',
`card_id` VARCHAR(32) NOT NULL COMMENT 'カード識別ID',
INDEX `idx_auth_token` (`id` ASC)
)
ENGINE = InnoDB
COMMENT = 'どのユーザがどの名刺をもつか保存';

---- insert ----
INSERT INTO `owned_cards` VALUES ("1","a","b");
INSERT INTO `owned_cards` VALUES ("1","a","c");
INSERT INTO `owned_cards` VALUES ("1","a","d");
INSERT INTO `owned_cards` VALUES ("1","a","e");
INSERT INTO `owned_cards` VALUES ("1","a","f");
INSERT INTO `owned_cards` VALUES ("1","b","a");
INSERT INTO `owned_cards` VALUES ("1","b","c");
INSERT INTO `owned_cards` VALUES ("1","b","d");
INSERT INTO `owned_cards` VALUES ("1","b","e");
INSERT INTO `owned_cards` VALUES ("1","b","f");
INSERT INTO `owned_cards` VALUES ("1","a","g");