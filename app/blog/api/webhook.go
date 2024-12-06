package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shenyue-gin/app/blog/model/api"
)

func Webhook(ctx *gin.Context) {
	var msg api.WebhookReq
	err := ctx.ShouldBindJSON(&msg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch msg.Event {
	case "SEND_MSG":
		var c *api.SendMessageData
		if err := json.Unmarshal(msg.Content, &c); err != nil {
			fmt.Println(err)
		}
		_, _ = Svc.WebHookSendMsg(ctx, c)

	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
