package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"wangting/http/middleWare"
	"wangting/http/route"
)

//入口程序
func main() {
	s := gin.Default()
	// 注册中间件
	s.Use(middleWare.Handel())
	//初始化路由
	route.Load(s)
	s.Run(":8000")
}


