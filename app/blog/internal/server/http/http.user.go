package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shenyue-gin/app/service/blog/model/api"
)

func find(ctx *gin.Context) {
	id := ctx.Query("id")
	name := ctx.Query("name")
	fmt.Println(id, name)
	ctx.JSON(200, name)
}

func register(ctx *gin.Context) {
	var user api.UserReq
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
	}
	err = svc.SaveUser(ctx.Request.Context(), &user)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(200, err)
	}
	ctx.JSON(200, user)
}
