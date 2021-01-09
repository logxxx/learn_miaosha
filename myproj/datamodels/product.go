package datamodels

type Product struct {
	// `json:"" sql:"" imooc:""`
	ID           int64  `json:"ID" sql:"id" imooc:"ID"`
	ProductName  string `json:"ProductName" sql:"product_name" imooc:"ProductName"`
	ProductNum   int64  `json:"ProductNum" sql:"product_num" imooc:"ProductNum"`
	ProductImage string `json:"ProductImage" sql:"product_image" imooc:"ProductImage"`
	ProductUrl   string `json:"ProductUrl" sql:"product_url" imooc:"ProductUrl"`
}
