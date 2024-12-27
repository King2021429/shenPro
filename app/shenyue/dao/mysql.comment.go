package dao

import "shenyue-gin/app/shenyue/model"

// 创建评论
func (d *Dao) CreateComment(comment model.Comment) error {
	return d.db.Create(&comment).Error
}

// 获取评论
func (d *Dao) GetComment(id uint) (model.Comment, error) {
	var comment model.Comment
	err := d.db.First(&comment, id).Error
	return comment, err
}

// 更新评论
func (d *Dao) UpdateComment(comment model.Comment) error {
	return d.db.Save(&comment).Error
}

// 删除评论
func (d *Dao) DeleteComment(id uint) error {
	return d.db.Delete(&model.Comment{}, id).Error
}
