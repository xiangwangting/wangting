package demoController

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"wangting/http/middleWare"
	"wangting/pkg/model"
)

type DemoController struct {
}


//demo获取信息
func (demoController *DemoController) DemoInfo(c *gin.Context) {
	var user model.User
	lib.GORMDefaultPool.First(&user,2)
	middleWare.ResponseSuccess(c, user)
}
//demo获取信息
func (demoController *DemoController) DemoError(c *gin.Context) {
	middleWare.ResponseError(c, 2001,nil)
}



