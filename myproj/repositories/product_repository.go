package repositories

import (
	"app/common"
	"app/datamodels"
	"database/sql"
	"fmt"
	"strconv"
)

type IProduct interface {
	Conn() (error) //链接数据库
	Insert(product *datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(product *datamodels.Product) error
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll()([]*datamodels.Product, error)
}

type ProductManager struct {
	table string
	mysqlConn *sql.DB
}

func NewProductManager(table string, db *sql.DB) (IProduct) {
	return &ProductManager{
		table:     table,
		mysqlConn: db,
	}
}

// 数据库连接
func (p *ProductManager) Conn() (err error) {
	if p.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		p.mysqlConn = mysql
	}
	if p.table == "" {
		p.table = "product"
	}
	return
}

//插入
func (p *ProductManager) Insert(product *datamodels.Product) (productId int64, err error) {
	//1.判断链接是否存在
	if err = p.Conn(); err != nil {
		return
	}

	//2.准备sql
	sql := "INSERT INTO  " + p.table + " SET product_name=?, product_num=?, product_image=?, product_url=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	//3.传入参数
	result, err := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return
	}
	productId, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}

//删除
func (p *ProductManager) Delete(productId int64) bool {
	//1.判断链接是否存在
	if err := p.Conn(); err != nil {
		return false
	}

	sql := "DELETE FROM  " + p.table + " where id =?"
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(productId)
	if err != nil {
		return false
	}
	return true
}

//商品更新
func (p *ProductManager) Update(product *datamodels.Product) (err error) {
	//1.判断链接是否存在
	if err = p.Conn(); err != nil {
		return
	}

	sql := "UPDATE  " + p.table + " SET product_name=?, product_num=?, product_image=?, product_url=? WHERE id="+strconv.FormatInt(product.ID, 10)

	fmt.Println("Update sql:", sql)
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return
	}

	return
}

//4.根据商品id查询商品
func (p *ProductManager) SelectByKey(productId int64) (productResult *datamodels.Product, err error) {
	productResult = new(datamodels.Product)
	//1.判断链接是否存在
	if err = p.Conn(); err != nil {
		return
	}

	sql := "SELECT * FROM " + p.table + " WHERE id=" + strconv.FormatInt(productId, 10)

	row, err := p.mysqlConn.Query(sql)
	if err != nil {
		return
	}
	defer row.Close()

	result := common.GetResultRow(row)
	if len(result) <= 0 {
		return
	}

	common.DataToStructByTagSql(result, productResult)

	return
}

//5.查询所有商品
func (p *ProductManager)SelectAll()(productArr []*datamodels.Product, err error) {
	productArr = make([]*datamodels.Product, 0)
	//1.判断链接是否存在
	if err = p.Conn(); err != nil {
		return
	}

	sql := "SELECT * FROM " + p.table

	rows, err := p.mysqlConn.Query(sql)
	if err != nil {
		return
	}
	defer rows.Close()

	results := common.GetResultRows(rows)

	for _, v := range results {
		product := new(datamodels.Product)
		common.DataToStructByTagSql(v, product)
		productArr = append(productArr, product)
	}

	return
}