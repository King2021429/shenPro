package dao

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"shenyue-gin/app/shenyue/model"
	"testing"
)

func TestDao_AIChat_(t *testing.T) {
	convey.Convey("TestDao_AIChat_", t, func() {
		history := []model.Message{
			{
				Role:    "system",
				Content: "你是二次元妹子，活泼可爱好动，名字叫江户川神月",
			},
		}
		firstResult := TestDao.AIChat("你叫什么名字？", &history)
		fmt.Println(firstResult)
		secondResult := TestDao.AIChat("今天上海天气是什么样子", &history)
		fmt.Println(secondResult)
		thirdResult := TestDao.AIChat("你能根据天气下一首诗吗", &history)
		fmt.Println(thirdResult)
	})

}

func TestDao_AIChatDeep_(t *testing.T) {
	convey.Convey("TestDao_AIChat_", t, func() {
		history := []model.Message{
			{
				Role:    "system",
				Content: "你是二次元妹子，活泼可爱好动，名字叫江户川神月",
			},
		}
		firstResult := TestDao.AIChatDeep("你叫什么名字？", &history)
		fmt.Println(firstResult)
		secondResult := TestDao.AIChatDeep("2 * 6 + 10 等于多少", &history)
		fmt.Println(secondResult)
		thirdResult := TestDao.AIChatDeep("9.11和9.7哪一个大", &history)
		fmt.Println(thirdResult)
	})

}
