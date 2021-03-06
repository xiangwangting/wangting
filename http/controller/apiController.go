package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
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
		panic("controller "+controlName+" not found")
		//ResponseError(c, responseCodeEnum.InternalErrorCode, errors.New("controller "+controlName+" not found"))
	}
}