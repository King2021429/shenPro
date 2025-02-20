package api

import (
	"fmt"
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
	//e.Use(commonAOP)
	//e.Use(testAOP)

	// 公共路由
	publicGroup := e.Group("/")
	{
		// 回掉 b站商业化客服回掉
		publicGroup.POST("webhook", Webhook)

		// 测试
		publicGroup.GET("test/id/:id", TestId)
		publicGroup.POST("test/path/*path", TestPath)
	}

	userGroup := e.Group("/user")
	{
		// 用户 注册 登陆
		userGroup.POST("/register", registerUser)
		userGroup.POST("/login", loginUser)
	}

	// 需要用户登陆态
	protectedGroup := e.Group("/protected", middleware.AuthMiddleware)
	{
		protectedGroup.GET("/user/getUserInfo", getUserInfo)

		// 文章相关接口
		protectedGroup.POST("/article/create", CreateArticle)
		protectedGroup.POST("/article/delete", DeleteArticle)
		protectedGroup.POST("/article/edit", EditArticle)
		protectedGroup.POST("/article/getList", GetArticleList)
		protectedGroup.POST("/article/getInfo", GetArticleInfo)

		// 文章点赞相关接口
		protectedGroup.POST("/article/like/update", LikeArticle) //点赞 点踩
		protectedGroup.POST("/article/like/list", GetLikeList)   // 根据uid获取全部点赞/点踩文章

		// 文章收藏相关接口
		protectedGroup.POST("/article/favorite/update", FavoriteArticle) // 文章收藏/取消收藏
		protectedGroup.POST("/article/favorite/list", GetFavoriteList)   // 根据uid获取全部收藏文章

		// 评论相关接口
		protectedGroup.POST("/comment/create", CreateComment)   // 创建评论
		protectedGroup.POST("/comment/delete", DeleteComment)   // 删除评论
		protectedGroup.POST("/comment/edit", EditComment)       // 编辑评论
		protectedGroup.POST("/comment/getList", GetCommentList) // 根据文章id获取评论列表

		// AI对话
		protectedGroup.POST("/ai/conversation_start", AIConversationStart)      // 创建对话
		protectedGroup.POST("/ai/conversation_send_msg", AIConversationSendMsg) // 发送消息
		protectedGroup.POST("/ai/conversation_delete", AIConversationDelete)    // 删除对话
		protectedGroup.POST("/ai/conversation_list", AIConversationList)        // 获取对话列表

	}

	// 管理员路由
	adminGroup := e.Group("/admin", middleware.AdminAuthMiddleware)
	{
		adminGroup.GET("/id/:id", TestId)
		// 管理文章
		// 管理用户

	}

	return e
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许访问的域名，这里可以根据实际情况修改，例如改为具体的前端域名
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// 去除多余的空格
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization,accept,origin,Cache-Control,X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, gin.H{"message": "Options请求成功"})
			return
		}

		c.Next()
	}
}

// commonAOP 公共AOP 这里可以进行一些上报操作
func commonAOP(ctx *gin.Context) {

	ctx.Next()
	fmt.Println("test")
}

func testAOP(ctx *gin.Context) {
	//middleware.AuthMiddleware(ctx)
	fmt.Println("test")
	//if ctx.IsAborted() {
	//	return
	//}
	url := ctx.Request.URL.String()
	if url == "/protected/user/getUserInfo" {
		//AIConversationList(ctx)
		//ctx.JSON(http.StatusUnauthorized, gin.H{"error": "有趣"})

	}

	fmt.Println(url)
	ctx.Abort()

}
