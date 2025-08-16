package types

import (
    "time"

    "gorm.io/gorm"
)

type Post struct{
    ID uint `gorm:"primaryKey" json:"id"`
    Title string `gorm:"size:200;not null" json:"title"`
    //type:text 表示让 Gorm 在创建或迁移表结构时，将该字段的数据类型设置为数据库支持的 text 类型。 
    // 长文本用 text
    Content string `gorm:"type:text;not null" json:"content"`
    //autoCreateTime 创建时追踪当前时间
    CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
    // 创建/更新时追踪当前时间
    UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
    //index 创建索引  json:"-" 忽略该字段，无读写
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

    //关联关系
    // 外键，关联 User 表的 ID
    UserID   uint      `gorm:"index;not null" json:"user_id"`
    // 反向关联，一个帖子属于一个用户
    User     User      `gorm:"foreignKey:UserID" json:"author,omitempty"`
    Comments []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty"`
}