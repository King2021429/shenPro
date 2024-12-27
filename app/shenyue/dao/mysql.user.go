package dao

import (
	"context"
	"fmt"
	"shenyue-gin/app/shenyue/model"
)

// 创建用户
func (d *Dao) CreateUser(ctx context.Context, user *model.User) error {
	return d.db.Create(&user).Error
}

// 获取用户
func (d *Dao) GetUser(ctx context.Context, id int64) (model.User, error) {
	var user model.User
	err := d.db.First(&user, id).Error
	return user, err
}

func (d *Dao) SelectAllEmail(ctx context.Context) ([]string, error) {
	// 查询所有用户的email字段值
	var emails []string
	err := d.db.Model(&model.User{}).Select("email").Find(&emails).Error
	if err != nil {
		fmt.Println("查询失败：", err)
		return nil, err
	}
	return emails, nil
}

func (d *Dao) SelectByUsername(ctx context.Context, username string) (user *model.User, err error) {
	// 查询db
	err = d.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		fmt.Println()
	}
	return user, nil
}

// 更新用户
func (d *Dao) UpdateUser(ctx context.Context, user model.User) error {
	return d.db.Save(&user).Error
}

// 删除用户
func (d *Dao) DeleteUser(ctx context.Context, id int64) error {
	return d.db.Delete(&model.User{}, id).Error
}

// 创建用户关注关系
func (d *Dao) CreateUserFollow(ctx context.Context, userFollow model.UserFollow) error {
	return d.db.Create(&userFollow).Error
}

// 获取用户关注关系
func (d *Dao) GetUserFollow(ctx context.Context, id int64) (model.UserFollow, error) {
	var userFollow model.UserFollow
	err := d.db.First(&userFollow, id).Error
	return userFollow, err
}

// 更新用户关注关系
func (d *Dao) UpdateUserFollow(ctx context.Context, userFollow model.UserFollow) error {
	return d.db.Save(&userFollow).Error
}

// 删除用户关注关系
func (d *Dao) DeleteUserFollow(ctx context.Context, id int64) error {
	return d.db.Delete(&model.UserFollow{}, id).Error
}

// Delete方法：适用于软删除，只会将DeletedAt字段设置为当前时间戳，不会从数据库中完全删除记录。
// 如果模型中定义了DeletedAt字段（通常是gorm.DeletedAt类型），GORM会认为该模型支持软删除。
// Unscoped().Delete方法：适用于硬删除，会从数据库中完全删除记录，忽略软删除逻辑。
