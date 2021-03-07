package demoController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"wangting/http/controller"
	"wangting/pkg/model"
)

type DemoController struct {
}

func Index(c *gin.Context) {
	var user model.User
	controller.ResponseSuccess(c, user)
}

//demo获取信息
func DemoError(c *gin.Context) {
	controller.ResponseError(c, http.StatusInternalServerError, errors.New("手动抛出异常"))
}



