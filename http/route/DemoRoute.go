package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wangting/http/controller/demoController"
)

func (h *Helper) DemoUserRoute() (r *Router) {
	return &Router{
		Path:   "/user",                //路由路径
		Method: http.MethodGet,              //方法
		Handlers: []gin.HandlerFunc{
			demoController.Index,
		},
	}
}

func (h *Helper) DemoErrorRoute() (r *Router) {
	return &Router{
		Path:   "/error",                //路由路径
		Method: http.MethodGet,              //方法
		Handlers: []gin.HandlerFunc{
			demoController.DemoError,
		},
	}
}

func (h *Helper) PostRoute() (r *Router) {
	return &Router{
		Path:   "/post",                //路由路径
		Method: http.MethodPost,              //方法
		Handlers: []gin.HandlerFunc{
			demoController.Post,
		},
	}
}


