package dao

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDao_SelectAllEmail_(t *testing.T) {
	convey.Convey("TestDao_SelectAllEmail_", t, func() {
		for i := 0; i < 5; i++ {
			res, err := TestDao.SelectAllEmail(TestCtx)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(res)
		}
	})
}
