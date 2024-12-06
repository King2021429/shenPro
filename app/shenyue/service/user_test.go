package service

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"shenyue-gin/app/shenyue/model"
	"testing"
)

func TestService_SaveUser_(t *testing.T) {
	convey.Convey("TestService_SaveUser_", t, func() {
		err := TestSrv.SaveUser(TestCtx, &model.UserReq{
			Nick:     "100",
			Password: "TestService_SaveUser_",
			Email:    "test@gmail.com",
			QQ:       1296436311,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
	})
}

func TestService_SendUserEmail_(t *testing.T) {
	convey.Convey("TestService_SendUserEmail_", t, func() {
		err := TestSrv.SendUserEmail(TestCtx)
		if err != nil {
			fmt.Println(err)
			return
		}
	})
}
