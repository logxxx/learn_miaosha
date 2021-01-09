package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	//注册模板
	template := iris.HTML(
		"./myproj/backend/web/views", ".html").Layout(
			"shared/layout.html").Reload(
				true)
	app.RegisterView(template)

	//设置模板文件
	app.StaticWeb("/assets",
		"./backend/web/assets")

	//出现异常，跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message",
			ctx.Values().GetStringDefault("message", "访问的页面出错"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	//注册控制器

	//启动服务
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		)
}