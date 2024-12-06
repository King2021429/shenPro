package dao

import (
	"context"
	"fmt"
	"log"
	"shenyue-gin/app/shenyue/model/db"
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

func (d *Dao) SelectByUsername(ctx context.Context, username string) (user *db.User, err error) {
	// 查询db
	err = d.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		fmt.Println()
	}
	return user, nil
}

// Delete方法：适用于软删除，只会将DeletedAt字段设置为当前时间戳，不会从数据库中完全删除记录。
//如果模型中定义了DeletedAt字段（通常是gorm.DeletedAt类型），GORM会认为该模型支持软删除。

// Unscoped().Delete方法：适用于硬删除，会从数据库中完全删除记录，忽略软删除逻辑。
func (d *Dao) DeleteUser(ctx context.Context, id uint) error {
	var user *db.User
	// 根据ID查找用户
	if err := d.db.First(&user, id).Error; err != nil {
		return err // 如果没有找到用户，返回错误
	}

	// 执行软删除
	if err := d.db.Unscoped().Delete(&user, id).Error; err != nil {
		return err // 如果删除失败，返回错误
	}

	return nil // 删除成功
}
