package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shenyue-gin/app/shenyue/errorcode"
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
	resp, errCode := Svc.CreateArticle(ctx.Request.Context(), &article, uid)
	if errCode != 0 {
		fmt.Println(errorcode.GetErrMsg(errCode))
		ctx.JSON(http.StatusOK, gin.H{"error": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// DeleteArticle 删除文章
func DeleteArticle(ctx *gin.Context) {
	var deleteArticleReq model.DeleteArticleReq
	err := ctx.ShouldBindJSON(&deleteArticleReq)
	if err != nil {
		fmt.Println(err)
	}

	uidStr := ctx.GetString("userID")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	resp, err := Svc.DeleteArticle(ctx.Request.Context(), &deleteArticleReq, uid)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}
