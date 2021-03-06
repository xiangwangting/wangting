package route

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
	"wangting/http/controller"
	"wangting/http/middleWare"
)


type Router struct {
	Param    Parameter
	Path     string
	Method   string
	Handlers []gin.HandlerFunc
}

//路由加载
func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	gin := gin.Default()
	apiGroup := gin.Group("/api")
	apiGroup.Use(middleWare.Handel(),middleWare.RecoveryMiddleware())
	//Build(new(handler.Helper), apiGroup)


	//router.GET("/", func(c *gin.Context) {
	//	c.String(200, "hello go")
	//})

	{
		apiNormalRegister(apiGroup)
	}
	return gin
}

////api,普通路由注册
func apiNormalRegister(router *gin.RouterGroup) {
	//demoController := demoController.DemoController{}
	//router.Any("demoinfo", demoController.DemoInfo)
	apiController :=  controller.ApiController{}
	router.Any("/*any",apiController.RegisterApi)
}

func Build(h interface{}, r gin.IRoutes) {
	valueOfh := reflect.ValueOf(h)
	numMethod := valueOfh.NumMethod()
	for i := 0; i < numMethod; i++ {
		rt := valueOfh.Method(i).Call(nil)[0].Interface().(*Router)
		rt.AddHandler(r)
	}
}

func (rt *Router) AddHandler(r gin.IRoutes) {
	if rt.Param != nil {
		replace := false
		for i, handler := range rt.Handlers {
			if handler == nil {
				//如果内部存在GenHandlerFunc表示占位，那么就替换
				rt.Handlers[i] = rt.genHandlerFunc()
				replace = true
			}
		}
		if !replace {
			rt.Handlers = append(rt.Handlers, rt.genHandlerFunc())
		}
	}
	switch strings.ToUpper(rt.Method) {
	case "GET":
		r.GET(rt.Path, rt.Handlers...)
	case "POST":
		r.POST(rt.Path, rt.Handlers...)
	case "PUT":
		r.PUT(rt.Path, rt.Handlers...)
	case "PATCH":
		r.PATCH(rt.Path, rt.Handlers...)
	case "HEAD":
		r.HEAD(rt.Path, rt.Handlers...)
	case "OPTIONS":
		r.OPTIONS(rt.Path, rt.Handlers...)
	case "DELETE":
		r.DELETE(rt.Path, rt.Handlers...)
	case "ANY":
		r.Any(rt.Path, rt.Handlers...)
	default:
		panic("Method: " + rt.Method + " is error")
	}
}

func (rt *Router) genHandlerFunc() gin.HandlerFunc {
	// 取变量a的反射类型对象
	paramsType := reflect.TypeOf(rt.Param).Elem()
	// 根据反射类型对象创建类型实例
	handler := func(c *gin.Context) {
		newParam := reflect.New(paramsType).Interface().(Parameter)
		newParam.BeforeBind(c)
		newParam.Bind(c, newParam)
		if newParam.Error() == nil {
			newParam.Service(c)
		}
		newParam.Result(c)
	}
	return handler
}


