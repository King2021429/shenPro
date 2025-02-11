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
		ctx.JSON(http.StatusOK, errorcode.BuildErrorResponse(ctx, errCode))
		return
	}
	ctx.JSON(http.StatusOK, errorcode.BuildDataResponse(ctx, resp))
}

// EditArticle 编辑文章
func EditArticle(ctx *gin.Context) {
	var article model.EditArticleReq
	err := ctx.ShouldBindJSON(&article)
	if err != nil {
		fmt.Println(err)
	}

	uidStr := ctx.GetString("userID")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	resp, err := Svc.EditArticle(ctx.Request.Context(), &article, uid)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err})
		return
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
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// GetArticleList 获取文章列表
func GetArticleList(ctx *gin.Context) {
	var getArticleListReq model.GetArticleListReq
	err := ctx.ShouldBindJSON(&getArticleListReq)
	if err != nil {
		fmt.Println(err)
		return
	}

	uidStr := ctx.GetString("userID")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	if uid == 0 {
		ctx.JSON(http.StatusOK, gin.H{"error": "uid为0"})
		return
	}
	resp, err := Svc.GetArticleList(ctx.Request.Context(), &getArticleListReq)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// GetArticleInfo 获取文章详情
func GetArticleInfo(ctx *gin.Context) {
	var getArticleByIdReq model.GetArticleByIdReq
	err := ctx.ShouldBindJSON(&getArticleByIdReq)
	if err != nil {
		fmt.Println(err)
		return
	}

	uidStr := ctx.GetString("userID")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	resp, err := Svc.GetArticleById(ctx.Request.Context(), &getArticleByIdReq, uid)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// LikeArticle 点赞文章
func LikeArticle(ctx *gin.Context) {
	var likeArticleReq model.LikeArticleReq
	err := ctx.ShouldBindJSON(&likeArticleReq)
	if err != nil {
		fmt.Println(err)
		return
	}
	uidStr := ctx.GetString("userID")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	resp, errCode := Svc.LikeArticle(ctx.Request.Context(), &likeArticleReq, uid)
	if errCode != 0 {
		ctx.JSON(http.StatusOK, gin.H{"error": errCode})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

func GetLikeList(ctx *gin.Context) {

}
