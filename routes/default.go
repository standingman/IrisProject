package routes

import "iris-master"

func GetIndexHandler(ctx iris.Context) {
	ctx.ViewData("Title", "Index Page")
	ctx.View("index.html")
}
