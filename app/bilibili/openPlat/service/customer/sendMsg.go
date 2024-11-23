package service

import (
	"encoding/json"
	"fmt"
	"shenyue-gin/app/bilibili/openPlat/dao"
	"shenyue-gin/app/bilibili/openPlat/model"
)

// SendMsg 新增共享
func SendMsg() (resp model.BaseResp, err error) {
	queryReq := model.QueryReq{}
	queryReqJson, _ := json.Marshal(queryReq)
	resp, err = dao.ApiRequest(string(queryReqJson), model.ConversationSendMsgUrl)
	if err != nil {
		fmt.Printf("SendMsg err:%+v", err)
	}

	// 解析返回值
	queryResp := &model.QueryResp{}
	err = json.Unmarshal(resp.Data, &queryResp)
	if err != nil {
		fmt.Printf("SendMsg resp json unmarshal err:%+v", err)
	}
	if queryResp == nil {
		fmt.Println("queryResp Data nil")
	}
	fmt.Printf("queryResp:%+v\n", queryResp)
	return
}
