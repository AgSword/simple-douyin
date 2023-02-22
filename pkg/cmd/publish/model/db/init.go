package db

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

func Init() {
	sqlDB, err := sql.Open("mysql", "root:Sss15946768062@tcp(47.100.224.26:3306)/tiktok8062?parseTime=True")
	Db, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Panicln("err:", err.Error())
	}
}
