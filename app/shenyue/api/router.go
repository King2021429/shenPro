package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shenyue-gin/app/shenyue/middleware"
	"shenyue-gin/app/shenyue/service"
)

var Svc *service.Service

func InitHttpRouter(s *service.Service) (e *gin.Engine) {
	Svc = s
	e = gin.Default()
	// 允许所有来源的跨域请求
	e.Use(CORS())

	// 公共路由
	publicGroup := e.Group("/")
	{
		// 回掉 b站商业化客服回掉
		publicGroup.POST("webhook", Webhook)

		// 用户 注册 登陆
		publicGroup.POST("user/register", registerUser)
		publicGroup.POST("user/login", loginUser)

		// 测试
		publicGroup.GET("test/id/:id", TestId)
		publicGroup.POST("test/path/*path", TestPath)
	}

	// 受保护路由
	protectedGroup := e.Group("/protected", middleware.AuthMiddleware)
	{
		protectedGroup.GET("/user/getUserInfo", getUserInfo)

		// 文章相关接口
		protectedGroup.POST("/article/create", CreateArticle)
		protectedGroup.POST("/article/delete", DeleteArticle)
		protectedGroup.POST("/article/edit", EditArticle)
		protectedGroup.POST("/article/getList", GetArticleList)
		protectedGroup.POST("/article/getInfo", GetArticleInfo)

		// 点赞相关接口
		protectedGroup.POST("article/like", LikeArticle)

		// 文章相关接口
		protectedGroup.POST("/comment/create", CreateComment)
		protectedGroup.POST("/comment/delete", DeleteComment)
		protectedGroup.POST("/comment/edit", EditComment)
		protectedGroup.POST("/comment/getInfo", GetCommentInfo)

		// AI的两个
		protectedGroup.POST("/ai/conversation_start", AIConversationStart)
		protectedGroup.POST("/ai/conversation_send_msg", AIConversationSendMsg)
		protectedGroup.POST("/ai/conversation_delete", AIConversationDelete)

	}

	// 管理员路由
	adminGroup := e.Group("/admin", middleware.AdminAuthMiddleware)
	{
		adminGroup.GET("/id/:id", TestId)
	}

	return e
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content - Type, Content - Length, Accept - Encoding, X - CSRF - Token, Authorization, accept, origin, Cache - Control, X - Requested - With")
		// 只允许"POST, GET"请求
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, gin.H{"message": "Options请求成功"})
			return
		}

		c.Next()
	}
}
