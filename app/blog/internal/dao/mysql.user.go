package dao

import (
	"context"
	"log"
	"shenyue-gin/app/blog/model/db"
)

// Save 测试
func (d *Dao) Save(ctx context.Context, user *db.User) (err error) {
	reply := d.db.Select("nick", "password").Create(user)
	if reply.Error != nil {
		log.Println("insert fail : ", err)
	}
	return reply.Error
}
