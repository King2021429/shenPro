package model

type BaseResp struct {
	Code      int64       `json:"code"`
	Message   string      `json:"message"`
	RequestId string      `json:"request_id"`
	Data      interface{} `json:"data"`
}

// 域名
const (
	// OpenPlatformHttpHostUat 直播开放平台测试环境host
	OpenPlatformHttpHostUat  = "https://test-live-open.biliapi.net"
	OpenPlatformHttpHostUat2 = "https://uat-api.live.bilibili.com"
	// OpenPlatformHttpHost 直播开放平台正式环境host
	OpenPlatformHttpHost = "https://api.live.bilibili.com"

	// UatMainOpenPlatformHttpHost 主站开平
	UatMainOpenPlatformHttpHost = "https://uat-member.bilibili.com"
	// MainOpenPlatformHttpHost 主站开平
	MainOpenPlatformHttpHost = "https://member.bilibili.com"

	//LocalHost = "http://0.0.0.0:8000"
	LocalHost = "http://127.0.0.1:8000"
)

const (
	SceneCode = "ARC_APP_SHARE"
	BizCode   = ""
)

const (
	ClientIdUat    = ""
	AppSecretUat   = ""
	AccessTokenUat = ""

	ClientIdProd    = ""
	AppSecretProd   = ""
	AccessTokenProd = ""
)
