package api

import "encoding/json"

// 定义一个通用的Content结构体，使用interface{}来处理不同格式的content内容
type Content struct {
	Data interface{} `json:"data,omitempty"`

	Openid   string `json:"openid,omitempty"`
	ClientID string `json:"client_id,omitempty"`

	// 授权
	Permits string `json:"permits,omitempty"`

	// 稿件
	ResourceID string `json:"resource_id,omitempty"`
	State      int    `json:"state,omitempty"`
	StateDesc  string `json:"state_desc,omitempty"`
}

// SendMessageData 测试
type SendMessageData struct {
	ConversationID int64  `json:"conversation_id"`
	Extra          string `json:"extra"`
	MsgContent     string `json:"msg_content"`
	MsgKey         string `json:"msg_key"`
	MsgSource      int64  `json:"msg_source"`
	MsgStatus      int64  `json:"msg_status"`
	MsgType        int64  `json:"msg_type"`
	SendType       int64  `json:"send_type"`
	Ts             int64  `json:"ts"`
	MainOpenID     string `json:"main_open_id"`
	UserOpenID     string `json:"user_open_id"`
}

type WebhookResp struct {
	data interface{} `json:"data"`
}

type WebhookReq struct {
	Event     string          `json:"event"`
	Content   json.RawMessage `json:"content"`
	Timestamp string          `json:"timestamp"`
}
