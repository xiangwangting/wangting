package route

import (
	"github.com/gin-gonic/gin"
	"wangting/http/controller/demoController"
	"wangting/http/middleWare"
)

//路由加载
func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello go")
	})
	router.GET("/userinfo", func(c *gin.Context) {
		c.String(200, "userinfo")
	})

	apiNormalGroup := router.Group("/api")
	apiNormalGroup.Use(middleWare.Handel())
	apiNormalRegister(apiNormalGroup)
	return router
}

//api,普通路由注册
func apiNormalRegister(router *gin.RouterGroup) {
	demoController := demoController.DemoController{}
	router.Any("demoinfo", demoController.DemoInfo)
}

