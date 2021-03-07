package http

import (
	"context"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"wangting/http/route"
)

var (
	HttpSrvHandler *http.Server
)
//http服务启动
func HttpServerRun() {
	gin.SetMode(lib.ConfBase.DebugMode)
	r := route.InitRouter()
	HttpSrvHandler = &http.Server{
		Addr:           lib.GetStringConf("base.http.addr"),
		Handler:        r,
		ReadTimeout:    time.Duration(lib.GetIntConf("base.http.read_timeout")) * time.Second,
		WriteTimeout:   time.Duration(lib.GetIntConf("base.http.write_timeout")) * time.Second,
		MaxHeaderBytes: 1 << uint(lib.GetIntConf("base.http.max_header_bytes")),
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n",lib.GetStringConf("base.http.addr"))
		fmt.Println("\n      ___           ___           ___           ___           ___           ___           ___              \n     /\\__\\         /\\  \\         /\\__\\         /\\__\\         /\\__\\         /\\__\\         /\\__\\             \n    /:/ _/_        \\:\\  \\       /:/  /        /:/  /        /:/ _/_       /:/ _/_       /:/ _/_            \n   /:/ /\\  \\        \\:\\  \\     /:/  /        /:/  /        /:/ /\\__\\     /:/ /\\  \\     /:/ /\\  \\           \n  /:/ /::\\  \\   ___  \\:\\  \\   /:/  /  ___   /:/  /  ___   /:/ /:/ _/_   /:/ /::\\  \\   /:/ /::\\  \\          \n /:/_/:/\\:\\__\\ /\\  \\  \\:\\__\\ /:/__/  /\\__\\ /:/__/  /\\__\\ /:/_/:/ /\\__\\ /:/_/:/\\:\\__\\ /:/_/:/\\:\\__\\         \n \\:\\/:/ /:/  / \\:\\  \\ /:/  / \\:\\  \\ /:/  / \\:\\  \\ /:/  / \\:\\/:/ /:/  / \\:\\/:/ /:/  / \\:\\/:/ /:/  /         \n  \\::/ /:/  /   \\:\\  /:/  /   \\:\\  /:/  /   \\:\\  /:/  /   \\::/_/:/  /   \\::/ /:/  /   \\::/ /:/  /          \n   \\/_/:/  /     \\:\\/:/  /     \\:\\/:/  /     \\:\\/:/  /     \\:\\/:/  /     \\/_/:/  /     \\/_/:/  /           \n     /:/  /       \\::/  /       \\::/  /       \\::/  /       \\::/  /        /:/  /        /:/  /            \n     \\/__/         \\/__/         \\/__/         \\/__/         \\/__/         \\/__/         \\/__/             \n")
		if err := HttpSrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", lib.GetStringConf("base.http.addr"), err)
		}
	}()
}
//http服务停止
func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil{
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
