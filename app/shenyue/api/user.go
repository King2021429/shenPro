package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	// 检查用户名是否已存在
	//var existingUser model.User
	//if err := db.Where("username =?", user.Username).First(&existingUser).Error; err == nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "用户名已存在"})
	//	return
	//}
	//
	//// 创建新用户
	//if err := db.Create(&user).Error; err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}

	c.JSON(http.StatusCreated, gin.H{"message": "注册成功"})

}

// 用户登录
func loginUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//var existingUser model.User
	//if err := db.Where("username =? AND password =?", user.Username, user.Password).First(&existingUser).Error; err != nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
	//	return
	//}

	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成令牌失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

// 获取用户信息
func getUserInfo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("userID")
		id, err := strconv.Atoi(userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户 ID"})
			return
		}

		var user model.User
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "用户未找到"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

// 获取管理员仪表盘信息
func getAdminDashboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "这是管理员仪表盘"})
	}
}
