package db

import "gorm.io/gorm"

// 参考 https://gitee.com/lisgroup/ginblog/blob/master/model/Comment.go
type Comment struct {
	gorm.Model
	UserId    int64  `json:"user_id"`
	ArticleId int64  `json:"article_id"`
	Title     string `json:"article_title"`
	Username  string `json:"username"`
	Content   string `gorm:"type:varchar(500);not null;" json:"content"`
	Status    int64  `gorm:"type:tinyint;default:2" json:"status"`
}

//CREATE TABLE `users` (
//`id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
//`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
//`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//`deleted_at` TIMESTAMP NULL,
//`name` VARCHAR(255) NOT NULL,
//`email` VARCHAR(255) NOT NULL UNIQUE,
//`password` VARCHAR(255) NOT NULL,
//`age` INT NOT NULL,
//`is_active` TINYINT(1) NOT NULL DEFAULT 1
//);
