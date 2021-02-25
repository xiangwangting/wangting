package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
	"log"
	"wangting/conf"
	"wangting/http/route"
)

//系统配置
var configMap conf.Config

//入口程序
func main() {
	//环境，初始化配置
	configMap.ENV = "local"
	configMap.Init()
	//连接数据库
	db, err := gorm.Open(configMap.DB.CONNECTION, configMap.DB.USERNAME+":"+configMap.DB.PASSWORD+"@tcp("+configMap.DB.HOST+":"+configMap.DB.PORT+")/"+configMap.DB.DATABASE+"?charset=utf8")
	if err != nil{
		log.Fatalf("gorm链接错误: %v", err)
	}
	defer db.Close()

	s := gin.Default()
	route.Load(s)
	s.Run(":8000")
}

