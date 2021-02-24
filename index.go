package main

import (
	"fmt"
	"github.com/unknwon/goconfig"
	"github.com/syyongx/php2go"
	"wangting/app/constants"
)

const ENV  = "local"

func main() {
	php2go := php2go.Md5("123123")
	fmt.Println(php2go)
	fmt.Println(constants.EnvType{}.Dev)
	initDb()
}

func initDb() {
	cfg, err := goconfig.LoadConfigFile("config/dbconf." + ENV + ".ini")
	if err != nil {
		fmt.Println(err)
	}
	//db_host, err := cfg.GetValue("", "DB_HOST")
	//if err != nil {
	//	fmt.Println(err)
	//}
	db_connection, err := cfg.GetValue("", "DB_CONNECTION")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ENV, db_connection)
	//db_port, err := cfg.GetValue("", "DB_PORT")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//db_database, err := cfg.GetValue("", "DB_DATABASE")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//db_username, err := cfg.GetValue("", "DB_USERNAME")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//db_password, err := cfg.GetValue("", "DB_PASSWORD")
	//if err != nil {
	//	fmt.Println(err)
	//}

}
