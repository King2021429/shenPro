package service

import (
	"encoding/json"
	"fmt"
	"shenyue-gin/app/bilibili/openPlat/dao"
	"shenyue-gin/app/bilibili/openPlat/model"
)

// ConversationList 获取当前会话列表
func ConversationList() (resp model.BaseResp, err error) {
	queryReq := model.OpenMarketCustomerConversationListReq{
		StaffId: 1026006,
		BizId:   3,
		ShopId:  1003,
	}
	queryReqJson, _ := json.Marshal(queryReq)
	resp, err = dao.ApiRequest(string(queryReqJson), model.ConversationListUrl)
	if err != nil {
		fmt.Printf("ConversationList err:%+v", err)
		return resp, err
	}

	// 解析返回值
	queryResp := &model.OpenMarketCustomerConversationListResp{}
	err = json.Unmarshal(resp.Data, &queryResp)
	if err != nil {
		fmt.Printf("ConversationList resp json unmarshal err:%+v", err)
	}
	if queryResp == nil {
		fmt.Println("ConversationList queryResp Data nil")
	}
	fmt.Printf("ConversationList queryResp:%+v\n", queryResp)
	return
}
