package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"wangting/pkg/enum/envEnum"
)

var Config struct {
	ENV string
	App struct {
		NAME string `yaml:"APP_NAME"`
		PORT string `yaml:"APP_PORT"`
	}
	DB struct {
		CONNECTION string `yaml:"DB_CONNECTION"`
		HOST       string `yaml:"DB_HOST"`
		PORT       string `yaml:"DB_PORT"`
		DATABASE   string `yaml:"DB_DATABASE"`
		USERNAME   string `yaml:"DB_USERNAME"`
		PASSWORD   string `yaml:"DB_PASSWORD"`
	}
	Redis struct {
		HOST       string `yaml:"REDIS_HOST"`
		PASSWORD   string `yaml:"REDIS_PASSWORD"`
		PORT       string `yaml:"REDIS_PORT"`
		CACHE_DB   string `yaml:"REDIS_CACHE_DB"`
		SESSION_DB string `yaml:"REDIS_SESSION_DB"`
	}
}

func init(){
	Config.ENV = envEnum.LOCAL
	initApp()
	initDb()
	initRedis()
}
//
func initApp() {
	content, err := ioutil.ReadFile( "conf/app.yaml")
	if err != nil {
		log.Fatalf("解析dapp.yaml读取错误: %v", err)
	}
	if yaml.Unmarshal(content, &Config.App) != nil {
		log.Fatalf("解析app.yaml出错: %v", err)
	}
}
//
////初始化数据库配置
func initDb() {
	content, err := ioutil.ReadFile("conf/db/" + Config.ENV + ".yaml")
	if err != nil {
		log.Fatalf("解析db/config.yaml读取错误: %v", err)
	}
	if yaml.Unmarshal(content, &Config.DB) != nil {
		log.Fatalf("解析db/config.yaml出错: %v", err)
	}
}
//
func initRedis() {
	content, err := ioutil.ReadFile("conf/redis/" + Config.ENV + ".yaml")
	if err != nil {
		log.Fatalf("解析redis/config.yaml读取错误: %v", err)
	}
	if yaml.Unmarshal(content, &Config.Redis) != nil {
		log.Fatalf("解析redis/config.yaml出错: %v", err)
	}
}
