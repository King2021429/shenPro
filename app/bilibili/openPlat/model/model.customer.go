package model

// OpenMarketCustomerSendMsgReq 外部请求开平 消息发送请求
type OpenMarketCustomerSendMsgReq struct {
	BizId          int64  `json:"biz_id" form:"biz_id" validate:"required"`     // 业务方类型id: 1-商家，2-带货,3-经营号
	ConversationId int64  `json:"conversation_id" form:"conversation_id"`       // 会话id
	StaffId        int64  `json:"staff_id" form:"staff_id" validate:"required"` // 客服id
	UserOpenId     string `json:"user_open_id"`                                 // 用户openid
	ShopId         int64  `json:"shop_id" form:"shop_id" validate:"required"`   // 商家id
	MsgType        int64  `json:"msg_type" form:"msg_type" validate:"required"` // 消息类型
	Msg            string `json:"msg" form:"msg" validate:"required"`           // 消息内容
}

// OpenMarketCustomerSendMsgResp 开平对外 消息发送返回结构体
type OpenMarketCustomerSendMsgResp struct {
	MsgKey string `json:"msg_key"` // 消息key
}

// OpenMarketCustomerConversationListReq 外部请求开平 获取当前会话列表
type OpenMarketCustomerConversationListReq struct {
	StaffId int64 `json:"staff_id" form:"staff_id" validate:"required"` // 客服id
	BizId   int64 `json:"biz_id" form:"biz_id" validate:"required"`     // 业务方类型id: 1-商家，2-带货,3-经营号
	ShopId  int64 `json:"shop_id" form:"shop_id" validate:"required"`   // 商家id
}

// OpenMarketCustomerConversationListResp 开平对外 会话列表返回信息
type OpenMarketCustomerConversationListResp struct {
	ConversationDTOList []*OpenConversationDTO `json:"conversation_dto_list"` // 会话列表集合
}

// OpenConversationDTO 开平对外 会话对象
type OpenConversationDTO struct {
	ConversationId   int64                  `json:"conversation_id"`     // 会话id
	StaffId          int64                  `json:"staff_id"`            // 客服id
	SkillGroupId     int64                  `json:"skill_group_id"`      // 技能组id
	UserName         string                 `json:"user_name"`           // 用户名称
	UserFace         string                 `json:"user_face"`           // 用户头像
	IncomingTime     int64                  `json:"incoming_time"`       // 进线时间戳 单位秒
	SessTs           int64                  `json:"sess_ts"`             // 最新会话时间戳 单位秒
	LastStaffMsgTime int64                  `json:"last_staff_msg_time"` // 最近一次客服发送消息时间
	UserOpenId       string                 `json:"user_open_id"`
	LastMsg          OpenConversationMsgDTO `json:"last_msg"`
}

// OpenConversationMsgDTO 最后一条消息
type OpenConversationMsgDTO struct {
	MsgType      int64  `json:"msg_type"`   // 消息类型,1-文字,2-图片,17-视频, 10018-主动消息
	Content      string `json:"content"`    // 消息内容
	MsgKey       int64  `json:"msg_key"`    // 消息key
	MsgSeqNo     int64  `json:"msg_seq_no"` // 消息序号
	MsgTs        int64  `json:"msg_ts"`     //消息时间戳 单位秒
	SenderOpenId string `json:"sender_open_id"`
	SenderType   int64  `json:"sender_type"` // 发送者类型 1-商家；2-灰条发送；3-用户
	SenderName   string `json:"sender_name"` // 发送者名称
	SenderFace   string `json:"sender_face"` // 发送者头像
}

// OpenMarketCustomerMsgQueryReq 外部请求开平 获取聊天记录
type OpenMarketCustomerMsgQueryReq struct {
	BizId      int64  `json:"biz_id" form:"biz_id" validate:"required"` // 业务方类型id: 1-商家，3-经营号
	Size       int64  `json:"size" form:"size" validate:"required"`     // 条数，最大20
	UserOpenId string `json:"user_open_id" form:"user_open_id" validate:"required"`
	ShopId     int64  `json:"shop_id" form:"shop_id" validate:"required"` // 商家id
	MaxSeqNo   int64  `json:"max_seq_no"`                                 // 最大消息序列号
	MinSeqNo   int64  `json:"min_seq_no"`                                 // 最小消息序列号
}

// OpenMarketCustomerMsgQueryResp 聊天记录返回参数
type OpenMarketCustomerMsgQueryResp struct {
	MsgDTOList []*OpenConversationMsgDTO `json:"msg_dto_list"` // 消息列表
	MaxSeqNo   int64                     `json:"max_seq_no"`   // 最大消息序列号
	MinSeqNo   int64                     `json:"min_seq_no"`   // 最小消息序列号
}

