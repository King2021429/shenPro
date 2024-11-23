package model

const (
	Title1 = "可以加班点餐喽"
	Title2 = "加班点餐还有30min截止"
	Title3 = "还有10分钟就可以领加班餐了"
	Title4 = "早上好"
	Title5 = "下午好"
	// Url1 神月订餐机器人
	Url1 = ""
	// Url2 保安
	Url2 = ""

	url3 = ""
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
