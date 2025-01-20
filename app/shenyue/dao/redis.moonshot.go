package dao

import (
	"context"
	"fmt"
	"time"
)

const (
	_rcKeyConversationId = "conversation_id:%d:user_id:%d"
)

func (d *Dao) RcSetConversation(ctx context.Context, conversationId, uid int64, value string) (err error) {
	key := fmt.Sprintf(_rcKeyConversationId, conversationId, uid)
	err = d.RcSet(ctx, key, value, 1*time.Hour)
	if err != nil {
		return
	}
	return
}

func (d *Dao) RcGetConversation(ctx context.Context, conversationId, uid int64) (value string, err error) {
	key := fmt.Sprintf(_rcKeyConversationId, conversationId, uid)
	value, err = d.RcGet(ctx, key)
	if err != nil {
		return
	}
	return
}

func (d *Dao) RcDelConversation(ctx context.Context, conversationId, uid int64) (err error) {
	key := fmt.Sprintf(_rcKeyConversationId, conversationId, uid)
	err = d.RcDel(ctx, key)
	if err != nil {
		return
	}
	return
}
