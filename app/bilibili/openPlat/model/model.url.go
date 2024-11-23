package model

const (
	AccountInfoUrl = "/arcopen/fn/user/account/info"   // 获取用户公开信息 GET
	UserScopesUrl  = "/arcopen/fn/user/account/scopes" // 查询用户已授权权限列表 GET

	ArcInitUrl     = "/arcopen/fn/archive/video/init"                                                   //文件上传预处理
	ArcComplete    = "/arcopen/fn/archive/video/complete?upload_token=7213c24789bf42b3a3482b7c7d9a597f" // 文件分片合片
	arcAddUrl      = "/arcopen/fn/archive/add"                                                          // 稿件提交 POST
	arcTypeList    = "/arcopen/fn/archive/type/list"                                                    // 分区查询 GET
	arcCoverUpload = "/arcopen/fn/archive/cover/upload"                                                 // 封面上传 POST

	dataUserStatUrl = "/arcopen/fn/data/user/stat" // 获取用户数据 GET

	// 纯v2
	shopInfoGetUrl = "/arcopen/fn/v2/market/shop/info/get" // 获取店铺信息 GET

	itemListUrl = "/arcopen/fn/v2/market/commodity/item/list" // 查询商品列表 GET

	StockQuery  = "/arcopen/fn/v2/market/stock/query?sku_id_list=1002884494" // 库存查询 GET
	StockUpdate = "/arcopen/fn/v2/market/stock/update"                       // 库存更新 POST

	// 客服模块
	ConversationSendMsgUrl        = "/arcopen/fn/market/customer/conversation/send_msg"            // 客服系统 消息发送
	ConversationListUrl           = "/arcopen/fn/market/customer/conversation/list"                // 客服系统 获取当前会话列表
	ConversationCustomerMsgQuery  = "/arcopen/fn/market/customer/conversation/customer_msg_query"  // 客服系统 获取聊天记录
	conversationCustomerUserFrom  = "/arcopen/fn/market/customer/conversation/user_from"           // 客服系统 获取用户来源
	conversationClose             = "/arcopen/fn/market/customer/conversation/close"               // 客服系统 关闭
	conversationStaffStatusUpdate = "/arcopen/fn/market/customer/conversation/staff_status_update" // 客服系统 修改客服状态

	conversationVideoUpload = "/arcopen/fn/market/customer/video_upload"
	conversationVideoGet    = "/arcopen/fn/market/customer/video_file_get"
)

// 图片上传url模块
const (
	// ImageUploadArcUrl 稿件图片上传
	ImageUploadArcUrl = "https://member.bilibili.com/arcopen/fn/archive/cover/upload"
	// ImageUploadCommodityUrl 商品图片上传
	ImageUploadCommodityUrl = "https://member.bilibili.com/arcopen/fn/v2/market/commodity/image/upload"
	// ImageUploadCustomer 客服图片上传
	ImageUploadCustomer = "https://member.bilibili.com/arcopen/fn/market/customer/image_upload"
)
