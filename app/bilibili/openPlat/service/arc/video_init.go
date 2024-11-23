package arc

import (
	"encoding/json"
	"fmt"
	"shenyue-gin/app/bilibili/openPlat/dao"
	"shenyue-gin/app/bilibili/openPlat/model"
)

// VideoInit 文件上传预处理
func VideoInit() (resp model.BaseResp, err error) {
	url := model.ArcInitUrl
	videoInitReq := model.VideoInitReq{
		Name:  "test.mp4",
		Utype: "0",
	}
	videoInitReqJson, _ := json.Marshal(videoInitReq)
	resp, err = dao.ApiRequestV2(string(videoInitReqJson), url)
	if err != nil {
		fmt.Printf("VideoInit err:%+v", err)
	}
	return
}
