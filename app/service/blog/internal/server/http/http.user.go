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

func save(ctx *gin.Context) {
	id := ctx.Query("id")
	name := ctx.Query("name")
	address := ctx.DefaultQuery("address", "北京")

	ctx.JSON(200, gin.H{
		"id":      id,
		"name":    name,
		"address": address,
	})
}

func register(ctx *gin.Context) {
	var user api.User
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
