package demoController

import (
	"errors"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"net/http"
	"wangting/http/controller"
	"wangting/pkg/model"
)

type DemoController struct {
}

func Index(c *gin.Context) {
	controller.ResponseSuccess(c, "this is demo index")
}

//demo获取信息
func DemoError(c *gin.Context) {
	controller.ResponseError(c, http.StatusInternalServerError, errors.New("手动抛出异常"))
}

//demo获取信息
func Post(c *gin.Context) {
	string , ok := c.GetPostForm("aa")
	if ok {
		controller.ResponseSuccess(c, string)
	}else{
		controller.ResponseError(c, http.StatusInternalServerError, errors.New("缺少aa参数"))
	}
}

func Gorm(c *gin.Context) {
	var user model.User
	lib.GORMDefaultPool.First(&user,3)
	controller.ResponseSuccess(c, user)
}



