package dao

import (
	"context"
	"shenyue-gin/app/shenyue/model"
)

// 创建评论
func (d *Dao) CreateComment(ctx context.Context, comment model.Comment) error {
	return d.db.Create(&comment).Error
}

// 获取评论
func (d *Dao) GetComment(ctx context.Context, id uint) (model.Comment, error) {
	var comment model.Comment
	err := d.db.First(&comment, id).Error
	return comment, err
}

// 更新评论
func (d *Dao) UpdateComment(ctx context.Context, comment model.Comment) error {
	return d.db.Save(&comment).Error
}

// 删除评论
func (d *Dao) DeleteComment(ctx context.Context, id uint) error {
	return d.db.Delete(&model.Comment{}, id).Error
}
