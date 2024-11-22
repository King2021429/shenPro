package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.POST("/test/id/:id", func(ctx *gin.Context) {
		param := ctx.Param("id")
		ctx.JSON(200, param)
	})

	r.POST("/test/path/*path", func(ctx *gin.Context) {
		param := ctx.Param("path")
		ctx.JSON(200, param)
	})

	ug := r.Group("/user")
	{
		ug.GET("find", func(ctx *gin.Context) {
			ctx.JSON(200, "user find")
		})
		ug.GET("save", func(ctx *gin.Context) {
			id := ctx.Query("id")
			name := ctx.Query("name")
			address := ctx.DefaultQuery("address", "北京")
			ctx.JSON(200, gin.H{
				"id":      id,
				"name":    name,
				"address": address,
			})
		})
	}

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
