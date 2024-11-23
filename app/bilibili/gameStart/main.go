package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 该示例仅为demo，如需使用在生产环境需要自行按需调整

const (
	OpenPlatformHttpHost = "https://api.live.bilibili.com/xlive/open-platform"
	IdCode               = ""
	Color                = ""

	// 自己的线上
	AccessKey       = "tj9wtzXdZo6MokSaG6zMpTAn"
	AccessKeySecret = ""
	AppId           = 1698969378874

	//uat
	//AccessKey       = "YKj9iG8B7l68aZpiyxlXVcdY"
	//AccessKeySecret = ""
	//AppId           = 1720781816197
	//OpenPlatformHttpHost = "https://test-live-open.biliapi.net"
	//IdCode               = ""
	//Color                = ""
)

type AppHeartbeatReq struct {
	// 主播身份码
	GameId string `json:"game_id"`
}

type StartAppRequest struct {
	// 主播身份码
	Code string `json:"code"`
	// 项目id
	AppId int64 `json:"app_id"`
}

type StartAppRespData struct {
	// 场次信息
	GameInfo GameInfo `json:"game_info"`
	// 长连信息
	WebsocketInfo WebSocketInfo `json:"websocket_info"`
}

type GameInfo struct {
	GameId string `json:"game_id"`
}

type WebSocketInfo struct {
	//  长连使用的请求json体 第三方无需关注内容,建立长连时使用即可
	AuthBody string `json:"auth_body"`
	//  wss 长连地址
	WssLink []string `json:"wss_link"`
}

type EndAppRequest struct {
	// 场次id
	GameId string `json:"game_id"`
	// 项目id
	AppId int64 `json:"app_id"`
}

func main() {
	// 开启应用
	resp, err := StartApp(IdCode, AppId)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 解析返回值 wu
	startAppRespData := &StartAppRespData{}
	err = json.Unmarshal(resp.Data, &startAppRespData)
	if err != nil {
		fmt.Println(err)
		return
	}

	if startAppRespData == nil {
		log.Println("start app get msg err")
		return
	}

	go func(gameId string) {
		for {
			_, _ = AppHeart(gameId)
			time.Sleep(time.Second * 2)
		}
	}(startAppRespData.GameInfo.GameId)

	if len(startAppRespData.WebsocketInfo.WssLink) == 0 {
		return
	}

	//开启长连
	err = StartWebsocket(startAppRespData.WebsocketInfo.WssLink[0], startAppRespData.WebsocketInfo.AuthBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 退出
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:

			// 关闭应用
			_, err = EndApp(startAppRespData.GameInfo.GameId, AppId)
			if err != nil {
				fmt.Println(err)
				return
			}
			log.Println("WebsocketClient exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

// StartApp 启动app
func StartApp(code string, appId int64) (resp BaseResp, err error) {
	startAppReq := StartAppRequest{
		Code:  code,
		AppId: appId,
	}
	reqJson, _ := json.Marshal(startAppReq)
	return ApiRequest(string(reqJson), "/v2/app/start")
}

// EndApp 关闭app
func EndApp(gameId string, appId int64) (resp BaseResp, err error) {
	endAppReq := EndAppRequest{
		GameId: gameId,
		AppId:  appId,
	}
	reqJson, _ := json.Marshal(endAppReq)
	return ApiRequest(string(reqJson), "/v2/app/end")
}

// AppHeart app心跳
func AppHeart(gameId string) (resp BaseResp, err error) {
	appHeartbeatReq := AppHeartbeatReq{
		GameId: gameId,
	}
	reqJson, _ := json.Marshal(appHeartbeatReq)
	log.Println("AppHeart")
	return ApiRequest(string(reqJson), "/v2/app/heartbeat")
}
