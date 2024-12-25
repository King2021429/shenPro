package model

type ConversationStartReq struct {
}

type ConversationStartResp struct {
	ConversationId int64 `json:"conversation_id"`
}
type ConversationSendMsgReq struct {
	Content        string `json:"content"`
	ConversationId int64  `json:"conversation_id"`
}

type ConversationSendMsgResp struct {
	Reply string `json:"reply"`
}
