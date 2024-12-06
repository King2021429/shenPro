package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shenyue-gin/app/blog/service"
)

var Svc *service.Service

func InitHttpRouter(s *service.Service) (e *gin.Engine) {
	Svc = s
	e = gin.Default()
	// 允许所有来源的跨域请求
	e.Use(CORS())
	e.POST("/test/id/:id", TestId)
	e.POST("/test/path/*path", TestPath)
	e.POST("/webhook", Webhook)
	ug := e.Group("/user")
	{
		ug.GET("find", Find)
		ug.POST("register", Register)
	}
	return e
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content - Type, Content - Length, Accept - Encoding, X - CSRF - Token, Authorization, accept, origin, Cache - Control, X - Requested - With")
		// 只允许"POST, GET"请求
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, gin.H{"message": "Options请求成功"})
			return
		}
		c.Next()
	}
}
