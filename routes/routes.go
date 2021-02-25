package routes

import (
	"github.com/gin-gonic/gin"
	"wangting/app/conrtoller/demoController"
)

func Load(router *gin.Engine)  {
	api := router.Group("api/v1")
	api.GET("login", demoController.Login)
}
