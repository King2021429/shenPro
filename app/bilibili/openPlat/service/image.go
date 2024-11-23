package service

import (
	"shenyue-gin/app/bilibili/openPlat/dao"
	"shenyue-gin/app/bilibili/openPlat/model"
)

func ImageUpload() (resp model.BaseResp, err error) {
	url := model.ImageUploadCustomer
	return dao.DaoFormRequest(url)
}
