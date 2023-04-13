-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Apr 13, 2023 at 06:44 PM
-- Server version: 10.4.25-MariaDB
-- PHP Version: 8.1.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `kstyle_be`
--

-- --------------------------------------------------------

--
-- Table structure for table `like_review`
--

CREATE TABLE `like_review` (
  `review_product_id_review` bigint(20) UNSIGNED NOT NULL,
  `member_id_member` bigint(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `members`
--

CREATE TABLE `members` (
  `id_member` bigint(20) UNSIGNED NOT NULL,
  `username` varchar(255) NOT NULL,
  `gender` varchar(255) DEFAULT NULL,
  `skintype` varchar(255) DEFAULT NULL,
  `skincolor` varchar(255) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `members`
--

INSERT INTO `members` (`id_member`, `username`, `gender`, `skintype`, `skincolor`, `created_at`, `updated_at`, `deleted_at`) VALUES
(22, 'fitrinurmalasari', 'female', 'normal', 'pale ivory', '2023-04-13 20:22:47.917', '2023-04-13 20:22:47.917', NULL),
(23, 'nuraisyah', 'female', 'acne', 'porcelain', '2023-04-13 20:23:33.793', '2023-04-13 20:23:33.793', NULL),
(24, 'nabilahm', 'female', 'sensitive', 'ivory', '2023-04-13 20:24:26.519', '2023-04-13 20:24:26.519', NULL),
(25, 'darasalsabila', 'female', 'oily', 'sand', '2023-04-13 20:25:02.714', '2023-04-13 20:25:02.714', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id_product` bigint(20) UNSIGNED NOT NULL,
  `name_product` longtext DEFAULT NULL,
  `price` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id_product`, `name_product`, `price`, `created_at`, `updated_at`, `deleted_at`) VALUES
(5, 'TOVEGAN - Red Remedy Toner', 29, '2023-04-13 18:47:26.499', '2023-04-13 18:47:26.499', NULL),
(6, 'Theonleaf - Vegan Greenery Repair Cica', 31, '2023-04-13 18:48:30.998', '2023-04-13 18:48:30.998', NULL),
(7, 'GINGER6 - Active Water Cream 50ml', 20, '2023-04-13 18:48:59.604', '2023-04-13 18:48:59.604', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `review_products`
--

CREATE TABLE `review_products` (
  `id_review` bigint(20) UNSIGNED NOT NULL,
  `product_id` bigint(20) UNSIGNED NOT NULL,
  `member_id` bigint(20) UNSIGNED NOT NULL,
  `desc_review` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `review_products`
--

INSERT INTO `review_products` (`id_review`, `product_id`, `member_id`, `desc_review`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 5, 22, 'I\'ve been using this cosmetic for a few weeks now and I\'m already seeing a noticeable difference in the appearance of my skin.', '2023-04-13 20:34:47.326', '2023-04-13 20:34:47.326', NULL),
(2, 5, 24, 'I have sensitive skin and this cosmetic has been a lifesaver. It doesn\'t cause any irritation and it looks great all day.', '2023-04-13 22:43:10.884', '2023-04-13 22:43:10.884', NULL),
(3, 6, 23, 'This cosmetic has been a game changer for me. It\'s helped clear up my acne and leaves my skin looking flawless.', '2023-04-13 23:05:32.969', '2023-04-13 23:05:32.969', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `like_review`
--
ALTER TABLE `like_review`
  ADD PRIMARY KEY (`review_product_id_review`,`member_id_member`),
  ADD KEY `fk_like_review_member` (`member_id_member`);

--
-- Indexes for table `members`
--
ALTER TABLE `members`
  ADD PRIMARY KEY (`id_member`),
  ADD UNIQUE KEY `idx_members_username` (`username`),
  ADD UNIQUE KEY `username` (`username`),
  ADD UNIQUE KEY `username_2` (`username`),
  ADD KEY `idx_members_deleted_at` (`deleted_at`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id_product`),
  ADD KEY `idx_products_deleted_at` (`deleted_at`);

--
-- Indexes for table `review_products`
--
ALTER TABLE `review_products`
  ADD PRIMARY KEY (`id_review`),
  ADD KEY `idx_review_products_deleted_at` (`deleted_at`),
  ADD KEY `fk_members_reviews` (`member_id`),
  ADD KEY `fk_products_reviews` (`product_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `members`
--
ALTER TABLE `members`
  MODIFY `id_member` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=26;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id_product` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `review_products`
--
ALTER TABLE `review_products`
  MODIFY `id_review` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `like_review`
--
ALTER TABLE `like_review`
  ADD CONSTRAINT `fk_like_review_member` FOREIGN KEY (`member_id_member`) REFERENCES `members` (`id_member`),
  ADD CONSTRAINT `fk_like_review_review_product` FOREIGN KEY (`review_product_id_review`) REFERENCES `review_products` (`id_review`);

--
-- Constraints for table `review_products`
--
ALTER TABLE `review_products`
  ADD CONSTRAINT `fk_members_reviews` FOREIGN KEY (`member_id`) REFERENCES `members` (`id_member`),
  ADD CONSTRAINT `fk_products_reviews` FOREIGN KEY (`product_id`) REFERENCES `products` (`id_product`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
