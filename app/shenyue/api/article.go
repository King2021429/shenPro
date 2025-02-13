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

	uidStr := ctx.GetString("Uid")
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

	uidStr := ctx.GetString("Uid")
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

	uidStr := ctx.GetString("Uid")
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

	uidStr := ctx.GetString("Uid")
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

	uidStr := ctx.GetString("Uid")
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
	uidStr := ctx.GetString("Uid")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	resp, errCode := Svc.LikeArticle(ctx.Request.Context(), &likeArticleReq, uid)
	if errCode != 0 {
		ctx.JSON(http.StatusOK, gin.H{"error": errCode})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// GetLikeList 根据用户id获取点赞列表
func GetLikeList(ctx *gin.Context) {
	var likeArticleListReq model.LikeArticleListReq
	err := ctx.ShouldBindJSON(&likeArticleListReq)
	if err != nil {
		fmt.Println(err)
		return
	}
	uidStr := ctx.GetString("Uid")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	if uid == 0 {
		ctx.JSON(http.StatusOK, gin.H{"error": "uid为0"})
		return
	}
	resp, errCode := Svc.GetLikeList(ctx.Request.Context(), &likeArticleListReq)
	if errCode != 0 {
		ctx.JSON(http.StatusOK, gin.H{"error": errCode})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})

}

// FavoriteArticle 收藏文章
func FavoriteArticle(ctx *gin.Context) {
	var favoriteArticleReq model.FavoriteArticleReq
	err := ctx.ShouldBindJSON(&favoriteArticleReq)
	if err != nil {
		fmt.Println(err)
		return
	}
	uidStr := ctx.GetString("Uid")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	resp, errCode := Svc.FavoriteArticle(ctx.Request.Context(), &favoriteArticleReq, uid)
	if errCode != 0 {
		ctx.JSON(http.StatusOK, gin.H{"error": errCode})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// GetFavoriteList 根据用户id获取收藏列表
func GetFavoriteList(ctx *gin.Context) {
	var favoriteArticleListReq model.FavoriteArticleListReq
	err := ctx.ShouldBindJSON(&favoriteArticleListReq)
	if err != nil {
		fmt.Println(err)
		return
	}
	uidStr := ctx.GetString("Uid")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	if uid == 0 {
		ctx.JSON(http.StatusOK, gin.H{"error": "uid为0"})
		return
	}
	resp, errCode := Svc.GetFavoriteList(ctx.Request.Context(), &favoriteArticleListReq)
	if errCode != 0 {
		ctx.JSON(http.StatusOK, gin.H{"error": errCode})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}
