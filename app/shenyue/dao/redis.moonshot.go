package dao

import (
	"context"
	"fmt"
	"time"
)

const (
	_rcKeyConversationId       = "conversation_id:%d"
	_rcGetConversationIdExpire = 60 * 60
)

func (d *Dao) RcSetConversation(ctx context.Context, conversationId int64, value string) (err error) {
	key := fmt.Sprintf(_rcKeyConversationId, conversationId)
	//expire := _rcGetConversationIdExpire
	err = d.RcSet(ctx, key, value, 1*time.Hour)
	if err != nil {
		return
	}
	return
}

func (d *Dao) RcGetConversation(ctx context.Context, conversationId int64) (value string, err error) {
	key := fmt.Sprintf(_rcKeyConversationId, conversationId)
	value, err = d.RcGet(ctx, key)
	if err != nil {
		return
	}
	return
}
