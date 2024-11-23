package main

import (
	"context"
	"flag"
	"fmt"
	_ "go.uber.org/automaxprocs"
	"os"
	"os/signal"
	"shenyue-gin/app/service/vrgo/internal/server/http"
	"shenyue-gin/app/service/vrgo/internal/service"
	"syscall"
	"time"
)

func main() {
	flag.Parse()

	// 设置环境变量来禁用CPU统计
	os.Setenv("DISABLE_CPU_STATS", "true")

	svc := service.New()
	// init http server
	httpSrv := http.New(svc)
	// init grpc server
	//grpcSrv := grpc.New(svc)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		fmt.Println("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
			if err := httpSrv.Shutdown(ctx); err != nil {
			}
			//if err := grpcSrv.Shutdown(ctx); err != nil {
			//}
			svc.Close()
			cancel()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
