package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shenyue-gin/app/shenyue/errorcode"
	"shenyue-gin/app/shenyue/model"
	"strconv"
)

// CreateComment 创建评论
func CreateComment(ctx *gin.Context) {
	var comment model.CreateCommentReq
	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		fmt.Println(err)
	}

	uidStr := ctx.GetString("uid")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	resp, errCode := Svc.CreateComment(ctx.Request.Context(), &comment, uid)
	if errCode != 0 {
		ctx.JSON(http.StatusOK, errorcode.BuildErrorResponse(ctx, errCode))
		return
	}
	ctx.JSON(http.StatusOK, errorcode.BuildDataResponse(ctx, resp))
}

// EditComment 编辑评论
func EditComment(ctx *gin.Context) {
	var comment model.EditCommentReq
	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		fmt.Println(err)
	}

	uidStr := ctx.GetString("uid")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	resp, err := Svc.EditComment(ctx.Request.Context(), &comment, uid)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// DeleteComment 删除评论
func DeleteComment(ctx *gin.Context) {
	var deleteCommentReq model.DeleteCommentReq
	err := ctx.ShouldBindJSON(&deleteCommentReq)
	if err != nil {
		fmt.Println(err)
	}

	uidStr := ctx.GetString("uid")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	resp, err := Svc.DeleteComment(ctx.Request.Context(), &deleteCommentReq, uid)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// GetCommentList 根据文章内容获取评论列表
func GetCommentList(ctx *gin.Context) {
	var getCommentListReq model.GetCommentListReq
	err := ctx.ShouldBindJSON(&getCommentListReq)
	if err != nil {
		fmt.Println(err)
	}
	resp, errCode := Svc.GetCommentsByArticleId(ctx.Request.Context(), &getCommentListReq)
	if errCode != 0 {
		ctx.JSON(http.StatusOK, errorcode.BuildErrorResponse(ctx, errCode))
		return
	}
	ctx.JSON(http.StatusOK, errorcode.BuildDataResponse(ctx, resp))
}
