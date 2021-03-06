package controller

import (
	"encoding/json"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"strings"
)

type ResponseCode int

//1000以下为通用码，1000以上为用户自定义码
const (
	SuccessCode       ResponseCode = 200 //成功
	UndefErrorCode    ResponseCode = 400 //为定义的错误
	InternalErrorCode ResponseCode = 500 //内部错误
	InvalidRequestErrorCode ResponseCode = 401  //无效请求
	CustomizeCode           ResponseCode = 1000 //业务自定义错误码
	GROUPALL_SAVE_FLOWERROR ResponseCode = 2001
)

//http,响应结构体
type Response struct {
	Code    ResponseCode `json:"code"` //状态吗
	Message string       `json:"message"` //提示信息
	Data    interface{}  `json:"data"` //数据
	Stack   interface{}  `json:"stack"` //堆栈
}

//http，响应错误
func ResponseError(c *gin.Context, code ResponseCode, err error) {
	stack := ""
	if c.Query("is_debug") == "1" || lib.ConfBase.DebugMode == "debug" {
		stack = strings.Replace(fmt.Sprintf("%+v", err), err.Error()+"\n", "", -1)
	}

	resp := &Response{Code: code, Message: err.Error(), Data: "", Stack: stack}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	c.AbortWithError(200, err)
}

//http，响应成功
func ResponseSuccess(c *gin.Context, data interface{}) {
	resp := &Response{Code: 200, Message: "成功", Data: data}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
