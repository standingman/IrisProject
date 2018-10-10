package main

import (
	"IrisProject/routes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)


func newApp() *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	tmpl := iris.HTML("./view/", ".html")
	app.StaticWeb("/public", "./view/public/")
	app.RegisterView(tmpl)
	app.Configure(routes.Configure)
	return app
}

func main() {
	db, _ := gorm.Open("mysql", "root:mysql.com2017@tcp(106.14.96.72:3306)/common_db?charset=utf8&parseTime=True&loc=Local")
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	defer db.Close()
	app := newApp()
	app.Run(iris.Addr(":8080"))
}


