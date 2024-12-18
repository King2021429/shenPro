package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shenyue-gin/app/shenyue/middleware"
	"shenyue-gin/app/shenyue/model"
)

// 注册 新增
// 审批 更新
// 注销 删除

func Find(ctx *gin.Context) {
	id := ctx.Query("id")
	name := ctx.Query("name")
	fmt.Println(id, name)
	err := Svc.SendUserEmail(ctx.Request.Context())
	if err != nil {
		fmt.Println(err)
		ctx.JSON(200, err)
	}
	ctx.JSON(200, name)
}

func Register(ctx *gin.Context) {
	var user model.UserReq
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
	}
	err = Svc.SaveUser(ctx.Request.Context(), &user)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(200, err)
	}
	ctx.JSON(200, user)
}

// 定义登录接口
func Login(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	// 从请求体中获取登录信息
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 假设这里进行了数据库验证等操作，验证用户名和密码是否正确，这里只是示例，直接返回成功
	if user.Username == "testuser" && user.Password == "testpassword" {
		// 生成Token
		token, err := middleware.GenerateToken(user.Username, user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		//c.JSON(
		//	http.StatusOK, gin.H{
		//		"status":  code,
		//		"data":    data,
		//		"message": errmsg.GetErrMsg(code),
		//	},
		//)
	}

}
