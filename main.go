package main

import (
	"context"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"wangting/conf"
	"wangting/http/middleWare"
	"wangting/http/route"
)

//入口程序
func main() {
	setup()
	s := gin.Default()
	// 注册中间件
	s.Use(middleWare.Handel())
	//初始化路由
	route.Load(s)
	logo := "██╗  ██╗██╗    ██╗██╗    ██╗████████╗\n╚██╗██╔╝██║    ██║██║    ██║╚══██╔══╝\n ╚███╔╝ ██║ █╗ ██║██║ █╗ ██║   ██║   \n ██╔██╗ ██║███╗██║██║███╗██║   ██║   \n██╔╝ ██╗╚███╔███╔╝╚███╔███╔╝   ██║   \n╚═╝  ╚═╝ ╚══╝╚══╝  ╚══╝╚══╝    ╚═╝   \n                                     "
	fmt.Println(logo)
	serverRun(s)
}

var initOnce sync.Once

//启动服务
func serverRun(s *gin.Engine) {
	//服务
	srv := &http.Server{
		Addr:    ":" + lib.ViperConfMap["app"].GetString("port"),
		Handler: s,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutdown server ...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(lib.ViperConfMap["app"].GetInt("read_timeout"))*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown:", err)
	}
	log.Println("server exiting")
}

func setup() {
	initOnce.Do(func() {
		if err:=lib.Init("conf/local/");err!=nil{
			log.Fatal(err)
		}
		appconf := &conf.AppConf{}
		err:=lib.ParseLocalConfig("app.toml",appconf)
		if err!=nil{
			log.Fatal(err)
		}
		defer lib.Destroy()
	})
}
