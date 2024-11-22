package http

import "github.com/gin-gonic/gin"

func testId(ctx *gin.Context) {
	param := ctx.Param("id")
	ctx.JSON(200, param)
}

func testPath(ctx *gin.Context) {
	param := ctx.Param("path")
	ctx.JSON(200, param)
}
