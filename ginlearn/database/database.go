package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() (db *gorm.DB, err error) {
	//开启数据库连接
	dsn := "root:root123@tcp(localhost:3306)/mygo?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err1 := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err1 != nil {
		fmt.Println("连接数据库异常")
		return nil, err1
	}
	return DB, err1
}
