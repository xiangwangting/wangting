package middleWare

import (
	"errors"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"wangting/http/controller"
	"wangting/pkg/compent"
	"wangting/pkg/enum/envEnum"
	"wangting/pkg/enum/responseCodeEnum"
)

// RecoveryMiddleware捕获所有panic，并且返回错误信息
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//先做一下日志记录
				fmt.Println(string(debug.Stack()))
				compent.ComLogNotice(c, "_com_panic", map[string]interface{}{
					"error": fmt.Sprint(err),
					"stack": string(debug.Stack()),
				})
				if lib.GetConfEnv() == envEnum.PRODUCT {
					controller.ResponseError(c, responseCodeEnum.InternalErrorCode, errors.New("内部错误"))
					return
				} else {
					controller.ResponseError(c, responseCodeEnum.InternalErrorCode, errors.New(fmt.Sprint(err)))
					return
				}
			}
		}()
		c.Next()
	}
}