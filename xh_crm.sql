-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        10.2.9-MariaDB - mariadb.org binary distribution
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  9.4.0.5125
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- 导出 xh_crm 的数据库结构
CREATE DATABASE IF NOT EXISTS `xh_crm` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `xh_crm`;

-- 导出  表 xh_crm.xh_color 结构
CREATE TABLE IF NOT EXISTS `xh_color` (
  `color_id` varchar(32) NOT NULL COMMENT '颜色id',
  `color_name` varchar(64) NOT NULL COMMENT '名称',
  `color_note` varchar(1024) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`color_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='颜色';

-- 正在导出表  xh_crm.xh_color 的数据：~4 rows (大约)
/*!40000 ALTER TABLE `xh_color` DISABLE KEYS */;
INSERT INTO `xh_color` (`color_id`, `color_name`, `color_note`) VALUES
	('104fa48ed8c448878973c82e46682d4a', 'blue', 'blue2'),
	('86ddcd1de48342feb8c4db4160dbde5d', 'green', 'green'),
	('91e1fbed58594c08803eafa5544cbab2', 'black', 'black'),
	('c7be79bd313641a49bf88e67988e2d8a', 'red', 'red3');
/*!40000 ALTER TABLE `xh_color` ENABLE KEYS */;

-- 导出  表 xh_crm.xh_customer 结构
CREATE TABLE IF NOT EXISTS `xh_customer` (
  `customer_id` varchar(32) NOT NULL,
  `customer_name` varchar(256) NOT NULL COMMENT '名称',
  `customer_note` varchar(1024) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`customer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='客户';

-- 正在导出表  xh_crm.xh_customer 的数据：~2 rows (大约)
/*!40000 ALTER TABLE `xh_customer` DISABLE KEYS */;
INSERT INTO `xh_customer` (`customer_id`, `customer_name`, `customer_note`) VALUES
	('dc399eaf406946a0a0a97f1493b4b315', '客户2', '22'),
	('fe7567e6078d4f55be16bb9d907d7dfe', '客户1', '11');
/*!40000 ALTER TABLE `xh_customer` ENABLE KEYS */;

-- 导出  表 xh_crm.xh_design 结构
CREATE TABLE IF NOT EXISTS `xh_design` (
  `design_id` varchar(32) NOT NULL COMMENT '花型id',
  `design_name` varchar(64) NOT NULL COMMENT '花型名称',
  `design_quantity` float NOT NULL DEFAULT 0,
  `design_note` varchar(1024) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`design_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='花型';

-- 正在导出表  xh_crm.xh_design 的数据：~1 rows (大约)
/*!40000 ALTER TABLE `xh_design` DISABLE KEYS */;
INSERT INTO `xh_design` (`design_id`, `design_name`, `design_quantity`, `design_note`) VALUES
	('86abe10cfb6c415395736a0b1c4c0cd0', '小菊花', 0, '小菊花'),
	('c246e690332649b3adf4d5cdfd78774e', '花型1', 0, 'http://recordcdn.quklive.co');
/*!40000 ALTER TABLE `xh_design` ENABLE KEYS */;

-- 导出  表 xh_crm.xh_design_color 结构
CREATE TABLE IF NOT EXISTS `xh_design_color` (
  `design_id` varchar(32) NOT NULL,
  `color_id` varchar(32) NOT NULL,
  `pic_url` varchar(1024) NOT NULL DEFAULT '' COMMENT '图片地址',
  `design_color_note` varchar(1024) NOT NULL DEFAULT '',
  PRIMARY KEY (`design_id`),
  KEY `FK_xh_design_color_xh_color` (`color_id`),
  CONSTRAINT `FK_xh_design_color_xh_color` FOREIGN KEY (`color_id`) REFERENCES `xh_color` (`color_id`),
  CONSTRAINT `FK_xh_design_color_xh_design` FOREIGN KEY (`design_id`) REFERENCES `xh_design` (`design_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='花型颜色关联表';

-- 正在导出表  xh_crm.xh_design_color 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `xh_design_color` DISABLE KEYS */;
INSERT INTO `xh_design_color` (`design_id`, `color_id`, `pic_url`, `design_color_note`) VALUES
	('c246e690332649b3adf4d5cdfd78774e', 'c7be79bd313641a49bf88e67988e2d8a', '/upload/designColor/b49c74f53ba942589683c05061371fc4.png', '');
/*!40000 ALTER TABLE `xh_design_color` ENABLE KEYS */;

-- 导出  表 xh_crm.xh_inbound 结构
CREATE TABLE IF NOT EXISTS `xh_inbound` (
  `inbound_id` varchar(32) NOT NULL,
  `inbound_cloth_id` varchar(32) NOT NULL,
  PRIMARY KEY (`inbound_id`,`inbound_cloth_id`),
  KEY `FK_xh_inbound_xh_inbound_cloth` (`inbound_cloth_id`),
  CONSTRAINT `FK_xh_inbound_xh_inbound_cloth` FOREIGN KEY (`inbound_cloth_id`) REFERENCES `xh_inbound_cloth` (`inbound_cloth_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  xh_crm.xh_inbound 的数据：~7 rows (大约)
/*!40000 ALTER TABLE `xh_inbound` DISABLE KEYS */;
INSERT INTO `xh_inbound` (`inbound_id`, `inbound_cloth_id`) VALUES
	('0e0fc7e74a3b498899829a59c2cda6ca', '2db27afa438d45a4965d11a987f29175'),
	('0e0fc7e74a3b498899829a59c2cda6ca', 'dcd1d06a19384907a91d2f3ffc111fe8'),
	('950d91c7b32c4cfdbde5b2c13dbb24fd', '1e80e9a41b0545c39a47ab073021e027'),
	('950d91c7b32c4cfdbde5b2c13dbb24fd', '248513cfd0434e24992c2d01208bc110'),
	('950d91c7b32c4cfdbde5b2c13dbb24fd', '83f7f97e57ec474fb78581e75b864c85'),
	('950d91c7b32c4cfdbde5b2c13dbb24fd', 'f5e0670c031043c681c174c41e0a7b7e'),
	('950d91c7b32c4cfdbde5b2c13dbb24fd', 'f8e02407a7f24175b1c46d6abd06b438');
/*!40000 ALTER TABLE `xh_inbound` ENABLE KEYS */;

-- 导出  表 xh_crm.xh_inbound_cloth 结构
CREATE TABLE IF NOT EXISTS `xh_inbound_cloth` (
  `inbound_cloth_id` varchar(32) NOT NULL,
  `design_id` varchar(32) NOT NULL,
  `color_id` varchar(32) NOT NULL,
  `supplier_id` varchar(32) NOT NULL COMMENT '供货商id',
  `inbound_cloth_quantity` float NOT NULL COMMENT '数量',
  `remain_quantity` float NOT NULL COMMENT '库存数量',
  `inbound_cloth_price` float NOT NULL COMMENT '价格',
  `create_time` timestamp NOT NULL DEFAULT current_timestamp(),
  `inbound_cloth_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `inbound_cloth_note` varchar(1024) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`inbound_cloth_id`),
  KEY `FK_xh_inbound_xh_supplier` (`supplier_id`),
  KEY `FK_xh_inbound_xh_design` (`design_id`),
  KEY `FK_xh_inbound_xh_color` (`color_id`),
  CONSTRAINT `FK_xh_inbound_xh_color` FOREIGN KEY (`color_id`) REFERENCES `xh_color` (`color_id`),
  CONSTRAINT `FK_xh_inbound_xh_design` FOREIGN KEY (`design_id`) REFERENCES `xh_design` (`design_id`),
  CONSTRAINT `FK_xh_inbound_xh_supplier` FOREIGN KEY (`supplier_id`) REFERENCES `xh_supplier` (`supplier_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='入库布匹';

-- 正在导出表  xh_crm.xh_inbound_cloth 的数据：~7 rows (大约)
/*!40000 ALTER TABLE `xh_inbound_cloth` DISABLE KEYS */;
INSERT INTO `xh_inbound_cloth` (`inbound_cloth_id`, `design_id`, `color_id`, `supplier_id`, `inbound_cloth_quantity`, `remain_quantity`, `inbound_cloth_price`, `create_time`, `inbound_cloth_time`, `inbound_cloth_note`) VALUES
	('1e80e9a41b0545c39a47ab073021e027', 'c246e690332649b3adf4d5cdfd78774e', 'c7be79bd313641a49bf88e67988e2d8a', '8288b5d67e424478a65df800dc6ae7f1', 144, 144, 100.01, '2018-08-15 09:19:26', '2018-08-15 09:19:31', ''),
	('248513cfd0434e24992c2d01208bc110', 'c246e690332649b3adf4d5cdfd78774e', 'c7be79bd313641a49bf88e67988e2d8a', '8288b5d67e424478a65df800dc6ae7f1', 135, 135, 150, '2018-08-15 09:19:26', '2018-08-15 09:19:31', ''),
	('2db27afa438d45a4965d11a987f29175', 'c246e690332649b3adf4d5cdfd78774e', 'c7be79bd313641a49bf88e67988e2d8a', '8288b5d67e424478a65df800dc6ae7f1', 10, 3.5, 100, '2018-08-10 17:23:11', '2018-08-10 17:23:30', ''),
	('83f7f97e57ec474fb78581e75b864c85', 'c246e690332649b3adf4d5cdfd78774e', 'c7be79bd313641a49bf88e67988e2d8a', '8288b5d67e424478a65df800dc6ae7f1', 125, 125, 150, '2018-08-15 09:19:26', '2018-08-15 09:19:31', ''),
	('dcd1d06a19384907a91d2f3ffc111fe8', 'c246e690332649b3adf4d5cdfd78774e', 'c7be79bd313641a49bf88e67988e2d8a', '8288b5d67e424478a65df800dc6ae7f1', 20, 4.5, 100, '2018-08-10 17:23:11', '2018-08-10 17:23:30', ''),
	('f5e0670c031043c681c174c41e0a7b7e', 'c246e690332649b3adf4d5cdfd78774e', 'c7be79bd313641a49bf88e67988e2d8a', '8288b5d67e424478a65df800dc6ae7f1', 133, 133, 100.01, '2018-08-15 09:19:26', '2018-08-15 09:19:31', ''),
	('f8e02407a7f24175b1c46d6abd06b438', 'c246e690332649b3adf4d5cdfd78774e', 'c7be79bd313641a49bf88e67988e2d8a', '8288b5d67e424478a65df800dc6ae7f1', 122, 122, 100.01, '2018-08-15 09:19:26', '2018-08-15 09:19:31', '');
/*!40000 ALTER TABLE `xh_inbound_cloth` ENABLE KEYS */;

-- 导出  表 xh_crm.xh_outbound 结构
CREATE TABLE IF NOT EXISTS `xh_outbound` (
  `outbound_id` varchar(32) NOT NULL,
  `outbound_cloth_id` varchar(32) NOT NULL,
  PRIMARY KEY (`outbound_id`,`outbound_cloth_id`),
  KEY `FK__xh_outbound_cloth` (`outbound_cloth_id`),
  CONSTRAINT `FK__xh_outbound_cloth` FOREIGN KEY (`outbound_cloth_id`) REFERENCES `xh_outbound_cloth` (`outbound_cloth_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  xh_crm.xh_outbound 的数据：~4 rows (大约)
/*!40000 ALTER TABLE `xh_outbound` DISABLE KEYS */;
INSERT INTO `xh_outbound` (`outbound_id`, `outbound_cloth_id`) VALUES
	('217629469e8f4ed08e36ff057ad90bd6', '33f4211d2e7a4b048479deaae000d52b'),
	('217629469e8f4ed08e36ff057ad90bd6', 'e46f089641c34bc4842a8eb2534fcb5c'),
	('a53125ef40c04d4892bb816acb22a1e0', 'a345e1aa97694cb5b3a90ff74ea8818b'),
	('a53125ef40c04d4892bb816acb22a1e0', 'ded17d83cea9479cae1cbeb75206f0be');
/*!40000 ALTER TABLE `xh_outbound` ENABLE KEYS */;

-- 导出  表 xh_crm.xh_outbound_cloth 结构
CREATE TABLE IF NOT EXISTS `xh_outbound_cloth` (
  `outbound_cloth_id` varchar(32) NOT NULL,
  `customer_id` varchar(32) NOT NULL COMMENT '客户id',
  `inbound_cloth_id` varchar(32) NOT NULL,
  `outbound_cloth_quantity` float NOT NULL COMMENT '数量',
  `outbound_cloth_price` float NOT NULL COMMENT '价格',
  `create_time` timestamp NOT NULL DEFAULT current_timestamp(),
  `outbound_cloth_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `outbound_cloth_note` varchar(1024) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`outbound_cloth_id`),
  KEY `FK_xh_outbound_xh_customer` (`customer_id`),
  KEY `FK_xh_outbound_cloth_xh_inbound_cloth` (`inbound_cloth_id`),
  CONSTRAINT `FK_xh_outbound_cloth_xh_inbound_cloth` FOREIGN KEY (`inbound_cloth_id`) REFERENCES `xh_inbound_cloth` (`inbound_cloth_id`),
  CONSTRAINT `FK_xh_outbound_xh_customer` FOREIGN KEY (`customer_id`) REFERENCES `xh_customer` (`customer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='出库布匹';

-- 正在导出表  xh_crm.xh_outbound_cloth 的数据：~4 rows (大约)
/*!40000 ALTER TABLE `xh_outbound_cloth` DISABLE KEYS */;
INSERT INTO `xh_outbound_cloth` (`outbound_cloth_id`, `customer_id`, `inbound_cloth_id`, `outbound_cloth_quantity`, `outbound_cloth_price`, `create_time`, `outbound_cloth_time`, `outbound_cloth_note`) VALUES
	('33f4211d2e7a4b048479deaae000d52b', 'fe7567e6078d4f55be16bb9d907d7dfe', 'dcd1d06a19384907a91d2f3ffc111fe8', 10, 100, '2018-08-10 17:23:47', '2018-08-10 17:24:06', ''),
	('a345e1aa97694cb5b3a90ff74ea8818b', 'fe7567e6078d4f55be16bb9d907d7dfe', 'dcd1d06a19384907a91d2f3ffc111fe8', 5.5, 100, '2018-08-10 17:24:35', '2018-08-10 17:24:54', ''),
	('ded17d83cea9479cae1cbeb75206f0be', 'fe7567e6078d4f55be16bb9d907d7dfe', '2db27afa438d45a4965d11a987f29175', 1.5, 100, '2018-08-10 17:24:35', '2018-08-10 17:24:54', ''),
	('e46f089641c34bc4842a8eb2534fcb5c', 'fe7567e6078d4f55be16bb9d907d7dfe', '2db27afa438d45a4965d11a987f29175', 5, 100, '2018-08-10 17:23:48', '2018-08-10 17:24:06', '');
/*!40000 ALTER TABLE `xh_outbound_cloth` ENABLE KEYS */;

-- 导出  表 xh_crm.xh_scheduling 结构
CREATE TABLE IF NOT EXISTS `xh_scheduling` (
  `scheduling_id` varchar(32) NOT NULL,
  `design_id` varchar(32) NOT NULL,
  `color_id` varchar(32) NOT NULL,
  `supplier_id` varchar(32) NOT NULL COMMENT '供货商id',
  `scheduling_quantity` float NOT NULL COMMENT '数量',
  `scheduling_price` float NOT NULL COMMENT '价格',
  `create_time` timestamp NOT NULL DEFAULT current_timestamp(),
  `scheduling_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '排单时间',
  `work_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '上机时间',
  `estimate_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '预计时间',
  `scheduling_note` varchar(1024) NOT NULL DEFAULT '' COMMENT '备注',
  `scheduling_status` int(11) NOT NULL DEFAULT 0 COMMENT '1-完成',
  PRIMARY KEY (`scheduling_id`),
  KEY `FK_xh_scheduling_xh_supplier` (`supplier_id`),
  KEY `FK_xh_scheduling_xh_design` (`design_id`),
  KEY `FK_xh_scheduling_xh_color` (`color_id`),
  CONSTRAINT `FK_xh_scheduling_xh_color` FOREIGN KEY (`color_id`) REFERENCES `xh_color` (`color_id`),
  CONSTRAINT `FK_xh_scheduling_xh_design` FOREIGN KEY (`design_id`) REFERENCES `xh_design` (`design_id`),
  CONSTRAINT `FK_xh_scheduling_xh_supplier` FOREIGN KEY (`supplier_id`) REFERENCES `xh_supplier` (`supplier_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='排单';

-- 正在导出表  xh_crm.xh_scheduling 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `xh_scheduling` DISABLE KEYS */;
INSERT INTO `xh_scheduling` (`scheduling_id`, `design_id`, `color_id`, `supplier_id`, `scheduling_quantity`, `scheduling_price`, `create_time`, `scheduling_time`, `work_time`, `estimate_time`, `scheduling_note`, `scheduling_status`) VALUES
	('5257c6d44c384afcaf92467aa17c01d8', 'c246e690332649b3adf4d5cdfd78774e', 'c7be79bd313641a49bf88e67988e2d8a', '8288b5d67e424478a65df800dc6ae7f1', 100.01, 100.01, '2018-08-07 19:44:56', '2018-08-07 19:45:11', '0000-00-00 00:00:00', '0000-00-00 00:00:00', '', 0);
/*!40000 ALTER TABLE `xh_scheduling` ENABLE KEYS */;

-- 导出  表 xh_crm.xh_supplier 结构
CREATE TABLE IF NOT EXISTS `xh_supplier` (
  `supplier_id` varchar(32) NOT NULL COMMENT '供货商id',
  `supplier_name` varchar(256) NOT NULL COMMENT '名称',
  `supplier_note` varchar(1024) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`supplier_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='供应商';

-- 正在导出表  xh_crm.xh_supplier 的数据：~1 rows (大约)
/*!40000 ALTER TABLE `xh_supplier` DISABLE KEYS */;
INSERT INTO `xh_supplier` (`supplier_id`, `supplier_name`, `supplier_note`) VALUES
	('8288b5d67e424478a65df800dc6ae7f1', 'supplier', '33');
/*!40000 ALTER TABLE `xh_supplier` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
