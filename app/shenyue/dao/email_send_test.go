package dao

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDao_SendEmail_(t *testing.T) {
	convey.Convey("TestDao_SendEmail_", t, func() {
		for i := 0; i < 5; i++ {
			err := TestDao.SendEmail(TestCtx, "1964363113@qq.com", fmt.Sprintf("内容:%d", i), fmt.Sprintf("标题:%d", i))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	})
}
