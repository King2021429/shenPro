package grpc

import (
	pb "shenyue-gin/app/service/vrgo/api"

	"go-common/library/net/rpc/warden"
)

// New new a grpc server.
func New(svc pb.DemoServer) (ws *warden.Server) {
	//var (
	//	cfg warden.ServerConfig
	//	ct  paladin.TOML
	//)
	cfg := &warden.ServerConfig{
		Addr:    "0.0.0.0:8000",
		Timeout: 1,
	}
	//if err := paladin.Get("grpc.toml").Unmarshal(&ct); err != nil {
	//	panic(err)
	//}
	//if err := ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
	//	panic(err)
	//}
	//cfg = &warden.ServerConfig{}
	ws = warden.NewServer(cfg)
	pb.RegisterDemoServer(ws.Server(), svc)
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return
}
