package demoHandel

import (
	"errors"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"net/http"
	"wangting/http/handel"
	"wangting/pkg/model"
)

type DemoController struct {
}

func Index(c *gin.Context) {
	handel.ResponseSuccess(c, "this is demo index")
}

//demo获取信息
func DemoError(c *gin.Context) {
	handel.ResponseError(c, http.StatusInternalServerError, errors.New("手动抛出异常"))
}

//demo获取信息
func Post(c *gin.Context) {
	if string, ok := c.GetPostForm("aa"); ok {
		handel.ResponseSuccess(c, string)
	} else {
		handel.ResponseError(c, http.StatusInternalServerError, errors.New("缺少aa参数"))
	}
}

func Gorm(c *gin.Context) {
	////获取链接池
	user := model.User{}
	lib.GORMDefaultPool.First(&user,2)
	fmt.Println(user)
	user2 := model.User{}
	lib.GORMDefaultPool.First(&user2,3)
	fmt.Println(user2)
	handel.ResponseSuccess(c, user)
}
