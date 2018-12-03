package routes

import (
	"github.com/kataras/iris"
)

func GetIndexHandler(ctx iris.Context) {
	ctx.ViewData("Title", "Index Page")
	ctx.View("index.html")
}
