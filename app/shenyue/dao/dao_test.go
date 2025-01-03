package dao

import (
	"context"
	"os"
	"shenyue-gin/app/shenyue/configs"
)

var TestDao *Dao
var TestCtx context.Context

func init() {
	os.Setenv("env", "local")
	// 初始化配置文件
	configs.InitConfig()
	config := configs.GetConfig()
	//初始化总的dao
	TestDao = NewDao(config)
	TestCtx = context.Background()
}
