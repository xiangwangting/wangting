package service
//
//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"wangting/conf"
//)
//
//type Service struct {
//	Srv http.Server
//}
//
//func (receiver *Service) Init(handeler *gin.Engine)  {
//	config := conf.Config
//	receiver.Srv := &http.Server{
//		Addr : ":"+config.App.PORT,
//		Handler: handeler,
//	}
//}
////启动服务
//func (receiver *Service) Star(){
//	go func() {
//
//	}()
//}