package main

import (
	"shenyue-gin/app/service/blog/internal/server/http"
	"shenyue-gin/app/service/blog/internal/service"
)

func main() {
	// 初始化service
	s := service.NewService()
	// 初始化http路由
	e := http.InitHttpRouter(s)
	// 监听并在 0.0.0.0:8080 上启动服务
	err := e.Run()
	if err != nil {
		return
	}
}
