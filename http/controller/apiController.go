package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
	"wangting/pkg/enum/responseCodeEnum"
)

type ApiController struct {

}

type app struct {
	w http.ResponseWriter
}


func (api *ApiController) RegisterApi(c *gin.Context){
	path := strings.Split(c.Request.URL.Path,"/")
	controlName := path[2]
	methodBase := "Index"
	if len(path) >= 4{
		methodBase = path[3]
	}
	if methodBase == "" {
		methodBase = "Index"
	}
	app := &app{c.Writer}
	rType := reflect.TypeOf(app)
	rValue := reflect.ValueOf(app)
	_, exist := rType.MethodByName(controlName)
	fmt.Println(rValue,rType)
	if exist {
		ResponseSuccess(c,methodBase+"存在")
		//args := []reflect.Value{rValue}
		//method.Func.Call(args)
	} else {
		ResponseError(c, responseCodeEnum.InternalErrorCode, errors.New("controller "+controlName+" not found"))
	}
}