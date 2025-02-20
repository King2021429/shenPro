package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func TestId(ctx *gin.Context) {
	param := ctx.Param("id")
	uid, _ := ctx.Get("uid")
	fmt.Println(uid)
	ctx.JSON(200, param)
}

func TestPath(ctx *gin.Context) {
	param := ctx.Param("path")
	ctx.JSON(200, param)
}
