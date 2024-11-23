package service

import (
	"encoding/json"
	"fmt"
	"shenyue-gin/app/bilibili/openPlat/dao"
	"shenyue-gin/app/bilibili/openPlat/model"
)

// CommonAddShare 新增共享
func CommonAddShare() (resp model.BaseResp, err error) {
	// 创建一个 CommonMsg 对象
	commonMsg := model.CommonMsg{
		Source:  "",
		Cover:   "",
		Title:   "投稿标题",
		TypeID:  172,
		TopicID: 1173894,
		VideoMaterialURL: []string{
			"https://1400335750.vod2.myqcloud.com/ff539f7evodcq1400335750/4d9970f01397757888517111156/QGeorKmhB2kA.mp4",
		},
	}

	// 将对象编码为 JSON 字符串
	jsonString, err := json.Marshal(commonMsg)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(jsonString)
	commonAddShareReq := model.CommonAddShareReq{
		CommonMsg: string(jsonString),
		BizCode:   model.BizCode,
		SceneCode: model.SceneCode,
	}
	commonAddShareReqJson, _ := json.Marshal(commonAddShareReq)
	return dao.ApiRequest(string(commonAddShareReqJson), "/arcopen/fn/resource/add_share")
}
