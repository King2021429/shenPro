package service

import (
	"fmt"
	"github.com/robfig/cron"
)

func (s *Service) NewCorn() {
	c := cron.New()
	err := c.AddFunc("0 0 * * *", func() {
		fmt.Println("hello")
	})
	if err != nil {
		fmt.Println(err)
	}
	c.Start()
	fmt.Println("定时任务启动")
}
