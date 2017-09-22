CREATE TABLE IF NOT EXISTS `user` (
  `uid` varchar(32) NOT NULL, # this is the user id
  `eth_address` varchar(41) DEFAULT "",
  `first_name` varchar(50) DEFAULT "",
  `last_name` varchar(50) DEFAULT "",
  `email_address` varchar(100) NOT NULL,
  `password` varchar(30) NOT NULL,
  `telephone_number` varchar(100) DEFAULT "",
  `address` varchar(250) DEFAULT "",
  `account_verified` bool DEFAULT false,
  `eth_verification` varchar(64) DEFAULT "",
  `country` varchar(30) DEFAULT "",
  `city` varchar(30) DEFAULT "",
  `state` varchar(20) DEFAULT "",
  `avatar` varchar(35) DEFAULT "", 
  `reg_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`)
) DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `contacts` (
  `uid` varchar(32) NOT NULL,
  `eth_address` varchar(41) DEFAULT "",
  `first_name` varchar(50) DEFAULT "",
  `last_name` varchar(50) DEFAULT "",
  `email_address` varchar(100) NOT NULL,
  `password` varchar(30) NOT NULL,
  `telephone_number` varchar(100) DEFAULT "",
  `address` varchar(250) DEFAULT "",
  `account_verified` bool DEFAULT false,
  `eth_verification` varchar(64) DEFAULT "",
  `country` varchar(30) DEFAULT "",
  `city` varchar(30) DEFAULT "",
  `state` varchar(20) DEFAULT "",
  `avatar` varchar(35) DEFAULT "", 
  `reg_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`)
) DEFAULT CHARSET=utf8;