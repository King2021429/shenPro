package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	server "shenyue-gin/app/shenyue/api"
	"shenyue-gin/app/shenyue/model"
	"shenyue-gin/app/shenyue/model/api"
)

func Webhook(ctx *gin.Context) {
	var msg api.WebhookReq
	err := ctx.ShouldBindJSON(&msg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch msg.Event {
	case model.EventSendMsg:
		var c *api.SendMsg
		if err = json.Unmarshal(msg.Content, &c); err != nil {
			fmt.Println(err)
		}
		_, _ = server.Svc.WebHookSendMsg(ctx, c)
		ctx.JSON(http.StatusOK, gin.H{"error": "test"})

	case model.EventDirectMsg:
		var c *api.EnterDirectMsg
		if err = json.Unmarshal(msg.Content, &c); err != nil {
			fmt.Println(err)
		}
		_, _ = server.Svc.WebHookEnterDirectMsg(ctx, c)
		ctx.JSON(http.StatusOK, gin.H{"error": "test"})

	case model.EventCloseMSG:
		var c *api.CloseMsg
		if err = json.Unmarshal(msg.Content, &c); err != nil {
			fmt.Println(err)
		}
		_, _ = server.Svc.WebHookCloseMsg(ctx, c)
		ctx.JSON(http.StatusOK, gin.H{"error": "test"})

	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "test"})
		return
	}

}
