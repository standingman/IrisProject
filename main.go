package main

import (
	"./global"
	"./routes"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func newApp() *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	tmpl := iris.HTML("./view/", ".html")
	app.StaticWeb("/view", "./view/static/")
	app.RegisterView(tmpl)
	app.Configure(routes.Configure)
	return app
}


func main() {
	global.InitDB()
	app := newApp()
	//global.DB.Create(&models.Searchinfo{Type:1,Value:"公司",TypeId:2})
	app.Run(iris.Addr(":8080"))
	//defer global.DB.Close()
}


