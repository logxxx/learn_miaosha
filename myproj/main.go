package main

import (
	"app/web/controllers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./web/views", ".html"))
	//注册控制器
	mvc.New(app.Party("/hello")).Handle(new(
		controllers.MovieController))
	app.Run(
		iris.Addr("localhost:8080"),
		)
}