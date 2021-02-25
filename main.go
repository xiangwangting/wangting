package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"wangting/conf"
)

//系统配置
var configMap conf.Config

//入口程序
func main() {
	//环境，初始化配置
	configMap.ENV = "local"
	configMap.Init()
	initOrm()
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	//s := gin.Default()
	//routes.Load(s)
	//// 路由组1 ，处理GET请求
	//client1 := r.Group("/api/v1")
	//client2 := r.Group("/api/v2")
	//{
	//	//client1.GET("/", demoController.Login)
	//	client1.GET("login", demoController.Login)
	//	client1.GET("submit", demoController.Submit)
	//}
	//{
	//	client2.POST("login", demoController.Login)
	//	client2.POST("submit", demoController.Submit)
	//}
	//s.Run(":8000")
}

func initOrm() {
	//连接数据库
	db, err := gorm.Open(configMap.DB.CONNECTION, configMap.DB.USERNAME+":"+configMap.DB.PASSWORD+"@tcp("+configMap.DB.HOST+":"+configMap.DB.PORT+")/"+configMap.DB.DATABASE+"?charset=utf8")
	if err != nil{
		log.Fatalf("gorm链接错误: %v", err)
	}
	defer db.Close()
}
