package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shenyue-gin/app/shenyue/model"
)

// AIConversationStart 开启会话
func AIConversationStart(ctx *gin.Context) {
	resp, _ := Svc.AIChatStart(ctx.Request.Context())
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// AIConversationSendMsg 输入对话
func AIConversationSendMsg(ctx *gin.Context) {
	var conversationSendMsgReq model.ConversationSendMsgReq
	err := ctx.ShouldBindJSON(&conversationSendMsgReq)
	if err != nil {
		fmt.Println(err)
	}
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
	err := ctx.ShouldBindJSON(&conversationDeleteReq)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := Svc.AIChatDelete(ctx.Request.Context(), &conversationDeleteReq)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, gin.H{"error": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}
