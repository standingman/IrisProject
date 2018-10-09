package routes

import "github.com/kataras/iris"

func Configure(a *iris.Application){
	a.Get("/",GetIndexHandler)
}