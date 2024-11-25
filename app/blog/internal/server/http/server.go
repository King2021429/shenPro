package http

import (
	"github.com/gin-gonic/gin"
	"shenyue-gin/app/blog/internal/service"
)

var svc *service.Service

func InitHttpRouter(s *service.Service) (e *gin.Engine) {
	svc = s
	e = gin.Default()
	e.POST("/test/id/:id", testId)
	e.POST("/test/path/*path", testPath)
	ug := e.Group("/user")
	{
		ug.GET("find", find)
		ug.POST("register", register)
	}
	return e
}
