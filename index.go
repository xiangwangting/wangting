package main

import (
	"fmt"
	"github.com/unknwon/goconfig"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	initDb()
}

func initDb() {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
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
	DB, err := gorm.Open(sqlite.Open(db_connection), &gorm.Config{})
	if err == nil {
		fmt.Println("open db sucess", DB)
	} else {
		fmt.Println("open db error ", err)
	}

}
