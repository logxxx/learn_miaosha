package repositories

import (
	"app/common"
	"app/datamodels"
	"database/sql"
	"fmt"
)

type IOrderRepository interface {
	Conn() error
	Insert(order *datamodels.Order) (int64, error)
	Delete(int64) bool
	Update(order *datamodels.Order) error
	SelectByKey(int64) (*datamodels.Order, error)
	SelectAll() ([]*datamodels.Order, error)
	SelectAllWithInfo() (map[int]map[string]string, error)
}

type OrderManagerRepository struct {
	table string
	mysqlConn *sql.DB
}

func NewOrderManagerRepository(table string, sql *sql.DB) IOrderRepository {
	return &OrderManagerRepository{
		table: table,
		mysqlConn: sql,
	}
}

//连接函数
func (o *OrderManagerRepository) Conn() error {
	if o.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		o.mysqlConn = mysql
	}
	if o.table == "" {
		o.table = "order"
	}
	return nil
}

//插入订单
func (o *OrderManagerRepository) Insert(order *datamodels.Order) (int64, error) {
	if err := o.Conn(); err != nil {
		return 0, err
	}

	sql := fmt.Sprintf("INSERT INTO %s SET user_id=?, product_id=?, order_status=?", o.table)
	stmt, err := o.mysqlConn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(order.UserId, order.ProductId, order.OrderStatus)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}

func (o *OrderManagerRepository) Delete(orderId int64) bool {
	if err := o.Conn(); err != nil {
		return false
	}

	sql := fmt.Sprintf("DELETE FROM %v where id=?", o.table)

	stmt, err := o.mysqlConn.Prepare(sql)
	if err != nil {
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(orderId)
	if err != nil {
		return false
	}

	return true
}

func (o *OrderManagerRepository) Update(order *datamodels.Order) error {
	if err := o.Conn(); err != nil {
		return err
	}

	sql := fmt.Sprintf("UPDATE %v SET user_id=?, product_id=?, order_status=? WHERE id=%v", o.table, order.Id)

	stmt, err := o.mysqlConn.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.UserId, order.ProductId, order.OrderStatus)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderManagerRepository) SelectByKey(orderId int64) (*datamodels.Order, error) {
	if err := o.Conn(); err != nil {
		return nil, err
	}

	sql := fmt.Sprintf("SELECT * from %v WHERE id=%v", o.table, orderId)

	row, err := o.mysqlConn.Query(sql)
	if err != nil {
		return nil, err
	}

	order := new(datamodels.Order)
	result := common.GetResultRow(row)
	if len(result) <= 0 {
		return order, nil
	}

	common.DataToStructByTagSql(result, order)

	return order, nil
}

func (o *OrderManagerRepository) SelectAll() ([]*datamodels.Order, error) {
	if err := o.Conn(); err != nil {
		return nil, err
	}

	sql := fmt.Sprintf("SELECT * FROM %v", o.table)

	rows, err := o.mysqlConn.Query(sql)
	if err != nil {
		return nil, err
	}

	orders := make([]*datamodels.Order, 0)
	results := common.GetResultRows(rows)
	if len(results) <= 0 {
		return orders, nil
	}

	for i := range results {
		order := new(datamodels.Order)
		common.DataToStructByTagSql(results[i], orders)
		orders = append(orders, order)
	}

	return orders, nil

}

func (o *OrderManagerRepository) SelectAllWithInfo() (map[int]map[string]string, error) {
	if err := o.Conn(); err != nil {
		return nil, err
	}

	sql := "SELECT o.id, p.product_name, o.order_status FROM order as o" +
		" LEFT JOIN product as p ON o.product_id=p.id"

	rows, err := o.mysqlConn.Query(sql)
	if err != nil {
		return nil, err
	}

	return common.GetResultRows(rows), nil
}