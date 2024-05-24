#!/bin/zsh

# 重置数据库 mynote
# 删除 mynote 中所有的表和数据，然后重新创建一个空的数据库

# MySQL 连接信息
MYSQL_USER="root"
MYSQL_PASSWORD="Android"
MYSQL_HOST="localhost"
MYSQL_DATABASE="mynote"

# 删除数据库
mysql -u "$MYSQL_USER" -p"$MYSQL_PASSWORD" -h "$MYSQL_HOST" -e "DROP DATABASE IF EXISTS $MYSQL_DATABASE;"

# 创建数据库
mysql -u "$MYSQL_USER" -p"$MYSQL_PASSWORD" -h "$MYSQL_HOST" -e "CREATE DATABASE $MYSQL_DATABASE;"
