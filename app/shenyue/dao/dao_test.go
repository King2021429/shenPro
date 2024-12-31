package dao

import (
	"context"
	"shenyue-gin/app/shenyue/configs"
)

var TestDao *Dao
var TestCtx context.Context

func init() {
	// 初始化配置文件
	configs.InitConfig()
	config := configs.GetConfig()
	//初始化总的dao
	TestDao = NewDao(config)
	TestCtx = context.Background()
}
