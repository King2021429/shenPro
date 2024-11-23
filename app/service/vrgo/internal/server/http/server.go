package http

import (
	"fmt"
	"net/http"

	"go-common/library/log"
	bm "go-common/library/net/http/blademaster"
	pb "shenyue-gin/app/service/vrgo/api"
	"shenyue-gin/app/service/vrgo/internal/model"
	"shenyue-gin/app/service/vrgo/internal/service"
)

var svc *service.Service

// New new a bm server.
func New(s *service.Service) (engine *bm.Engine) {
	//var (
	//	cfg bm.ServerConfig
	//	ct  paladin.TOML
	//)
	//if err := paladin.Get("http.toml").Unmarshal(&ct); err != nil {
	//	panic(err)
	//}
	cfg := &bm.ServerConfig{
		Addr:    "0.0.0.0:8000",
		Timeout: 1,
	}
	//if _, err := toml.DecodeFile("http.toml", &ct); err != nil {
	//}
	//if err := ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
	//	panic(err)
	//}

	svc = s
	engine = bm.DefaultServer(cfg)
	pb.RegisterDemoBMServer(engine, s)
	initRouter(engine)
	err := engine.Start()
	if err != nil {
		fmt.Println(err)
	}
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/vrgo")
	{
		g.GET("/start", howToStart)
	}
}

func ping(ctx *bm.Context) {
	if _, err := svc.Ping(ctx, nil); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

// example for http request handler.
func howToStart(c *bm.Context) {
	k := &model.Kratos{
		Hello: "Hello Kratos!",
	}
	c.JSON(k, nil)
}
