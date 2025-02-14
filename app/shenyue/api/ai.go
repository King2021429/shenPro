package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shenyue-gin/app/shenyue/model"
	"strconv"
)

// AIConversationStart 开启会话
func AIConversationStart(ctx *gin.Context) {
	uidStr := ctx.GetString("uid")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	resp, _ := Svc.AIChatStart(ctx.Request.Context(), uid)
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// AIConversationSendMsg 输入对话
func AIConversationSendMsg(ctx *gin.Context) {
	var conversationSendMsgReq model.ConversationSendMsgReq
	uidStr := ctx.GetString("uid")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	err := ctx.ShouldBindJSON(&conversationSendMsgReq)
	if err != nil {
		fmt.Println(err)
	}
	conversationSendMsgReq.Uid = uid
	resp, err := Svc.AIChatSendMsg(ctx.Request.Context(), &conversationSendMsgReq)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, gin.H{"error": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// AIConversationDelete 删除对话
func AIConversationDelete(ctx *gin.Context) {
	var conversationDeleteReq model.ConversationDeleteReq
	uidStr := ctx.GetString("uid")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	err := ctx.ShouldBindJSON(&conversationDeleteReq)
	if err != nil {
		fmt.Println(err)
	}
	conversationDeleteReq.Uid = uid
	resp, err := Svc.AIChatDelete(ctx.Request.Context(), &conversationDeleteReq)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, gin.H{"error": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// AIConversationList 获取对话列表
func AIConversationList(ctx *gin.Context) {
	var conversationListReq model.ConversationListReq
	uidStr := ctx.GetString("uid")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	//err := ctx.ShouldBindJSON(&conversationListReq)
	//if err != nil {
	//	fmt.Println(err)
	//}
	conversationListReq.Uid = uid
	resp, err := Svc.AIChatList(ctx.Request.Context(), &conversationListReq)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, gin.H{"error": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}
