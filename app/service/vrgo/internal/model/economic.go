package model

// 定义队伍和角色的结构体
type Team struct {
	Name   string
	Heroes [5]Hero
}

type Hero struct {
	Name   string
	Wealth int
}
