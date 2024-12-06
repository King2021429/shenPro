package service

import (
	"context"
	"fmt"
	"shenyue-gin/app/blog/model/api"
)

// WebHookSendMsg
func (s *Service) WebHookSendMsg(ctx context.Context, req *api.SendMessageData) (resp *api.WebhookResp, err error) {
	fmt.Printf("成功转换为ContentType1结构体，%+v", req)
	return
}

// WebHookEnterDirectMsg ENTER_DIRECT_MSG
func (s *Service) WebHookEnterDirectMsg(ctx context.Context, req *api.SendMessageData) (resp *api.WebhookResp, err error) {

	fmt.Printf("成功转换为ContentType1结构体，%+v", req)
	return
}

// WebHookCloseMsg CLOSE_MSG
func (s *Service) WebHookCloseMsg(ctx context.Context, req *api.SendMessageData) (resp *api.WebhookResp, err error) {

	fmt.Printf("成功转换为ContentType1结构体，%+v", req)

	return
}
