# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.10)
# Database: interop2016
# Generation Time: 2016-04-25 05:24:12 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table browsing
# ------------------------------------------------------------

DROP TABLE IF EXISTS `browsing`;

CREATE TABLE `browsing` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `src_ip` varchar(255) NOT NULL,
  `dst_ip` varchar(255) NOT NULL,
  `src_port` int(11) NOT NULL,
  `dst_port` int(11) NOT NULL,
  `timestamp` datetime NOT NULL,
  `title` text,
  `url` text,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `browsing_time` float DEFAULT NULL,
  `download` int(11) DEFAULT NULL,
  `domain` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_browsing_on_created_at` (`created_at`),
  KEY `index_browsing_on_src_ip` (`src_ip`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `browsing` WRITE;
/*!40000 ALTER TABLE `browsing` DISABLE KEYS */;

INSERT INTO `browsing` (`id`, `src_ip`, `dst_ip`, `src_port`, `dst_port`, `timestamp`, `title`, `url`, `created_at`, `browsing_time`, `download`, `domain`)
VALUES
	(1,'1.1.1.1','2.2.2.2',123,80,'2016-04-01 00:00:00','Jojo','http://jojo.com','2016-04-25 13:58:30',10,0,'jojo.com'),
	(2,'10.24.1.10','2.2.2.2',1233,80,'2016-04-01 00:00:00','Jojo','http://jojo.com','2016-04-25 14:02:27',10,0,'jojo.com'),
	(3,'1.1.1.1','3.3.3.3',123,80,'2016-04-01 00:00:00','scryed','http://www.s-cry-ed.net','2016-04-25 14:05:11',10,0,'www.s-cry-ed.net'),
	(4,'1.1.1.1','5.5.5.5',123,80,'2016-04-01 00:00:00','pixiv','http://www.pixiv.net','2016-04-25 14:07:18',10,0,'www.pixiv.net'),
	(5,'10.24.1.11','3.3.3.3',123,80,'2016-04-01 00:00:00','scryed','http://www.s-cry-ed.net','2016-04-25 14:21:33',10,0,'www.s-cry-ed.net'),
	(6,'10.24.1.111','2.2.2.2',123,80,'2016-04-01 00:00:00','Jojo','http://jojo.com','2016-04-25 14:22:46',10,0,'jojo.com');

/*!40000 ALTER TABLE `browsing` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table meta
# ------------------------------------------------------------

DROP TABLE IF EXISTS `meta`;

CREATE TABLE `meta` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `value` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_meta_on_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table word
# ------------------------------------------------------------

DROP TABLE IF EXISTS `word`;

CREATE TABLE `word` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `count` text,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `word` WRITE;
/*!40000 ALTER TABLE `word` DISABLE KEYS */;

INSERT INTO `word` (`id`, `name`, `count`, `created_at`)
VALUES
	(16,'foo','1','2016-04-19 06:49:56'),
	(17,'var','2','2016-04-19 06:49:56'),
	(18,'baz','13','2016-04-19 06:49:56'),
	(19,'foo','12','2016-04-19 06:49:56'),
	(20,'bar','2','2016-04-19 06:49:56'),
	(21,'baz','1','2016-04-19 06:49:56'),
	(22,'foo','1','2016-04-19 06:50:20'),
	(23,'var','2','2016-04-19 06:50:20'),
	(24,'baz','12','2016-04-19 06:50:20'),
	(25,'foo','1','2016-04-19 06:50:20'),
	(26,'bar','2','2016-04-19 06:50:20'),
	(27,'baz','123','2016-04-19 06:50:20');

/*!40000 ALTER TABLE `word` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
