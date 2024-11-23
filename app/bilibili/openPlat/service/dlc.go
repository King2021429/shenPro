package service

import (
	"encoding/json"
	"fmt"
	"shenyue-gin/app/bilibili/openPlat/dao"
	"shenyue-gin/app/bilibili/openPlat/model"
)

// QueryDlc 新增共享
func QueryDlc() (resp model.BaseResp, err error) {
	queryReq := model.QueryReq{
		DlcId:    "",
		OrderIds: []string{"", "", ""},
	}
	queryReqJson, _ := json.Marshal(queryReq)
	resp, err = dao.ApiRequest(string(queryReqJson), "/xlive/open-platform/v1/appDlc/queryOrderList")
	if err != nil {
		fmt.Printf("ApiRequest err:%+v", err)
	}

	return
}
