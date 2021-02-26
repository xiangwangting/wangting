package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/syyongx/php2go"
)

//加载路由
func Load(router *gin.Engine) {
	initApiRoute(router)
	router.GET("/", handelApi)
}

func initApiRoute(router *gin.Engine) {
	//接口
	api := router.Group("api/v1")
	api.GET("/*any", handelApi)
	api.POST("/*any", handelApi)
	api.PUT("/*any", handelApi)
	api.DELETE("/*any", handelApi)
	api.PATCH("/*any", handelApi)
	//静态页面
	template := router.Group("template")
	template.GET("/*any", hendelTemplate)
}

func handelApi(c *gin.Context) {
	fmt.Println("接口开始执行")
	path := c.Request.URL.Path
	path = php2go.StrReplace("/api/v1/", "", path, 9999999999)
	//pathArr := php2go.Explode("/", path)
	//routePackage := pathArr[0]
	//function := pathArr[1]
	//routePackage.function()
	c.String(200, path)
	fmt.Println("接口执行完毕")
}

func hendelTemplate(c *gin.Context){

}
