package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wangting/conf"
	"wangting/http/middleWare"
	"wangting/http/route"
	"github.com/e421083458/golang_common/lib"
)

//入口程序
func main() {
	lib.InitModule("conf/local/",[]string{"mysql","base","redis"})
	defer lib.Destroy()
	lib.Log.TagInfo(lib.NewTrace(), lib.DLTagUndefind, map[string]interface{}{"message": "todo sth"})
	time.Sleep(time.Second)
	return
	s := gin.Default()
	// 注册中间件
	s.Use(middleWare.Handel())
	//初始化路由
	route.Load(s)
	logo := "██╗  ██╗██╗    ██╗██╗    ██╗████████╗\n╚██╗██╔╝██║    ██║██║    ██║╚══██╔══╝\n ╚███╔╝ ██║ █╗ ██║██║ █╗ ██║   ██║   \n ██╔██╗ ██║███╗██║██║███╗██║   ██║   \n██╔╝ ██╗╚███╔███╔╝╚███╔███╔╝   ██║   \n╚═╝  ╚═╝ ╚══╝╚══╝  ╚══╝╚══╝    ╚═╝   \n                                     "
	fmt.Println(logo)
	serverRun(s)
}

//启动服务
func serverRun(s *gin.Engine) {
	confMap := conf.Config
	//服务
	srv := &http.Server{
		Addr:    ":" + confMap.App.PORT,
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(confMap.App.TIMEOUT) * time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown:", err)
	}
	log.Println("server exiting")
}

