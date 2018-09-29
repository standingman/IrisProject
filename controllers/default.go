package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type ExampleController struct{}

func (c *ExampleController) GetApi() mvc.Result {
	/*	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Welcome</h1>",
	}*/
	return mvc.View{
		Name: "index.html",
	}
}

func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
	anyMiddlewareHere := func(ctx iris.Context) {
		ctx.Application().Logger().Warnf("Inside /custom_path")
		ctx.Next()
	}
	b.Handle("GET", "/custom_path", "CustomHandlerWithoutFollowingTheNamingGuide", anyMiddlewareHere)

}

func (c *ExampleController) CustomHandlerWithoutFollowingTheNamingGuide() string {
	return "hello from the custom handler without following the naming guide"
}
