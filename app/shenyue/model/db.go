package model

import (
	"gorm.io/gorm"
)

// User 用户表结构体
// Status 0 用户 1 admin
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Avatar   string `gorm:"not null"`
	Status   int64  `gorm:"default:0"`
}

// Article 文章表结构体
type Article struct {
	gorm.Model
	Uid     int64  `gorm:"not null"`
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	Cover   string `gorm:"not null"`
}

// Comment 评论结构体
type Comment struct {
	gorm.Model
	ArticleID int64      `gorm:"not null" json:"article_id"`
	Uid       int64      `gorm:"not null" json:"Uid"`
	Content   string     `gorm:"type:text;not null" json:"content"`
	ParentID  int64      `gorm:"default:0" json:"parent_id"`
	Replies   []*Comment `gorm:"-" json:"replies"` // 使用 gorm:"-" 标记该字段不映射到数据库
}

// UserFollow 用户关注表结构体
// status 0取消 1关注 2拉黑
type UserFollow struct {
	gorm.Model
	FollowerID int64 `gorm:"not null"`
	FolloweeID int64 `gorm:"not null"`
	Status     int64 `gorm:"default:0"`
}

// ArticleFavorite 文章收藏表结构体
type ArticleFavorite struct {
	gorm.Model
	Uid       int64 `gorm:"not null"`
	ArticleID int64 `gorm:"not null"`
	Status    int64 `gorm:"default:0"`
}

// ArticleLike 文章点赞表结构体
// status 0取消 1点赞 2点踩
type ArticleLike struct {
	gorm.Model
	Uid       int64 `gorm:"not null"`
	ArticleID int64 `gorm:"not null"`
	Status    int64 `gorm:"default:0"`
}

//GORM 通过结构体字段标签和内部约定来识别和处理
//gorm.Model 中的各个字段。
//在创建、更新和删除记录时，GORM 会根据这些规则自动完成相应的操作
//并与数据库进行交互，确保数据的一致性和完整性。
