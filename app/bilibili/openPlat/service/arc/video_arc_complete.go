package arc

import (
	"fmt"
	"shenyue-gin/app/bilibili/openPlat/dao"
	"shenyue-gin/app/bilibili/openPlat/model"
)

// VideoArcComplete 合片
func VideoArcComplete() (resp model.BaseResp, err error) {
	url := model.ArcComplete

	resp, err = dao.ApiRequestV2("", url)
	if err != nil {
		fmt.Printf("VideoArcComplete err:%+v", err)
	}
	return
}
