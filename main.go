package main

import (
	"github.com/e421083458/golang_common/lib"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"wangting/http/route"
)

//入口程序
func main() {
	setup()
	route.HttpServerRun()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	route.HttpServerStop()
}

var initOnce sync.Once

//加载资源配置
func setup() {
	initOnce.Do(func() {
		if err:=lib.Init("conf/local/");err!=nil{
			log.Fatal(err)
		}
		defer lib.Destroy()
	})
}
