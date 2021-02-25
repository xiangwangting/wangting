package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	ENV string
	DB struct {
		CONNECTION string `yaml:"DB_CONNECTION"`
		HOST       string `yaml:"DB_HOST"`
		PORT       string `yaml:"DB_PORT"`
		DATABASE   string `yaml:"DB_DATABASE"`
		USERNAME   string `yaml:"DB_USERNAME"`
		PASSWORD   string `yaml:"DB_PASSWORD"`
	}
}

//构造函数初始化
func (config *Config) Init() *Config{
	initDb(config)
	return config
}
//初始化数据库配置
func initDb(config *Config){
	content, err := ioutil.ReadFile("conf/db/"+ config.ENV +".yaml")
	if err != nil {
		log.Fatalf("解析config.yaml读取错误: %v", err)
	}
	if yaml.Unmarshal(content, &config.DB) != nil {
		log.Fatalf("解析config.yaml出错: %v", err)
	}
}
