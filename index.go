package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/goconfig"
)

const ENV  = "local"

func main() {
	//initDb()
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 路由组1 ，处理GET请求
	v1 := r.Group("/api/v1")
	// {} 是书写规范
	{
		v1.GET("login", login)
		v1.GET("submit", submit)
	}
	r.Run(":8000")
}

func initDb() {
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

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
