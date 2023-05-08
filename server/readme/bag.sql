--
-- 表的结构 `bag_0000`
--

DROP TABLE IF EXISTS `bag_0000`;
CREATE TABLE IF NOT EXISTS `bag_0000` (
  `uid` bigint(20) UNSIGNED NOT NULL COMMENT 'uid',
  `item` text NOT NULL COMMENT '普通道具',
  `expire` text NOT NULL COMMENT '时效性道具',
  `itime` int(10) NOT NULL DEFAULT '0' COMMENT 'timestamp',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='背包';

--
-- 转存表中的数据 `bag_0000`
--

INSERT INTO `bag_0000` (`uid`, `item`, `expire`, `itime`) VALUES
(2, '{\"200\":2000}', '{\"item\":{\"300\":30}}', 1658480131);

-- --------------------------------------------------------

--
-- 表的结构 `bag_0001`
--

DROP TABLE IF EXISTS `bag_0001`;
CREATE TABLE IF NOT EXISTS `bag_0001` (
  `uid` bigint(20) UNSIGNED NOT NULL COMMENT 'uid',
  `item` text NOT NULL COMMENT '普通道具',
  `expire` text NOT NULL COMMENT '时效性道具',
  `itime` int(10) NOT NULL DEFAULT '0' COMMENT 'timestamp',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='背包';

--
-- 转存表中的数据 `bag_0001`
--

INSERT INTO `bag_0001` (`uid`, `item`, `expire`, `itime`) VALUES
(1, '{\"100000\":100000}', '{\"item\":{\"3600\":1200}}', 1658480131);
COMMIT;