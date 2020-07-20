CREATE DATABASE  IF NOT EXISTS `simple_oa_system` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `simple_oa_system`;
-- MySQL dump 10.13  Distrib 5.7.20, for Linux (x86_64)
--
-- Host: localhost    Database: simple_oa_system
-- ------------------------------------------------------
-- Server version	5.7.30-0ubuntu0.16.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `bulletin`
--

DROP TABLE IF EXISTS `bulletin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `bulletin` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '公告id',
  `title` varchar(45) NOT NULL DEFAULT '' COMMENT '公告标题',
  `content` text COMMENT '公共内容',
  `post_staff_id` char(12) NOT NULL DEFAULT '' COMMENT '发表人id',
  `author` varchar(45) NOT NULL DEFAULT '管理员' COMMENT '发表人姓名',
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='公告表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `goods`
--

DROP TABLE IF EXISTS `goods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goods` (
  `id` tinyint(3) unsigned NOT NULL AUTO_INCREMENT COMMENT '物品id',
  `name` varchar(45) NOT NULL DEFAULT '' COMMENT '物品名称',
  `description` varchar(100) NOT NULL DEFAULT '' COMMENT '物品描述',
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='物品表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goods`
--

LOCK TABLES `goods` WRITE;
/*!40000 ALTER TABLE `goods` DISABLE KEYS */;
INSERT INTO `goods` VALUES (1,'笔','一只笔','2020-07-14 09:08:31','2020-07-14 09:08:31'),(2,'笔记本','一个笔记本','2020-07-14 09:08:31','2020-07-14 09:08:31'),(3,'公作牌','公牌换取','2020-07-16 10:32:18','2020-07-16 10:32:33');
/*!40000 ALTER TABLE `goods` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goods_apply`
--

DROP TABLE IF EXISTS `goods_apply`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goods_apply` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '申请表id',
  `goods_id` tinyint(4) NOT NULL COMMENT '物品id',
  `goods_name` varchar(45) NOT NULL COMMENT '物品名字',
  `staff_id` char(12) NOT NULL COMMENT '申请人员工id',
  `staff_name` varchar(45) NOT NULL COMMENT '员工姓名',
  `apply_reason` varchar(255) NOT NULL DEFAULT '' COMMENT '请假原因',
  `approve_staff_id` char(12) NOT NULL DEFAULT '' COMMENT '审核人id\n',
  `approve_staff_name` varchar(45) NOT NULL DEFAULT '' COMMENT '审核人姓名',
  `approve_time` datetime DEFAULT NULL,
  `approve_reason` varchar(100) NOT NULL DEFAULT '' COMMENT '拒绝理由',
  `status` varchar(45) NOT NULL DEFAULT 'pending' COMMENT '审核状态；pending:待审核；reject:已拒绝；agree:通过；cancel ：取消',
  `is_get` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否领取 0：未领取 1 ：已领取 ',
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='物品申请表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `leave_apply`
--

DROP TABLE IF EXISTS `leave_apply`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `leave_apply` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '申请表id',
  `leave_type_id` tinyint(4) NOT NULL COMMENT '请假类型id',
  `leave_type_name` varchar(45) NOT NULL COMMENT '请假类型名字',
  `start_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '请假时间',
  `end_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '结束时间',
  `staff_id` char(12) NOT NULL COMMENT '申请人员工id',
  `staff_name` varchar(45) NOT NULL COMMENT '员工姓名',
  `apply_reason` varchar(255) NOT NULL DEFAULT '' COMMENT '请假原因',
  `approve_staff_id` char(12) NOT NULL DEFAULT '' COMMENT '审核人id\n',
  `approve_staff_name` varchar(45) NOT NULL DEFAULT '' COMMENT '审核人姓名',
  `approve_time` datetime DEFAULT NULL,
  `approve_reason` varchar(100) NOT NULL DEFAULT '' COMMENT '拒绝理由',
  `status` varchar(45) NOT NULL DEFAULT 'pending' COMMENT '审核状态；pending:代审核；reject:已拒绝；agree:通过；cancel ：取消',
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COMMENT='假期申请表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `leave_type`
--

DROP TABLE IF EXISTS `leave_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `leave_type` (
  `id` tinyint(3) unsigned NOT NULL AUTO_INCREMENT COMMENT '请假类型id\n',
  `name` varchar(45) NOT NULL DEFAULT '' COMMENT '请假类型名称',
  `rule` varchar(255) NOT NULL DEFAULT '' COMMENT '规则描述',
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='请假类型表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `leave_type`
--

LOCK TABLES `leave_type` WRITE;
/*!40000 ALTER TABLE `leave_type` DISABLE KEYS */;
INSERT INTO `leave_type` VALUES (1,'心情假','一次休息半天','2020-07-14 09:08:31','2020-07-14 09:08:31'),(2,'健康假','一次休息半天 健康假','2020-07-14 09:08:31','2020-07-14 09:08:31'),(3,'端午节','休息三天， 需要补班','2020-07-16 10:05:17','2020-07-16 10:05:17'),(5,'国庆节','七天','2020-07-19 15:59:42','2020-07-19 15:59:42');
/*!40000 ALTER TABLE `leave_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menu`
--

DROP TABLE IF EXISTS `menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `menu` (
  `id` tinyint(3) unsigned NOT NULL AUTO_INCREMENT COMMENT '菜单id',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '菜单名字',
  `parent_id` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '上级菜单id 默认0顶级',
  `menu_code` varchar(30) DEFAULT NULL COMMENT '菜单标识',
  `url` varchar(45) NOT NULL DEFAULT '' COMMENT '跳转连接',
  `icon` varchar(30) NOT NULL DEFAULT 'fa fa-home' COMMENT '图标icon',
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `menu_code_UNIQUE` (`menu_code`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8 COMMENT='菜单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu`
--

LOCK TABLES `menu` WRITE;
/*!40000 ALTER TABLE `menu` DISABLE KEYS */;
INSERT INTO `menu` VALUES (1,'首页',0,'home','/admin/home/index','fa fa-home','2020-07-14 09:08:31','2020-07-14 09:08:31'),(2,'假期申请',0,'leaveapply','/admin/leaveapply/showlist','fa fa-home','2020-07-14 09:08:31','2020-07-14 09:08:31'),(3,'办公室预约',0,'officeapply','/admin/officeapply/showlist','fa fa-home','2020-07-14 09:08:31','2020-07-14 09:08:31'),(4,'物品申请',0,'goodsapply','/admin/goodsapply/showlist','fa fa-home','2020-07-14 09:08:31','2020-07-14 09:08:31'),(5,'审批管理',0,'approvalmanage','','fa fa-home','2020-07-14 09:08:31','2020-07-14 09:08:31'),(6,'假期审批',5,'leaveapprove','/admin/leaveapprove/showlist','fa fa-home','2020-07-14 09:08:31','2020-07-14 09:08:31'),(7,'办公室审批',5,'officeapprove','/admin/officeapprove/showlist','fa fa-home','2020-07-14 09:08:31','2020-07-14 09:08:31'),(8,'物品审批',5,'goodsapprove','/admin/goodsapprove/showlist','fa fa-home','2020-07-14 09:08:31','2020-07-14 09:08:31'),(9,'员工管理',0,'staffmanage','/admin/staffmanage/showlist','fa fa-home','2020-07-14 09:08:31','2020-07-14 09:08:31'),(10,'权限管理',0,'permissionmanage','/admin/permissionmanage/showrolelist','fa fa-home','2020-07-14 09:08:31','2020-07-14 09:08:31'),(11,'公告管理',0,'bulletinmanage','/admin/bulletinmanage/showlist','fa fa-home','2020-07-14 09:08:31','2020-07-14 09:08:31'),(12,'设置',0,'setting','','fa fa-home','2020-07-14 09:08:31','2020-07-14 09:08:31'),(13,'假期设置',12,'leavesetting','/admin/leavesetting/showlist','fa fa-home','2020-07-14 09:08:31','2020-07-14 09:08:31'),(14,'办公室设置',12,'officesetting','/admin/officesetting/showlist','fa fa-home','2020-07-14 09:08:31','2020-07-14 09:08:31'),(15,'物品设置',12,'goodssetting','/admin/goodssetting/showlist','fa fa-home','2020-07-14 09:08:31','2020-07-14 09:08:31');
/*!40000 ALTER TABLE `menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `office`
--

DROP TABLE IF EXISTS `office`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `office` (
  `id` tinyint(3) unsigned NOT NULL AUTO_INCREMENT,
  `serial_number` varchar(12) NOT NULL DEFAULT '' COMMENT '办公室编号',
  `name` varchar(45) NOT NULL DEFAULT '会议室' COMMENT '办公室名称\n',
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='办公室表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `office`
--

LOCK TABLES `office` WRITE;
/*!40000 ALTER TABLE `office` DISABLE KEYS */;
INSERT INTO `office` VALUES (1,'104','一楼104会议室','2020-07-14 09:08:31','2020-07-14 09:08:31'),(2,'303','三楼大会仪式','2020-07-14 09:08:31','2020-07-14 09:09:31'),(3,'522','五楼会议室1','2020-07-16 10:18:43','2020-07-16 10:19:08');
/*!40000 ALTER TABLE `office` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `office_apply`
--

DROP TABLE IF EXISTS `office_apply`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `office_apply` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '申请表id',
  `office_id` tinyint(4) NOT NULL COMMENT '办公表id',
  `office_name` varchar(45) NOT NULL COMMENT '办公室名字',
  `start_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '开始时间',
  `end_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '结束时间',
  `staff_id` char(12) NOT NULL COMMENT '申请人员工id',
  `staff_name` varchar(45) NOT NULL COMMENT '员工姓名',
  `apply_reason` varchar(255) NOT NULL DEFAULT '' COMMENT '请假原因',
  `approve_staff_id` char(12) NOT NULL DEFAULT '' COMMENT '审核人id\n',
  `approve_staff_name` varchar(45) NOT NULL DEFAULT '' COMMENT '审核人姓名',
  `approve_time` datetime DEFAULT NULL,
  `approve_reason` varchar(100) NOT NULL DEFAULT '' COMMENT '审核理由',
  `status` varchar(45) NOT NULL DEFAULT 'pending' COMMENT '审核状态；pending:以预约；cancel ：取消;close:关闭请求',
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_time` (`start_time`,`end_time`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='办公室申请表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `operate_menu`
--

DROP TABLE IF EXISTS `operate_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `operate_menu` (
  `id` tinyint(3) unsigned NOT NULL AUTO_INCREMENT,
  `operate_code` varchar(20) NOT NULL COMMENT '菜单权限操作code',
  `operate_name` varchar(45) NOT NULL COMMENT '菜单操作名字',
  `menu_code` varchar(30) NOT NULL COMMENT '菜单code',
  `methods` varchar(200) NOT NULL COMMENT '对应操作具有的方法',
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8 COMMENT='菜单权限操作表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `operate_menu`
--

LOCK TABLES `operate_menu` WRITE;
/*!40000 ALTER TABLE `operate_menu` DISABLE KEYS */;
INSERT INTO `operate_menu` VALUES (1,'apply','申请假期','leaveapply','applyLeave','2020-07-14 09:08:31','2020-07-14 09:08:31'),(2,'cancel','取消申请','leaveapply','cancelLeave','2020-07-14 09:08:31','2020-07-14 09:08:31'),(3,'apply','预约办公室','officeapply','applyOffice','2020-07-14 09:08:31','2020-07-14 09:08:31'),(4,'cancel','取消预约','officeapply','cancelOffice','2020-07-14 09:08:31','2020-07-14 09:08:31'),(5,'apply','申请物品','goodsapply','applyGoods','2020-07-14 09:08:31','2020-07-14 09:08:31'),(6,'cancel','取消申请','goodsapply','cancelGoods','2020-07-14 09:08:31','2020-07-14 09:08:31'),(7,'agree','同意申请','leaveapprove','agreeLeave','2020-07-14 09:08:31','2020-07-14 09:08:31'),(8,'reject','拒绝申请','leaveapprove','rejectLeave','2020-07-14 09:08:31','2020-07-14 09:08:31'),(9,'close','关闭预约','officeapprove','closeOffice','2020-07-14 09:08:31','2020-07-14 09:08:31'),(10,'confirm','确认领取','goodsapprove','confirmGoods','2020-07-14 09:08:31','2020-07-14 09:08:31'),(11,'agree','同意申请','goodsapprove','agreeGoods','2020-07-14 09:08:31','2020-07-14 09:08:31'),(12,'reject','拒绝申请','goodsapprove','rejectGoods','2020-07-14 09:08:31','2020-07-14 09:08:31'),(13,'add','添加员工','staffamanage','addStaff','2020-07-14 09:08:31','2020-07-14 09:08:31'),(14,'freeze','冻结解冻','staffmanage','isFreeze','2020-07-14 09:08:31','2020-07-14 09:08:31'),(15,'delete','删除员工','staffmanage','deleteStaff','2020-07-14 09:08:31','2020-07-14 09:08:31'),(16,'add','添加角色','permissionmanage','addRole','2020-07-14 09:08:31','2020-07-14 09:08:31'),(17,'modify','修改角色','permissionmanage','modifyRole','2020-07-14 09:08:31','2020-07-14 09:08:31'),(18,'add','添加公告','bulletinmanage','addBulletion','2020-07-14 09:08:31','2020-07-14 09:08:31'),(19,'delete','删除角色','permissionmanage','deleteRole','2020-07-14 09:08:31','2020-07-14 09:08:31'),(20,'modify','修改角色','staffmanage','modifyStaffRole','2020-07-14 09:08:31','2020-07-14 09:08:31'),(21,'set','设置审核','staffmanage','isApprover','2020-07-14 09:08:31','2020-07-14 09:08:31'),(22,'delete','删除公告','bulletinmanage','deleteBulletion','2020-07-14 09:08:31','2020-07-14 09:08:31'),(23,'add','添加假期','leavesetting','addLeave','2020-07-14 09:08:31','2020-07-14 09:08:31'),(24,'modify','修改假期','leavesetting','modifyLeave','2020-07-14 09:08:31','2020-07-14 09:08:31'),(25,'delete','删除假期','leavesetting','deleteLeave','2020-07-14 09:08:31','2020-07-14 09:08:31'),(26,'add','添加办公室','officesetting','addOffice','2020-07-14 09:08:31','2020-07-14 09:08:31'),(27,'modify','修改办公室','officesetting','modifyOffice','2020-07-14 09:08:31','2020-07-14 09:08:31'),(28,'delete','删除办公室','officesetting','deleteOffice','2020-07-14 09:08:31','2020-07-14 09:08:31'),(29,'add','添加物品','goodssetting','addGoods','2020-07-14 09:08:31','2020-07-14 09:08:31'),(30,'modify','修改物品','goodssetting','modifyGoods','2020-07-14 09:08:31','2020-07-14 09:08:31'),(31,'delete','删除物品','goodssetting','deleteGoods','2020-07-14 09:08:31','2020-07-14 09:08:31');
/*!40000 ALTER TABLE `operate_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role` (
  `id` tinyint(3) unsigned NOT NULL AUTO_INCREMENT COMMENT '角色id',
  `name` varchar(45) NOT NULL DEFAULT '' COMMENT '角色名称',
  `menu_permission` text NOT NULL COMMENT 'json数据：菜单权限',
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='角色表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `staff`
--

DROP TABLE IF EXISTS `staff`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `staff` (
  `id` char(12) NOT NULL COMMENT '员工编号',
  `name` varchar(45) NOT NULL DEFAULT '' COMMENT '员工姓名',
  `gender` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '性别:1:男 0:女',
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号码',
  `phone_code` varchar(10) NOT NULL DEFAULT '+86' COMMENT '手机区号',
  `address` varchar(100) NOT NULL DEFAULT '' COMMENT '地址',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(10) DEFAULT '' COMMENT '密码盐',
  `email` varchar(45) NOT NULL DEFAULT '' COMMENT '邮箱',
  `profile_pictrue` varchar(50) NOT NULL DEFAULT '' COMMENT '头像',
  `title` varchar(60) NOT NULL DEFAULT '员工' COMMENT '职务姓名\n',
  `is_valid` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否有效 1：有效，0：无效',
  `is_freeze` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否冻结 1 冻结，0：正常',
  `is_approver` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否有审批权限1：有；0：无',
  `last_login_time` datetime DEFAULT NULL COMMENT '上次登陆时间',
  `last_login_ip` varchar(20) NOT NULL DEFAULT '' COMMENT '上次登陆ip',
  `role_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '角色id,默认0 没有菜单权限',
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  KEY `role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='员工表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `staff`
--

LOCK TABLES `staff` WRITE;
/*!40000 ALTER TABLE `staff` DISABLE KEYS */;
INSERT INTO `staff` VALUES ('000000','超级管理员',1,'','+86','','26b272636faa7890e6f022f56b361079','sa','','','员工',1,0,1,'2020-07-20 08:30:00','127.0.0.1',0,'2020-07-14 09:08:31','2020-07-20 08:30:00');
/*!40000 ALTER TABLE `staff` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-07-20  8:31:21
