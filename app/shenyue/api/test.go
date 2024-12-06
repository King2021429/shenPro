package api

import "github.com/gin-gonic/gin"

func TestId(ctx *gin.Context) {
	param := ctx.Param("id")
	ctx.JSON(200, param)
}

func TestPath(ctx *gin.Context) {
	param := ctx.Param("path")
	ctx.JSON(200, param)
}
