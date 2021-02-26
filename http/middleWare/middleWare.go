package middleWare

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 定义中间
func Handel() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		loginInit(c)
		// 执行函数
		c.Next()
		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
	}
}

//初始化登陆信息
func loginInit(c *gin.Context){
	fmt.Println("初始化登陆信息")
}
