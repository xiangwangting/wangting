package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/syyongx/php2go"
	"wangting/conf"
	"wangting/pkg/compent/DB"
)

//加载路由
func Load(router *gin.Engine) {
	initApiRoute(router)
	initWeb(router)
}

//初始化接口路由
func initApiRoute(router *gin.Engine) {
	//接口
	api := router.Group("api/v1")
	api.GET("/*any", handelApi)
	api.POST("/*any", handelApi)
	api.PUT("/*any", handelApi)
	api.DELETE("/*any", handelApi)
	api.PATCH("/*any", handelApi)

}

//初始化网站
func initWeb(router *gin.Engine){
	//静态页面
	template := router.Group("web")
	template.GET("/*any", hendelWeb)
	router.GET("/", hendelWeb)
}
//处理接口
func handelApi(c *gin.Context) {
	fmt.Println("接口开始执行")
	db := DB.GORM
	fmt.Println(db)
	path := c.Request.URL.Path
	path = php2go.StrReplace("/api/v1/", "", path, 9999999999)
	//pathArr := php2go.Explode("/", path)
	//routePackage := pathArr[0]
	//function := pathArr[1]
	//routePackage.function()
	c.String(200, path)
}
//处理页面
func hendelWeb(c *gin.Context){
	path := c.Request.URL.Path
	c.String(200, "["+conf.Config.App.NAME+"]站点:"+path)
}
