package main

import (
	"fmt"
	"shenyue-gin/app/service/robot/model"
	"shenyue-gin/app/service/robot/service"
	"shenyue-gin/app/service/robot/utils"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			// 获取当前时间
			now := time.Now()
			// 判断是否是工作日
			if utils.IsWeekday(now) {
				if now.Hour() == 12 && now.Minute() == 00 && now.Second() == 0 {
					service.SendOrder(model.Title1, model.Url1)
					//service.SendOrder(model.Title1, model.Url2)
				}
				if now.Hour() == 16 && now.Minute() == 30 && now.Second() == 0 {
					service.SendOrder(model.Title2, model.Url1)
					//service.SendOrder(model.Title2, model.Url2)
				}
				if now.Hour() == 20 && now.Minute() == 20 && now.Second() == 0 {
					service.SendOrder(model.Title3, model.Url1)
					//service.SendOrder(model.Title3, model.Url2)
				}
			}
		}
	}()
	// 等待用户输入退出命令
	fmt.Println("Press Ctrl+C to exit")

	select {}
}
