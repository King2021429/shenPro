package dao

import (
	"errors"
	"gorm.io/gorm"
	"shenyue-gin/app/shenyue/model"
)

// CreateAIChat 创建AIChat记录
func (d *Dao) CreateAIChat(chat *model.AIChat) error {
	return d.db.Create(chat).Error
}

// DeleteAIChat 删除AIChat记录
func (d *Dao) DeleteAIChat(uid, conversationId int64) error {
	return d.db.Where("uid = ? AND conversation_id = ?", uid, conversationId).Delete(&model.AIChat{}).Error
}

// UpdateConversationContent 更新AIChat记录的ConversationContent字段
func (d *Dao) UpdateConversationContent(uid, conversationId int64, newContent string) error {
	return d.db.Model(&model.AIChat{}).Where("uid = ? AND conversation_id = ?", uid, conversationId).Update("ConversationContent", newContent).Error
}

// GetAIChatList 获取AIChat记录列表
func (d *Dao) GetAIChatList(uid int64) ([]model.AIChat, error) {
	var chats []model.AIChat
	if err := d.db.Where("uid = ?", uid).Find(&chats).Error; err != nil {
		return nil, err
	}
	return chats, nil
}

// GetAIChatByUidAndConversationId 根据 Uid 和 conversationId 查询记录
func (d *Dao) GetAIChatByUidAndConversationId(uid, conversationId int64) (*model.AIChat, error) {
	var chat model.AIChat
	result := d.db.Where("uid = ? AND conversation_id = ?", uid, conversationId).First(&chat)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // 未找到记录
		}
		return nil, result.Error // 其他错误
	}
	return &chat, nil
}
