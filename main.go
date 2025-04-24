package main

import (
	"gohttp101/data"
	"gohttp101/server"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(192.168.114.129:3306)/gethttp101?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&data.Todo{})
	if err != nil {
		return
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "../gohttp101/dao",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(data.Todo{})

	// Generate the code
	g.Execute()

	svr := server.GetServer(db)
	svr.RegisterRoute(server.Get)
	svr.RegisterRoute(server.Post)
	svr.RegisterRoute(server.Update)
	//svr.RegisterRoute(server.Post)
	svr.Start()

}
