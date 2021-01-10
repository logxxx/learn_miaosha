package controllers

import (
	"app/common"
	"app/datamodels"
	"app/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"strconv"
)

type ProductController struct {
	Ctx iris.Context
	ProductService services.IProductService
}

func (p *ProductController) GetAll() mvc.View {
	productArr, _ := p.ProductService.GetAllProduct()
	return mvc.View{
		Name: "product/view.html", //渲染模板
		Data: iris.Map{ //变量
			"productArray": productArr,
		},
	}
}

// 修改商品
func (p *ProductController) PostUpdate() {
	product := &datamodels.Product{}
	p.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName:"imooc"})
	if err := dec.Decode(p.Ctx.Request().Form, product); err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	err := p.ProductService.UpdateProduct(product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}

func (p *ProductController) GetAdd() mvc.View {
	return mvc.View {
		Name: "product/add.html", //要渲染的模板

	}
}

func (p *ProductController) PostAdd() {
	product := &datamodels.Product{}
	p.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName:"imooc"})
	if err := dec.Decode(p.Ctx.Request().Form, product); err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	_, err := p.ProductService.InsertProduct(product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}

func (p *ProductController) GetManager() mvc.View {
	 idString := p.Ctx.URLParam("id")
	 id, err := strconv.ParseInt(idString, 10, 64)
	 if err != nil {
	 	p.Ctx.Application().Logger().Debug(err)
	 }
	 product, err := p.ProductService.GetProductById(id)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	return mvc.View{
		Name:"product/manager.html",
		Data: iris.Map{
			"product": product,
		},
	}
}

func (p *ProductController) GetDelete() {
	idString := p.Ctx.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	ok := p.ProductService.DeleteProductById(id)
	if !ok {
		p.Ctx.Application().Logger().Debug("删除商品id=%失败", id)
	}
	p.Ctx.Application().Logger().Debug("删除商品id=%成功", id)
	p.Ctx.Redirect("/product/all")
}