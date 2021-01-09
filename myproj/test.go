package main

import (
	"app/common"
	"app/datamodels"
	"fmt"
)

func main() {
	data := map[string]string{
		"id":"1",
		"product_name": "测试结构体",
		"product_num": "1000",
		"product_image": "www.baidu.com",
		"product_url": "www.google.com",
	}

	product := &datamodels.Product{}

	common.DataToStructByTagSql(data, product)

	fmt.Printf("%+v", product)
}
