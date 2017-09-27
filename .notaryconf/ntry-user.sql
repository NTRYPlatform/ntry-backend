CREATE TABLE IF NOT EXISTS `user` (
  `uid` varchar(32) NOT NULL, # this is the user id
  `eth_address` varchar(41) DEFAULT "",
  `first_name` varchar(50) DEFAULT "",
  `last_name` varchar(50) DEFAULT "",
  `email_address` varchar(100) NOT NULL,
  `password` varchar(30) NOT NULL,
  `telephone_number` varchar(20) DEFAULT "",
  `address` varchar(250) DEFAULT "",
  `account_verified` bool DEFAULT false,
  `eth_verification` varchar(66) DEFAULT "",
  `country` varchar(30) DEFAULT "",
  `city` varchar(30) DEFAULT "",
  `state` varchar(20) DEFAULT "",
  `avatar` varchar(35) DEFAULT "", 
  `reg_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`)
) DEFAULT CHARSET=utf8;

# for contacts - user-to-user junction table
CREATE TABLE IF NOT EXISTS `user_to_user` (
  `p_uid` varchar(32) NOT NULL,
  `s_uid` varchar(32) NOT NULL,
  PRIMARY KEY (`p_uid`,`s_uid`),
  FOREIGN KEY (`p_uid`) REFERENCES `user` (`uid`) ON DELETE CASCADE,
  FOREIGN KEY (`s_uid`) REFERENCES `user` (`uid`) ON DELETE CASCADE
) DEFAULT CHARSET=utf8;

# for car contracts
CREATE TABLE IF NOT EXISTS `car_contract` (
  `cid` varchar(32) NOT NULL,
  `year` int NOT NULL,
  `make` varchar(20) NOT NULL,
  `model` varchar(20) NOT NULL,
  `vin` varchar(20) NOT NULL,
  `type` varchar(20) NOT NULL,
  `color` varchar(20) NOT NULL, 
  `engine_no` varchar(25) DEFAULT "",
  `mileage` int NOT NULL,
  `total_price` int NOT NULL,
  `down_payment` int NOT NULL,
  `remaining_payment` int NOT NULL,
  `remaining_payment_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP, # might want to change this
  `creation_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP, # might want to change this
  `last_updated_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP, # might want to change this
PRIMARY KEY (`cid`)
) DEFAULT CHARSET=utf8;

# for contracts and users - junction table
CREATE TABLE IF NOT EXISTS `car_contract_user` (
  `buyer` varchar(32) NOT NULL,
  `seller` varchar(32) NOT NULL,
  `cid` varchar(32) NOT NULL,
  PRIMARY KEY (`buyer`, `seller`, `cid`),
  FOREIGN KEY (`buyer`) REFERENCES `user` (`uid`) ON DELETE CASCADE,
  FOREIGN KEY (`seller`) REFERENCES `user` (`uid`) ON DELETE CASCADE,
  FOREIGN KEY (`cid`) REFERENCES `car_contract` (`cid`) ON DELETE CASCADE
) DEFAULT CHARSET=utf8;
