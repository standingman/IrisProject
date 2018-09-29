package main

import (
	"IrisProject/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
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
	mvc.New(app).Handle(new(controllers.ExampleController))
	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}


