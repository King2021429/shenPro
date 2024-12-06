package db

import (
	"time"
)

// User 定义User模型，绑定users表，ORM库操作数据库，需要定义一个struct类型和MYSQL表进行绑定或者叫映射，struct字段和MYSQL表字段一一对应
type User struct {
	ID int64 // 主键
	//通过在字段后面的标签说明，定义golang字段和表字段的关系
	//例如 `gorm:"column:username"` 标签说明含义是: Mysql表的列名（字段名)为username
	Nick     string    `gorm:"column:nick"`
	Password string    `gorm:"column:password"`
	Email    string    `gorm:"column:email"`
	QQ       int64     `gorm:"column:qq"`
	Wechat   string    `gorm:"column:wechat"`
	Phone    int64     `gorm:"column:phone"`
	Status   int64     `gorm:"column:status"`
	Mtime    time.Time `gorm:"column:mtime"`
	Ctime    time.Time `gorm:"column:ctime"`
}

func (u User) TableName() string {
	//绑定MYSQL表名为users
	return "user"
}
