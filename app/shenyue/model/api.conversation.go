package model

type ConversationStartReq struct {
	UserId int64 `json:"user_id"`
}

type ConversationStartResp struct {
	ConversationId int64 `json:"conversation_id"`
}

type ConversationSendMsgReq struct {
	UserId         int64  `json:"user_id"`
	Content        string `json:"content"`
	ConversationId int64  `json:"conversation_id"`
}

type ConversationSendMsgResp struct {
	Reply string `json:"reply"`
}

type ConversationDeleteReq struct {
	UserId         int64 `json:"user_id"`
	ConversationId int64 `json:"conversation_id"`
}

type ConversationDeleteResp struct {
}
