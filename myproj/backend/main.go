package main

import (
	"app/backend/web/controllers"
	"app/common"
	"app/repositories"
	"app/services"
	"context"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"os"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	//注册模板
	wd, _ := os.Getwd()
	app.Logger().Infof("wd:%v", wd)
	template := iris.HTML(
		"./backend/web/views", ".html").Layout(
			"shared/layout.html").Reload(
				true)
	app.RegisterView(template)

	//设置模板文件
	app.HandleDir("/assets",
		"./backend/web/assets")

	//出现异常，跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message",
			ctx.Values().GetStringDefault("message", "访问的页面出错"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	//连接Mysql
	db, err := common.NewMysqlConn()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//注册控制器
	productRepository := repositories.NewProductManager("product", db)
	productService := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productService)
	product.Handle(new(controllers.ProductController))

	//启动服务
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		)
}