package types

import (
	"ginlearn/utils"
	"time"

	"gorm.io/gorm"
)

type User struct {
	//创建标签，primaryKey 当前列定义为主键， json:"id" 在进行json序列化与反序列化使用字段名
	ID uint `gorm:"primaryKey" json:"id"`
	// uniqueIndex 与index类似，但是创建的是唯一索引，size：定义当前列的长度，not null指定列不为空
	Username string `gorm:"uniqueIndex,size:50;not null" json:"username"`
	Password string `gorm:"size:255;not null" json:"password"`
	Email    string `gorm:"size:100;not null" json:"email"`
	//autoCreateTime 创建时追踪当前时间
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	// 创建/更新时追踪当前时间
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	//index 创建索引  json:"-" 忽略该字段，无读写
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	//foreignKey:UserID 一个用户有多个帖子，外键是 Post 结构体中的 UserID（作用场景：用于定义结构体之间的数据库关联关系（通常是一对多关系））
	//json:"posts" 指定该字段在 JSON 中对应的键名为 posts（若结构体字段名与 JSON 键名一致，可省略此配置）
	//omitempty 是可选参数，表示：若该字段的值为零值（如空切片、空指针、0、"" 等），则在 JSON 序列化时忽略该字段，不生成对应的键值对
	//关联关系可以同时参考 post中字段解释
	Posts []Post `gorm:"foreignKey:UserID" json:"posts,omitempty"`
	//与上面类同
	Comments []Comment `gorm:"foreignKey:UserID" json:"comments,omitempty"`
}

// BeforeCreate 钩子函数----创建用户前加密密码
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	//使用utils包中的密码加密函数
	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}

// CheckPassword 验证密码是否正确
func (u *User) CheckPassword(password string) bool {
	return utils.CheckPasswordHash(password, u.Password)
}
