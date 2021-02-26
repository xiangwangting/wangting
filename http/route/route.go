package route

import (
	"github.com/gin-gonic/gin"
	"github.com/syyongx/php2go"
)

//加载路由
func Load(router *gin.Engine) {
	initApiRoute(router)
	router.GET("/",handel)
}

func initApiRoute(router *gin.Engine) {
	api := router.Group("api/v1")
	api.GET("/*any", handel)
	api.POST("/*any", handel)
	api.PUT("/*any", handel)
	api.DELETE("/*any", handel)
	api.PATCH("/*any", handel)
}

func handel(c *gin.Context) {
	path := c.Request.URL.Path
	path = php2go.StrReplace("/api/v1/", "", path, 9999999999)
	pathArr := php2go.Explode("/", path)
	routePackage := pathArr[0]
	function := pathArr[1]
	//routePackage.function()
	c.String(200, routePackage+"/"+function)
}
