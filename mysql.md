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
