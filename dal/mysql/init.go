package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() {
	var err error
	//DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN)) //&gorm.Config{
	//	PrepareStmt:            true,
	//	SkipDefaultTransaction: true,
	//},
	DB, err = gorm.Open(mysql.Open("root:Sss15946768062@tcp(47.100.224.26:3306)/tiktok8062?charset=utf8&parseTime=True&loc=Local"))
	println(DB == nil)
	if err != nil {
		panic(err)
	}
}
