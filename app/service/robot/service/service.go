package service

import (
	"encoding/json"
	"fmt"
	"shenyue-gin/app/service/robot/dao"
	"shenyue-gin/app/service/robot/model"
	"time"
)

func SendHello(title string, url string) {
	description := fmt.Sprintf("讲文明，懂礼貌！现在时间是%s,保安祝您工作愉快！(PS:我是I机器人，不会回消息哦～）", time.Now().Format("15:04:05"))
	article := &model.ArticleStruct{
		Title:       title,
		Description: description,
		Url:         "",
		Picture:     "https://www.aikobo.cn/draw/_next/image?url=https%3A%2F%2Fboss.aikobo.cn%2Faigc-public%2F538269e09e3a3596470fd50db85358c74624dd7b9db4c406bd5f49ff4582eae5.png&w=3840&q=75",
	}
	news := &model.News{
		Articles: article,
	}
	callBackPic := &model.CallBackPic{
		MsgType: "news",
		News:    news,
	}
	sendStr, _ := json.Marshal(callBackPic)
	dao.Send(sendStr, url)
}

func SendOrder(title string, url string) {
	description := fmt.Sprintf("今天星期%d,现在时间是%s！加班餐虽好，请不要经常加班哦", time.Now().Weekday(), time.Now().Format("15:04:05"))
	article := &model.ArticleStruct{
		Title:       title,
		Description: description,
		Url:         "https://h5.platform.xixiang000.com/",
		Picture:     "",
	}
	news := &model.News{
		Articles: article,
	}
	callBackPic := &model.CallBackPic{
		MsgType: "news",
		News:    news,
	}
	sendStr, _ := json.Marshal(callBackPic)
	dao.Send(sendStr, url)
}
