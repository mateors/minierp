## How to get database structure in MySQL via query
> mysqldump database_name --compact --no-data

## How to check mysql verison
> mysql -V
```sample output:
mysql  Ver 15.1 Distrib 10.5.12-MariaDB, for debian-linux-gnu (x86_64) using readline 5.2
```

## Mysql default character set
If you’re using MySQL 5.7, the default MySQL collation is generally latin1_swedish_ci because MySQL uses latin1 as its default character set. \
If you’re using MySQL 8.0, the default charset is utf8mb4.

## Which charset should i use?
If you elect to use UTF-8 as your collation, always use **utf8mb4** (specifically utf8mb4_unicode_ci). You should not use UTF-8 because MySQL’s UTF-8 is different from proper UTF-8 encoding. This is the case because it doesn’t offer full unicode support which can lead to data loss or security issues. 

Keep in mind that **utf8mb4_general_ci** is a simplified set of sorting rules which **takes shortcuts designed to improve speed** while utf8mb4_unicode_ci sorts accurately in a wide range of languages. In general, **utf8mb4 is the “safest” character set** as it also supports 4-byte unicode while utf8 only supports up to 3.

## Example SQL CREATE statement:

```
CREATE TABLE `records` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `type` varchar(10) DEFAULT NULL,
  `content` text DEFAULT NULL,
  `ttl` int(11) DEFAULT NULL,
  `status` tinyint(1) DEFAULT 1,
  PRIMARY KEY (`id`),
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

**Same table with different charset and collate.**

CREATE TABLE `records` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `type` varchar(10) DEFAULT NULL,
  `content` text DEFAULT NULL,
  `ttl` int(11) DEFAULT NULL,
  `status` tinyint(1) DEFAULT 1,
  PRIMARY KEY (`id`),
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;

```
