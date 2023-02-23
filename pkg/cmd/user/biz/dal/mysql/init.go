package mysql

import (
	// "github.com/AgSword/simpleDouyin/conf"

	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	// DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN)) //&gorm.Config{
	//	PrepareStmt:            true,
	//	SkipDefaultTransaction: true,
	//},
	sqlDB, err := sql.Open("mysql", "root:Sss15946768062@tcp(47.100.224.26:3306)/tiktok8062?parseTime=True")
	DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	println(DB == nil)
	if err != nil {
		panic(err)
	}
}
