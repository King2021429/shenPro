package dao

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestDao_RcGet_(t *testing.T) {
	convey.Convey("TestDao_RcGet_", t, func() {
		value, err := TestDao.RcGet(TestCtx, "key2")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(value)
	})
}

func TestDao_RcSet_(t *testing.T) {
	convey.Convey("TestDao_RcSet_", t, func() {
		err := TestDao.RcSet(TestCtx, "key2", "value2", 30*time.Second)
		if err != nil {
			fmt.Println(err)
			return
		}
	})
}
