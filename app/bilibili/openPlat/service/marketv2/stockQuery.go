package service

import (
	"fmt"
	"shenyue-gin/app/bilibili/openPlat/dao"
	"shenyue-gin/app/bilibili/openPlat/model"
)

// StockQuery 新增共享
func StockQuery() (resp model.BaseResp, err error) {
	url := model.StockUpdate
	resp, err = dao.ApiGetRequest("", url)
	if err != nil {
		fmt.Printf("StockQuery err:%+v", err)
	}

	// 解析返回值
	queryResp := &model.StockInfo{}
	//var queryResp interface{}

	if queryResp == nil {
		fmt.Println("StockQuery Data nil")
	}
	fmt.Printf("StockQuery:%+v\n", queryResp)
	return
}
