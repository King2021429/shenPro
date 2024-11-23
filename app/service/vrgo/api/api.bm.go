// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: shenyue-gin/app/service/vrgo/api/api.proto

/*
Package api is a generated blademaster stub package.
This code was generated with kratos/tool/protobuf/protoc-gen-bm v0.1.

package 命名使用 {appid}.{version} 的方式, version 形如 v1, v2 ..

It is generated from these files:

	shenyue-gin/app/service/vrgo/api/api.proto
*/
package api

import (
	"context"

	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathDemoMatchList = "/vr/match/list"
var PathDemoGradeInfo = "/vr/grade/info"
var PathDemoEconomic = "/vr/economic/info"
var PathDemoInfo = "/vr/now/info"
var PathDemoWebhook = "/vr/webhook"

// DemoBMServer is the server API for Demo service.
type DemoBMServer interface {
	MatchList(ctx context.Context, req *MatchListReq) (resp *MatchListResp, err error)

	GradeInfo(ctx context.Context, req *GradeInfoReq) (resp *GradeInfoResp, err error)

	Economic(ctx context.Context, req *EconomicReq) (resp *EconomicResp, err error)

	Info(ctx context.Context, req *InfoReq) (resp *InfoResp, err error)

	Webhook(ctx context.Context, req *WebhookReq) (resp *WebhookResp, err error)
}

var DemoSvc DemoBMServer

func demoMatchList(c *bm.Context) {
	p := new(MatchListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := DemoSvc.MatchList(c, p)
	c.JSON(resp, err)
}

func demoGradeInfo(c *bm.Context) {
	p := new(GradeInfoReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := DemoSvc.GradeInfo(c, p)
	c.JSON(resp, err)
}

func demoEconomic(c *bm.Context) {
	p := new(EconomicReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := DemoSvc.Economic(c, p)
	c.JSON(resp, err)
}

func demoInfo(c *bm.Context) {
	p := new(InfoReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := DemoSvc.Info(c, p)
	c.JSON(resp, err)
}

func demoWebhook(c *bm.Context) {
	p := new(WebhookReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := DemoSvc.Webhook(c, p)
	c.JSON(resp, err)
}

// RegisterDemoBMServer Register the blademaster route
func RegisterDemoBMServer(e *bm.Engine, server DemoBMServer) {
	DemoSvc = server
	e.GET("/vr/match/list", demoMatchList)
	e.GET("/vr/grade/info", demoGradeInfo)
	e.GET("/vr/economic/info", demoEconomic)
	e.GET("/vr/now/info", demoInfo)
	e.POST("/vr/webhook", demoWebhook)
}
