CREATE TABLE IF NOT EXISTS `user` (
  `uid` varchar(32) NOT NULL, # this is the user id
  `eth_address` varchar(41) DEFAULT NULL,
  `first_name` varchar(50) DEFAULT NULL,
  `last_name` varchar(50) DEFAULT NULL,
  `email_address` varchar(100) NOT NULL,
  `password` varchar(30) NOT NULL,
  `telephone_number` varchar(100) DEFAULT NULL,
  `address` varchar(250) DEFAULT NULL,
  `account_verified` bool DEFAULT NULL,
  `eth_verification` varchar(64) DEFAULT NULL,
  `reg_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`)
) DEFAULT CHARSET=utf8;