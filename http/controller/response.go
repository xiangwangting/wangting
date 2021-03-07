package controller

import (
	"encoding/json"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"wangting/pkg/enum/envEnum"
)


//http,响应结构体
type Response struct {
	Code    int `json:"code"` //状态吗
	Message string       `json:"message"` //提示信息
	Data    interface{}  `json:"data"` //数据
	Stack   interface{}  `json:"stack"` //堆栈
}

//http，响应错误
func ResponseError(c *gin.Context, code int, err error) {
	stack := ""
	fmt.Println()
	if c.Query("is_debug") == "1" || lib.GetConfEnv() != envEnum.PRODUCT {
		stack = strings.Replace(fmt.Sprintf("%+v", err), err.Error()+"\n", "", -1)
	}

	resp := &Response{Code: code, Message: err.Error(), Data: "", Stack: stack}
	c.JSON(http.StatusOK, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	c.AbortWithError(http.StatusOK, err)
}

//http，响应成功
func ResponseSuccess(c *gin.Context, data interface{}) {
	resp := &Response{Code: http.StatusOK, Message: "成功", Data: data}
	c.JSON(http.StatusOK, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
