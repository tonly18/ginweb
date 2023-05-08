DROP TABLE IF EXISTS `item_0000`;
CREATE TABLE IF NOT EXISTS `item_0000` (
  `uid` bigint(20) UNSIGNED NOT NULL COMMENT 'uid',
  `content` varchar(21840) NOT NULL DEFAULT '{}' COMMENT 'content',
  `stime` int(10) NOT NULL DEFAULT 0 COMMENT 'timestamp',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='item';

DROP TABLE IF EXISTS `item_0001`;
CREATE TABLE IF NOT EXISTS `item_0001` (
  `uid` bigint(20) UNSIGNED NOT NULL COMMENT 'uid',
  `content` varchar(21840) NOT NULL DEFAULT '{}' COMMENT 'content',
  `stime` int(11) NOT NULL DEFAULT 0 COMMENT 'timestamp',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='item';
