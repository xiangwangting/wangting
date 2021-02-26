package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"wangting/conf"
	"wangting/http/middleWare"
	"wangting/http/route"
	"wangting/pkg/enum/envEnum"
)

//系统配置
var ConfigMap conf.Config

//公共初始化
func init(){
	//环境，初始化配置
	ConfigMap.ENV = envEnum.LOCAL
	ConfigMap.Init()
	//连接数据库
	initDB()
}

//入口程序
func main() {
	s := gin.Default()
	// 注册中间件
	s.Use(middleWare.Handel())
	//初始化路由
	route.Load(s)
	s.Run(":8000")
}

//初始化数据库操作
func initDB(){
	DB, _ := gorm.Open(ConfigMap.DB.CONNECTION, ConfigMap.DB.USERNAME+":"+ConfigMap.DB.PASSWORD+"@tcp("+ConfigMap.DB.HOST+":"+ConfigMap.DB.PORT+")/"+ConfigMap.DB.DATABASE+"?charset=utf8")
	defer DB.Close()
}


