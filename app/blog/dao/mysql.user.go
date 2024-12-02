package dao

import (
	"context"
	"fmt"
	"log"
	"shenyue-gin/app/blog/model/db"
)

// Save 测试
func (d *Dao) Save(ctx context.Context, user *db.User) (err error) {
	reply := d.db.Select("nick", "password", "email", "qq").Create(user)
	if reply.Error != nil {
		log.Println("insert fail : ", err)
	}
	return reply.Error
}

func (d *Dao) SelectAllEmail(ctx context.Context) ([]string, error) {
	// 查询所有用户的email字段值
	var emails []string
	err := d.db.Model(&db.User{}).Select("email").Find(&emails).Error
	if err != nil {
		fmt.Println("查询失败：", err)
		return nil, err
	}
	return emails, nil
}
