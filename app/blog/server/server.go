package server

import (
	"github.com/gin-gonic/gin"
	"shenyue-gin/app/blog/api"
	"shenyue-gin/app/blog/service"
)

var Svc *service.Service

func InitHttpRouter(s *service.Service) (e *gin.Engine) {
	Svc = s
	e = gin.Default()
	e.POST("/test/id/:id", api.TestId)
	e.POST("/test/path/*path", api.TestPath)
	ug := e.Group("/user")
	{
		ug.GET("find", api.Find)
		ug.POST("register", api.Register)
	}
	return e
}
