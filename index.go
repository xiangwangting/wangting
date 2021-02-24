package main

import (
	"fmt"
	"github.com/unknwon/goconfig"
)

func main() {
	initDb()
	//// 1.创建路由
	//// 默认使用了2个中间件Logger(), Recovery()
	//r := gin.Default()
	//// 路由组1 ，处理GET请求
	//v1 := r.Group("/v1")
	//// {} 是书写规范
	//{
	//	v1.GET("/login", login)
	//	v1.GET("submit", submit)
	//}
	//v2 := r.Group("/v2")
	//{
	//	v2.POST("/login", login)
	//	v2.POST("/submit", submit)
	//}
	//r.Run(":8000")
}

func initDb() {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		fmt.Println(err)
	}
	db_host, err := cfg.GetValue("", "DB_HOST")
	if err != nil {
		fmt.Println(err)
	}
	db_port, err := cfg.GetValue("", "DB_PORT")
	if err != nil {
		fmt.Println(err)
	}
	db_database, err := cfg.GetValue("", "DB_DATABASE")
	if err != nil {
		fmt.Println(err)
	}
	db_username, err := cfg.GetValue("", "DB_USERNAME")
	if err != nil {
		fmt.Println(err)
	}
	db_password, err := cfg.GetValue("", "DB_PASSWORD")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db_host, db_port, db_database, db_username, db_password)
}
