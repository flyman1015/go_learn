package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
1. 题目1：模型定义
  - 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
    - 要求 ：
      - 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
      - 编写Go代码，使用Gorm创建这些模型对应的数据库表。
2. 题目2：关联查询
  - 基于上述博客系统的模型定义。
    - 要求 ：
      - 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
      - 编写Go代码，使用Gorm查询评论数量最多的文章信息。
3. 题目3：钩子函数
  - 继续使用博客系统的模型。
    - 要求 ：
      - 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
      - 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

*/

func main() {

	//连接数据库
	db, err := ConnenDb()
	if err != nil {
		fmt.Println("连接数据库异常")
		return
	}

	// //创建表插入数据
	// CreateDbInfo(db)

	// 查询文章相关信息
	// QueryPostsInfo(db)

	fmt.Println("=================================")
}

// 连接数据库方法
func ConnenDb() (db *gorm.DB, err error) {
	//开启数据库连接
	dsn := "root:root123456@tcp(localhost:3306)/mygo?charset=utf8mb4&parseTime=True&loc=Local"
	db1, err1 := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db1, err1
}

// 文章钩子函数
func (p *Post) AfterCreate(db *gorm.DB) (err error) {

	result := db.Model(&User{}).Where("id = ?", p.AuthorID).Update("posts_count", gorm.Expr("posts_count + ?", 1))
	if result.Error != nil {
		fmt.Println("文章更新时用户文章统计数更新异常")
		return result.Error
	}
	return nil
}

// 评论钩子函数
func (c *Comment) AfterDele(db *gorm.DB) (err error) {
	var commentCount int64
	result := db.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount)
	if result.Error != nil {
		fmt.Println("统计评论数量异常")
		return result.Error
	}

	status := "有评论"
	if commentCount == 0 {
		status = "无评论"
		result = db.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", status)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func HookFunc(db *gorm.DB) {

	//创建
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	// 创建用户
	user := User{Username: "张三122UUU", Email: "zhangsan@example.com111UUU"}
	db.Create(&user)

	// 创建文章 (触发AfterCreate钩子)
	post := Post{
		Title:    "我的第一篇博客",
		Content:  "欢迎来到我的博客...",
		AuthorID: user.ID,
	}
	db.Create(&post) // 此时用户的PostsCount自动+1

	// 创建评论
	comment := Comment{
		Content:  "很好的文章!",
		PostID:   post.ID,
		AuthorID: user.ID,
	}
	db.Create(&comment)

	// 删除评论 (触发AfterDelete钩子)
	db.Delete(&comment) // 删除后检查文章状态
}

// 文章查询函数
func QueryPostsInfo(db *gorm.DB) {
	var use User
	userRuslt := db.Where("username =?", "code_master").First(&use)
	if userRuslt.Error != nil {
		fmt.Println("查询用户数据异常")
		return
	}

	//查询用户的文章
	var postList []Post
	postListRuslt := db.Where("author_id =?", use.ID).Find(&postList)
	if postListRuslt.Error != nil {
		fmt.Println("查询用户文章数据异常")
		return
	}

	for _, p := range postList {
		fmt.Println(p)
		var comList []Comment
		postListRuslt := db.Where("post_id =?", p.ID).Find(&comList)
		if postListRuslt.Error != nil {
			fmt.Println("查询评论")
			return
		}
		fmt.Println("查询评论", p.ID, comList)
	}

	//查询评论最多的文章
	subQuery := db.Model(&Comment{}).
		Select("post_id,COUNT(*) AS comment_count").
		Group("post_id").
		Order("comment_count desc").
		Limit(1)

	var maxCommentPost Post

	maxResult := db.Model(&Post{}).Select("id", "title", "content").Joins("JOIN (?) AS max_post ON posts.id =max_post.post_id", subQuery).First(&maxCommentPost)

	if maxResult.Error != nil {
		fmt.Println("查询评论最多文章异常")
		return
	}
	fmt.Println("查询最多文章：", maxCommentPost)
}

// 创建表插入数据
func CreateDbInfo(db *gorm.DB) {
	//定义表
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	//插入数据
	db.Create(GetSampleUsers())
	db.Create(GetSamplePosts())
	db.Create(GetSampleComments())
}

type User struct {
	ID         uint   `gorm:"primaryKey"`                   //创建唯一建
	Username   string `gorm:"uniqueIndex;size:50;not null"` //创建唯一索引
	Email      string `gorm:"uniqueIndex;size:100;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`               //创建所以
	Posts      []Post         `gorm:"foreignKey:AuthorID"` // 一对多关系
	PostsCount int            `gorm:"default:0"`           // 用户文章数量统计字段

}

type Post struct {
	ID            uint   `gorm:"primaryKey"`
	Title         string `gorm:"size:200;not null"`
	Content       string `gorm:"type:text;not null"`
	AuthorID      uint   `gorm:"index;not null"` // 外键
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Comments      []Comment      `gorm:"foreignKey:PostID"`                    // 一对多关系
	CommentStatus string         `gorm:"type:enum('无评论','有评论');default:'无评论'"` // 评论状态
}

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
	AuthorID  uint   `gorm:"index;not null"` // 外键
	PostID    uint   `gorm:"index;not null"` // 外键
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func GetSampleUsers() []User {
	return []User{
		{
			ID:       1,
			Username: "tech_guru",
			Email:    "guru@tech.com",
		},
		{
			ID:       2,
			Username: "code_master",
			Email:    "master@code.dev",
		},
		{
			ID:       3,
			Username: "dev_learner",
			Email:    "learner@dev.net",
		},
	}
}

func GetSamplePosts() []Post {
	return []Post{
		{
			ID:       1,
			Title:    "深入理解Golang并发模型",
			Content:  "Go语言的并发模型是其最强大的特性之一...",
			AuthorID: 1,
		},
		{
			ID:       2,
			Title:    "GORM高级技巧大全",
			Content:  "本文将介绍GORM的各种高级用法和最佳实践...",
			AuthorID: 1,
		},
		{
			ID:       3,
			Title:    "从零构建RESTful API",
			Content:  "使用Go和Gin框架构建高性能API服务...",
			AuthorID: 2,
		},
		{
			ID:       4,
			Title:    "数据库优化实战",
			Content:  "如何优化SQL查询提升应用性能...",
			AuthorID: 2,
		},
		{
			ID:       5,
			Title:    "微服务架构设计模式",
			Content:  "微服务架构的常见模式和反模式...",
			AuthorID: 3,
		},
	}
}

func GetSampleComments() []Comment {
	return []Comment{
		{
			ID:       1,
			Content:  "非常有深度的文章！",
			AuthorID: 2,
			PostID:   1,
		},
		{
			ID:       2,
			Content:  "期待更多关于channel的内容",
			AuthorID: 3,
			PostID:   1,
		},
		{
			ID:       3,
			Content:  "GORM的关联查询确实很方便",
			AuthorID: 1,
			PostID:   2,
		},
		{
			ID:       4,
			Content:  "解决了我的实际问题",
			AuthorID: 3,
			PostID:   3,
		},
		{
			ID:       5,
			Content:  "优化后性能提升明显",
			AuthorID: 1,
			PostID:   4,
		},
		{
			ID:       6,
			Content:  "实例代码能否分享一下？",
			AuthorID: 3,
			PostID:   4,
		},
		{
			ID:       7,
			Content:  "架构设计思路很清晰",
			AuthorID: 2,
			PostID:   5,
		},
	}
}
