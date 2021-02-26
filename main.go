package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"wangting/conf"
	"wangting/http/route"
	"wangting/pkg/enum/envEnum"
)

//系统配置
var configMap conf.Config

//入口程序
func main() {
	//环境，初始化配置
	configMap.ENV = envEnum.LOCAL
	configMap.Init()
	fmt.Println(configMap)
	//连接数据库
	initDB()
	//初始化路由
	s := gin.Default()
	route.Load(s)
	s.Run(":8000")
}

func initDB(){
	db, _ := gorm.Open(configMap.DB.CONNECTION, configMap.DB.USERNAME+":"+configMap.DB.PASSWORD+"@tcp("+configMap.DB.HOST+":"+configMap.DB.PORT+")/"+configMap.DB.DATABASE+"?charset=utf8")
	defer db.Close()
}

