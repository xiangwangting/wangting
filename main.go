package main

import (
	"github.com/e421083458/golang_common/lib"
	"log"
	"os"
	"os/signal"
	"syscall"
	"wangting/pkg/service"
)

//入口程序
func main() {
	setup()
	service.HttpServerRun()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	service.HttpServerStop()
}

//加载资源配置
func setup() {
	if err := lib.Init("conf/local/"); err != nil {
		log.Fatal(err)
	}
}
