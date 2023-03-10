-- MySQL dump 10.13  Distrib 8.0.30, for Win64 (x86_64)
--
-- Host: localhost    Database: shop
-- ------------------------------------------------------
-- Server version	8.0.30

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `admin`
--

DROP TABLE IF EXISTS `admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(20) DEFAULT NULL,
  `password` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin`
--

LOCK TABLES `admin` WRITE;
/*!40000 ALTER TABLE `admin` DISABLE KEYS */;
INSERT INTO `admin` VALUES (1,'admin','admin');
/*!40000 ALTER TABLE `admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `adress`
--

DROP TABLE IF EXISTS `adress`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `adress` (
  `id` int NOT NULL AUTO_INCREMENT,
  `consignee` varchar(20) DEFAULT NULL,
  `postcode` char(6) DEFAULT NULL,
  `tel` char(11) DEFAULT NULL,
  `mid` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `adress`
--

LOCK TABLES `adress` WRITE;
/*!40000 ALTER TABLE `adress` DISABLE KEYS */;
/*!40000 ALTER TABLE `adress` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `buycar`
--

DROP TABLE IF EXISTS `buycar`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `buycar` (
  `id` int NOT NULL AUTO_INCREMENT,
  `pid` int DEFAULT NULL,
  `number` int DEFAULT NULL,
  `mid` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `buycar_FK` (`pid`),
  KEY `buycar_FK_1` (`mid`),
  CONSTRAINT `buycar_FK` FOREIGN KEY (`pid`) REFERENCES `product` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `buycar_FK_1` FOREIGN KEY (`mid`) REFERENCES `member` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `buycar`
--

LOCK TABLES `buycar` WRITE;
/*!40000 ALTER TABLE `buycar` DISABLE KEYS */;
INSERT INTO `buycar` VALUES (14,2,5,2);
/*!40000 ALTER TABLE `buycar` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `category` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT NULL,
  `fid` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` VALUES (1,'?????????????????????',NULL),(3,'????????????????????????',NULL),(4,'??????????????????????????????',NULL),(5,'??????????????????????????????',NULL),(6,'?????????????????????',NULL),(7,'???????????????????????????',NULL),(8,'??????????????????????????????',NULL),(9,'???????????????',NULL),(10,'???????????????????????????',NULL),(11,'???????????????',NULL);
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `member`
--

DROP TABLE IF EXISTS `member`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `member` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(20) DEFAULT NULL,
  `password` varchar(20) DEFAULT NULL,
  `name` varchar(20) DEFAULT NULL,
  `identity` char(18) DEFAULT NULL,
  `tel` char(11) DEFAULT NULL,
  `email` varchar(20) DEFAULT NULL,
  `balance` decimal(10,0) DEFAULT NULL,
  `state` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `member`
--

LOCK TABLES `member` WRITE;
/*!40000 ALTER TABLE `member` DISABLE KEYS */;
INSERT INTO `member` VALUES (1,'tom','123456',NULL,NULL,NULL,'123@qq.com',NULL,0),(2,'123','123',NULL,NULL,NULL,'1@q.c',NULL,0),(4,'124','123',NULL,NULL,NULL,'1@q.c',NULL,0),(5,'125','123',NULL,NULL,NULL,'1@q.c',NULL,0),(6,'126','123',NULL,NULL,NULL,'1@q.c',NULL,0),(7,'xyy','123456',NULL,NULL,NULL,'1310500986@qq.com',NULL,0);
/*!40000 ALTER TABLE `member` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `message`
--

DROP TABLE IF EXISTS `message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `message` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(20) DEFAULT NULL,
  `content` varchar(200) DEFAULT NULL,
  `date` datetime DEFAULT NULL,
  `mid` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `message`
--

LOCK TABLES `message` WRITE;
/*!40000 ALTER TABLE `message` DISABLE KEYS */;
INSERT INTO `message` VALUES (1,'??????','22222','2023-01-15 16:54:20',1),(2,'??????','???????????????????????????','2023-01-15 22:09:17',7);
/*!40000 ALTER TABLE `message` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order`
--

DROP TABLE IF EXISTS `order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order` (
  `id` int NOT NULL AUTO_INCREMENT,
  `time` datetime DEFAULT NULL,
  `state` int DEFAULT NULL,
  `mid` int DEFAULT NULL,
  `address` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order`
--

LOCK TABLES `order` WRITE;
/*!40000 ALTER TABLE `order` DISABLE KEYS */;
INSERT INTO `order` VALUES (5,'2023-01-15 17:51:53',1,1,'?????????,132222222,Yuhong'),(6,'2023-01-15 17:56:34',1,1,'11,22,3333'),(7,'2023-01-15 18:00:17',1,1,'222,333,4444'),(8,'2023-01-15 18:10:07',1,1,'111,222,567'),(9,'2023-01-15 18:32:08',1,1,'?????????,13066704022,Yuhong???XijiangJie'),(10,'2023-01-15 18:42:27',1,1,'??????,12345678,Yuhong'),(11,'2023-01-15 18:49:06',1,1,'??????,14325876,Shenyang'),(12,'2023-01-15 22:07:46',1,7,'xyy,13056789876,Yuhong shenyang'),(13,'2023-01-15 22:08:53',1,7,'xyy,13056789876,Yuhong');
/*!40000 ALTER TABLE `order` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orderdetail`
--

DROP TABLE IF EXISTS `orderdetail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orderdetail` (
  `id` int NOT NULL AUTO_INCREMENT,
  `pid` int DEFAULT NULL,
  `oid` int DEFAULT NULL,
  `num` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orderdetail`
--

LOCK TABLES `orderdetail` WRITE;
/*!40000 ALTER TABLE `orderdetail` DISABLE KEYS */;
INSERT INTO `orderdetail` VALUES (1,7,10,7),(2,15,10,1),(3,2,10,2),(4,1,11,3),(5,4,11,4),(6,6,12,5),(7,14,12,4),(8,8,12,2),(9,1,13,10);
/*!40000 ALTER TABLE `orderdetail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product`
--

DROP TABLE IF EXISTS `product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `type` varchar(20) DEFAULT NULL,
  `color` varchar(20) DEFAULT NULL,
  `price` decimal(10,0) DEFAULT NULL,
  `state` int DEFAULT NULL,
  `number` int DEFAULT NULL,
  `img` varchar(50) DEFAULT NULL,
  `imgPath` varchar(50) DEFAULT NULL,
  `cid` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `product_FK` (`cid`),
  CONSTRAINT `product_FK` FOREIGN KEY (`cid`) REFERENCES `category` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product`
--

LOCK TABLES `product` WRITE;
/*!40000 ALTER TABLE `product` DISABLE KEYS */;
INSERT INTO `product` VALUES (1,'????????????','??????10','??????',3000,0,77,'xiaomi10.jpg','../static/resources/images/',10),(2,'????????????','??????10','??????',3000,0,100,'xiaomi10.jpg','../static/resources/images/',10),(3,'????????????','??????10','??????',3000,0,100,'xiaomi10.jpg','../static/resources/images/',10),(4,'????????????','??????10','??????',3000,0,96,'xiaomi10.jpg','../static/resources/images/',10),(5,'????????????','??????10','??????',3000,0,100,'xiaomi10.jpg','../static/resources/images/',10),(6,'????????????','??????10','??????',3000,0,95,'xiaomi10.jpg','../static/resources/images/',10),(7,'????????????','??????10','??????',2000,0,100,'xiaomi10.jpg','../static/resources/images/',10),(8,'????????????','??????10','??????',1000,0,98,'xiaomi10.jpg','../static/resources/images/',10),(9,'????????????','??????10','??????',3000,0,100,'xiaomi10.jpg','../static/resources/images/',10),(10,'????????????','??????10','??????',3000,0,100,'xiaomi10.jpg','../static/resources/images/',10),(11,'????????????','??????10','??????',3000,0,100,'xiaomi10.jpg','../static/resources/images/',10),(12,'????????????','??????10','??????',3000,0,100,'xiaomi10.jpg','../static/resources/images/',10),(13,'????????????','??????10','??????',3000,0,100,'xiaomi10.jpg','../static/resources/images/',10),(14,'????????????','??????10','??????',3000,0,96,'xiaomi10.jpg','../static/resources/images/',10),(15,'????????????','??????10','??????',4000,0,100,'xiaomi10.jpg','../static/resources/images/',10);
/*!40000 ALTER TABLE `product` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'shop'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-01-17 14:13:29
