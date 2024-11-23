package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	// 染色，不染不用填
	Color = ""

	// 开平地址
	OpenPlatformHttpHost = "https://member.bilibili.com"
)

var (
	clients = []*ClientContext{
		{
			Name:        "开平长连应用",
			ClientId:    "93a0f774fae84e6c",
			AccessToken: "",
		},
		//{
		//	Name:        "开平长连应用",
		//	ClientId:    "fdf8b377a47c44ab",
		//	AccessToken: "",
		//},
	}
)

type ClientContext struct {
	Name        string // log前缀
	ClientId    string
	AccessToken string
}

func PJsonMarshal(a interface{}) string {
	rv, _ := json.MarshalIndent(a, "", "	")
	return string(rv)
}

func PJsonString(in []byte) string {
	buf := bytes.NewBuffer([]byte{})
	_ = json.Indent(buf, in, "", "		")
	return buf.String()
}

func main() {
	ctx, ctxCancel := context.WithCancel(context.Background())

	eg := errgroup.Group{}
	eg.SetLimit(100)

	for i := range clients {
		param := clients[i]

		eg.Go(func() (err error) {
			rv := RunWs(ctx, param)
			ctxCancel()
			return rv
		})

	}

	defer func() {
		ctxCancel()

		eg.Wait()
	}()

	// 退出
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case s := <-c:
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				return
			case syscall.SIGHUP:
			default:
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

func RunWs(ctx context.Context, cli *ClientContext) error {
	// 获取长链地址
	resp, err := WsStart(cli.ClientId, cli.AccessToken)
	if err != nil {
		fmt.Printf("call WsStart error %+v", err)
		return err
	}

	// 解析返回值
	wsStartRespData := &WsStartData{}
	err = json.Unmarshal(resp.Data, &wsStartRespData)
	if err != nil {
		fmt.Printf("json.Unmarshal error %+v", err)
		return err
	}
	if wsStartRespData == nil {
		fmt.Printf("WsStart Response is nil")
		return fmt.Errorf("WsStart response is nil")
	}

	fmt.Printf("WsStart Response %s", PJsonMarshal(wsStartRespData))

	err = StartWebsocket(cli, wsStartRespData.WebsocketInfo.WssLink[0], wsStartRespData.WebsocketInfo.AuthBody)
	if err != nil {
		fmt.Printf("StartWebsocket err:%+v", err)
		return err
	}

	// 维持心跳
	go func() {
		for {
			resp, err = AppHeart(cli.ClientId, wsStartRespData.ConnId)
			if err != nil || resp.Code != 0 {
				fmt.Printf("AppHeart Err:%+v Resp:%+v", err, resp)
			}
			time.Sleep(time.Second * 2)
		}
	}()

	<-ctx.Done()
	return nil
}

// 获取连接，启动心跳
func WsStart(clientId, aToken string) (resp BaseResp, err error) {
	req := map[string]interface{}{}
	reqJson, _ := json.Marshal(req)
	return ApiRequest(string(reqJson), fmt.Sprintf("/arcopen/fn/live/room/ws-start?client_id=%s&access_token=%s", clientId, aToken))
}

// 心跳一次
func AppHeart(clientId, connId string) (resp BaseResp, err error) {
	req := map[string]interface{}{
		"conn_id": connId,
	}
	reqJson, _ := json.Marshal(req)
	return ApiRequest(string(reqJson), fmt.Sprintf("/arcopen/fn/live/room/ws-heartbeat?client_id=%s", clientId))
}
