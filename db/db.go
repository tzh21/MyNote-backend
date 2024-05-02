package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// 将 user, password, dbname 替换为实际数据库用户名、密码和数据库名称
	dsn := "root:Android@tcp(127.0.0.1:3306)/mynote?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
