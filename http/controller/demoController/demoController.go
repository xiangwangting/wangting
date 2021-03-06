package demoController

import (
	"errors"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"wangting/http/controller"
	"wangting/pkg/model"
)

type DemoController struct {
}


//demo获取信息
func (demoController *DemoController) DemoInfo(c *gin.Context) {
	var user model.User
	lib.GORMDefaultPool.First(&user,1)
	controller.ResponseSuccess(c, user)
}
//demo获取信息
func (demoController *DemoController) DemoError(c *gin.Context) {
	controller.ResponseError(c, 500, errors.New("手动抛出异常"))
}



