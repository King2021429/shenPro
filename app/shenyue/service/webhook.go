package service

import (
	"context"
	"fmt"
	"shenyue-gin/app/shenyue/model/api"
	"shenyue-gin/app/shenyue/model/errorcode"
)

// WebHookSendMsg SEND_MSG
func (s *Service) WebHookSendMsg(ctx context.Context, req *api.SendMsg) (resp *api.WebhookResp, code int64) {
	resp = &api.WebhookResp{}
	fmt.Printf("成功转换为ContentType1结构体，%+v", req)
	return resp, errorcode.ERROR_USERNAME_USED
}

// WebHookEnterDirectMsg ENTER_DIRECT_MSG
func (s *Service) WebHookEnterDirectMsg(ctx context.Context, req *api.EnterDirectMsg) (resp *api.WebhookResp, code int64) {
	resp = &api.WebhookResp{}
	fmt.Printf("成功转换为ContentType1结构体，%+v", req)
	return
}

// WebHookCloseMsg CLOSE_MSG
func (s *Service) WebHookCloseMsg(ctx context.Context, req *api.CloseMsg) (resp *api.WebhookResp, code int64) {
	resp = &api.WebhookResp{}
	fmt.Printf("成功转换为ContentType1结构体，%+v", req)
	return
}
