package service

import (
	"context"
)

var TestSrv *Service
var TestCtx context.Context

func init() {
	//_ = flag.Set("conf", "../../../configs")
	//初始化总的dao
	TestSrv = NewService()
	TestCtx = context.Background()
}
