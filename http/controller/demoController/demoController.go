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
	//获取链接池
	dbpool, err := lib.GetGormPool("default")
	if err != nil {
		controller.ResponseError(c, http.StatusBadGateway,err)
	}
	user := model.User{}
	dbpool.Find(&user)
	controller.ResponseSuccess(c, dbpool)
}



