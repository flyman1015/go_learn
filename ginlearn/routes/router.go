package routes

import (
	"ginlearn/controllers"
	"ginlearn/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// 创建控制器实例
	authController := controllers.NewUserController()       // 假设 UserController 包含注册和登录方法
	postController := controllers.NewPostController()       // 假设 PostController 包含文章相关方法
	commentController := controllers.NewCommentController() // 假设 CommentController 包含评论相关方法

	// ===== 不需要认证的路由 =====
	apiPublic := router.Group("/api")
	{
		// 用户认证
		apiPublic.POST("/register", authController.Register)
		apiPublic.POST("/login", authController.Login)

		// 公开内容访问
		apiPublic.GET("/posts", postController.GetPosts)
		apiPublic.GET("/posts/:id", postController.GetPost)
		apiPublic.GET("/posts/:id/comments", commentController.GetComments)
	}

	// ===== 需要认证的路由 =====
	apiPrivate := apiPublic.Use(middleware.JWTUserMiddleware())
	{
		// 文章管理
		apiPrivate.POST("/posts", postController.CreatePost)
		apiPrivate.PUT("/posts/:id", middleware.PostOwnershipMiddleware(), postController.UpdatePost)
		apiPrivate.DELETE("/posts/:id", middleware.PostOwnershipMiddleware(), postController.DeletePost)

		// 评论管理
		apiPrivate.POST("/posts/:id/comments", middleware.PostOwnershipMiddleware(), commentController.CreateComment)

	}
}
