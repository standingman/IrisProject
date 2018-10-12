package main

import (
	"IrisProject/global"
	"IrisProject/routes"
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
	global.InitDB()
	app := newApp()
	app.Run(iris.Addr(":8080"))
	defer global.DB.Close()
}


