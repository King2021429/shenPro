package dao

import (
	"context"
	"shenyue-gin/app/shenyue/model"
)

// CreateComment 方法创建新评论并返回 ID 和错误
func (d *Dao) CreateComment(ctx context.Context, comment model.Comment) (uint, error) {
	// 使用 GORM 的 Create 方法创建新记录
	result := d.db.Create(&comment)
	// 返回新创建记录的 ID 和可能出现的错误
	return comment.ID, result.Error
}

// GetCommentsByArticleID 根据文章 ID 查询该文章下的所有评论
func (d *Dao) GetCommentsByArticleID(ctx context.Context, articleID int64) ([]model.Comment, error) {
	var comments []model.Comment
	// 查询指定文章 ID 的所有评论
	result := d.db.WithContext(ctx).Where("article_id = ?", articleID).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
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
