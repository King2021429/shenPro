package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shenyue-gin/app/shenyue/model"
	"strconv"
)

// CreateArticle 创建文章
func CreateArticle(ctx *gin.Context) {
	var article model.CreateArticleReq
	err := ctx.ShouldBindJSON(&article)
	if err != nil {
		fmt.Println(err)
	}

	uidStr := ctx.GetString("userID")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	resp, err := Svc.CreateArticle(ctx.Request.Context(), &article, uid)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, gin.H{"error": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}
