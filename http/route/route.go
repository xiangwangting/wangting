package route

import (
	"github.com/gin-gonic/gin"
	"github.com/syyongx/php2go"
)

//加载路由
func Load(router *gin.Engine)  {
	api := router.Group("api/v1")
	api.GET("/*any",handel)
}

func handel(c *gin.Context){
	path := c.Request.URL.Path
	path = php2go.StrReplace("/api/v1/","",path,9999999999)
	routePackage := php2go.Explode("/",path)[0]
	c.String(200, routePackage)
}