// OpenMarketCustomerUserFromReq 获取用户来源 请求参数
type OpenMarketCustomerUserFromReq struct {
	ConversationId int64 `json:"conversation_id" form:"conversation_id" validate:"required"` // 会话id
}

// OpenMarketCustomerUserFromResp 获取用户来源 返回参数
type OpenMarketCustomerUserFromResp struct {
	UserBasicDTO    OpenUserBasicDTO    `json:"user_basic_dto"`     // 用户来源基本信息
	SourceUpInfoDTO OpenSourceUpInfoDTO `json:"source_up_info_dto"` // 会话up来源信息
}

// OpenUserBasicDTO 用户来源基本信息
type OpenUserBasicDTO struct {
	DataType int64  `json:"data_type"` // 数据类型：1 订单，2 商品，3店铺，13私信历史会话、11经营号个人空间品牌tab、12稿件评论区蓝链
	ViewTime string `json:"view_time"` // 来访时间
}

// OpenSourceUpInfoDTO 会话up来源信息
type OpenSourceUpInfoDTO struct {
	SourceUpOpenId string `json:"source_up_open_id"`
	SourceUpName   string `json:"source_up_name"` // up主名称
	SourceBvid     string `json:"source_bvid"`    // 视频稿件bvid
	SourceBtitle   string `json:"source_btitle"`  // 视频稿件名称
}

// OpenMarketCustomerConversationCloseReq 客服关闭会话请求
type OpenMarketCustomerConversationCloseReq struct {
	ConversationId int64 `json:"conversation_id" form:"conversation_id" validate:"required"` // 会话id
	StaffId        int64 `json:"staff_id" form:"staff_id" validate:"required"`               // 客服id
}

// OpenMarketCustomerConversationCloseResp 客服关闭会话返回参数
type OpenMarketCustomerConversationCloseResp struct{}

// OpenMarketCustomerStaffStatusUpdateReq 客服修改状态 请求参数
type OpenMarketCustomerStaffStatusUpdateReq struct {
	Status              int64 `json:"status" form:"status" validate:"required"`     // 客服状态,1-上线,2-忙碌,3-离线
	StaffId             int64 `json:"staff_id" form:"staff_id" validate:"required"` // 客服id
	IsCloseConversation int64 `json:"is_close_conversation"`                        // 离线是否关闭会话  1-是 0-否,默认0
}

// OpenMarketCustomerStaffStatusUpdateResp 客服修改状态 返回参数
type OpenMarketCustomerStaffStatusUpdateResp struct{}

// MarketCustomerImageUploadReq 客服系统 图片上传
type MarketCustomerImageUploadReq struct {
	StaffId int64 `json:"staff_id" form:"staff_id" validate:"required"` // 客服id
}

// MarketCustomerImageUploadResp 客服系统 图片上传 返回url
type MarketCustomerImageUploadResp struct {
	Url string `json:"url"` // 图片url
}

// MarketCustomerVideoUploadReq 客服系统 视频上传请求参数
type MarketCustomerVideoUploadReq struct {
	FileName string `json:"file_name" form:"file_name"` // 文件名字
	StaffId  int64  `json:"staff_id" form:"staff_id"`   // 客服id
}

// MarketCustomerVideoUploadResp 客服系统 视频上传返回参数
type MarketCustomerVideoUploadResp struct {
	BossKey          string `json:"bossKey"`
	UploadPreSignUrl string `json:"uploadPreSignUrl"` // 上传签名链接
}

// OpenMarketCustomerVideoUploadResp 客服系统 视频上传返回参数 对外
type OpenMarketCustomerVideoUploadResp struct {
	BossKey          string `json:"boss_key"`
	UploadPreSignUrl string `json:"upload_pre_sign_url"` // 上传签名链接
}

// MarketCustomerVideoFileGetReq 客服系统 获取视频信息 请求参数
type MarketCustomerVideoFileGetReq struct {
	FileName string `json:"file_name" form:"file_name"` // bosskey
}

// MarketCustomerVideoFileGetResp 客服系统 获取视频信息 返回参数
type MarketCustomerVideoFileGetResp struct {
	FileUrl string `json:"fileUrl"`
}

// OpenMarketCustomerVideoFileGetResp 客服系统 获取视频信息 返回参数 对外
type OpenMarketCustomerVideoFileGetResp struct {
	FileUrl string `json:"file_url"`
}
