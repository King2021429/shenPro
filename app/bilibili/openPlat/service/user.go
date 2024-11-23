package service

import (
	"encoding/json"
	"shenyue-gin/app/bilibili/openPlat/dao"
	"shenyue-gin/app/bilibili/openPlat/model"
)

// ArticleAdd 请求测试
func ArticleAdd() (resp model.BaseResp, err error) {
	var articleAddReq = model.ArticleAddReq{
		Title:      "神月",
		Category:   1,
		TemplateId: 5,
		Summary:    "神月的开放平台稿件测试简介",
		Content:    "神月的开放平台正文测试",
	}
	articleAddReqJson, _ := json.Marshal(articleAddReq)
	return dao.ApiRequest(string(articleAddReqJson), "/arcopen/fn/article/add")
}
