package main

import (
	"log"
	"os"
	"os/signal"
	"shenyue-gin/app/shenyue/api"
	"shenyue-gin/app/shenyue/service"
	"syscall"
)

func main() {
	// 初始化service service里面会初始化dao
	newService := service.NewService()
	// 初始化http路由
	e := api.InitHttpRouter(newService)
	// 监听并在 0.0.0.0:8001 上启动服务
	err := e.Run(":8001")
	if err != nil {
		return
	}
	// 退出
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			newService.Close()
			// 关闭应用
			log.Println("blog exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
