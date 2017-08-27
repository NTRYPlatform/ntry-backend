CREATE TABLE IF NOT EXISTS `user` (
  `eth_address` varchar(41) NOT NULL,
  `pub_key` varchar(64) DEFAULT NULL,
  `first_name` varchar(50) DEFAULT NULL,
  `last_name` varchar(50) DEFAULT NULL,
  `email_address` varchar(100) NOT NULL,
  `password` varchar(30) NOT NULL,
  `telephone_number` varchar(100) DEFAULT NULL,
  `address` varchar(250) DEFAULT NULL,
  `email_verified` bool DEFAULT NULL,
  `eth_verification` varchar(41) DEFAULT NULL,
  PRIMARY KEY (`eth_address`)
) DEFAULT CHARSET=utf8;

