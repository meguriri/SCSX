/*
Navicat MySQL Data Transfer

Source Server         : teach
Source Server Version : 50051
Source Host           : 127.0.0.1:3306
Source Database       : shop

Target Server Type    : MYSQL
Target Server Version : 50051
File Encoding         : 65001

Date: 2022-04-27 09:34:40
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` int(11) NOT NULL auto_increment,
  `username` varchar(20) default NULL,
  `password` varchar(20) default NULL,
  PRIMARY KEY  (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin` VALUES ('1', 'admin', 'admin');

-- ----------------------------
-- Table structure for buycar
-- ----------------------------
DROP TABLE IF EXISTS `buycar`;
CREATE TABLE `buycar` (
  `id` int(11) NOT NULL auto_increment,
  `pid` int(11) default NULL,
  `num` int(11) default NULL,
  `mid` int(11) default NULL,
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of buycar
-- ----------------------------

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` int(11) NOT NULL auto_increment,
  `name` varchar(20) default NULL,
  PRIMARY KEY  (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES ('1', '进口食品、生鲜');
INSERT INTO `category` VALUES ('2', '食品、饮料、酒');
INSERT INTO `category` VALUES ('3', '母婴、玩具、童装');
INSERT INTO `category` VALUES ('4', '家居、家庭清洁、纸品');
INSERT INTO `category` VALUES ('5', '美妆、个人护理、洗护');
INSERT INTO `category` VALUES ('6', '女装、内衣、中老年');
INSERT INTO `category` VALUES ('7', '鞋靴、箱包、腕表配饰');
INSERT INTO `category` VALUES ('8', '男装、运动');
INSERT INTO `category` VALUES ('9', '手机、小家电、电脑');
INSERT INTO `category` VALUES ('10', '礼品、充值');

-- ----------------------------
-- Table structure for member
-- ----------------------------
DROP TABLE IF EXISTS `member`;
CREATE TABLE `member` (
  `id` int(11) NOT NULL auto_increment,
  `username` varchar(20) default NULL,
  `password` varchar(20) default NULL,
  `name` varchar(20) default NULL,
  `identity` char(18) default NULL,
  `tel` char(11) default NULL,
  `email` varchar(20) default NULL,
  `balance` decimal(10,2) default NULL,
  `state` int(11) default NULL,
  PRIMARY KEY  (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of member
-- ----------------------------
INSERT INTO `member` VALUES ('1', 'tom', '111111', null, null, null, '261749311@qq.com', null, '0');
INSERT INTO `member` VALUES ('2', 'lily', '111111', null, null, null, '263@qq.com', null, '0');
INSERT INTO `member` VALUES ('3', 'a', 'a', null, null, null, 'a@qq.com', null, '0');
INSERT INTO `member` VALUES ('4', 'lucy', '111111', null, null, null, 'lucy@qq.com', null, '0');

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `id` int(11) NOT NULL auto_increment,
  `title` varchar(20) default NULL,
  `content` varchar(200) default NULL,
  `date` datetime default NULL,
  `mid` int(11) default NULL,
  PRIMARY KEY  (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of message
-- ----------------------------
INSERT INTO `message` VALUES ('1', null, null, '2022-03-02 00:00:00', null);
INSERT INTO `message` VALUES ('2', '投诉', '发货太慢!!', '2022-04-24 15:37:32', '1');

-- ----------------------------
-- Table structure for myorder
-- ----------------------------
DROP TABLE IF EXISTS `myorder`;
CREATE TABLE `myorder` (
  `id` int(11) NOT NULL auto_increment,
  `time` datetime default NULL,
  `state` int(11) default NULL,
  `mid` int(11) default NULL,
  `address` varchar(255) default NULL,
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of myorder
-- ----------------------------
INSERT INTO `myorder` VALUES ('7', '2022-04-22 15:50:07', '0', '1', 'tom,13456787878,沈阳市');
INSERT INTO `myorder` VALUES ('8', '2022-04-22 15:52:54', '3', '1', 'aaa,13456565656,大连市');
INSERT INTO `myorder` VALUES ('21', '2022-04-24 15:51:16', '1', '1', '李四,13889898989,沈阳');
INSERT INTO `myorder` VALUES ('22', '2022-04-24 15:52:38', '1', '4', '赵钱,18989898989,北京');
INSERT INTO `myorder` VALUES ('23', '2022-04-24 15:55:22', '1', '2', '王先生,13456565656,鞍山');
INSERT INTO `myorder` VALUES ('24', '2022-04-24 15:57:17', '1', '2', '李女士,15656565656,哈尔滨');
INSERT INTO `myorder` VALUES ('25', '2022-04-24 15:58:56', '1', '2', '小王,18989898989,大连市');

-- ----------------------------
-- Table structure for orderdetail
-- ----------------------------
DROP TABLE IF EXISTS `orderdetail`;
CREATE TABLE `orderdetail` (
  `id` int(11) NOT NULL auto_increment,
  `pid` int(11) default NULL,
  `oid` int(11) default NULL,
  `num` int(11) default NULL,
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=121 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of orderdetail
-- ----------------------------
INSERT INTO `orderdetail` VALUES ('106', '10', '7', '1');
INSERT INTO `orderdetail` VALUES ('107', '1', '8', '1');
INSERT INTO `orderdetail` VALUES ('108', '10', '8', '1');
INSERT INTO `orderdetail` VALUES ('109', '3', '21', '2');
INSERT INTO `orderdetail` VALUES ('110', '7', '21', '1');
INSERT INTO `orderdetail` VALUES ('111', '10', '22', '1');
INSERT INTO `orderdetail` VALUES ('112', '1', '22', '1');
INSERT INTO `orderdetail` VALUES ('113', '2', '22', '5');
INSERT INTO `orderdetail` VALUES ('114', '1', '23', '3');
INSERT INTO `orderdetail` VALUES ('115', '3', '23', '3');
INSERT INTO `orderdetail` VALUES ('116', '7', '23', '1');
INSERT INTO `orderdetail` VALUES ('117', '2', '24', '1');
INSERT INTO `orderdetail` VALUES ('118', '4', '25', '1');
INSERT INTO `orderdetail` VALUES ('119', '8', '25', '3');
INSERT INTO `orderdetail` VALUES ('120', '5', '25', '2');

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product` (
  `id` int(11) NOT NULL auto_increment,
  `name` varchar(50) default NULL,
  `type` varchar(20) default NULL,
  `color` varchar(20) default NULL,
  `price` decimal(10,2) default NULL,
  `state` int(11) default NULL,
  `number` int(11) default NULL,
  `img` varchar(50) default NULL,
  `imgPath` varchar(50) default NULL,
  `cid` int(11) default NULL,
  PRIMARY KEY  (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of product
-- ----------------------------
INSERT INTO `product` VALUES ('1', '小米手机1', 'note5', '红色', '3000.00', null, '95', 'tel_b2.jpg', '/upload', '9');
INSERT INTO `product` VALUES ('2', '荣耀301', 'H30', '黑色', '3000.00', null, '194', 't1.jpg', '/upload', '9');
INSERT INTO `product` VALUES ('3', 'DELL笔记本1', 'T14', '红色', '3000.00', null, '15', 't3.jpg', '/upload', '9');
INSERT INTO `product` VALUES ('4', '苹果手机1', 'IPhone6', '黑色', '3000.00', null, '29', 'tel_r.jpg', '/upload', '9');
INSERT INTO `product` VALUES ('5', '小米电视1', '50寸', '红色', '3000.00', null, '38', 'tel_3.jpg', '/upload', '9');
INSERT INTO `product` VALUES ('6', '电饭锅', '2-3人容量', '黑色', '3000.00', null, '50', 'n_img1.jpg', '/upload', '9');
INSERT INTO `product` VALUES ('7', '运动手表1', '小米', '红色', '3000.00', null, '58', 'tel_6.jpg', '/upload', '9');
INSERT INTO `product` VALUES ('8', '手机支架1', '小米', '黑色', '3000.00', null, '69', 'tel_4.jpg', '/upload', '9');
INSERT INTO `product` VALUES ('9', '小米手机91', '小米9', '红色', '4000.00', null, '80', 'tel_b2.jpg', '/upload', '9');
INSERT INTO `product` VALUES ('10', '小米手机10', '小米10', '黑色', '5000.00', null, '86', 'tel_b1.jpg', '/upload', '8');
