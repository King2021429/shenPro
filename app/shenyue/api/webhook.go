package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shenyue-gin/app/shenyue/model"
)

func Webhook(ctx *gin.Context) {
	var msg model.WebhookReq
	err := ctx.ShouldBindJSON(&msg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch msg.Event {
	case model.EventSendMsg:
		var c *model.SendMsg
		if err = json.Unmarshal(msg.Content, &c); err != nil {
			fmt.Println(err)
		}
		_, _ = Svc.WebHookSendMsg(ctx, c)
		ctx.JSON(http.StatusOK, gin.H{"message": model.EventSendMsg})

	case model.EventDirectMsg:
		var c *model.EnterDirectMsg
		if err = json.Unmarshal(msg.Content, &c); err != nil {
			fmt.Println(err)
		}
		_, _ = Svc.WebHookEnterDirectMsg(ctx, c)
		ctx.JSON(http.StatusOK, gin.H{"message": model.EventDirectMsg})

	case model.EventCloseMSG:
		var c *model.CloseMsg
		if err = json.Unmarshal(msg.Content, &c); err != nil {
			fmt.Println(err)
		}
		_, _ = Svc.WebHookCloseMsg(ctx, c)
		ctx.JSON(http.StatusOK, gin.H{"message": model.EventCloseMSG})

	default:
		ctx.JSON(http.StatusBadRequest, msg.Content)
		return
	}

}
