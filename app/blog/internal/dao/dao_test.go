package dao

import (
	"context"
)

var TestDao *Dao
var TestCtx context.Context

func init() {
	//_ = flag.Set("conf", "../../../configs")
	//初始化总的dao
	TestDao = NewDao()
	TestCtx = context.Background()
}
