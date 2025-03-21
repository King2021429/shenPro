package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shenyue-gin/app/shenyue/middleware"
	"shenyue-gin/app/shenyue/model"
	"strconv"
)

// 注册 新增
// 审批 更新
// 注销 删除

// 注册用户
func registerUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码不能为空"})
		return
	}
	err := Svc.SaveUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "注册成功"})

}

// 用户登录
func loginUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码不能为空"})
		return
	}
	uid, err := Svc.LoginUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	token, err := middleware.GenerateToken(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成令牌失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

// 获取用户信息
func getUserInfo(c *gin.Context) {
	uidStr := c.GetString("uid")
	uid, err := strconv.Atoi(uidStr)
	if err != nil || uid <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户 ID"})
		return
	}
	user, err := Svc.FindUserInfo(c.Request.Context(), int64(uid))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户未找到"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// 获取管理员仪表盘信息
func getAdminDashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "这是管理员仪表盘"})
}
