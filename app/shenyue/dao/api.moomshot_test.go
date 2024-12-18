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
		firstResult := TestDao.AIChat("地球的自转周期是多少？", history)
		fmt.Println(firstResult)
		secondResult := TestDao.AIChat("月球呢？", history)
		fmt.Println(secondResult)
		thirdResult := TestDao.AIChat("你叫什么名字？", history)
		fmt.Println(thirdResult)
	})
}
