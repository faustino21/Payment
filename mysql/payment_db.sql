-- --------------------------------------------------------
-- Host:                         localhost
-- Server version:               8.0.28 - MySQL Community Server - GPL
-- Server OS:                    Win64
-- HeidiSQL Version:             11.3.0.6295
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for payment
CREATE DATABASE IF NOT EXISTS `payment` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `payment`;

-- Dumping structure for table payment.customers
CREATE TABLE IF NOT EXISTS `customers` (
  `customer_id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `password` varchar(50) NOT NULL,
  `saldo` int NOT NULL DEFAULT '0',
  `address` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `token` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`customer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table payment.customers: ~2 rows (approximately)
/*!40000 ALTER TABLE `customers` DISABLE KEYS */;
INSERT INTO `customers` (`customer_id`, `name`, `password`, `saldo`, `address`, `token`, `updated_at`, `created_at`, `deleted_at`) VALUES
	(1, 'tino', '123456', 980000, 'Jakarta', NULL, '2022-04-25 23:10:20', '2022-04-25 11:11:25', NULL),
	(2, 'Gary', 'wiefnwie', 1000000, 'Jakarta', NULL, '2022-04-25 13:12:25', '2022-04-25 11:11:49', NULL);
/*!40000 ALTER TABLE `customers` ENABLE KEYS */;

-- Dumping structure for table payment.merchant
CREATE TABLE IF NOT EXISTS `merchant` (
  `merchant_id` int NOT NULL AUTO_INCREMENT,
  `merchant_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `saldo` int NOT NULL DEFAULT '0',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`merchant_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table payment.merchant: ~3 rows (approximately)
/*!40000 ALTER TABLE `merchant` DISABLE KEYS */;
INSERT INTO `merchant` (`merchant_id`, `merchant_name`, `saldo`, `update_at`, `created_at`, `deleted_at`) VALUES
	(1, 'tokopakedi', 11500000, '2022-04-25 19:02:43', '2022-04-25 13:20:24', NULL),
	(2, 'BurungDara Indonesia', 50000000, '2022-04-25 13:23:10', '2022-04-25 13:22:42', NULL),
	(3, 'Zalorant', 8020232, '2022-04-25 23:10:20', '2022-04-25 13:30:09', NULL);
/*!40000 ALTER TABLE `merchant` ENABLE KEYS */;

-- Dumping structure for table payment.transfer
CREATE TABLE IF NOT EXISTS `transfer` (
  `transfer_id` int NOT NULL AUTO_INCREMENT,
  `id_customer` int DEFAULT NULL,
  `id_merchant` int DEFAULT NULL,
  `cost` int DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`transfer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table payment.transfer: ~3 rows (approximately)
/*!40000 ALTER TABLE `transfer` DISABLE KEYS */;
INSERT INTO `transfer` (`transfer_id`, `id_customer`, `id_merchant`, `cost`, `updated_at`, `created_at`, `deleted_at`) VALUES
	(3, 1, 1, 500000, '2022-04-25 22:34:25', '2022-04-25 22:34:24', NULL),
	(4, 1, 3, 500000, '2022-04-25 19:09:42', '2022-04-25 19:09:42', NULL),
	(5, 1, 3, 20000, '2022-04-25 23:10:20', '2022-04-25 23:10:20', NULL);
/*!40000 ALTER TABLE `transfer` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
