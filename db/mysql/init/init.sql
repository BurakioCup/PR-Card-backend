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
DROP TABLE IF EXISTS `owned_cards`;
DROP TABLE IF EXISTS `card_qrs`;

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
`name` VARCHAR(32) NOT NULL COMMENT '名前',
`name_image` VARCHAR(64) COMMENT '名前画像のパス',
`tag_image` VARCHAR(64) COMMENT 'タグ画像のパス',
`face_image` VARCHAR(64) COMMENT '顔画像のパス',
`status_image` VARCHAR(64) COMMENT 'ステータス画像のパス',
`free_image` TEXT COMMENT '自由記述欄の画像パス',
PRIMARY KEY (`id`),
INDEX idx_auth_token (`id` ASC)
)
ENGINE = InnoDB
COMMENT = '名刺の要素を保存';

CREATE TABLE IF NOT EXISTS `pr_card`.`owned_cards` (
`id` VARCHAR(64) NOT NULL COMMENT '識別ID',
`user_id` VARCHAR(32) NOT NULL COMMENT 'ユーザ識別ID',
`card_id` VARCHAR(32) NOT NULL COMMENT 'カード識別ID',
INDEX `idx_auth_token` (`id` ASC)
)
ENGINE = InnoDB
COMMENT = 'どのユーザがどの名刺をもつか保存';

CREATE TABLE IF NOT EXISTS `pr_card`.`card_qrs` (
`card_id` VARCHAR(32) NOT NULL COMMENT 'カード識別ID',
`card_image` VARCHAR(64) NOT NULL COMMENT 'qrコードURL',
INDEX `idx_auth_token` (`card_id` ASC)
)
ENGINE = InnoDB
COMMENT = 'カードのQRコードののURLを保存';

-- insert --
INSERT INTO `users` VALUES ("a","a","uuid");

INSERT INTO `card_my_word` VALUES ("1","a","sample1");
INSERT INTO `card_my_word` VALUES ("2","a","sample2");
INSERT INTO `card_my_word` VALUES ("3","a","sample3");
INSERT INTO `card_my_word` VALUES ("4","a","sample4");
INSERT INTO `card_my_word` VALUES ("5","b","sample1");
INSERT INTO `card_my_word` VALUES ("6","b","sample2");
INSERT INTO `card_my_word` VALUES ("7","b","sample3");
INSERT INTO `card_my_word` VALUES ("8","b","sample4");
INSERT INTO `card_my_word` VALUES ("9","c","sample1");
INSERT INTO `card_my_word` VALUES ("10","c","sample2");
INSERT INTO `card_my_word` VALUES ("11","c","sample3");
INSERT INTO `card_my_word` VALUES ("12","c","sample4");

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
INSERT INTO `users` values ("a","a","a","a");
INSERT INTO `card_qrs` values ("a","https://qrコードだよー");
