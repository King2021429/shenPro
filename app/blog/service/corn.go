package service

import (
	"fmt"
	"github.com/robfig/cron"
)

func (s *Service) NewCorn() {
	c := cron.New()
	err := c.AddFunc("0 0 * * *", func() {
		fmt.Println("执行每天 0 点的任务")
	})
	if err != nil {
		fmt.Println(err)
	}
	c.Start()
}
