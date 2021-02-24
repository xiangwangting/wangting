package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/goconfig"
	"wangting/app/conrtoller/demoController"
)

const ENV  = "local"

func main() {
	initOrm()
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 路由组1 ，处理GET请求
	client1 := r.Group("/api/v1")
	client2 := r.Group("/api/v1")
	{
		//client1.GET("/", demoController.Login)
		client1.GET("login", demoController.Login)
		client1.GET("submit", demoController.Submit)
	}
	{
		client2.POST("login", demoController.Login)
		client2.POST("submit", demoController.Submit)
	}
	r.Run(":8000")
}

func initOrm() {
	cfg, err := goconfig.LoadConfigFile("config/dbconf." + ENV + ".ini")
	if err != nil {
		fmt.Println(err)
	}
	//db_host, err := cfg.GetValue("", "DB_HOST")
	//if err != nil {
	//	fmt.Println(err)
	//}
	db_connection, err := cfg.GetValue("", "DB_CONNECTION")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ENV, db_connection)
	//db_port, err := cfg.GetValue("", "DB_PORT")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//db_database, err := cfg.GetValue("", "DB_DATABASE")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//db_username, err := cfg.GetValue("", "DB_USERNAME")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//db_password, err := cfg.GetValue("", "DB_PASSWORD")
	//if err != nil {
	//	fmt.Println(err)
	//}

}

