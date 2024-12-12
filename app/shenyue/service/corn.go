package service

import (
	"fmt"
	"github.com/robfig/cron"
	"shenyue-gin/app/shenyue/model"
	"shenyue-gin/app/shenyue/utils"
)

func (s *Service) NewCorn() {
	c := cron.New()
	err := c.AddFunc("0 0 12 * * *", func() {
		if utils.IsWorkingDay() {
			fmt.Println("定时任务触发，当前是工作日且非节假日，现在是12:00")
			s.SendOrder(model.Title1, model.Url1)
		}
	})
	err = c.AddFunc("0 30 16 * * *", func() {
		if utils.IsWorkingDay() {
			fmt.Println("定时任务触发，当前是工作日且非节假日，现在是16:30")
			s.SendOrder(model.Title2, model.Url1)
		}
	})
	err = c.AddFunc("0 0 21 * * *", func() {
		if utils.IsWorkingDay() {
			fmt.Println("定时任务触发，当前是工作日且非节假日，现在是21:00")
			s.SendOrder(model.Title3, model.Url1)
		}
	})
	if err != nil {
		fmt.Println(err)
	}
	c.Start()
	fmt.Println("定时任务启动")
}
