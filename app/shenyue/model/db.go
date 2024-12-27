package model

import (
	"gorm.io/gorm"
	"time"
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
	UserID    int64     `gorm:"not null"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	Published time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	User      User      `gorm:"references:ID"`
}

// 评论表结构体
type Comment struct {
	gorm.Model
	UserID          int64 `gorm:"not null"`
	ArticleID       int64 `gorm:"not null"`
	ParentCommentID *int64
	Content         string  `gorm:"not null"`
	User            User    `gorm:"references:ID"`
	Article         Article `gorm:"references:ID"`
}

// 用户关注表结构体
// status 0取消 1关注 2拉黑
type UserFollow struct {
	gorm.Model
	FollowerID int64 `gorm:"not null"`
	FolloweeID int64 `gorm:"not null"`
	Status     int64 `gorm:"default:0"`
	Follower   User  `gorm:"references:ID"`
	Followee   User  `gorm:"references:ID"`
}

// 文章收藏表结构体
type ArticleCollection struct {
	gorm.Model
	UserID    int64   `gorm:"not null"`
	ArticleID int64   `gorm:"not null"`
	User      User    `gorm:"references:ID"`
	Article   Article `gorm:"references:ID"`
}

// 文章点赞表结构体
// status 0取消 1点赞 2点踩
type ArticleLike struct {
	gorm.Model
	UserID    int64   `gorm:"not null"`
	ArticleID int64   `gorm:"not null"`
	Status    int64   `gorm:"default:0"`
	User      User    `gorm:"references:ID"`
	Article   Article `gorm:"references:ID"`
}
