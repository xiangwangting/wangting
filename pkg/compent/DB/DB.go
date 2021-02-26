package DB

import (
	"github.com/jinzhu/gorm"
	"wangting/conf"
)

//gorm链接
var GORM *gorm.DB

//初始化
func init() {
	ConfigMap := conf.Config
	client, _ := gorm.Open(ConfigMap.DB.CONNECTION, ConfigMap.DB.USERNAME+":"+ConfigMap.DB.PASSWORD+"@tcp("+ConfigMap.DB.HOST+":"+ConfigMap.DB.PORT+")/"+ConfigMap.DB.DATABASE+"?charset=utf8")
	defer client.Close()
	GORM = client
}
