package model

type CommonMsg struct {
	Source           string   `json:"source"`
	Cover            string   `json:"cover"`
	Title            string   `json:"title"`
	TypeID           int      `json:"type_id"`
	TopicID          int      `json:"topic_id"`
	VideoMaterialURL []string `json:"video_material_url"`
}

type CommonAddShareReq struct {
	// common_msg 通用信息 大json串
	CommonMsg string `json:"common_msg"`
	// biz_code
	BizCode string `json:"biz_code"`
	// scene_code 场景码
	SceneCode string `json:"scene_code"`
}

type CommonAddShareRespData struct {
	// link_url
	LinkUrl string `json:"link_url"`
}
