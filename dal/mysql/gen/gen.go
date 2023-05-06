package main

import (
	"fmt"
	"github.com/AgSword/simpleDouyin/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Querier Dynamic SQL
type Querier interface {
	// FilterWithNameAndRole SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
	db, err := gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:      "dal/mysql/dao",
		ModelPkgPath: "",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(db) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	User := g.GenerateModel("user")
	Comment := g.GenerateModel("comment")
	Favorite := g.GenerateModel("favorite")
	Relation := g.GenerateModel("relation")
	Token := g.GenerateModel("token")
	Video := g.GenerateModel("video")
	g.ApplyBasic(User, Comment, Favorite, Token, Relation, Video)

	// Generate Type Safe API with Dynamic SQL defined on Querier interface
	g.ApplyInterface(func(Querier) {}, User, Comment, Favorite, Token, Relation, Video)

	// Generate the code
	g.Execute()
}
