CREATE DATABASE  IF NOT EXISTS `restuarant` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `restuarant`;
-- MySQL dump 10.13  Distrib 8.0.42, for Win64 (x86_64)
--
-- Host: localhost    Database: restuarant
-- ------------------------------------------------------
-- Server version	8.0.42

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `customers`
--

DROP TABLE IF EXISTS `customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `customers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `address` varchar(500) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  KEY `idx_customers_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customers`
--

LOCK TABLES `customers` WRITE;
/*!40000 ALTER TABLE `customers` DISABLE KEYS */;
INSERT INTO `customers` VALUES (1,'พี่อุ้ม','oam@example.com','+1-555-1001','100 Customer St, Residential Area','2025-07-23 12:56:47','2025-07-23 13:11:58'),(2,'Jane Smith','jane.smith@example.com','+1-555-1002','200 Buyer Ave, Suburb','2025-07-23 12:56:47','2025-07-23 12:56:47'),(3,'Bob Wilson','bob.wilson@example.com','+1-555-1003','300 Client Rd, Downtown','2025-07-23 12:56:47','2025-07-23 12:56:47'),(4,'Alice Johnson','alice.johnson@example.com','+1-555-1004','400 User Blvd, Midtown','2025-07-23 12:56:47','2025-07-23 12:56:47'),(5,'Charlie Brown','charlie.brown@example.com','+1-555-1005','500 Guest Lane, Uptown','2025-07-23 12:56:47','2025-07-23 12:56:47'),(6,'น้องต่อ','tor@demo.com','09998899','MRT','2025-07-23 13:18:16','2025-07-23 13:18:16');
/*!40000 ALTER TABLE `customers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menu_items`
--

DROP TABLE IF EXISTS `menu_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `menu_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `restaurant_id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` text,
  `price` decimal(10,2) NOT NULL,
  `category` varchar(100) DEFAULT NULL,
  `is_available` tinyint(1) DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_menu_items_restaurant_id` (`restaurant_id`),
  KEY `idx_menu_items_category` (`category`),
  CONSTRAINT `menu_items_ibfk_1` FOREIGN KEY (`restaurant_id`) REFERENCES `restaurants` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu_items`
--

LOCK TABLES `menu_items` WRITE;
/*!40000 ALTER TABLE `menu_items` DISABLE KEYS */;
INSERT INTO `menu_items` VALUES (1,1,'Margherita Pizza','Fresh tomato sauce, mozzarella cheese, fresh basil',12.99,'Pizza',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(2,1,'Pepperoni Pizza','Tomato sauce, mozzarella, pepperoni slices',14.99,'Pizza',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(3,1,'Supreme Pizza','Pepperoni, sausage, bell peppers, onions, mushrooms',17.99,'Pizza',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(4,1,'Caesar Salad','Romaine lettuce, croutons, parmesan, caesar dressing',8.99,'Salad',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(5,1,'Garlic Bread','Fresh baked bread with garlic butter',4.99,'Appetizer',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(6,1,'Tiramisu','Classic Italian dessert with coffee and mascarpone',6.99,'Dessert',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(7,2,'Classic Burger','Beef patty, lettuce, tomato, onion, pickles',9.99,'Burger',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(8,2,'Cheeseburger','Classic burger with american cheese',10.99,'Burger',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(9,2,'BBQ Bacon Burger','Beef patty, bacon, BBQ sauce, onion rings',12.99,'Burger',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(10,2,'Chicken Sandwich','Grilled chicken breast, mayo, lettuce',8.99,'Sandwich',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(11,2,'French Fries','Crispy golden fries',3.99,'Side',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(12,2,'Onion Rings','Beer-battered onion rings',4.99,'Side',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(13,2,'Milkshake','Vanilla, chocolate, or strawberry',4.99,'Beverage',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(14,3,'Salmon Roll','Fresh salmon, avocado, cucumber',8.99,'Roll',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(15,3,'Tuna Roll','Fresh tuna, avocado',9.99,'Roll',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(16,3,'California Roll','Crab, avocado, cucumber',7.99,'Roll',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(17,3,'Salmon Sashimi','Fresh salmon slices (6 pieces)',12.99,'Sashimi',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(18,3,'Tuna Sashimi','Fresh tuna slices (6 pieces)',14.99,'Sashimi',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(19,3,'Miso Soup','Traditional soybean soup',3.99,'Soup',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(20,3,'Edamame','Steamed and salted soybeans',4.99,'Appetizer',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(21,4,'Beef Tacos','Seasoned ground beef, lettuce, cheese (3 tacos)',8.99,'Tacos',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(22,4,'Chicken Tacos','Grilled chicken, salsa, cheese (3 tacos)',9.99,'Tacos',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(23,4,'Fish Tacos','Grilled fish, cabbage slaw, lime (3 tacos)',11.99,'Tacos',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(24,4,'Beef Burrito','Large flour tortilla with beef, beans, rice',10.99,'Burrito',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(25,4,'Chicken Quesadilla','Grilled chicken and cheese in tortilla',8.99,'Quesadilla',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(26,4,'Guacamole & Chips','Fresh guacamole with tortilla chips',5.99,'Appetizer',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(27,4,'Churros','Fried pastry with cinnamon sugar',4.99,'Dessert',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(28,5,'Pad Thai','Stir-fried rice noodles with shrimp or chicken',12.99,'Noodles',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(29,5,'Green Curry','Coconut curry with vegetables and choice of protein',13.99,'Curry',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(30,5,'Tom Yum Soup','Spicy and sour soup with shrimp',8.99,'Soup',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(31,5,'Spring Rolls','Fresh vegetables wrapped in rice paper (4 rolls)',6.99,'Appetizer',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(32,5,'Mango Sticky Rice','Sweet sticky rice with fresh mango',5.99,'Dessert',1,'2025-07-23 12:56:47','2025-07-23 12:56:47'),(33,5,'Thai Iced Tea','Sweet tea with condensed milk',3.99,'Beverage',1,'2025-07-23 12:56:47','2025-07-23 12:56:47');
/*!40000 ALTER TABLE `menu_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_items`
--

DROP TABLE IF EXISTS `order_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_id` int NOT NULL,
  `menu_item_id` int NOT NULL,
  `quantity` int NOT NULL,
  `unit_price` decimal(10,2) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_order_items_order_id` (`order_id`),
  KEY `idx_order_items_menu_item_id` (`menu_item_id`),
  CONSTRAINT `order_items_ibfk_1` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`) ON DELETE CASCADE,
  CONSTRAINT `order_items_ibfk_2` FOREIGN KEY (`menu_item_id`) REFERENCES `menu_items` (`id`) ON DELETE RESTRICT,
  CONSTRAINT `order_items_chk_1` CHECK ((`quantity` > 0))
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_items`
--

LOCK TABLES `order_items` WRITE;
/*!40000 ALTER TABLE `order_items` DISABLE KEYS */;
INSERT INTO `order_items` VALUES (1,1,1,1,12.99),(2,1,4,1,8.99),(3,2,8,1,10.99),(4,2,11,1,3.99),(5,2,13,1,4.99),(6,3,14,1,8.99),(7,3,18,1,14.99),(8,3,19,1,3.99),(9,4,22,1,9.99),(10,4,26,1,5.99),(11,5,28,1,12.99),(12,5,33,1,3.99),(13,6,7,1,9.99),(14,6,11,1,3.99);
/*!40000 ALTER TABLE `order_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `customer_id` int NOT NULL,
  `restaurant_id` int NOT NULL,
  `total_amount` decimal(10,2) NOT NULL,
  `status` enum('pending','confirmed','preparing','ready','delivered','cancelled') DEFAULT 'pending',
  `order_date` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `delivery_address` varchar(500) DEFAULT NULL,
  `notes` text,
  PRIMARY KEY (`id`),
  KEY `idx_orders_customer_id` (`customer_id`),
  KEY `idx_orders_restaurant_id` (`restaurant_id`),
  KEY `idx_orders_status` (`status`),
  KEY `idx_orders_date` (`order_date`),
  CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`) ON DELETE RESTRICT,
  CONSTRAINT `orders_ibfk_2` FOREIGN KEY (`restaurant_id`) REFERENCES `restaurants` (`id`) ON DELETE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` VALUES (1,1,1,21.98,'confirmed','2025-07-23 12:56:47','100 Customer St, Residential Area','Please ring doorbell, leave at door'),(2,2,2,19.97,'preparing','2025-07-23 12:56:47','200 Buyer Ave, Suburb','Extra sauce on burger, no pickles'),(3,3,3,25.97,'ready','2025-07-23 12:56:47','300 Client Rd, Downtown','Hold the wasabi, extra ginger'),(4,4,4,18.98,'delivered','2025-07-23 12:56:47','400 User Blvd, Midtown','Delivered successfully'),(5,5,5,16.98,'pending','2025-07-23 12:56:47','500 Guest Lane, Uptown','Call when arriving'),(6,1,2,15.98,'cancelled','2025-07-23 12:56:47','100 Customer St, Residential Area','Customer cancelled order');
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `restaurants`
--

DROP TABLE IF EXISTS `restaurants`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `restaurants` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `address` varchar(500) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `cuisine_type` varchar(100) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `restaurants`
--

LOCK TABLES `restaurants` WRITE;
/*!40000 ALTER TABLE `restaurants` DISABLE KEYS */;
INSERT INTO `restaurants` VALUES (1,'Pizza Palace','123 Main St, Downtown','+1-555-0101','info@pizzapalace.com','Italian','2025-07-23 12:56:47','2025-07-23 12:56:47'),(2,'Burger Barn','456 Oak Ave, Midtown','+1-555-0102','orders@burgerbarn.com','American','2025-07-23 12:56:47','2025-07-23 12:56:47'),(3,'Sushi Zen','789 Pine Rd, Uptown','+1-555-0103','hello@sushizen.com','Japanese','2025-07-23 12:56:47','2025-07-23 12:56:47'),(4,'Taco Fiesta','321 Elm St, Downtown','+1-555-0104','contact@tacofiesta.com','Mexican','2025-07-23 12:56:47','2025-07-23 12:56:47'),(5,'Thai Garden','654 Maple Ave, Eastside','+1-555-0105','info@thaigarden.com','Thai','2025-07-23 12:56:47','2025-07-23 12:56:47');
/*!40000 ALTER TABLE `restaurants` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-07-23 22:07:15
