package main

import (
	"log"
	"os"
	"os/signal"
	"shenyue-gin/app/bilibili/openPlat/service/user"

	"syscall"
)

// 入口启动函数
func main() {
	// 调用方法
	//resp, err := service.QueryDlc()
	//resp, err := customer.ConversationList()
	//resp, err := marketV2.StockQuery()
	//resp, err := service.CommonAddShare()
	//_, _ = service.ImageUpload()
	_, _ = user.AccountScope()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("WebsocketClient exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
