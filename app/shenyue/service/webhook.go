package service

import (
	"context"
	"fmt"
	"shenyue-gin/app/shenyue/errorcode"
	"shenyue-gin/app/shenyue/model"
)

// WebHookSendMsg SEND_MSG
func (s *Service) WebHookSendMsg(ctx context.Context, req *model.SendMsg) (resp *model.WebhookResp, code int64) {
	resp = &model.WebhookResp{}
	fmt.Printf("成功转换为ContentType1结构体，%+v", req)
	return resp, errorcode.ERROR_USERNAME_USED
}

// WebHookEnterDirectMsg ENTER_DIRECT_MSG
func (s *Service) WebHookEnterDirectMsg(ctx context.Context, req *model.EnterDirectMsg) (resp *model.WebhookResp, code int64) {
	resp = &model.WebhookResp{}
	fmt.Printf("成功转换为ContentType1结构体，%+v", req)
	return
}

// WebHookCloseMsg CLOSE_MSG
func (s *Service) WebHookCloseMsg(ctx context.Context, req *model.CloseMsg) (resp *model.WebhookResp, code int64) {
	resp = &model.WebhookResp{}
	fmt.Printf("成功转换为ContentType1结构体，%+v", req)
	return
}
