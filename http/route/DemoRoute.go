package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wangting/http/handel/demoHandel"
)

func (h *Helper) DemoUserRoute() (r *Router) {
	return &Router{
		Path:   "/index",                //路由路径
		Method: http.MethodGet,              //方法
		Handlers: []gin.HandlerFunc{
			demoHandel.Index,
		},
	}
}

func (h *Helper) DemoErrorRoute() (r *Router) {
	return &Router{
		Path:   "/error",                //路由路径
		Method: http.MethodGet,              //方法
		Handlers: []gin.HandlerFunc{
			demoHandel.DemoError,
		},
	}
}

func (h *Helper) PostRoute() (r *Router) {
	return &Router{
		Path:   "/post",                //路由路径
		Method: http.MethodPost,              //方法
		Handlers: []gin.HandlerFunc{
			demoHandel.Post,
		},
	}
}

func (h *Helper) GormRoute() (r *Router) {
	return &Router{
		Path:   "/gorm",                //路由路径
		Method: http.MethodGet,              //方法
		Handlers: []gin.HandlerFunc{
			demoHandel.Gorm,
		},
	}
}


