package handler

import "wangting/http/route"

func (h *Helper) DemoHandler() (r *route.Router) {
	return &route.Router{
		Param:  nil, //所需要的参数
		Path:   "/",                //路由路径
		Method: "GET",              //方法
	}
}