package user

import (
	"fmt"
	"shenyue-gin/app/bilibili/openPlat/dao"
	"shenyue-gin/app/bilibili/openPlat/model"
)

// AccountScope 查询用户已授权权限列表
func AccountScope() (resp model.BaseResp, err error) {
	url := model.UserScopesUrl
	resp, err = dao.ApiGetRequest("", url)
	if err != nil {
		fmt.Printf("AccountScope err:%+v", err)
	}
	return
}
