/*CREATE DATABASE `service_config` ;*/

/*Table structure for table `userinfo` */

DROP TABLE IF EXISTS `order`;

CREATE TABLE `porder` (
         `id` INT(10) NOT NULL AUTO_INCREMENT,
         `order_number` VARCHAR(16) NOT NULL,
         `user_id` INT(10) NOT NULL,
         `payment_id` INT(10) DEFAULT 0,
         `status` VARCHAR(16) NOT NULL,
         `created_time` DATETIME NOT NULL,
         `updated_time` DATETIME DEFAULT NULL,
         PRIMARY KEY  (`id`)
) ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
