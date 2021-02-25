package demoController

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func Submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func Any(c *gin.Context){
	c.String(200, c.Request.RequestURI)
}
