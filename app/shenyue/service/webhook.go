package service

import (
	"context"
	"fmt"
	"shenyue-gin/app/shenyue/model"
)

// WebHookSendMsg SEND_MSG
func (s *Service) WebHookSendMsg(ctx context.Context, req *model.SendMsg, webhook *model.WebhookReq) (resp *model.WebhookResp, code int64) {
	resp = &model.WebhookResp{}
	fmt.Println("new WebHookSendMsg req:")
	fmt.Println(webhook.Event)
	fmt.Printf("成功转换为ContentType1结构体，%+v\n", req)
	fmt.Println(webhook.Timestamp)
	return
}

// WebHookEnterDirectMsg ENTER_DIRECT_MSG
func (s *Service) WebHookEnterDirectMsg(ctx context.Context, req *model.EnterDirectMsg, webhook *model.WebhookReq) (resp *model.WebhookResp, code int64) {
	resp = &model.WebhookResp{}
	fmt.Println("new ENTER_DIRECT_MSG req:")
	fmt.Println(webhook.Event)
	fmt.Printf("成功转换为ContentType1结构体，%+v\n", req)
	fmt.Println(webhook.Timestamp)
	return
}

// WebHookCloseMsg CLOSE_MSG
func (s *Service) WebHookCloseMsg(ctx context.Context, req *model.CloseMsg, webhook *model.WebhookReq) (resp *model.WebhookResp, code int64) {
	resp = &model.WebhookResp{}
	fmt.Println("new CLOSE_MSG req:")
	fmt.Println(webhook.Event)
	fmt.Printf("成功转换为ContentType1结构体，%+v\n", req)
	fmt.Println(webhook.Timestamp)
	return
}
