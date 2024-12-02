package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shenyue-gin/app/blog/model/api"
)

func Find(ctx *gin.Context) {
	id := ctx.Query("id")
	name := ctx.Query("name")
	fmt.Println(id, name)
	err := Svc.SendUserEmail(ctx.Request.Context())
	if err != nil {
		fmt.Println(err)
		ctx.JSON(200, err)
	}
	ctx.JSON(200, name)
}

func Register(ctx *gin.Context) {
	var user api.UserReq
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
	}
	err = Svc.SaveUser(ctx.Request.Context(), &user)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(200, err)
	}
	ctx.JSON(200, user)
}
