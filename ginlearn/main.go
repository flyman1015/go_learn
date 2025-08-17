package main

import (
	"fmt"
	"ginlearn/database"
	"ginlearn/logger"
	"ginlearn/middleware"
	"ginlearn/routes"
	"ginlearn/types"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Printf("===============")

	//链接数据库初始化表结构
	db, err := database.InitDb()
	err = db.AutoMigrate(&types.User{}, &types.Post{}, &types.Comment{})
	fmt.Printf("==============+++++======")
	if err != nil {
		logger.Log.Fatalf("数据库迁移失败: %v", err)
	}

	// 创建纯净的 Gin 引擎
	router := gin.New()
	fmt.Printf("==========****====+++++======")
	//自定panic异常中间
	router.Use(middleware.RecoveryMiddleware())

	// router.Use(middleware.JWTUserMiddleware(&cfg))

	// 设置路由
	routes.SetupRoutes(router)
	fmt.Printf("==========****===zzz=+++++======")
	//启动服务
	if err := router.Run(":8090"); err != nil {
		logger.Log.Fatalf("服务启动失败: %v", err)
	}
	fmt.Printf("=======dddd===****====+++++======")
}
