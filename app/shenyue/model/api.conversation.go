package model

type ConversationStartReq struct {
	Uid int64 `json:"uid"`
}

type ConversationStartResp struct {
	ConversationId int64 `json:"conversation_id"`
}

type ConversationSendMsgReq struct {
	Uid            int64  `json:"uid"`
	Content        string `json:"content"`
	ConversationId int64  `json:"conversation_id"`
}

type ConversationSendMsgResp struct {
	Reply string `json:"reply"`
}

type ConversationDeleteReq struct {
	Uid            int64 `json:"uid"`
	ConversationId int64 `json:"conversation_id"`
}

type ConversationDeleteResp struct {
}
