package service

import (
	"context"
	"shenyue-gin/app/shenyue/configs"
)

var TestSrv *Service
var TestCtx context.Context

func init() {
	// 初始化配置文件
	configs.InitConfig()
	config := configs.GetConfig()
	//初始化总的dao
	TestSrv = NewService(config)
	TestCtx = context.Background()
}
