package route

import (
	"github.com/gin-gonic/gin"
)

//路由加载
func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello go")
	})
	return router
}
