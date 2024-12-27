package model

import (
	"gorm.io/gorm"
)

// 用户表结构体
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string
	Avatar   string
}

// 文章表结构体
type Article struct {
	gorm.Model
	UID     int64  `gorm:"not null"`
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
}

// 评论表结构体
type Comment struct {
	gorm.Model
	UID             int64 `gorm:"not null"`
	ArticleID       int64 `gorm:"not null"`
	ParentCommentID *int64
	Content         string `gorm:"not null"`
}

// UserFollow 用户关注表结构体
// status 0取消 1关注 2拉黑
type UserFollow struct {
	gorm.Model
	FollowerID int64 `gorm:"not null"`
	FolloweeID int64 `gorm:"not null"`
	Status     int64 `gorm:"default:0"`
}

// ArticleCollection 文章收藏表结构体
type ArticleCollection struct {
	gorm.Model
	UID       int64 `gorm:"not null"`
	ArticleID int64 `gorm:"not null"`
}

// ArticleLike 文章点赞表结构体
// status 0取消 1点赞 2点踩
type ArticleLike struct {
	gorm.Model
	UserID    int64 `gorm:"not null"`
	ArticleID int64 `gorm:"not null"`
	Status    int64 `gorm:"default:0"`
}
