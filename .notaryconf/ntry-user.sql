CREATE TABLE IF NOT EXISTS `user` (
  `eth_address` varchar(41) NOT NULL,
  `secondary_address` varchar(64) DEFAULT NULL,
  `first_name` varchar(50) DEFAULT NULL,
  `last_name` varchar(50) DEFAULT NULL,
  `email_address` varchar(100) NOT NULL,
  `password` varchar(30) NOT NULL,
  `telephone_number` varchar(100) DEFAULT NULL,
  `address` varchar(250) DEFAULT NULL,
  `email_verified` bool DEFAULT NULL,
  `eth_verification` varchar(41) DEFAULT NULL,
  `verification_code` varchar(40) DEFAULT NULL,
  `reg_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`eth_address`)
) DEFAULT CHARSET=utf8;

