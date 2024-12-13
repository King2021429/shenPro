package model

const (
	Title1 = "可以点加班餐喽"
	Title2 = "加班餐还有30min截止点餐"
	Title3 = "可以拿加班餐了"
	// Url1 神月订餐机器人
	Url1 = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=daf43e42-b437-400f-9c32-b3f7bd6a58fb"
)

type CallBackPic struct {
	MsgType string `json:"msgtype"`
	News    *News  `json:"news"`
}

type News struct {
	Articles *ArticleStruct `json:"articles"`
}

type ArticleStruct struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Picture     string `json:"picurl"`
}
