package model

import "encoding/json"

//{
//    "content": {
//        "conversation_id": 115987617718784,
//        "close_type": 3,
//        "close_time": 1733405507,
//        "main_open_id": "d4a2b0d73cba4ceab6b835c5fc00a46a",
//        "user_open_id": "86345135749f4f36b3210bb0e517e431"
//    },
//    "event": "CLOSE_MSG",
//    "timestamp": "2024-12-01 00:00:00"
//}

type CloseMsg struct {
	ConversationID int64  `json:"conversation_id"`
	CloseType      int64  `json:"close_type"`
	CloseTime      int64  `json:"close_time"`
	MainOpenID     string `json:"main_open_id"`
	UserOpenID     string `json:"user_open_id"`
}

type EnterDirectMsg struct {
	ConversationID int64  `json:"conversation_id"`
	UserOpenID     string `json:"user_open_id"`
	UserNick       string `json:"user_nick"`
	UserFace       string `json:"user_face"`
	Extra          string `json:"extra"`
	MsgContent     struct {
		Message string `json:"message"`
		Type    int    `json:"type"`
	} `json:"msg_content"`
	MsgKey     int64 `json:"msg_key"`
	MsgSource  int   `json:"msg_source"`
	MsgStatus  int   `json:"msg_status"`
	MsgType    int   `json:"msg_type"`
	SenderType int   `json:"sender_type"`
	Ts         int64 `json:"ts"`
}

// SendMsg
type SendMsg struct {
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
