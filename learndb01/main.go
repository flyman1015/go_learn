package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func main() {

	// _, err := InitDb()
	// if err != nil {
	// 	fmt.Println("连接数据库异常")
	// 	return
	// }
	// fmt.Println("数据库连接成功")

	dsn := "root:root123@tcp(localhost:3306)/mygo?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

}

func InitDb() (db *gorm.DB, err error) {

	//开启数据库连接
	dsn := "root:root123456@tcp(localhost:3306)/mygo?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err1 := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 日志模式：Info 级别会打印 SQL
	})
	if err1 != nil {
		fmt.Println("连接数据库异常", err1)
		return nil, err1
	}
	return DB, err1
}
