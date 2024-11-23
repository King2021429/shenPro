package user

import (
	"fmt"
	"shenyue-gin/app/bilibili/openPlat/dao"
	"shenyue-gin/app/bilibili/openPlat/model"
)

// AccountInfo 查询用户已授权权限列表
func AccountInfo() (resp model.BaseResp, err error) {
	url := model.AccountInfoUrl
	resp, err = dao.ApiGetRequest("", url)
	if err != nil {
		fmt.Printf("AccountScope err:%+v", err)
	}
	return
}
