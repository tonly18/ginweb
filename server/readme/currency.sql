DROP TABLE IF EXISTS `currency_0000`;
CREATE TABLE IF NOT EXISTS `currency_0000` (
  `uid` bigint(20) UNSIGNED NOT NULL COMMENT 'uid',
  `gold` int(11) NOT NULL DEFAULT 0 COMMENT '金币',
  `diamond` int(11) NOT NULL DEFAULT 0 COMMENT '钻石',
  `cash` int(11) NOT NULL DEFAULT 0 COMMENT '现票',
  `life` int(11) NOT NULL DEFAULT 0 COMMENT '生命',
  `stime` int(10) NOT NULL DEFAULT 0 COMMENT '时间戳',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='货币';

DROP TABLE IF EXISTS `currency_0001`;
CREATE TABLE IF NOT EXISTS `currency_0001` (
  `uid` bigint(20) UNSIGNED NOT NULL COMMENT 'uid',
  `gold` int(11) NOT NULL DEFAULT 0 COMMENT '金币',
  `diamond` int(11) NOT NULL DEFAULT 0 COMMENT '钻石',
  `cash` int(11) NOT NULL DEFAULT 0 COMMENT '现票',
  `life` int(11) NOT NULL DEFAULT 0 COMMENT '生命',
  `stime` int(10) NOT NULL DEFAULT 0 COMMENT '时间戳',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='货币';
