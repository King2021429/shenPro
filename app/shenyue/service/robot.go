package service

import (
	"encoding/json"
	"fmt"
	"shenyue-gin/app/shenyue/model"
	"time"
)

func (s *Service) SendOrder(title string, url string) {
	description := fmt.Sprintf("今天星期%d,现在时间是%s！加班餐虽好，请不要经常加班哦", time.Now().Weekday(), time.Now().Format("15:04:05"))
	article := &model.ArticleStruct{
		Title:       title,
		Description: description,
		Url:         "https://h5.platform.xixiang000.com/",
		Picture:     "https://p9-flow-imagex-sign.byteimg.com/ocean-cloud-tos/image_skill/1b77c088-e543-4dcc-9c3f-2eb2410197a7_1732601313567252452~tplv-a9rns2rl98-web-watermark-v2.png?rk3s=b14c611d&x-expires=1764137313&x-signature=sk2sCQpO8m0bs53L5kXN80Ks5BE%3D",
	}
	news := &model.News{
		Articles: article,
	}
	callBackPic := &model.CallBackPic{
		MsgType: "news",
		News:    news,
	}
	sendStr, _ := json.Marshal(callBackPic)
	s.dao.Send(sendStr, url)
}
