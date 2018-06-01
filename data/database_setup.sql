CREATE DATABASE `url` /*!40100 DEFAULT CHARACTER SET utf8 */;

CREATE TABLE `url`.`category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `category_name` varchar(100) NOT NULL,
  `category_logo` varchar(100) NOT NULL,
  `parent_category_id` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `url`.`url` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `url` varchar(200) NOT NULL,
  `logo` varchar(100) NOT NULL,
  `introduce` text(100) NOT NULL,
  `is_star` int(11) NOT NULL DEFAULT '0',
  `order_number` int(11) NOT NULL DEFAULT '0',
  `category_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `category_id` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#
INSERT INTO `url`.`category` (category_name,category_logo,parent_category_id) VALUES ("未分类", "", 0);

## ----------***----------
## ----------***----------
# INSERT INTO url (url,logo,introduce,is_star,order_number,category_id) VALUES ("xx", "xx", "xxx",0,0,0);

# INSERT INTO category (category_name,category_logo,parent_category_id) VALUES ("golang", "xx.png",0);
