package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	ENV string
	App struct {
		NAME string `yaml:"APP_NAME"`
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

//构造函数初始化
func (config *Config) Init() *Config {
	initApp(config)
	initDb(config)
	initRedis(config)
	return config
}

func initApp(config *Config) {
	content, err := ioutil.ReadFile( "conf/app.yaml")
	if err != nil {
		log.Fatalf("解析dapp.yaml读取错误: %v", err)
	}
	if yaml.Unmarshal(content, &config.App) != nil {
		log.Fatalf("解析app.yaml出错: %v", err)
	}
}

//初始化数据库配置
func initDb(config *Config) {
	content, err := ioutil.ReadFile("conf/db/" + config.ENV + ".yaml")
	if err != nil {
		log.Fatalf("解析db/config.yaml读取错误: %v", err)
	}
	if yaml.Unmarshal(content, &config.DB) != nil {
		log.Fatalf("解析db/config.yaml出错: %v", err)
	}
}

func initRedis(config *Config) {
	content, err := ioutil.ReadFile("conf/redis/" + config.ENV + ".yaml")
	if err != nil {
		log.Fatalf("解析redis/config.yaml读取错误: %v", err)
	}
	if yaml.Unmarshal(content, &config.Redis) != nil {
		log.Fatalf("解析redis/config.yaml出错: %v", err)
	}
}
