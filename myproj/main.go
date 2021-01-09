package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./web/views", ".html"))
	//注册控制器
	app.Run(
		iris.Addr("localhost:8080"),
		)
}