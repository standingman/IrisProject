package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"irisPorject/controllers"
)

func newApp() *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	mvc.New(app).Handle(new(controllers.ExampleController))
	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}


